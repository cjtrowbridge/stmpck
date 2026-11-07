# Territory — Map

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Map:** Offer several live map previews generated from non-grid geography, with irregular coastlines and truly varied territory shapes, sizes, and adjacency. Let the player scramble map geometry independently of starting positions. Selecting a map begins the match; omit setup Start/Back buttons.

## Implementation and verification

Replace the grid mesh with seeded non-grid geography and shared curved borders. Show three live previews; independently scramble geometry and starting positions. A map card launches directly.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.
