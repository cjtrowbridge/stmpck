#!/usr/bin/env python3
"""Install the pinned Linux Android build prerequisites locally; does not build.

Canonical open-source toolchain policy: the Java tools (java/javac/jar/keytool)
come from Debian's open-source OpenJDK 21 package, and ADB comes from Debian's
`adb` package. The setup helper verifies both, prints the exact apt remedy when
either is missing, and exits non-zero until they are installed. OpenJDK 21 is
the last LTS whose javac still accepts `-source 8 -target 8`, which the APK
packaging step relies on; a newer major is reported as fatal rather than used.
"""
import concurrent.futures
import hashlib
from pathlib import Path
import re
import shutil
import stat
import subprocess
import urllib.request
import zipfile

ROOT = Path(__file__).resolve().parent.parent / '.tools' / 'android-sdk'
# Official repository2-1.xml checksums; preserve Google's included license files.
PACKAGES = [
    ('platform-35_r02.zip', '0bb560a90a7a2cbd0dd8348224d518b638fe7949', 'platforms/android-35'),
    ('build-tools_r35_linux.zip', '2cfaa0bbb2336e9ec18ed3ecea84fa2e2af607bc', 'build-tools/35.0.0'),
    ('android-ndk-r27c-linux.zip', '090e8083a715fdb1a3e402d0763c388abb03fb4e', 'ndk/27.2.12479018'),
]
# Canonically required Debian packages: open-source OpenJDK 21 and ADB.
APT_PACKAGES = ('openjdk-21-jdk-headless', 'adb')
JDK_TOOLS = ('java', 'javac', 'jar', 'keytool')
# OpenJDK 21 is the pinned major: last release supporting javac -source/-target 8.
REQUIRED_JDK_MAJOR = 21

def install(package):
    name, checksum, relative = package
    destination = ROOT / relative
    if (destination / '.stmpck-installed').exists():
        print('Already installed:', relative, flush=True)
        return
    archive = ROOT / name
    print('Downloading', name, flush=True)
    with urllib.request.urlopen('https://dl.google.com/android/repository/' + name, timeout=120) as response, archive.open('wb') as output:
        digest = hashlib.sha1()
        while chunk := response.read(1024 * 1024):
            digest.update(chunk)
            output.write(chunk)
    if digest.hexdigest() != checksum:
        raise RuntimeError('Checksum mismatch: ' + name)
    stage = ROOT / (name + '.extract')
    shutil.rmtree(stage, ignore_errors=True)
    stage.mkdir()
    with zipfile.ZipFile(archive) as bundle:
        for entry in bundle.infolist():
            target = (stage / entry.filename).resolve()
            if not target.is_relative_to(stage.resolve()):
                raise RuntimeError('Invalid archive path')
            mode = entry.external_attr >> 16
            if stat.S_ISLNK(mode):
                link = bundle.read(entry).decode()
                if not (target.parent / link).resolve().is_relative_to(stage.resolve()):
                    raise RuntimeError('Invalid archive symlink')
                target.parent.mkdir(parents=True, exist_ok=True)
                target.symlink_to(link)
            else:
                bundle.extract(entry, stage)
                if mode and not entry.is_dir():
                    target.chmod(mode & 0o777)
    children = list(stage.iterdir())
    if len(children) != 1 or not children[0].is_dir():
        raise RuntimeError('Unexpected archive structure: ' + name)
    destination.parent.mkdir(parents=True, exist_ok=True)
    if destination.exists():
        raise RuntimeError('Partial installation exists: ' + str(destination))
    children[0].rename(destination)
    stage.rmdir()
    (destination / '.stmpck-installed').write_text(checksum + '\n')
    archive.unlink()
    print('Installed', relative, flush=True)

def verify_open_source_toolchain():
    """Verify the canonical open-source toolchain (OpenJDK 21, adb) is present.

    Returns a list of human-readable problems; empty means the toolchain is
    ready. Prints the exact apt remedy for each problem.
    """
    problems = []
    for tool in JDK_TOOLS:
        if not shutil.which(tool):
            problems.append('Java tool not found on PATH: ' + tool)
    if not problems:
        try:
            version = subprocess.run(['java', '-version'], capture_output=True, text=True, timeout=10)
        except (FileNotFoundError, subprocess.TimeoutExpired) as exc:
            problems.append('Could not run java -version: ' + str(exc))
        else:
            match = re.search(r'version "(\d+)', (version.stderr or '') + (version.stdout or ''))
            if not match:
                problems.append('Could not determine the Java major version from java -version')
            elif int(match.group(1)) != REQUIRED_JDK_MAJOR:
                problems.append('Java major version must be ' + str(REQUIRED_JDK_MAJOR)
                                + ' (found ' + match.group(1) + '); OpenJDK 21 is the pinned open-source toolchain')
    if not shutil.which('adb'):
        problems.append('adb not found on PATH')
    if problems:
        print('Canonical open-source toolchain is incomplete:', flush=True)
        for problem in problems:
            print('  -', problem, flush=True)
        print('Remedy: sudo apt-get install -y ' + ' '.join(APT_PACKAGES)
              + ' (install packages yourself; this helper never runs sudo)', flush=True)
    else:
        print('Canonical open-source toolchain present: OpenJDK 21 on PATH, adb on PATH', flush=True)
    return problems

if __name__ == '__main__':
    ROOT.mkdir(parents=True, exist_ok=True)
    with concurrent.futures.ThreadPoolExecutor(max_workers=3) as pool:
        list(pool.map(install, PACKAGES))
    raise SystemExit(1 if verify_open_source_toolchain() else 0)
