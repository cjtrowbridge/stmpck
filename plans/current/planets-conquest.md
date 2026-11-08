# planets: Conquest

This plan owns **Conquest** under [planets](../../ROADMAP.md#planets). It addresses that single requirement by implementing: Each player starts with one planet and ten units. Every owned body produces one unit per three seconds; AI follows identical orders. Deterministic one-for-one combat, survivor capture, victory only when hostile bodies and fleets are gone.

## Implementation and validation

Implemented in the game's deterministic model and shared additions controls/rendering, save format and gallery. Focused rules and save tests, long-running validity checks, and navigation/storage/presentation tests pass, including the race detector. Android debug and WebAssembly builds succeed. See the [deliverable evidence](planets.md#evidence-and-remaining-work) and [journal](../../journal/2026-11-08.md). Phone touch checks are deferred because the test phone is disconnected; do not claim device verification.
