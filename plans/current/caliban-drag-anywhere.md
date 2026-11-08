# Caliban: Drag anywhere

This plan owns **Drag anywhere** under [Caliban](../../ROADMAP.md#caliban). The requirement addresses the requested refinement: Relative horizontal dragging anywhere on the screen moves clockwise/right or counterclockwise/left from the lane at gesture start. One screen width spans twelve lanes. Keep tap-to-select in the arena and keyboard/controller alternatives; dragging never changes intensity.

## Implementation and checks

Implemented 2026-11-08 using the defaults above. Browser checks drag below the arena in both directions and verify lane wrapping.

Rules/navigation/storage/presentation checks pass with the race detector. Browser evidence: reports/game-refinements-2/checks.json and screenshots. Android and WebAssembly builds pass: reports/build/build-report-20261108T215620Z.json; Linux remains blocked by missing X11 headers. Phone, physical-controller, native screen-reader and balance verification remain pending.
