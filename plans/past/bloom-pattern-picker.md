# Bloom: Pattern picker and placement

This plan owns **Pattern picker and placement** under [Bloom](../../ROADMAP.md#bloom). The requirement addresses the requested refinement: A Patterns button, P key or controller right shoulder opens a visual two-column popup with named thumbnails. Pause while choosing/placing. Preview a centered stamp at the selected grid cell with toroidal wrapping; Place/confirm adds live cells without erasing existing cells. Cancel/Back leaves the board unchanged. Keep normal paint, step, rate and clear controls.

## Implementation and checks

Implemented 2026-11-08 using the defaults above. Model and browser checks cover wrapped additive stamps, save continuity, thumbnails, ghost placement, cancellation without mutation, pausing and modal Back.

Rules/navigation/storage/presentation checks pass with the race detector. Browser evidence: reports/game-refinements-2/checks.json and screenshots. Android and WebAssembly builds pass: reports/build/build-report-20261108T215620Z.json; Linux remains blocked by missing X11 headers. Phone, physical-controller, native screen-reader and balance verification remain pending.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
