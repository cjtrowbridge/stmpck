# Android debug build

For the [release-target foundation deliverable](../../ROADMAP.md#2-choose-the-first-release-target), address **Build a locally signed Android debug APK of the opening menu for the first phone test.** Produce an installable native APK sharing the existing menu implementation, so the first phone test exercises the actual app.

## Implementation

- Install pinned SDK platform, build tools, and NDK locally under ignored `.tools/`.
- Move reusable menu rendering/input into an importable package and retain a small desktop entrypoint.
- Bind that package with the pinned Ebitengine mobile tool; add a minimal Android activity with lifecycle handling and private preference storage.
- Build the AAR, compile the activity, package, align, and sign a development APK through `scripts/build.py` and its direct helper. Use development-only signing credentials under `.tools/` and package ID `dev.cjtrowbridge.stmpck.debug`.
- Add Android arm64 to the no-flag target matrix with honest prerequisite checks. Do not request internet permission or enable store publishing.

## Verification

Verify APK signing, manifest package/permissions, and build output. Run existing navigation/preference tests after sharing the menu code. Device installation, launch, interaction, screenshot, lifecycle, and crash-log checks belong to the linked [ADB plan](adb-setup.md).

## Result — 2026-11-07

Complete for Android arm64 development builds. The canonical no-flag build produces `releases/android/arm64/app/latest.apk` (4,583,898 bytes for this revision). Signature verification succeeds; manifest inspection confirms the development package ID and no requested permissions. The APK was installed and launched successfully on the Pixel 10 Pro XL.

The shared menu now lives in `internal/app/`; desktop/web and Android entrypoints reuse it. Existing navigation/preference tests and the WebAssembly build pass. Native Linux compilation still needs X11 development headers and therefore the full matrix verification reports failure even when Android and WebAssembly succeed.

Local tools: Go 1.26.4, Ebitengine/ebitenmobile 2.9.9, SDK API 35, Build-Tools 35.0.0, NDK 27.2.12479018, and OpenJDK 21 (originally Ubuntu's package extracted under `.tools/`; since 2026-11-12 canonically Debian's `openjdk-21-jdk-headless` package — see [open-source-build-toolchain](../current/open-source-build-toolchain.md)). `scripts/setup_android.py` installs the pinned SDK/NDK archives with checksum verification and preserved symlinks and verifies the OpenJDK 21 and adb packages on `PATH`; it never runs `sudo` itself. APK packaging is handled by `scripts/build_android.py`, called by `scripts/build.py`. `scripts/deploy_android.py` installs an existing APK on one explicitly selected or uniquely authorized device.

The debug key is local and ignored. This is a device-testing build, not a store release. Reference: [Ebitengine mobile integration](https://ebitengine.org/en/documents/mobile.html).
