# Shift: Sliding and solvability

This plan owns **Sliding and solvability** under the [Shift roadmap deliverable](../../ROADMAP.md#shift). The requirement addresses 4×4 numbered sliding puzzle; generate each scramble by legal moves from the solved board, avoiding immediate reversals and a solved start. The intended outcome is a small, playable implementation consistent with shared offline, sensory-control and accessibility commitments.

## Work and checks

Implement in the game's deterministic model and shared app integration. Verify relevant rule boundaries, saved-state continuity, gallery behavior and keyboard/touch controls. Keep gameplay timing independent of intensity and preserve existing game identifiers and saves. Initial parameters are defaults for playtesting, not a claim of finished balancing. See the [deliverable plan](shift.md) for coordinated evidence; device and broader release checks remain pending.

## Status

Implemented. Focused rules/save tests and race checks pass; Android and WebAssembly compile. Browser checks exercise the game's controls and preserve its save across reload. See the [deliverable's verification](shift.md#verification). Phone, physical-controller, broader accessibility and balance checks remain pending.
