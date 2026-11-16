# Streamline the implementation

This plan owns **Streamline the implementation** under [Foundation 4](../../ROADMAP.md#4-simplification-and-demoscenization). The requirement covers simplifying shared/game code while preserving rules, saves, accessible controls and sensory choices. The first active scope is the save path: establish ownership, ordering and cost evidence before choosing a runtime change. Broader cleanup is unstarted.

## Review — 2026-11-15

`internal/app/play.go` publishes immutable encoded collection snapshots. `Flush` loads the snapshot before calling `Store.Save`; `internal/storage/storage.go` serializes writes with a mutex. This avoids simultaneous file writes but does not order snapshot selection with those writes.

A permitted interleaving is: caller A loads old snapshot S1 and is delayed; the game loop publishes S2; caller B loads S2, acquires the storage lock and saves; A resumes, acquires the same lock and saves S1 last. This is a source-level logical counterexample, not an observed device failure or a Go interleaving-test result. A clean race detector alone would not rule it out. The smallest candidate correction is serialization of Flush before snapshot selection, so persistence order follows selection order. Evaluate save-failure status ordering with the same ownership boundary. Do not add a worker queue or change the save format to address this.

Cost observations: changed gameplay updates encode the entire collection; several continuous games mark changed every active update. Flush validates new bytes, reads/validates the old save, writes the backup, then writes the primary with synchronous file syncs. These are actual operations visible in source, not proof of an unacceptable cost. Existing tests cover save validity/recovery but not the app-level concurrent selection interleaving.

## Dispositions and next evidence

- [x] Review ownership and identify the stale-write interleaving and candidate minimal correction.
- [ ] Reproduce the ordering failure with a controlled Go interleaving; implement and verify the smallest correction without changing gameplay/snapshot semantics.
- [ ] Measure full-collection encoding allocations/time and save latency/frequency on representative populated saves; correlate with device frame stalls before optimizing.
- [ ] Complete focused regression and applicable canonical build/preview validation before claiming a runtime change complete.

No Go executable is available in this session's environment (`command -v go` and local toolchain inspection). No benchmark, race-test pass, runtime correction or performance improvement is claimed. Defer runtime edits until they can be checked; source review and framework guidance can complete independently. Preserve existing snapshot freshness across suspension: reducing snapshot cadence is a behavioral decision, not a free optimization. Phone evidence remains distinct from host measurements.
