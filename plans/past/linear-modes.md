# Linear — Modes

This plan implements the corresponding [roadmap](../../ROADMAP.md) requirement: **Modes:** Play against the computer only for now. Move local two-player mode to the future PvP roadmap; existing local two-player saves retain their board and continue against the computer.

## Implementation and verification

Remove the local two-player setup option for now; Linear plays against the computer. Preserve existing saved boards while converting their mode to computer play. PvP is deferred.

Implemented. Focused rules, geometry, color, navigation, and save-migration tests pass. Android and WebAssembly builds pass. Device verification is recorded separately; broader accessibility and controller checks remain open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
