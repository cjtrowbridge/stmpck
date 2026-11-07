# Territory — Reinforcement

For the [Territory deliverable](../../ROADMAP.md#3-territory), implement **Reinforcement:** Each owned territory produces one unit every three seconds of active simulation; neutral territories do not produce units. Initial rate is a balancing default.

## Implementation

Replace the prior turn-based implementation with the real-time model and shared renderer. Preserve existing maps and armies on save migration. Validate deterministic movement, production, combat, persistence, and rendered borders; deploy through the canonical build pipeline.

Status: implemented. Rules, geometry/border coverage, save migration, and mid-flight continuation tests pass, including the race detector. Android and WebAssembly builds pass. Final phone verification is recorded in reports/territory-realtime/. Broader accessibility and longer balance playtests remain release work.
