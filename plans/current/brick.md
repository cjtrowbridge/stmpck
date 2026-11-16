# Brick

This plan owns the [Brick deliverable](../../ROADMAP.md#brick): add a small, offline Brick game to the collection. Its requirement plans below define the rules, direct controls, readable presentation and save/preview integration together; this parent coordinates them without replacing their one-to-one mappings.

## Initial implementation

Normalized paddle/ball physics with substeps and one-hit bricks in three layouts. Horizontal gestures anywhere and left/right alternatives. Deliberate serve, three lives, next board after clearing. Persist ball, paddle, velocity, bricks, lives and level; pause/restart; genuine rally preview with readable silhouettes.

- [Arena and collisions](../past/brick-arena-and-collisions.md)
- [Controls and rounds](../past/brick-controls-and-rounds.md)
- [Integration](../past/brick-integration.md)

## Evidence and remaining work

Implemented in `internal/play/brick.go`, with shared additions input/rendering, collection saves and actual-model demonstrations. The original five identifiers remain unchanged and older saves load with the six appended games unstarted. Tests in `internal/play/additions_test.go` cover rules, long-running validity, corruption rejection and save/resume; navigation/storage/presentation regressions and race-detector checks pass. Android debug and WebAssembly builds succeed. Linux is blocked by missing host X11 headers.

Phone testing is explicitly deferred because the device is disconnected. Initial gameplay balance, physical controllers, screen readers and wider device/accessibility coverage remain release work. No new assets, runtime dependencies, audio output or network requests were added.

## Remaining work and closure — 2026-11-16

Complete the previously recorded phone, physical-controller, native accessibility and hands-on balance acceptance for this game, including its implemented refinements. Archive when that evidence satisfies the mapped deliverable and shared completion criteria; implementation requirement plans are archived separately without claiming these checks passed.
