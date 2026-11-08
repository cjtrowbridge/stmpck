# flow: Controls and goal

This plan owns **Controls and goal** under [flow](../../ROADMAP.md#flow). It addresses that single requirement by implementing: Tap to rotate clockwise; grid cursor and confirm for keyboard/controller. Completion is visible through connected markers and a result.

## Implementation and validation

Implemented in the game's deterministic model and shared additions controls/rendering, save format and gallery. Focused rules and save tests, long-running validity checks, and navigation/storage/presentation tests pass, including the race detector. Android debug and WebAssembly builds succeed. See the [deliverable evidence](flow.md#evidence-and-remaining-work) and [journal](../../journal/2026-11-08.md). Phone touch checks are deferred because the test phone is disconnected; do not claim device verification.
