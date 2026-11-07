# Territory — Goal

For the [Territory deliverable](../../ROADMAP.md#3-territory), implement **Goal:** Capture every territory and defeat any remaining enemy forces in transit. Show human elimination when neither owned territory nor a force in transit remains.

## Implementation

Replace the prior turn-based implementation with the real-time model and shared renderer. Preserve existing maps and armies on save migration. Validate deterministic movement, production, combat, persistence, and rendered borders; deploy through the canonical build pipeline.

Status: implemented. Rules, geometry/border coverage, save migration, and mid-flight continuation tests pass, including the race detector. Android and WebAssembly builds pass. Final phone verification is recorded in reports/territory-realtime/. Broader accessibility and longer balance playtests remain release work.
