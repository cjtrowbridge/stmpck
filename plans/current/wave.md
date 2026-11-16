# Wave

This plan owns the [Wave deliverable](../../ROADMAP.md#wave): add the selected game with deterministic offline rules, clear controls, local saves and actual-game previews. It coordinates the separate requirements below without replacing their one-to-one plans.

## Initial scope

Eight columns by four rows of procedural enemies, horizontal formations descending at boundaries, automatic player fire and seeded enemy shots. Drag horizontally anywhere or use keyboard/controller movement.

Three lives, three waves, deliberate Next/restart and pause; persist enemies/projectiles/clocks/RNG. Difficulty grows with wave, independent of intensity; game-driven gallery.

- [Movement and action](../past/wave-movement-and-action.md)
- [Waves and integration](../past/wave-waves-and-integration.md)

## Verification

Implemented in `internal/play/wave.go`, with shared expansion controls, procedural rendering and game-driven previews. Rule tests, corrupted-save rejection, old eleven-game save compatibility and deterministic continuation pass, including race checks. Android debug and WebAssembly builds succeed through `scripts/build.py --verify`; Linux still lacks host X11 headers.

Local Firefox interaction and reload-persistence checks pass at a 480×914 viewport. See `reports/seven-games/checks.json`, its screenshots and the [journal](../../journal/2026-11-08.md). Phone/controller/accessibility verification and hands-on balancing remain separate acceptance work. No phone deployment, external assets or runtime dependencies were introduced.

## Remaining work and closure — 2026-11-16

Complete the previously recorded phone, physical-controller, native accessibility and hands-on balance acceptance for this game, including its implemented refinements. Archive when that evidence satisfies the mapped deliverable and shared completion criteria; implementation requirement plans are archived separately without claiming these checks passed.
