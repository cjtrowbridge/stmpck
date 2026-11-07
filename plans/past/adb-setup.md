# ADB connection and first phone deployment

For the [release-target foundation deliverable](../../ROADMAP.md#2-choose-the-first-release-target), address **Set up ADB, connect and authorize the test phone, deploy the debug APK, and verify the first successful on-device launch.** The outcome is the opening menu running on the actual test phone, with installation, input, lifecycle, and crash checks recorded.

## Work and validation

- Use `.tools/platform-tools/adb`; install host USB rules if needed.
- Connect the unlocked test phone, enable USB debugging, accept authorization, and require `adb devices -l` to report `device`.
- Read the selected phone's Android version, ABI, and display size before deployment.
- Install the APK from the [Android debug build plan](android-debug-build.md), targeting that device explicitly and preserving existing app data on subsequent installs.
- Launch `dev.cjtrowbridge.stmpck.debug`, confirm its process and resumed activity, and inspect a screenshot.
- Exercise scrolling, preview selection/back, theme and motion controls, and suspend/resume. Check app crash logs and preference persistence.
- Record evidence and leave the roadmap item unchecked until the connected-device checks succeed.

## Result — 2026-11-07

Complete for the first phone deployment. The Pixel 10 Pro XL is authorized over USB and runs Android 17 / arm64. The development APK installed successfully and launched with `Status: ok` (584 ms reported cold activity launch). Screenshots confirm the opening gallery, Merge preview, return navigation, and scrolled gallery. Theme and motion preference changes, process-restart persistence, and suspend/resume were exercised. No crash entry was found for the checked app process. Starting preferences were restored.

Evidence is in the ignored `reports/android-test/` directory, including `checks.json` and screenshots. Some automated touches were missed around lock/startup transitions; broader input reliability remains follow-up work in the menu input plan. This milestone verifies first deployment, not store readiness or complete accessibility.

Repeat deployment with `python scripts/deploy_android.py`; specify `--serial` if multiple devices are connected. Keep the debug signing key in `.tools/android-debug.keystore` for subsequent updates.
