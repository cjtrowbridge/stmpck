"""Resolve and validate the same JDK for Android setup and packaging.

The PATH compiler selects one JDK home; the existing local JDK is a fallback
only when PATH has no compiler. All Java tools come from that home, and both
runtime and compiler must match the project's pinned major. No installation
or package-provenance claim is made here.
"""
from pathlib import Path
import os
import re
import shutil
import subprocess

ROOT = Path(__file__).resolve().parent.parent
REQUIRED_JDK_MAJOR = 21
REMEDY = 'sudo apt-get install -y openjdk-21-jdk-headless adb'


def java_environment(environment=None):
    """Return the validated build environment, or a useful prerequisite error."""
    env = dict(os.environ if environment is None else environment)
    compiler = shutil.which('javac', path=env.get('PATH', ''))
    if compiler:
        home = Path(compiler).resolve().parent.parent
    else:
        home = ROOT / '.tools/jdk'
    tools = {name: home / 'bin' / name for name in ('java', 'javac', 'jar', 'keytool')}
    for name, path in tools.items():
        if not path.is_file() or not os.access(path, os.X_OK):
            raise RuntimeError(f'Selected JDK is missing executable {name}: {path}; remedy: {REMEDY}')
    env['JAVA_HOME'] = str(home)
    env['PATH'] = str(home / 'bin') + os.pathsep + env.get('PATH', '')
    for name in ('java', 'javac'):
        try:
            result = subprocess.run([str(tools[name]), '-version'], env=env,
                                    capture_output=True, text=True, timeout=10)
        except (OSError, subprocess.TimeoutExpired) as error:
            raise RuntimeError(f'Cannot check selected {name}: {error}; remedy: {REMEDY}') from error
        output = result.stdout + result.stderr
        pattern = r'(?:openjdk|java) version "(\d+)' if name == 'java' else r'javac\s+(\d+)'
        version = re.search(pattern, output)
        if result.returncode != 0 or not version or int(version.group(1)) != REQUIRED_JDK_MAJOR:
            raise RuntimeError(f'Selected {name} must report major {REQUIRED_JDK_MAJOR}: '
                               f'{output.strip() or "no version output"}; remedy: {REMEDY}')
    return env
