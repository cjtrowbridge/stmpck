# Bounce — Scoring

For [Bounce](../../ROADMAP.md#5-bounce), implement the roadmap requirement: **Scoring:** Use a short fixed match—first to five points. Represent progress with simple markers where practical.

## Work and validation

Implement this requirement in the Bounce state model and shared app integration as appropriate. Keep presentation independent of game rules, preserve state across navigation, and verify relevant behavior with deterministic model tests and the canonical Android/WebAssembly build. Physical input and accessibility validation are recorded separately, not inferred from compilation.

## Status

Implemented; coordinate through [bounce.md](bounce.md); preserve this one-to-one roadmap link.

## Implementation evidence

Implemented in `internal/play/bounce.go`: normalized rectangular arena, human/computer paddles, direct and directional controls, fixed-step substeps/crossing collision tests, three speeds, bounded opponent reaction/movement, deliberate serves, first-to-five scoring, and restart. Tests cover serve gating, paddle/wall collisions, scoring, and opponent movement bounds. Phone play and Back-to-paused saving pass.

This requirement is implemented within the playable game. Relevant pure-Go model/storage tests and Android/WebAssembly builds pass. Phone smoke checks cover the main touch path, Back, and saved sessions, not every possible rule outcome or assistive technology. Shared release verification remains open.
