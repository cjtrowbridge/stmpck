# Stack: Pace and integration

This plan owns **Pace and integration** under the [Stack roadmap deliverable](../../ROADMAP.md#stack). The requirement addresses Explicit slow/medium/fast gravity, deliberate hard drop and a 40-line goal. Persist board, falling/next pieces, bag/RNG and clock. Preview uses actual placements. The intended outcome is a small, playable implementation consistent with shared offline, sensory-control and accessibility commitments.

## Work and checks

Implement in the game's deterministic model and shared app integration. Verify relevant rule boundaries, saved-state continuity, gallery behavior and keyboard/touch controls. Keep gameplay timing independent of intensity and preserve existing game identifiers and saves. Initial parameters are defaults for playtesting, not a claim of finished balancing. See the [deliverable plan](stack.md) for coordinated evidence; device and broader release checks remain pending.

## Status

Implemented. Focused rules/save tests and race checks pass; Android and WebAssembly compile. Browser checks exercise the game's controls and preserve its save across reload. See the [deliverable's verification](stack.md#verification). Phone, physical-controller, broader accessibility and balance checks remain pending.
