# Native Windows build artifact

This requirement plan belongs to [Choose the first release target](../../ROADMAP.md#2-choose-the-first-release-target), coordinated by [release-target.md](../current/release-target.md). It maps exactly to “Produce a native Windows executable automatically when the canonical build runs on Windows.” The existing matrix lists Windows but preview source hashing depends on host path separators, and desktop builds enable CGO indiscriminately. The outcome is a verified native executable at `releases/windows/amd64/app/latest.exe` on this Windows/amd64 host, through the ordinary build entrypoint without target flags or a C compiler.

## Work and verification

- Preserve the existing native-host target policy and explicit unsupported-target reporting.
- Use the Windows backend's pure-Go build and retain other platforms' settings.
- Make preview fingerprints portable across path separators and checkout line endings; regenerate through the canonical driver when capture sources change and retain asset validation.
- Add focused regressions for Windows build environment, artifact paths and portable freshness checks.
- Run canonical build verification, preview asset tests and browser gallery checks; record actual results here and in the journal. No phone deployment or store packaging is part of this requirement.

## Status

Implemented and verified on Windows/amd64. Windows interactive runtime, accessibility and store release verification remain pending. The existing native-host policy is unchanged; Windows/arm64 is not in the matrix.

Windows builds now use `CGO_ENABLED=0`. Ordinary builds return nonzero if any attempted target fails, including without `--verify`; stale artifacts cannot satisfy verification. Source fingerprints use POSIX paths and normalize CRLF, preserving detection of actual source edits. Regeneration initially exhausted WASM memory on this host, so development capture now flushes each frame's draw commands with pixel readback instead of queuing a whole animation. Gameplay rules and the player's rendering code are unchanged.

## Evidence — 2026-11-08

- `python -m unittest discover -s scripts -p test_build.py`: four regressions pass (Windows environment/path/selection, stale executable rejection, ordinary failure exit status, checkout line endings versus real source changes).
- `python scripts/build.py --generate-previews`: all eighteen games and ninety images regenerated successfully; manifest checksums and freshness pass. Images total 2,239,187 bytes. Local Firefox was added to PATH for capture and browser checks only. Initial pinned toolchain/module downloads required network access; app operation remains offline.
- Visually reviewed all wireframe/default/vivid posters and first/middle/last animation samples. Boards, palette distinctions, representative actions and replay endpoints remain recognizable without clipping or blank frames. Contact sheets: `reports/windows-build/posters.png` and `animation-samples.png`. Generated PNGs and manifest accompany this change; capture timing and seeds are unchanged.
- `python scripts/build.py --verify`: both feasible targets pass in `reports/build/build-report-20261109T030240Z.json` (UTC filename; November 8 local). Windows executable: 13,435,392 bytes; WASM: 14,424,098 bytes. Independently checked MZ/PE signatures and amd64 machine type. Other targets are explicitly reported unsupported on this host. No C compiler was detected or required.
- `go test ./internal/preview ./internal/play ./internal/menu ./internal/storage ./internal/presentation`, using the driver's isolated Go environment: all pass.
- `python scripts/check_gallery.py --label windows --checks`: passes still freeze, monochrome, held-frame return, animation resume, maximum intensity, scrolling, settings, entry/return and no preview save/count writes. Evidence: `reports/sprite-previews/windows/`; headless Firefox viewport 500×915. Median interval about 33.34 ms at default and maximum, with no sampled intervals over 50 ms; host/browser diagnostic evidence only.
- `git diff --check`: passes. Roadmap citation and parent-plan links resolve. No native interactive play, phone deployment or store publishing performed.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
