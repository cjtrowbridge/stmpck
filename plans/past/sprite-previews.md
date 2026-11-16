# Sprite gallery previews

This plan owns **Sprite gallery previews** under [the shared experience](../../ROADMAP.md#7-define-the-smallest-shared-experience). Live demonstrations currently simulate every game and search possible moves while browsing, causing reported stalls. Replace gallery execution with bundled captures made from real game rules and rendering, and document a repeatable process for future games.

Use deterministic seeds and scripted inputs, capture 96-pixel frames at 12 fps for eight seconds (twelve for Bloom's six patterns), and export default/vivid PNG sheets plus wireframe/default/vivid posters. Preserve board proportions, intensity autoplay, shared breathing/sparkles, sorting, navigation and player saves. Decode sheets away from the update loop, upload at most one per update, and bound decoded GPU sheet pixels to 64 MiB. Posters remain available during loading. Only visible cards advance playback; settings, gameplay and backgrounding suspend it.

The generator is development-only, uses the existing renderer through a tagged WebAssembly capture program, and runs through scripts/build.py. Check in images and a source fingerprint manifest; ordinary builds verify freshness without needing a browser. Adding or changing a game requires regeneration, visual review and asset checks. Document costs, build/test evidence, and remaining device verification before completion.

## Implemented and verified — 2026-11-08

The [required game-update and deployment gate](../../docs/previews.md#required-game-update-and-deployment-gate) is now canonical in AGENTS, README and the roadmap's shared completion criteria. It applies to every added/updated game and every deployment, including already-implemented games. Regeneration is required for changed capture inputs; accurate unchanged assets may be reused after validation. Review evidence, asset tests and intended-target build validation must accompany the revision being completed/deployed. Automated freshness checks and manual visual review are explicitly distinguished.

The workflow is documented in [docs/previews.md](../../docs/previews.md), required by AGENTS and linked in the shared game-completion criteria. Run `python scripts/build.py --generate-previews`. A two-second simulated lead-in produces useful posters. Ninety bundled PNGs cover all eighteen games, with a source/checksum manifest. Generation stages all captures before publishing them; successful repeated runs produced byte-identical PNGs on this host. Capture failure is reported without silently accepting incomplete assets. Rule/render changes invalidate the ordinary build until regeneration.

`internal/app/preview.go` now owns runtime playback. Only the tagged development capture program calls demonstration code from the app package. The menu loads posters, decodes one sheet at a time off the update loop and uploads at most one per update. Sheet pixels are capped at 64 MiB; visible sheets are protected, and excess cards use posters instead of thrashing. Poster pixels add about 1.9 MiB and one temporary decoded sheet up to about 5.1 MiB, excluding driver/allocator overhead. Still/default switching preserves the cached frame. Breathing/rainbow borders and sparkles remain shared effects.

Evidence:

- All rule/navigation/storage/presentation/preview tests pass with the race detector. Asset tests verify catalog/variant completeness, PNG decoding/dimensions, frame bounds, animation changes, checksums and per-sheet budget. Manual checks verify stale-source and corrupt-PNG rejection.
- Browser interaction checks pass for still and wireframe freezing, same-frame return, animation resume, maximum intensity, scrolling, settings, game entry/return and absence of preview play counts/saves. Screenshots and results: `reports/sprite-previews/after/`.
- Same-host headless Firefox, 480×914, ten-second warm default samples: old live gallery median 83.50 ms, p95 99.76 ms, p99 149.88 ms, 124 frames over 50 ms; sprites median 17.04 ms, p95 17.22 ms, p99 17.26 ms, zero over 50 ms. Maximum-intensity sprites median 17.06 ms, p95 17.16 ms, maximum 34.14 ms. These are browser animation-frame observations, not GPU timings or phone/battery measurements. Baseline: `reports/sprite-previews/before-isolated/checks.json`.
- Images total 2,261,134 bytes. Build report `reports/build/build-report-20261108T223321Z.json`: Android APK 6,779,354 bytes and WASM 14,413,370 bytes. Native Linux still fails for missing X11 headers. Android's temporary gomobile build now disables VCS stamping consistently with desktop/WASM; no deployment was performed.

Phone performance, cold-loading behavior on constrained devices, GPU/driver memory, energy use, wider resolutions and native accessibility remain release checks. The 96-pixel captures deliberately favor small assets; review larger presentation sizes before implementing the future 1–8-column setting.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
