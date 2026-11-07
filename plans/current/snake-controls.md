# Snake — Controls

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Controls:** Use screen swipes for touch steering, with keys/controller equivalents and no on-screen directional arrow buttons. Prevent direct reversal into the snake's neck and handle rapid direction changes consistently.

## Implementation and verification

Remove touch arrow buttons. Screen swipes steer; keyboard/controller directional equivalents remain. Swipes never change in-game intensity.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.
