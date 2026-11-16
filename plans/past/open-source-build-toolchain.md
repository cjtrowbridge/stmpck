# Canonical open-source build toolchain

For [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target), address **Require and use an open-source alternative to the JDK (Debian's OpenJDK 21 package) and Debian's adb package as the canonical local build toolchain, verified by the setup helper.** The earlier setup path left the Java toolchain and Platform-Tools as unowned manual steps, which made fresh-host setup unreproducible and silently dependent on a proprietary-default JDK.

## Implemented outcome

`scripts/android_toolchain.py` resolves the JDK containing the PATH compiler, or the existing local `.tools/jdk` fallback only when no PATH compiler exists. It requires executable java/javac/jar/keytool from that home and successful, parseable major-21 runtime/compiler version commands. Both setup and packaging call it; packaging sets JAVA_HOME and prioritizes that JDK on PATH. Setup checks ADB and Java before SDK downloads and never runs sudo. Executable validation does not prove Debian package provenance.

## Verification and remaining scope

Ten Python regressions pass, including runtime/compiler mismatch, failed/unparseable commands, missing sibling tools, PATH versus local fallback, and shared setup/build rejection. The canonical build verifies Linux/amd64 and WebAssembly with local test prerequisites. Android SDK/NDK, Java and ADB are unavailable in this environment; actual Android packaging/phone launch remains pending. This requirement remains unchecked in the roadmap until that evidence exists.

## Authorized alignment — 2026-11-15

The user approved aligning setup and packaging. Resolve one JDK directory (canonical PATH compiler, with the documented local fallback only if no PATH compiler exists); verify java and javac major 21 plus jar/keytool availability from that directory. Use the same resolved directory as JAVA_HOME and first on PATH during packaging. Reject malformed/nonzero version results and mixed Java/compiler majors. Setup checks this before downloading SDK archives. Add focused mocked regressions without downloading tools in tests. This validates the selected executable toolchain; it does not prove Debian package provenance.

## Remaining work and closure — 2026-11-16

Build and launch Android with the selected canonical toolchain and record actual packaging/device evidence. Archive when the unchecked roadmap requirement has that evidence; mocked toolchain tests alone do not close it.

## Closure — 2026-11-16

The remaining Android packaging/launch evidence now exists. Debian OpenJDK 21.0.13+9 packages were unpacked under ignored `.tools/jdk-unpack`, with the documented `.tools/jdk` fallback and repo-local resolution of packaged configuration links. Setup/build selection validates that runtime/compiler and sibling tools; the canonical Android helper uses the same JDK. Existing SDK/NDK tooling is reused read-only. Debian `/usr/bin/adb` provides the device workflow. Ten build/toolchain Python checks pass, including mismatch/rejection regressions.

Canonical Android/Linux/WebAssembly build and fresh install/launch pass in `reports/build/build-report-20261116T222812Z.json`; final version-3 packaging passes in `reports/build/build-report-20261116T223000Z.json`. Pixel 4a subsequently updated from version 2 through the app's public feed, with saved game bytes preserved. See [updater evidence](android-updates.md). No system Java installation or sudo operation was performed. The requirement's executable-selection and Android evidence scope is complete; archive it without claiming all target-platform verification.
