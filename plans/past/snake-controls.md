# Snake — Controls

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Controls:** Use screen swipes for touch steering, with keys/controller equivalents and no on-screen directional arrow buttons. Prevent direct reversal into the snake's neck and handle rapid direction changes consistently.

## Implementation and verification

Remove touch arrow buttons. Screen swipes steer; keyboard/controller directional equivalents remain. Swipes never change in-game intensity.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
