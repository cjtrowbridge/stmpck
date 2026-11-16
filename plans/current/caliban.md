# Caliban

This plan owns the [Caliban deliverable](../../ROADMAP.md#caliban): add a small, offline Caliban game to the collection. Its requirement plans below define the rules, direct controls, readable presentation and save/preview integration together; this parent coordinates them without replacing their one-to-one mappings.

## Initial implementation

Twelve procedural lanes; rim movement, automatic inward shots and approaching enemies. Arena taps choose a lane; relative horizontal dragging anywhere moves around the rim; left/right alternatives; three lives, deterministic waves; fixed gameplay speed independent of intensity. Save wave, enemies, bullets, clocks and RNG. Pause/resume and restart; model preview and outlined threats.

- [Arena and action](../past/caliban-arena-and-action.md)
- [Controls and progression](../past/caliban-controls-and-progression.md)
- [Integration](../past/caliban-integration.md)

- [Drag anywhere](../past/caliban-drag-anywhere.md)

## Refinements implemented 2026-11-08

One screen width spans twelve lanes, relative to the lane at gesture start. Focused rules/save/race tests, browser interactions and Android/WebAssembly builds pass. Evidence is recorded in the refinement requirement plans linked above; phone verification remains pending.

## Evidence and remaining work

Implemented in `internal/play/caliban.go`, with shared additions input/rendering, collection saves and actual-model demonstrations. The original five identifiers remain unchanged and older saves load with the six appended games unstarted. Tests in `internal/play/additions_test.go` cover rules, long-running validity, corruption rejection and save/resume; navigation/storage/presentation regressions and race-detector checks pass. Android debug and WebAssembly builds succeed. Linux is blocked by missing host X11 headers.

Phone testing is explicitly deferred because the device is disconnected. Initial gameplay balance, physical controllers, screen readers and wider device/accessibility coverage remain release work. No new assets, runtime dependencies, audio output or network requests were added.

## Remaining work and closure — 2026-11-16

Complete the previously recorded phone, physical-controller, native accessibility and hands-on balance acceptance for this game, including its implemented refinements. Archive when that evidence satisfies the mapped deliverable and shared completion criteria; implementation requirement plans are archived separately without claiming these checks passed.
