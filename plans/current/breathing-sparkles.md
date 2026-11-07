# Breathing borders and progressive sparkles

For [presentation and accessibility rules](../../ROADMAP.md#5-establish-presentation-and-accessibility-rules), address **extend the vivid view with breathing borders and random sparkles, then three progressively more intense rainbow/sparkle levels**. Add four opt-in levels to the right while preserving the existing views and default.

## Implementation

Use procedural strokes and tiny sparkles with smooth fade envelopes, no new assets. First level uses existing palette colors, a four-second breath (15 cycles per minute), and two new sparkles per second. Higher levels use five, ten, and eighteen sparkles per second with progressively stronger/faster rainbow accents and breathing. Keep backgrounds stable and controls readable. Random positions are seeded per session; time determines bounded sparkles rather than an ever-growing particle list. Effects stop when leaving these levels and never change gameplay state or add sound.

## Validation

Check navigation across all eight levels, unchanged launch/autoplay defaults, bounded effects cost, and Android/WebAssembly builds. Deploy through the canonical test-phone launch command. Visual/device checks remain separate from build verification.

## Results

Navigation/autoplay tests pass across all eight levels. Android and WebAssembly build successfully, and the fresh APK installs and launches on the Pixel through the canonical command. Phone captures in `reports/sparkles/` show breathing borders, changed accent colors, and scattered sparkles; the high-intensity screenshot was visually inspected. The effect loop is bounded to at most 29 live sparkle candidates. Sustained performance and accessibility testing remain pending. Linux still lacks X11 headers.

Roadmap audit: the implemented timing constants are now explicit in the mapped requirement (breathing rates 15/45/120/180 cycles per minute, rainbow cycles 24/16/10 seconds). Sustained performance and accessibility validation remain separate unchecked roadmap work.

## Breathing tempo revision

The breathing-rate progression is 15, 45, 120, and 180 complete cycles per minute: resting-style breathing at the first effect level, then accelerating to happy-hardcore tempo. A cycle is one full expansion/contraction, not a half-cycle. The continuous cosine envelope controls border width and accent blending; backgrounds remain stable. These are aesthetic animation settings, not a guided breathing exercise. Android/WebAssembly builds and test-phone installation/launch pass for this timing revision (`.tools/breathing-tempo-build.log`). Visual assessment of the revised tempo remains pending.
