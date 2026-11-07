#!/usr/bin/env python3
"""Install and launch an existing STMPCK debug APK on one authorized phone."""
import argparse
from pathlib import Path
import shutil
import subprocess
import time

ROOT = Path(__file__).resolve().parent.parent
PACKAGE = 'dev.cjtrowbridge.stmpck.debug'
APK = ROOT / 'releases/android/arm64/app/latest.apk'


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--serial', help='Required when more than one device is authorized')
    args = parser.parse_args()
    local = ROOT / '.tools/platform-tools/adb'
    adb = str(local) if local.exists() else shutil.which('adb')
    if not adb:
        parser.error('ADB is missing; install Android SDK Platform-Tools')
    if not APK.exists():
        parser.error('APK is missing; run python scripts/build.py --verify first')
    output = subprocess.check_output([adb, 'devices'], text=True)
    devices = [line.split()[0] for line in output.splitlines()[1:]
               if len(line.split()) >= 2 and line.split()[1] == 'device']
    serial = args.serial
    if serial is None and len(devices) == 1:
        serial = devices[0]
    if serial not in devices:
        parser.error('Select one authorized device with --serial; check adb devices -l')
    command = [adb, '-s', serial]
    print('Installing STMPCK debug APK (existing app data is preserved).', flush=True)
    subprocess.run(command + ['install', '--no-incremental', '-r', str(APK)], check=True)
    print('Launching STMPCK.', flush=True)
    # Android can briefly foreground a system package-update activity after install.
    for attempt in range(12):
        result = subprocess.check_output(command + ['shell', 'am', 'start', '-W', '-n', PACKAGE + '/.MainActivity'], text=True)
        process = subprocess.run(command + ['shell', 'pidof', PACKAGE], capture_output=True, text=True)
        if 'Status: ok' in result and ('Activity: ' + PACKAGE + '/') in result and process.returncode == 0:
            print(result)
            break
        time.sleep(0.5)
    else:
        raise RuntimeError('Android did not confirm the app process/activity after installation: ' + result)
    print('Launched. Verify the menu on the phone; this does not publish a release.', flush=True)


if __name__ == '__main__':
    main()
