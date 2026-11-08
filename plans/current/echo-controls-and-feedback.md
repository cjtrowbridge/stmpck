# echo: Controls and feedback

This plan owns **Controls and feedback** under [echo](../../ROADMAP.md#echo). It addresses that single requirement by implementing: Large pads have distinct shapes, positions and numbers; touch or cursor/confirm; silent play.

## Implementation and validation

Implemented in the game's deterministic model and shared additions controls/rendering, save format and gallery. Focused rules and save tests, long-running validity checks, and navigation/storage/presentation tests pass, including the race detector. Android debug and WebAssembly builds succeed. See the [deliverable evidence](echo.md#evidence-and-remaining-work) and [journal](../../journal/2026-11-08.md). Phone touch checks are deferred because the test phone is disconnected; do not claim device verification.
