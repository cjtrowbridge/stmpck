# caliban: Controls and progression

This plan owns **Controls and progression** under [caliban](../../ROADMAP.md#caliban). It addresses that single requirement by implementing: Touch/drag chooses lane; left/right alternatives; three lives, deterministic waves; fixed gameplay speed independent of intensity.

## Implementation and validation

Implemented in the game's deterministic model and shared additions controls/rendering, save format and gallery. Focused rules and save tests, long-running validity checks, and navigation/storage/presentation tests pass, including the race detector. Android debug and WebAssembly builds succeed. See the [deliverable evidence](caliban.md#evidence-and-remaining-work) and [journal](../../journal/2026-11-08.md). Phone touch checks are deferred because the test phone is disconnected; do not claim device verification.
