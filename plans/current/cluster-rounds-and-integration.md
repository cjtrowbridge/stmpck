# Cluster: Rounds and integration

This plan owns **Rounds and integration** under the [Cluster roadmap deliverable](../../ROADMAP.md#cluster). The requirement addresses Five starting rows; clear all to win, reaching bottom row loses. Drag to aim and release to fire, directional aim/confirm and Fire button. Save board, shot and queued colors; deliberate replay and real-model preview. The intended outcome is a small, playable implementation consistent with shared offline, sensory-control and accessibility commitments.

## Work and checks

Implement in the game's deterministic model and shared app integration. Verify relevant rule boundaries, saved-state continuity, gallery behavior and keyboard/touch controls. Keep gameplay timing independent of intensity and preserve existing game identifiers and saves. Initial parameters are defaults for playtesting, not a claim of finished balancing. See the [deliverable plan](cluster.md) for coordinated evidence; device and broader release checks remain pending.

## Status

Implemented. Focused rules/save tests and race checks pass; Android and WebAssembly compile. Browser checks exercise the game's controls and preserve its save across reload. See the [deliverable's verification](cluster.md#verification). Phone, physical-controller, broader accessibility and balance checks remain pending.
