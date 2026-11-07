# Merge

Implement the [Merge deliverable](../../ROADMAP.md#1-merge) as a playable offline game with rules, input, presentation, shared controls, and local saves. Its individual requirement plans own the entries below.

## Work

Use small deterministic state models in `internal/play`, shared rendering for live games and demonstrations, and the existing menu/intensity system. Keep mechanical time separate from decorative animation. Verify rules with pure Go tests and build/deploy through the canonical command. Preserve pending accessibility/device verification explicitly.

## Status

Playable implementation complete; shared release acceptance remains open.

Implemented in `internal/play/merge.go` and shared app rendering/input: 4×4/5×5/6×6 setup, a fixed 2,048 goal, 90% 2 / 10% 4 spawns, single-merge-per-move rules, score, movement feedback, numeric/dot-exponent presentation, and one deterministic undo including RNG. State tests cover no-op moves, merging, end conditions, target continuation, grid sizes, and undo replay. Phone touch moves and exact board/RNG undo pass.

Every game uses the same renderer for demonstrations and play, shares pause/new/back controls, and persists its model through `internal/storage`. Broader accessibility, physical controller verification, and extended playtesting remain pending. Evidence: `reports/five-games/checks.json`, game screenshots in the same directory, and `.tools/five-games-final.log`.
