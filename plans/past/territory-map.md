# Territory — Map

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Map:** Offer several live map previews generated from non-grid geography, with irregular coastlines and truly varied territory shapes, sizes, and adjacency. Let the player scramble map geometry independently of starting positions. Selecting a map begins the match; omit setup Start/Back buttons.

## Implementation and verification

Replace the grid mesh with seeded non-grid geography and shared curved borders. Show three live previews; independently scramble geometry and starting positions. A map card launches directly.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
