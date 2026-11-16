# Streamline the implementation

This plan owns **Streamline the implementation** under [Foundation 4](../../ROADMAP.md#4-simplification-and-demoscenization). The requirement covers simplifying shared/game code while preserving rules, saves, accessible controls and sensory choices. The first active scope is the save path: correct flush ordering, suppress unchanged Echo/Cluster persistence, and measure the resulting costs. Broader cleanup is unstarted.

## Review — 2026-11-15

`internal/app/play.go` publishes immutable encoded collection snapshots. `Flush` loads the snapshot before calling `Store.Save`; `internal/storage/storage.go` serializes writes with a mutex. This avoids simultaneous file writes but does not order snapshot selection with those writes.

A permitted interleaving is: caller A loads old snapshot S1 and is delayed; the game loop publishes S2; caller B loads S2, acquires the storage lock and saves; A resumes, acquires the same lock and saves S1 last. This is a source-level logical counterexample, not an observed device failure or a Go interleaving-test result. A clean race detector alone would not rule it out. The smallest candidate correction is serialization of Flush before snapshot selection, so persistence order follows selection order. Evaluate save-failure status ordering with the same ownership boundary. Do not add a worker queue or change the save format to address this.

Cost observations: changed gameplay updates encode the entire collection; several continuous games mark changed every active update. Flush validates new bytes, reads/validates the old save, writes the backup, then writes the primary with synchronous file syncs. These are actual operations visible in source, not proof of an unacceptable cost. Existing tests cover save validity/recovery but not the app-level concurrent selection interleaving.

## Dispositions and next evidence

- [x] Review ownership and identify the stale-write interleaving and candidate minimal correction.
- [x] Reproduce the serialization-boundary failure against the prior implementation; serialize selection/write/reporting and verify queued flushes select the latest published snapshot.
- [x] Measure host full-collection encoding, save cost and idle update overhead against the previous revision using populated saves.
- [ ] Correlate save cost with frame stalls and energy use on the test phone.
- [x] Complete focused race regressions, preview regeneration/review and canonical Linux/WebAssembly verification. Android packaging and phone checks remain pending prerequisites.

No Go executable is available in this session's environment (`command -v go` and local toolchain inspection). No benchmark, race-test pass, runtime correction or performance improvement is claimed. Defer runtime edits until they can be checked; source review and framework guidance can complete independently. Preserve existing snapshot freshness across suspension: reducing snapshot cadence is a behavioral decision, not a free optimization. Phone evidence remains distinct from host measurements.

## Authorized implementation — 2026-11-15

The user approved the review's corrections and measurement work. Serialize Game.Flush before snapshot selection through write/error reporting. Keep immutable snapshots and synchronous storage; no worker queue or format change. Make Echo and Cluster updates report whether state changed, so stationary updates do not snapshot or periodically write. Keep other game behavior unchanged and verify terminal transitions still save. Add focused ordering and update checks plus representative serialization/save benchmarks. Regenerate previews through the canonical driver because game/capture sources change, review representative captures and navigation, and run feasible target validation. Attempt local toolchain setup in ignored .tools and report unavailable phone evidence explicitly. Prior environment limitations above describe the review checkpoint, not a permanent prohibition on implementation.

## Implemented and locally verified — 2026-11-15

`Game.Flush` now holds `flushMu` across snapshot selection, storage and error reporting. Echo/Cluster ticks report state changes through the existing app update functions. Unchanged updates neither encode nor advance the save clock; the transition to idle flushes the last changed snapshot once. No queue, new save format, reduced active snapshot cadence or gameplay rule change was introduced.

Regression evidence: the prior `c99952a` implementation fails the flush-boundary and stationary-save checks in an isolated copy (only the test mutex field was added to compile the boundary test; baseline Flush never acquires it). The new implementation passes app, rule, storage, menu, presentation and preview tests with the race detector. Checks include write failure/recovery, stationary snapshots/writes, changed-tick persistence, final idle-transition flush, and Echo/Cluster transition reporting. The queued-flush test uses a bounded scheduling window; this is a targeted regression, not exhaustive interleaving proof.

Three sequential benchmark samples per revision, after capture completed, on Linux/amd64, Go 1.26.4, AMD Custom CPU 1772, populated eighteen-game collection, same local filesystem/cache and virtual display, 300ms samples:

| Workload | Before median | After median | Allocation evidence |
| --- | --- | --- | --- |
| Idle Echo update | 51,189 ns | 29.71 ns | ~18.5 KB / 58 allocations -> zero |
| Idle Cluster update | 49,626 ns | 32.47 ns | ~18.5 KB / 58 allocations -> zero |
| Collection encoding | 25,851 ns | 25,707 ns | ~9.5 KB / 1 allocation, unchanged |
| Store save including backup | 573,452 ns | 557,662 ns | ~266 KB / 1,704 allocations, unchanged |

Retain the change: it removes redundant idle work and passes correctness checks. Do not interpret small encode/save timing differences as optimization gains. These microbenchmarks omit rendering, OS scheduling variation and phone storage/energy behavior. The initial overlapping exploratory samples were replaced by the sequential samples above. Logs: `reports/save-refinement/{before,after}.txt`, `baseline-regressions.log`, `regressions.log`, `app-final.log`.

Canonical `scripts/build.py --generate-previews` regenerated all 90 images and the checked source manifest (2,218,104 compressed bytes). Reviewed wireframe/default/vivid posters and first/middle/last default/vivid frames across all eighteen games: recognizable boards, aspect ratios, palette differences and sequence progression. Diagnostic sheets: `reports/save-refinement/preview-review-{0,1}.png`. Updated captures are not byte-identical to the earlier environment's assets; no cross-environment byte-determinism claim is made.

`python3 scripts/check_gallery.py --label save-refinement --checks` passes still freeze, wireframe, same-frame return, autoplay resume, maximum intensity, scrolling, settings, entry/return and absence of preview saves/counts. Host Firefox viewport 500x914: default and maximum median 17.08ms, p99 17.14ms, no frames above 50ms in the samples. This is a smoke/performance observation, not a before/after gameplay FPS comparison or phone guarantee.

Canonical `scripts/build.py --verify` passes Linux/amd64 (13,888,280 bytes) and WASM (14,371,824 bytes); report `reports/build/build-report-20261116T015937Z.json`. Native verification used locally unpacked X11 headers/libraries and Xvfb under ignored `.tools`, without installing system packages. Go 1.26.4 was downloaded and verified against the official release manifest. Android is infeasible here: SDK/NDK, Java and ADB are missing. Setup now fails before downloads with the canonical apt remedy. No deployment occurred. Phone memory, cold loading, sustained frame stalls and energy measurements remain explicitly open.
