# Planets: Opponent reaction progression

This plan owns **Opponent reaction progression** under [Planets](../../ROADMAP.md#planets). The requirement addresses the requested refinement: New matches begin with a six-second AI decision interval, decreasing linearly toward two seconds over three minutes of active simulation after the first order. Schedule each interval at the previous decision; pause/closed time does not count. Persist elapsed time, interval and clock. Legacy matches keep their existing two-second cadence through migrated ramp-complete timing.

## Implementation and checks

Implemented 2026-11-08 using the defaults above. Model checks cover the six-to-two-second curve, first decision boundary, timing validation and deterministic save continuation. Browser checks verify initial delay, faster subsequent intervals and pause/reload continuity. Legacy matches retain their prior cadence.

Rules/navigation/storage/presentation checks pass with the race detector. Browser evidence: reports/game-refinements-2/checks.json and screenshots. Android and WebAssembly builds pass: reports/build/build-report-20261108T215620Z.json; Linux remains blocked by missing X11 headers. Phone, physical-controller, native screen-reader and balance verification remain pending.
