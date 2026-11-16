# Merge — Rules

For [Merge](../../ROADMAP.md#1-merge), implement the roadmap requirement: **Rules:** Slide every tile in one direction. Matching tiles combine into the next value; each tile can merge only once per move.

## Work and validation

Implement this requirement in the Merge state model and shared app integration as appropriate. Keep presentation independent of game rules, preserve state across navigation, and verify relevant behavior with deterministic model tests and the canonical Android/WebAssembly build. Physical input and accessibility validation are recorded separately, not inferred from compilation.

## Status

Implemented; coordinate through [merge.md](merge.md); preserve this one-to-one roadmap link.

## Implementation evidence

Implemented in `internal/play/merge.go` and shared app rendering/input: 4×4/5×5/6×6 setup, a fixed 2,048 goal, 90% 2 / 10% 4 spawns, single-merge-per-move rules, score, movement feedback, numeric/dot-exponent presentation, and one deterministic undo including RNG. State tests cover no-op moves, merging, end conditions, target continuation, grid sizes, and undo replay. Phone touch moves and exact board/RNG undo pass.

This requirement is implemented within the playable game. Relevant pure-Go model/storage tests and Android/WebAssembly builds pass. Phone smoke checks cover the main touch path, Back, and saved sessions, not every possible rule outcome or assistive technology. Shared release verification remains open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
