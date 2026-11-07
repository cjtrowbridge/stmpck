"""Android debug packaging, called only by the canonical build driver."""
from pathlib import Path
import os
import shutil
import subprocess
import zipfile

ROOT = Path(__file__).resolve().parent.parent
SDK = ROOT / '.tools/android-sdk'
NDK = SDK / 'ndk/27.2.12479018'
BUILD_TOOLS = SDK / 'build-tools/35.0.0'
PLATFORM = SDK / 'platforms/android-35/android.jar'
JDK = ROOT / '.tools/jdk'


def unavailable_reason():
    if not PLATFORM.exists() or not (NDK / 'source.properties').exists() or not (BUILD_TOOLS / 'aapt2').exists():
        return 'Android SDK/NDK missing; run python scripts/setup_android.py on Linux'
    for command in ('java', 'javac', 'jar', 'keytool'):
        if not (JDK / 'bin' / command).exists() and not shutil.which(command):
            return 'Android debug packaging requires a JDK: missing ' + command
    return ''


def build(output: Path, go: str, env: dict):
    env = dict(env)
    binary = ROOT / '.tools/bin'
    binary.mkdir(exist_ok=True)
    env.update(ANDROID_HOME=str(SDK), ANDROID_NDK_HOME=str(NDK),
               GOPATH=str(ROOT / '.tools/go-work'), GOBIN=str(binary))
    if (JDK / 'bin/java').exists():
        env['JAVA_HOME'] = str(JDK)
        env['PATH'] = str(JDK / 'bin') + os.pathsep + env['PATH']
    env['PATH'] = str(Path(go).parent) + os.pathsep + str(binary) + os.pathsep + env['PATH']
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
         '-classpath', extracted / 'classes.jar', '-d', classes, ROOT / 'android/MainActivity.java'])
    activity_jar = work / 'activity.jar'
    run(['jar', 'cf', activity_jar, '-C', classes, '.'])
    run([BUILD_TOOLS / 'd8', '--lib', PLATFORM, '--min-api', '23', '--output', dex,
         extracted / 'classes.jar', activity_jar])
    unsigned = work / 'unsigned.apk'
    run([BUILD_TOOLS / 'aapt2', 'link', '-I', PLATFORM, '--manifest',
         ROOT / 'android/AndroidManifest.xml', '-o', unsigned])
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
        run(['keytool', '-genkeypair', '-keystore', key, '-storepass', 'android',
             '-keypass', 'android', '-alias', 'androiddebugkey', '-keyalg', 'RSA',
             '-keysize', '2048', '-validity', '3650', '-dname', 'CN=STMPCK Debug,O=Development,C=US'])
        key.chmod(0o600)
    output.parent.mkdir(parents=True, exist_ok=True)
    run([BUILD_TOOLS / 'apksigner', 'sign', '--ks', key, '--ks-pass', 'pass:android',
         '--key-pass', 'pass:android', '--out', output, aligned])
    run([BUILD_TOOLS / 'apksigner', 'verify', '--verbose', output])
