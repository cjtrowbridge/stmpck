# planets: Unrestricted destinations

This plan owns **Unrestricted destinations** under [planets](../../ROADMAP.md#planets). It addresses that single requirement by implementing: Any owned body can send to any other body; no adjacency restriction.

## Implementation and validation

Implemented in the game's deterministic model and shared additions controls/rendering, save format and gallery. Focused rules and save tests, long-running validity checks, and navigation/storage/presentation tests pass, including the race detector. Android debug and WebAssembly builds succeed. See the [deliverable evidence](planets.md#evidence-and-remaining-work) and [journal](../../journal/2026-11-08.md). Phone touch checks are deferred because the test phone is disconnected; do not claim device verification.
