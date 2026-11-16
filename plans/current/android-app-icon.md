# Android app icon from the root logo

This plan owns **Android app icon** under [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target). The Android package has no app icon configured. Use the existing root `logo.svg` as the launcher icon without redesigning it, preserving package identity, saves and the existing build entrypoint.

## Outcome and checks

Render the logo into bundled launcher PNGs at Android density sizes, reference them in the manifest, and compile/link resources through the existing Android packaging helper. Inspect the rendered icon and verify dimensions, manifest/resource wiring and available build checks. Android packaging and launcher appearance require the Android toolchain/device; record unavailable checks explicitly. No game/capture source changes or deployment are in scope.

## Status

Implemented: five density-specific PNGs rendered from the unchanged root SVG, manifest icon reference and AAPT2 resource compile/link in the existing packaging helper. Visually reviewed the 192px render with the stimpack, rainbow and 128 text intact; the unsuitable ImageMagick internal SVG render was discarded. Conversion uses system librsvg and Pillow only during asset preparation, with no new build/runtime dependency.

Checks: all five PNGs decode as RGBA at 48/72/96/144/192px; manifest reference resolves. Four existing build-driver regressions pass. Canonical `python3 scripts/build.py --verify` passes preview freshness/integrity and WebAssembly (14,371,824 bytes), but exits nonzero because Linux lacks X11 headers. Android is unavailable because the SDK/NDK prerequisites are missing. Report: `reports/build/build-report-20261116T221504Z.json`. No APK or deployment is claimed.

Remaining work and closure: build Android with AAPT2 to verify packaged resources, then inspect the installed launcher icon on the phone. Archive after those checks pass. No game/capture sources changed, and no deployment was attempted.
