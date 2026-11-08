# Stack

This plan owns the [Stack deliverable](../../ROADMAP.md#stack): add the selected game with deterministic offline rules, clear controls, local saves and actual-game previews. It coordinates the separate requirements below without replacing their one-to-one plans.

## Initial scope

10×20 board, seven four-cell shapes, shuffled seven-piece bags, clockwise rotation with small horizontal kicks, row clearing and blocked-spawn failure.

Explicit slow/medium/fast gravity, deliberate hard drop and a 40-line goal. Persist board, falling/next pieces, bag/RNG and clock. Preview uses actual placements.

- [Placement and clearing](stack-placement-and-clearing.md)
- [Pace and integration](stack-pace-and-integration.md)

## Verification

Implemented in `internal/play/stack.go`, with shared expansion controls, procedural rendering and game-driven previews. Rule tests, corrupted-save rejection, old eleven-game save compatibility and deterministic continuation pass, including race checks. Android debug and WebAssembly builds succeed through `scripts/build.py --verify`; Linux still lacks host X11 headers.

Local Firefox interaction and reload-persistence checks pass at a 480×914 viewport. See `reports/seven-games/checks.json`, its screenshots and the [journal](../../journal/2026-11-08.md). Phone/controller/accessibility verification and hands-on balancing remain separate acceptance work. No phone deployment, external assets or runtime dependencies were introduced.
