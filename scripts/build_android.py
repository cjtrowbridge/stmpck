"""Android debug packaging, called only by the canonical build driver."""
from pathlib import Path
import os
import re
import shutil
import subprocess
import zipfile

from android_toolchain import java_environment

ROOT = Path(__file__).resolve().parent.parent
SDK = ROOT / '.tools/android-sdk'
NDK = SDK / 'ndk/27.2.12479018'
BUILD_TOOLS = SDK / 'build-tools/35.0.0'
PLATFORM = SDK / 'platforms/android-35/android.jar'


def unavailable_reason():
    if not PLATFORM.exists() or not (NDK / 'source.properties').exists() or not (BUILD_TOOLS / 'aapt2').exists():
        return 'Android SDK/NDK missing; run python scripts/setup_android.py on Linux'
    try:
        java_environment()
    except RuntimeError as error:
        return str(error)
    return ''


def build(output: Path, go: str, env: dict):
    env = java_environment(env)
    key = ROOT / '.tools/android-debug.keystore'
    if not key.exists():
        raise RuntimeError('STMPCK signing key missing: restore .tools/android-debug.keystore from the original build host')
    expected_signer = (ROOT / 'android/update-signer.sha256').read_text().strip()
    # Match desktop/WASM builds: temporary gomobile packages do not need Git stamping.
    env['GOFLAGS'] = (env.get('GOFLAGS', '') + ' -buildvcs=false').strip()
    binary = ROOT / '.tools/bin'
    binary.mkdir(exist_ok=True)
    env.update(ANDROID_HOME=str(SDK), ANDROID_NDK_HOME=str(NDK),
               GOPATH=str(ROOT / '.tools/go-work'), GOBIN=str(binary))
    env['PATH'] = str(Path(env['JAVA_HOME']) / 'bin') + os.pathsep + str(Path(go).parent) + os.pathsep + str(binary) + os.pathsep + env['PATH']
    work = ROOT / '.tools/android-build'
    work.mkdir(exist_ok=True)

    def run(args):
        print('  Android:', Path(str(args[0])).name, ' '.join(str(a) for a in args[1:3]), flush=True)
        subprocess.run([str(a) for a in args], cwd=ROOT, env=env, check=True, timeout=900)

    tool = binary / 'ebitenmobile'
    if not tool.exists():
        run([go, 'install', 'github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@v2.9.9'])
    aar = work / 'stmpck.aar'
    run([tool, 'bind', '-target', 'android/arm64', '-androidapi', '23', '-javapkg',
         'dev.cjtrowbridge.stmpck', '-o', aar, './mobile'])
    extracted = work / 'aar'
    shutil.rmtree(extracted, ignore_errors=True)
    with zipfile.ZipFile(aar) as bundle:
        bundle.extractall(extracted)
    classes = work / 'classes'
    dex = work / 'dex'
    for directory in (classes, dex):
        shutil.rmtree(directory, ignore_errors=True)
        directory.mkdir()
    run(['javac', '-source', '8', '-target', '8', '-bootclasspath', PLATFORM,
         '-classpath', extracted / 'classes.jar', '-d', classes, *sorted((ROOT / 'android').glob('*.java'))])
    activity_jar = work / 'activity.jar'
    run(['jar', 'cf', activity_jar, '-C', classes, '.'])
    run([BUILD_TOOLS / 'd8', '--lib', PLATFORM, '--min-api', '23', '--output', dex,
         extracted / 'classes.jar', activity_jar])
    resources = work / 'resources.zip'
    run([BUILD_TOOLS / 'aapt2', 'compile', '--dir', ROOT / 'android/res', '-o', resources])
    unsigned = work / 'unsigned.apk'
    run([BUILD_TOOLS / 'aapt2', 'link', '-I', PLATFORM, '--manifest',
         ROOT / 'android/AndroidManifest.xml', '-R', resources, '-o', unsigned])
    with zipfile.ZipFile(unsigned, 'a', compression=zipfile.ZIP_DEFLATED) as bundle:
        for path in dex.glob('*.dex'):
            bundle.write(path, path.name)
        for path in (extracted / 'jni').rglob('*.so'):
            bundle.write(path, 'lib/' + path.relative_to(extracted / 'jni').as_posix())
        assets = extracted / 'assets'
        if assets.exists():
            for path in assets.rglob('*'):
                if path.is_file():
                    bundle.write(path, 'assets/' + path.relative_to(assets).as_posix())
    aligned = work / 'aligned.apk'
    run([BUILD_TOOLS / 'zipalign', '-f', '-P', '16', '4', unsigned, aligned])
    key = ROOT / '.tools/android-debug.keystore'
    if not key.exists():
        raise RuntimeError('STMPCK update signing key missing: restore .tools/android-debug.keystore; never generate a replacement for existing installs')
    output.parent.mkdir(parents=True, exist_ok=True)
    run([BUILD_TOOLS / 'apksigner', 'sign', '--ks', key, '--ks-pass', 'pass:android',
         '--key-pass', 'pass:android', '--out', output, aligned])
    verification = subprocess.run([str(BUILD_TOOLS / 'apksigner'), 'verify', '--verbose',
                                   '--print-certs', str(output)], env=env, check=True,
                                  capture_output=True, text=True, timeout=30)
    signers = re.findall(r'Signer #\d+ certificate SHA-256 digest: ([0-9a-fA-F]+)', verification.stdout)
    if signers != [expected_signer]:
        output.unlink(missing_ok=True)
        raise RuntimeError('APK signer differs from the original STMPCK update identity; refusing artifact')
    print('  Android: APK signature matches the original update identity', flush=True)
