# Crossing

This plan owns the [Crossing deliverable](../../ROADMAP.md#crossing): add the selected game with deterministic offline rules, clear controls, local saves and actual-game previews. It coordinates the separate requirements below without replacing their one-to-one plans.

## Initial scope

Nine-column crossing with three road lanes, five water lanes with moving logs, safe banks and median. Swipe or directional controls hop one cell; ride logs, avoid cars/water and reach three marked destinations.

Three lives and three levels, with pace determined by level. Save positions, lane phases, destinations, lives and clocks; pause hazards together. Deliberate retry/Next and model-driven preview.

- [Traversal](crossing-traversal.md)
- [Levels and integration](crossing-levels-and-integration.md)

## Verification

Implemented in `internal/play/crossing.go`, with shared expansion controls, procedural rendering and game-driven previews. Rule tests, corrupted-save rejection, old eleven-game save compatibility and deterministic continuation pass, including race checks. Android debug and WebAssembly builds succeed through `scripts/build.py --verify`; Linux still lacks host X11 headers.

Local Firefox interaction and reload-persistence checks pass at a 480×914 viewport. See `reports/seven-games/checks.json`, its screenshots and the [journal](../../journal/2026-11-08.md). Phone/controller/accessibility verification and hands-on balancing remain separate acceptance work. No phone deployment, external assets or runtime dependencies were introduced.
