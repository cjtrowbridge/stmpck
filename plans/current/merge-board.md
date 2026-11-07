# Merge — Board

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Board:** Let the player select a grid size before starting a new game; start the selected board with two tiles. Offer stacked 4×4, 5×5, and 6×6 choices. Selecting a size starts immediately; omit setup Start/Back buttons while retaining system/header Back.

## Implementation and verification

Stack 4×4, 5×5, and 6×6 choices vertically; choosing one starts immediately. Remove setup footer Start/Back while preserving header/system Back.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.
