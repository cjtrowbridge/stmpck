# Merge — Board

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Board:** Let the player select a grid size before starting a new game; start the selected board with two tiles. Offer stacked 4×4, 5×5, and 6×6 choices. Selecting a size starts immediately; omit setup Start/Back buttons while retaining system/header Back.

## Implementation and verification

Stack 4×4, 5×5, and 6×6 choices vertically; choosing one starts immediately. Remove setup footer Start/Back while preserving header/system Back.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
