# Crossing: Levels and integration

This plan owns **Levels and integration** under the [Crossing roadmap deliverable](../../ROADMAP.md#crossing). The requirement addresses Three lives and three levels, with pace determined by level. Save positions, lane phases, destinations, lives and clocks; pause hazards together. Deliberate retry/Next and model-driven preview. The intended outcome is a small, playable implementation consistent with shared offline, sensory-control and accessibility commitments.

## Work and checks

Implement in the game's deterministic model and shared app integration. Verify relevant rule boundaries, saved-state continuity, gallery behavior and keyboard/touch controls. Keep gameplay timing independent of intensity and preserve existing game identifiers and saves. Initial parameters are defaults for playtesting, not a claim of finished balancing. See the [deliverable plan](crossing.md) for coordinated evidence; device and broader release checks remain pending.

## Status

Implemented. Focused rules/save tests and race checks pass; Android and WebAssembly compile. Browser checks exercise the game's controls and preserve its save across reload. See the [deliverable's verification](crossing.md#verification). Phone, physical-controller, broader accessibility and balance checks remain pending.
