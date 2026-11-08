# Echo

This plan owns the [Echo deliverable](../../ROADMAP.md#echo): add a small, offline Echo game to the collection. Its requirement plans below define the rules, direct controls, readable presentation and save/preview integration together; this parent coordinates them without replacing their one-to-one mappings.

## Initial implementation

Seeded four-pad sequence grows one cue after success, up to 32 cues. Large pads have distinct shapes, positions and numbers; touch or cursor/confirm; silent play. Save sequence, phase, response and clocks. Pause and retry the same sequence after failure; paced demonstration and response preview.

- [Sequence](echo-sequence.md)
- [Controls and feedback](echo-controls-and-feedback.md)
- [Integration](echo-integration.md)

## Evidence and remaining work

Implemented in `internal/play/echo.go`, with shared additions input/rendering, collection saves and actual-model demonstrations. The original five identifiers remain unchanged and older saves load with the six appended games unstarted. Tests in `internal/play/additions_test.go` cover rules, long-running validity, corruption rejection and save/resume; navigation/storage/presentation regressions and race-detector checks pass. Android debug and WebAssembly builds succeed. Linux is blocked by missing host X11 headers.

Phone testing is explicitly deferred because the device is disconnected. Initial gameplay balance, physical controllers, screen readers and wider device/accessibility coverage remain release work. No new assets, runtime dependencies, audio output or network requests were added.
