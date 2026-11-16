# Bounce

Implement the [Bounce deliverable](../../ROADMAP.md#5-bounce) as a playable offline game with rules, input, presentation, shared controls, and local saves. Its individual requirement plans own the entries below.

## Work

Use small deterministic state models in `internal/play`, shared rendering for live games and demonstrations, and the existing menu/intensity system. Keep mechanical time separate from decorative animation. Verify rules with pure Go tests and build/deploy through the canonical command. Preserve pending accessibility/device verification explicitly.

## Status

Playable implementation complete; shared release acceptance remains open.

Implemented in `internal/play/bounce.go`: normalized rectangular arena, human/computer paddles, direct and directional controls, fixed-step substeps/crossing collision tests, three speeds, bounded opponent reaction/movement, deliberate serves, first-to-five scoring, and restart. Tests cover serve gating, paddle/wall collisions, scoring, and opponent movement bounds. Phone play and Back-to-paused saving pass.

Every game uses the same renderer for demonstrations and play, shares pause/new/back controls, and persists its model through `internal/storage`. Broader accessibility, physical controller verification, and extended playtesting remain pending. Evidence: `reports/five-games/checks.json`, game screenshots in the same directory, and `.tools/five-games-final.log`.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
