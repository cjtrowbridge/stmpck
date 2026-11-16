# Bounce — Controls

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Controls:** Move the paddle with relative up/down swipes or drags anywhere on the app screen, including outside the arena, so the finger need not obscure play. Keep keyboard/controller equivalents; a tap on a control still performs that control action.

## Implementation and verification

Relative vertical swipes anywhere on the app screen move the paddle, not absolute finger positioning inside the arena. Distinguish drags from control taps.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
