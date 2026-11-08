# Cluster

This plan owns the [Cluster deliverable](../../ROADMAP.md#cluster): add the selected game with deterministic offline rules, clear controls, local saves and actual-game previews. It coordinates the separate requirements below without replacing their one-to-one plans.

## Initial scope

Eight-column staggered hex grid, four marked colors, fixed-speed launched pieces with side-wall reflection. Snap to a free neighbor; clear connected matches of at least three and drop pieces disconnected from ceiling.

Five starting rows; clear all to win, reaching bottom row loses. Drag to aim and release to fire, directional aim/confirm and Fire button. Save board, shot and queued colors; deliberate replay and real-model preview.

- [Aiming and matching](cluster-aiming-and-matching.md)
- [Rounds and integration](cluster-rounds-and-integration.md)

## Verification

Implemented in `internal/play/cluster.go`, with shared expansion controls, procedural rendering and game-driven previews. Rule tests, corrupted-save rejection, old eleven-game save compatibility and deterministic continuation pass, including race checks. Android debug and WebAssembly builds succeed through `scripts/build.py --verify`; Linux still lacks host X11 headers.

Local Firefox interaction and reload-persistence checks pass at a 480×914 viewport. See `reports/seven-games/checks.json`, its screenshots and the [journal](../../journal/2026-11-08.md). Phone/controller/accessibility verification and hands-on balancing remain separate acceptance work. No phone deployment, external assets or runtime dependencies were introduced.
