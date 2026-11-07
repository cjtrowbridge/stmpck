# Territory — Committed forces

For the [Territory deliverable](../../ROADMAP.md#3-territory), implement **Committed forces:** Orders send all but one available unit and subtract that force immediately. Routes, destinations, and committed forces cannot be changed or recalled. Move at a fixed map-space speed of 0.12 units per second for every player.

## Implementation

Replace the prior turn-based implementation with the real-time model and shared renderer. Preserve existing maps and armies on save migration. Validate deterministic movement, production, combat, persistence, and rendered borders; deploy through the canonical build pipeline.

Status: implemented. Rules, geometry/border coverage, save migration, and mid-flight continuation tests pass, including the race detector. Android and WebAssembly builds pass. Final phone verification is recorded in reports/territory-realtime/. Broader accessibility and longer balance playtests remain release work.
