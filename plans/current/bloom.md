# Bloom

This plan owns the [Bloom deliverable](../../ROADMAP.md#bloom): add a small, offline Bloom game to the collection. Its requirement plans below define the rules, direct controls, readable presentation and save/preview integration together; this parent coordinates them without replacing their one-to-one mappings.

## Initial implementation

Conway B3/S23 on a wrapping 24×24 board. Select 2, 6 or 12 generations per second independently of intensity. Paint/toggle while paused, run/pause, single step and confirmed clearing. Grid cursor/confirm alternatives. Persist cells, generation, rate and clock; genuine evolution preview; no score or required ending.

- [Simulation](../past/bloom-simulation.md)
- [Interaction](../past/bloom-interaction.md)
- [Integration](../past/bloom-integration.md)

- [Recognizable preview patterns](../past/bloom-preview-patterns.md)

- [Pattern picker and placement](../past/bloom-pattern-picker.md)

## Refinements implemented 2026-11-08

A six-pattern catalog now drives gallery examples and a paused picker with wrapped additive placement, preview and cancellation. Focused rules/save/race tests, browser interactions and Android/WebAssembly builds pass. Evidence is recorded in the refinement requirement plans linked above; phone verification remains pending.

## Evidence and remaining work

Implemented in `internal/play/bloom.go`, with shared additions input/rendering, collection saves and actual-model demonstrations. The original five identifiers remain unchanged and older saves load with the six appended games unstarted. Tests in `internal/play/additions_test.go` cover rules, long-running validity, corruption rejection and save/resume; navigation/storage/presentation regressions and race-detector checks pass. Android debug and WebAssembly builds succeed. Linux is blocked by missing host X11 headers.

Phone testing is explicitly deferred because the device is disconnected. Initial gameplay balance, physical controllers, screen readers and wider device/accessibility coverage remain release work. No new assets, runtime dependencies, audio output or network requests were added.

## Remaining work and closure — 2026-11-16

Complete the previously recorded phone, physical-controller, native accessibility and hands-on balance acceptance for this game, including its implemented refinements. Archive when that evidence satisfies the mapped deliverable and shared completion criteria; implementation requirement plans are archived separately without claiming these checks passed.
