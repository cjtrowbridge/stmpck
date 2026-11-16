# Choose the first release target

This plan addresses the [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target) as a whole. The intended outcome and work are described below.

## Work and validation

Establish the target-device development workflow. ADB is now a prerequisite for Android testing; the store release platform remains undecided. Coordinate device setup and later packaging verification.

## Status

In progress. Android is now the first device-testing target: a signed arm64 debug APK runs on the authorized Pixel 10 Pro XL (Android 17). The ADB and debug-build requirement plans are completed under `plans/past/` and linked from the roadmap. Store release target selection, broader platform validation, and release signing remain open.

The [canonical VS Code launch](test-phone-launch.md) now builds and deploys the debug app to the test phone; its underlying command has been verified.

The [Windows build requirement](../past/windows-build.md) adds a native executable to ordinary Windows-host builds, without selecting a store release target.

## Remaining work and closure — 2026-11-16

Coordinate the outstanding toolchain and launch checks. Full deliverable closure requires platform selection, builds/runs and broader native/device verification plus the mapped packaging decisions. Unstarted requirements stay in the roadmap without placeholder plans.

The [Android app icon requirement](android-app-icon.md) packages the existing root logo; Android resource/launcher verification remains part of its own acceptance.

The [Android development updater](android-updates.md) is authorized for an artifacts-only public GitHub feed. It does not select a store release target or replace ADB diagnostics.
