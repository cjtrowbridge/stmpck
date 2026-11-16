# Android wrapper and launcher icon

`MainActivity.java` hosts the game and forwards lifecycle/Back actions. `AndroidManifest.xml` declares the development package and its launcher icon. `scripts/build_android.py`, called by the canonical build driver, compiles `res/` with AAPT2 and links those resources into the APK.

The five `res/mipmap-*/ic_launcher.png` files are renders of the root [logo.svg](../logo.svg), without redesign or cropping: mdpi 48px, hdpi 72px, xhdpi 96px, xxhdpi 144px and xxxhdpi 192px. The checked-in PNGs avoid adding an SVG renderer to ordinary builds.

When changing the logo, render the SVG at its 1024px intrinsic size with an SVG renderer supporting its transforms, text and gradients (these assets use librsvg with DejaVu Sans), then downsample with Lanczos to each size. Visually inspect the stimpack, rainbow and 128 text before replacing the PNGs. ImageMagick’s internal SVG renderer omitted transformed content during the original conversion and was not used for the final assets.

Run `python3 scripts/build.py --verify` for locally feasible targets. Android packaging and launcher checks require the Android toolchain and a device; successful desktop builds do not verify Android resources.

## Development updates

`AppUpdater.java` checks the fixed HTTPS artifact at `https://raw.githubusercontent.com/cjtrowbridge/stmpck-builds/main/latest.apk`. A Settings button (tap, Enter, U or controller A) checks explicitly; the wrapper also checks once per process startup. Startup equality/network failures are silent. The exact installed APK SHA-256 identifies identical bytes; different bytes must have the same package/signing identity and a strictly higher version code before a native confirmation dialog offers installation. Android verifies the APK during installation. A public hash is not authentication.

One background worker downloads at most 128 MiB with 15-second connection/read timeouts and a two-minute transfer window. Activity destruction cancels it. Partial downloads are never exposed to the installer; each worker has a unique temporary file. `UpdateApkProvider.java` grants read access only to the completed cached APK. Permission and installer screens remain Android-owned. Installer launch is reported as a handoff, not success; cancellation leaves the installed version intact. Ordinary play never waits for the update service.

The update package remains `dev.cjtrowbridge.stmpck.debug`. Keep the original `.tools/android-debug.keystore` private and backed up; the build refuses a missing key and verifies the output against the public certificate fingerprint in `update-signer.sha256`. Increment the manifest version code for every published change. Do not generate a new key, change package identity, or uninstall an existing app to get around an update rejection.

The public `cjtrowbridge/stmpck-builds` repository contains only APK/checksum/build information. Source stays private. Publish only after the canonical build and [preview gate](../docs/previews.md) pass for the same source revision; publish the APK and checksum atomically in one artifact-repository commit and verify the downloaded bytes. Build/publishing credentials are never embedded in the app. The initial updater-enabled APK can be installed from the phone browser; later versions use Settings or the startup offer. ADB remains an optional test/diagnostic tool, not a delivery requirement. This is a development channel, not store distribution.
