# Bloom: Recognizable preview patterns

This plan owns **Recognizable preview patterns** under [Bloom](../../ROADMAP.md#bloom). The requirement addresses the requested refinement: Use six compact pattern definitions (glider, blinker, toad, beacon, block, pulsar) for genuine simulation previews. Cycle isolated examples every two seconds of preview playback; still menu levels freeze them. Test oscillation periods and glider translation.

## Implementation and checks

Implemented 2026-11-08 using the defaults above. Model checks verify oscillator periods, glider translation and catalog patterns in demonstrations.

Rules/navigation/storage/presentation checks pass with the race detector. Browser evidence: reports/game-refinements-2/checks.json and screenshots. Android and WebAssembly builds pass: reports/build/build-report-20261108T215620Z.json; Linux remains blocked by missing X11 headers. Phone, physical-controller, native screen-reader and balance verification remain pending.
