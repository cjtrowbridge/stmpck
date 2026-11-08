# Territory Attack Selection

For the corresponding [roadmap game requirement](../../ROADMAP.md), implement **Attack selection:** Swipe or drag from an owned source to an adjacent destination to issue an order on release. Retain source-then-destination taps and keyboard/controller selection. Allow friendly transfers and neutral or hostile destinations; reject invalid endpoints without changing existing orders.

## Implementation and verification

Update shared pointer handling and the game renderer. Territory gestures use actual map hit testing and existing immutable movement orders; do not act during setup, pause, settings, or restart confirmation. Verify valid and invalid endpoints and immediate source deductions with model tests; build and deploy through the canonical pipeline. Physical controller and broader accessibility checks remain separate release work.

Status: implemented. Focused model tests pass, including the race detector. Android and WebAssembly builds pass and the APK was installed/launched on the authorized phone. The phone was locked during subsequent touch verification, so this revision’s live gesture/visual check remains pending.
