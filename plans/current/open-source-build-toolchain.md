# Canonical open-source build toolchain

For [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target), address **Require and use an open-source alternative to the JDK (Debian's OpenJDK 21 package) and Debian's adb package as the canonical local build toolchain, verified by the setup helper.** The earlier setup path left the Java toolchain and Platform-Tools as unowned manual steps, which made fresh-host setup unreproducible and silently dependent on a proprietary-default JDK.

## Outcome

- `scripts/setup_android.py` installs the pinned SDK/NDK archives as before, then verifies the canonical open-source toolchain: `java`, `javac`, `jar`, `keytool` on `PATH` at OpenJDK major version 21, and `adb` on `PATH`. A mismatching major (e.g. 25) is fatal, not a warning, because `javac -source 8 -target 8` (used by APK packaging) is removed after 21. The helper prints the exact remedy (`sudo apt-get install -y openjdk-21-jdk-headless adb`) and never runs `sudo` itself.
- `scripts/build_android.py` reports a missing Java tool with the canonical package name in its prerequisite message.
- README and AGENTS record Debian `openjdk-21-jdk-headless` and `adb` as canonical; the `.tools/jdk` fallback remains accepted but is no longer the documented path.

## Verification

Reviewed source diff; deterministic local checks run against the changed scripts. The operator installs the packages; a fresh canonical `build.py --verify --launch-android` run on this machine with the packages present is required before this entry is considered fully verified on-device. No app behavior changes: menu, games, saves, and sensory controls are untouched by this toolchain-policy change.

## Status

Implemented locally on Debian 13 (trixie, OpenJDK 21.0.x in main). Awaiting operator package installation and a canonical build/launch run for on-device evidence.
