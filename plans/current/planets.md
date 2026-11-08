# Planets

This plan owns the [Planets deliverable](../../ROADMAP.md#planets): add a small, offline Planets game to the collection. Its requirement plans below define the rules, direct controls, readable presentation and save/preview integration together; this parent coordinates them without replacing their one-to-one mappings.

## Initial implementation

Three selectable systems with five to seven planets and zero to two moons; orbital positions, owners and unit counts. Any owned body can send to any other body; no adjacency restriction. Tap source/destination or swipe; send all but one immediately. Fleets pursue a locked moving body at constant speed, faster than orbital motion; no recalling orders. Each player starts with one planet and ten units. Every owned body produces one unit per three seconds; AI follows identical orders. Deterministic one-for-one combat, survivor capture, victory only when hostile bodies and fleets are gone. Systems two and three include capturable moons orbiting parents, following the same unit and targeting rules. Persist system, phases, ownership, units and fleets; pause without offline catch-up. System selection and next-system progression; real-model previews and numeric owner cues.

- [Planetary systems](planets-planetary-systems.md)
- [Unrestricted destinations](planets-unrestricted-destinations.md)
- [Orders and travel](planets-orders-and-travel.md)
- [Conquest](planets-conquest.md)
- [Moons](planets-moons.md)
- [Integration](planets-integration.md)

- [Player starts the match](planets-player-start.md)

- [Opponent reaction progression](planets-reaction-progression.md)

## Refinements implemented 2026-11-08

New matches wait for a valid player order before production/opponents begin; opponents start at six-second intervals and approach two seconds over three active minutes. Timing persists; legacy matches retain their previous cadence. Focused rules/save/race tests, browser interactions and Android/WebAssembly builds pass. Evidence is recorded in the refinement requirement plans linked above; phone verification remains pending.

## Evidence and remaining work

Implemented in `internal/play/planets.go`, with shared additions input/rendering, collection saves and actual-model demonstrations. The original five identifiers remain unchanged and older saves load with the six appended games unstarted. Tests in `internal/play/additions_test.go` cover rules, long-running validity, corruption rejection and save/resume; navigation/storage/presentation regressions and race-detector checks pass. Android debug and WebAssembly builds succeed. Linux is blocked by missing host X11 headers.

Phone testing is explicitly deferred because the device is disconnected. Initial gameplay balance, physical controllers, screen readers and wider device/accessibility coverage remain release work. No new assets, runtime dependencies, audio output or network requests were added.
