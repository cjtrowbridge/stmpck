# Android development updates through GitHub

This plan owns **Android development updates** under [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target). Routine test-build delivery currently requires ADB. The authorized outcome is Apparat-style fixed-URL APK downloads, an actual Settings update button and one silent check per app process start, with Android installation confirmation and ordinary play remaining offline.

## Authorized outcome and acceptance

The user approved implementation and publishing the build to an artifacts-only public GitHub repository. Use `cjtrowbridge/stmpck-builds`, with a fixed raw HTTPS `latest.apk` URL, while keeping the source repository private. Preserve package identity and local saves. Compare downloaded bytes against the installed APK SHA-256, then require the same signing identity and a higher version code before offering installation. No background polling, gameplay telemetry or embedded GitHub credentials.

Use the existing Java wrapper, a small Gomobile callback and Settings control. Bound downloads to one worker, 128 MiB and two minutes with connection/read timeouts; cancel on Activity destruction and remove partial files. Silent startup checks suppress equal-version/network-error UI; only a valid newer update is offered. Manual checks expose useful status. Never report installer launch as successful installation.

Build only through `scripts/build.py`, regenerate/review affected captures, pass preview/rule/save checks and intended-target canonical validation. Publish only that APK, plus checksum/build metadata; verify the public bytes after publication. Phone in-place upgrade and offline persistence evidence remain separate if the device is unavailable.

## Implementation checkpoint and constraints

Implement, verify, build and publish. The original signing key has been recovered from the user-authorized Razer host, and local Android tooling is ready. Do not replace the installed app's signer or uninstall it. No source publication, store release, app-data reset or inference operation is authorized by this task.

## Pre-publication checks — 2026-11-16

Recovered the original STMPCK key from the user-authorized Razer host into ignored local storage; its public SHA-256 certificate fingerprint is pinned in `android/update-signer.sha256`. No key material is published. Debian OpenJDK 21 packages were unpacked locally; SDK/NDK tools are reused read-only from the neighboring Apparat tooling directory, without executing Apparat source or modifying its checkout.

Focused game/save/presentation/preview/update-bridge tests pass with the race detector. All ten Python build/toolchain regressions pass. Required preview generation produced byte-identical PNG assets for all eighteen games; the source manifest was refreshed. Visually reviewed all still styles and first/middle/last animation samples (`reports/android-updates/preview-review.png`). Browser checks pass still/wireframe freezing, held-frame return, autoplay, maximum intensity, scrolling, Settings, game entry/return and no preview save/count writes (`reports/sprite-previews/android-updates/checks.json`).

Canonical Android/Linux/WebAssembly verification passes. Version 2 was built and installed on the connected clean Pixel 4a through the canonical build/launch action, with the original signer; report `reports/build/build-report-20261116T222812Z.json`. Its startup check failed silently while the feed was absent. The manual Settings action reported failure without blocking play. Short held injected touches work; zero-duration injected taps were too short to reliably sample, so they are not evidence of physical-touch failure. Version 3 adds explicit native view keyboard focus and is the publication candidate for the in-app upgrade test.

Final publication candidate: canonical verification passes Android/Linux/WebAssembly in `reports/build/build-report-20261116T223000Z.json`, Android version code 3 / name 0.3-dev, 6,730,608 bytes. The Pixel 4a displays the root-logo icon correctly in Android App info. Publication and the real in-app update/save comparison follow this checkpoint.

## Publication and closure — 2026-11-16

Complete. Published version **0.3-dev (3)** at [the fixed APK URL](https://raw.githubusercontent.com/cjtrowbridge/stmpck-builds/main/latest.apk) in the public artifacts-only [stmpck-builds repository](https://github.com/cjtrowbridge/stmpck-builds). Artifact commit `5d31f3e`, source commit `a2a43a5`, APK 6,730,608 bytes, SHA-256 `b9c794251f6d8f6e0d8e093bc9a1affbac0454de35d1b0f2fa6e66ef9bef19e3`. Public downloaded bytes match the canonical artifact. Source remains private; no credentials/signing key were published.

On the Pixel 4a, version 2 detected the published update on process startup, offered Later/Update, and allowed deferral followed by a manual Settings retry. Android’s per-app install permission and system installer were exercised without programmatic permission bypass. One attempt returned `INSTALL_FAILED_VERIFICATION_FAILURE` after the test foregrounded STMPCK while Google verification was pending; retrying with the system UI left in front completed normally, with no code, key or security-setting change. Do not infer success from handoff alone.

Android confirms installed version 3 / 0.3-dev. The test game save remained byte-identical across the in-app upgrade (SHA-256 `7154a3a95487edfabec4db9d40d2bf0e4b51270ccf1e5161ae715346b97ff80f`). A subsequent startup produced no update prompt and manual Settings check reported **Up to date.** Screenshots and before/after saves are in `reports/android-updates/`; launcher and App info display the root-logo icon correctly. Browser keyboard mappings and touch on Android are checked; injected phone keyboard events did not establish native physical-keyboard/controller support, which remains shared release verification. No new claim is made for wider devices, screen readers or exhaustive interrupted-download/network fault coverage.

The scoped implementation, validation and publication are complete; archive this plan. Future development publishes must increment the version code and repeat the existing preview/build gate. Store distribution remains unconfigured.
