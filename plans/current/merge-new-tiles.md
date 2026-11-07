# Merge — New tiles

For [Merge](../../ROADMAP.md#1-merge), implement the roadmap requirement: **New tiles:** After a move that changes the board, spawn one tile in an empty cell. Spawn 2 with 90% probability and 4 with 10% probability.

## Work and validation

Implement this requirement in the Merge state model and shared app integration as appropriate. Keep presentation independent of game rules, preserve state across navigation, and verify relevant behavior with deterministic model tests and the canonical Android/WebAssembly build. Physical input and accessibility validation are recorded separately, not inferred from compilation.

## Status

Implemented; coordinate through [merge.md](merge.md); preserve this one-to-one roadmap link.

## Implementation evidence

Implemented in `internal/play/merge.go` and shared app rendering/input: 4×4/5×5/6×6 setup, a fixed 2,048 goal, 90% 2 / 10% 4 spawns, single-merge-per-move rules, score, movement feedback, numeric/dot-exponent presentation, and one deterministic undo including RNG. State tests cover no-op moves, merging, end conditions, target continuation, grid sizes, and undo replay. Phone touch moves and exact board/RNG undo pass.

This requirement is implemented within the playable game. Relevant pure-Go model/storage tests and Android/WebAssembly builds pass. Phone smoke checks cover the main touch path, Back, and saved sessions, not every possible rule outcome or assistive technology. Shared release verification remains open.
