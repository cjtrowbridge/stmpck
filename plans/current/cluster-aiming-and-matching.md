# Cluster: Aiming and matching

This plan owns **Aiming and matching** under the [Cluster roadmap deliverable](../../ROADMAP.md#cluster). The requirement addresses Eight-column staggered hex grid, four marked colors, fixed-speed launched pieces with side-wall reflection. Snap to a free neighbor; clear connected matches of at least three and drop pieces disconnected from ceiling. The intended outcome is a small, playable implementation consistent with shared offline, sensory-control and accessibility commitments.

## Work and checks

Implement in the game's deterministic model and shared app integration. Verify relevant rule boundaries, saved-state continuity, gallery behavior and keyboard/touch controls. Keep gameplay timing independent of intensity and preserve existing game identifiers and saves. Initial parameters are defaults for playtesting, not a claim of finished balancing. See the [deliverable plan](cluster.md) for coordinated evidence; device and broader release checks remain pending.

## Status

Implemented. Focused rules/save tests and race checks pass; Android and WebAssembly compile. Browser checks exercise the game's controls and preserve its save across reload. See the [deliverable's verification](cluster.md#verification). Phone, physical-controller, broader accessibility and balance checks remain pending.
