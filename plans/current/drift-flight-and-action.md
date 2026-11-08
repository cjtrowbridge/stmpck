# Drift: Flight and action

This plan owns **Flight and action** under the [Drift roadmap deliverable](../../ROADMAP.md#drift--provisional-name). The requirement addresses Wrapping space, inertial ship movement, automatic firing, and rocks splitting twice. Drag anywhere to steer/thrust; keyboard/controller steering and a thrust toggle. The intended outcome is a small, playable implementation consistent with shared offline, sensory-control and accessibility commitments.

## Work and checks

Implement in the game's deterministic model and shared app integration. Verify relevant rule boundaries, saved-state continuity, gallery behavior and keyboard/touch controls. Keep gameplay timing independent of intensity and preserve existing game identifiers and saves. Initial parameters are defaults for playtesting, not a claim of finished balancing. See the [deliverable plan](drift.md) for coordinated evidence; device and broader release checks remain pending.

## Status

Implemented. Focused rules/save tests and race checks pass; Android and WebAssembly compile. Browser checks exercise the game's controls and preserve its save across reload. See the [deliverable's verification](drift.md#verification). Phone, physical-controller, broader accessibility and balance checks remain pending.
