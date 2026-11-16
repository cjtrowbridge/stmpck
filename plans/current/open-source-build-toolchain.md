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
