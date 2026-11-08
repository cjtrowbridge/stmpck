# Shift

This plan owns the [Shift deliverable](../../ROADMAP.md#shift): add the selected game with deterministic offline rules, clear controls, local saves and actual-game previews. It coordinates the separate requirements below without replacing their one-to-one plans.

## Initial scope

4×4 numbered sliding puzzle; generate each scramble by legal moves from the solved board, avoiding immediate reversals and a solved start.

Tap a neighbor of the empty cell or use directional input; visual completion, one-step undo and confirmed new puzzle. Persist arrangement/initial/undo/move count; preview uses legal slides.

- [Sliding and solvability](shift-sliding-and-solvability.md)
- [Interaction and integration](shift-interaction-and-integration.md)

## Verification

Implemented in `internal/play/shift.go`, with shared expansion controls, procedural rendering and game-driven previews. Rule tests, corrupted-save rejection, old eleven-game save compatibility and deterministic continuation pass, including race checks. Android debug and WebAssembly builds succeed through `scripts/build.py --verify`; Linux still lacks host X11 headers.

Local Firefox interaction and reload-persistence checks pass at a 480×914 viewport. See `reports/seven-games/checks.json`, its screenshots and the [journal](../../journal/2026-11-08.md). Phone/controller/accessibility verification and hands-on balancing remain separate acceptance work. No phone deployment, external assets or runtime dependencies were introduced.
