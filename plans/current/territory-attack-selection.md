# Territory — Attack selection

For the [Territory deliverable](../../ROADMAP.md#3-territory), implement **Attack selection:** Select an owned source, then an adjacent territory to issue an immediate order. Allow friendly transfers and movement into neutral or hostile territory.

## Implementation

Replace the prior turn-based implementation with the real-time model and shared renderer. Preserve existing maps and armies on save migration. Validate deterministic movement, production, combat, persistence, and rendered borders; deploy through the canonical build pipeline.

Status: implemented. Rules, geometry/border coverage, save migration, and mid-flight continuation tests pass, including the race detector. Android and WebAssembly builds pass. Final phone verification is recorded in reports/territory-realtime/. Broader accessibility and longer balance playtests remain release work.
