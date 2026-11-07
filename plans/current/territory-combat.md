# Territory — Combat

For the [Territory deliverable](../../ROADMAP.md#3-territory), implement **Combat:** Resolve combat only on arrival: each side rolls one die per round, ties favor the defender, and one unit is lost per round until one force is exhausted. Friendly arrivals reinforce; empty neutral territory is occupied.

## Implementation

Replace the prior turn-based implementation with the real-time model and shared renderer. Preserve existing maps and armies on save migration. Validate deterministic movement, production, combat, persistence, and rendered borders; deploy through the canonical build pipeline.

Status: implemented. Rules, geometry/border coverage, save migration, and mid-flight continuation tests pass, including the race detector. Android and WebAssembly builds pass. Final phone verification is recorded in reports/territory-realtime/. Broader accessibility and longer balance playtests remain release work.
