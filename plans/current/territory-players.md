# Territory — Players

For the [Territory deliverable](../../ROADMAP.md#3-territory), implement **Players:** Choose one to three computer opponents; each player starts with exactly one territory and ten units, with all remaining territories neutral and empty. Show ownership immediately in previews and scramble starting positions independently.

## Implementation

Replace the prior turn-based implementation with the real-time model and shared renderer. Preserve existing maps and armies on save migration. Validate deterministic movement, production, combat, persistence, and rendered borders; deploy through the canonical build pipeline.

Status: implemented. Rules, geometry/border coverage, save migration, and mid-flight continuation tests pass, including the race detector. Android and WebAssembly builds pass. Final phone verification is recorded in reports/territory-realtime/. Broader accessibility and longer balance playtests remain release work.
