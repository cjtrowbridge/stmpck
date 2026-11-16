# Android system Back

For [Define the smallest shared experience](../../ROADMAP.md#7-define-the-smallest-shared-experience), address **Handle Android system Back one level at a time: close settings, return from a game to the home menu while preserving catalog position, and stay at home if already there. Share navigation with Escape/controller Back.** This plan owns the implemented menu requirement; unfinished integration and verification retain separate roadmap entries.

## Implementation and evidence

Modern Android Back callbacks and the legacy button route thread-safe requests into the game update loop. Nested settings → game → home and repeated Back at home pass state tests. Android/WebAssembly builds and phone install/launch pass; evidence is .tools/android-back-build.log. Physical system button/gesture verification remains explicitly unchecked because the phone locked during the attempted check. Extending navigation to playable games remains separate future work.

## Status

Complete for the stated implementation scope. Reviewed against the roadmap on 2026-11-07.

The former help overlay is replaced by the [settings mockup](settings-mockup.md). Optional control explanations remain in settings; Back dismisses settings first.

The five-game phone touch check now confirms system Back returns to the gallery and leaves Snake/Bounce paused; all sessions survive process restart. Evidence: `reports/five-games/checks.json`.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
