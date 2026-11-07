# Canonical test-phone launch

For [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target), address **Make VS Code launch.json build, install, and launch the current Android debug app on the connected test phone as the canonical development launch action.** Replace the inherited governance launch with the actual development feedback loop.

## Implementation

VS Code Run/Debug uses `scripts/build.py --verify --launch-android`. The ordinary no-flag build remains unchanged. The launch option builds through the canonical driver, requires this invocation to successfully build Android, then invokes `deploy_android.py` using the same Python interpreter. Never deploy a stale APK after an Android build failure. Other target failures remain visible but do not block a successful Android build from launching. Use the existing authorized-device selection and data-preserving debug install. No inference endpoint is involved.

## Validation

Verified: launch configuration parses; mocked failure-path checks prevent deploying after an unsuccessful Android build and allow a successful Android build despite other-target failures. The exact launch command built Android/WebAssembly, installed the fresh APK, and confirmed the test-phone activity/process launch. Linux remains blocked by missing X11 headers. VS Code’s button itself has not been exercised interactively. Evidence: `.tools/phone-launch.log`.
