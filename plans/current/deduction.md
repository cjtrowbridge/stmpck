# Deduction

This plan owns the [Deduction deliverable](../../ROADMAP.md#deduction): add the selected game with deterministic offline rules, clear controls, local saves and actual-game previews. It coordinates the separate requirements below without replacing their one-to-one plans.

## Initial scope

9×9 Sudoku with 3×3 regions. Shuffle a valid completed grid, then remove clues only while a solver proves uniqueness; aim for 36 clues.

Cell selection and visible digit palette, keyboard digits and controller cycling, erase, one-step undo, immutable clues and conflict/completion feedback. Save clues/entries/solution/undo; actual solving preview. Notes/hints deferred.

- [Puzzle rules](deduction-puzzle-rules.md)
- [Interaction and integration](deduction-interaction-and-integration.md)

## Verification

Implemented in `internal/play/deduction.go`, with shared expansion controls, procedural rendering and game-driven previews. Rule tests, corrupted-save rejection, old eleven-game save compatibility and deterministic continuation pass, including race checks. Android debug and WebAssembly builds succeed through `scripts/build.py --verify`; Linux still lacks host X11 headers.

Local Firefox interaction and reload-persistence checks pass at a 480×914 viewport. See `reports/seven-games/checks.json`, its screenshots and the [journal](../../journal/2026-11-08.md). Phone/controller/accessibility verification and hands-on balancing remain separate acceptance work. No phone deployment, external assets or runtime dependencies were introduced.
