# Merge Controls

For the corresponding [roadmap game requirement](../../ROADMAP.md), implement **Controls:** Use touch swipes or mouse drags, with keyboard/controller directional equivalents. Do not show on-screen directional arrow buttons.

## Implementation and verification

Update shared pointer handling and the game renderer. Territory gestures use actual map hit testing and existing immutable movement orders; do not act during setup, pause, settings, or restart confirmation. Verify valid and invalid endpoints and immediate source deductions with model tests; build and deploy through the canonical pipeline. Physical controller and broader accessibility checks remain separate release work.

Status: implemented. Focused model tests pass, including the race detector. Android and WebAssembly builds pass and the APK was installed/launched on the authorized phone. The phone was locked during subsequent touch verification, so this revision’s live gesture/visual check remains pending.
