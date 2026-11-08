# Planets: Player starts the match

This plan owns **Player starts the match** under [Planets](../../ROADMAP.md#planets). The requirement addresses the requested refinement: New matches move orbits but hold production and opponent orders until the first valid human fleet dispatch. Save Started and versioned timing state. Legacy saves migrate as already-started matches so existing fleets and progress continue.

## Implementation and checks

Implemented 2026-11-08 using the defaults above. Model and browser checks verify waiting orbits, unchanged production, selection without starting, first valid dispatch and saved states. Legacy saves preserve existing fleets.

Rules/navigation/storage/presentation checks pass with the race detector. Browser evidence: reports/game-refinements-2/checks.json and screenshots. Android and WebAssembly builds pass: reports/build/build-report-20261108T215620Z.json; Linux remains blocked by missing X11 headers. Phone, physical-controller, native screen-reader and balance verification remain pending.
