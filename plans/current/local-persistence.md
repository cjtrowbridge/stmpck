# Local persistence

This plan addresses the [Local persistence](../../ROADMAP.md#9-define-local-persistence) as a whole. The intended outcome and work are described below.

## Work and validation

Coordinate unfinished game saves, stable identifiers, and safe recovery. Sensory-preference persistence is explicitly deferred until after the five games and has its own future plan.

## Status

Versioned game-session persistence is implemented for all five games. Decode validation, native atomic replacement with a validated backup, and browser local storage are in place. Round-trip/corruption tests and phone process-restart persistence checks pass. Future schema migrations and broader lifecycle validation remain pending. Earlier menu preference loading/saving is disabled by user direction; new sessions use the default animated intensity, autoplay on, and sound off.
