# bloom: Simulation

This plan owns **Simulation** under [bloom](../../ROADMAP.md#bloom). It addresses that single requirement by implementing: Conway B3/S23 on a wrapping 24×24 board. Select 2, 6 or 12 generations per second independently of intensity.

## Implementation and validation

Implemented in the game's deterministic model and shared additions controls/rendering, save format and gallery. Focused rules and save tests, long-running validity checks, and navigation/storage/presentation tests pass, including the race detector. Android debug and WebAssembly builds succeed. See the [deliverable evidence](bloom.md#evidence-and-remaining-work) and [journal](../../journal/2026-11-08.md). Phone touch checks are deferred because the test phone is disconnected; do not claim device verification.
