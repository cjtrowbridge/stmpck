# Game Sessions

For the [foundation](../../ROADMAP.md#foundation), implement **Remember each game's unfinished session. Sensory-preference persistence is scheduled after the first five games below.** This requirement is needed by the first five playable games.

## Work and validation

Implement the shared behavior alongside the pure game models and native/web app integration. Verify deterministic state transitions, save/load round trips, and corruption handling where relevant. Device/accessibility verification remains explicit.

## Status

Implemented. Deterministic model tests, storage recovery tests, Android/WebAssembly builds, and the five-game phone smoke check pass.

`internal/play` owns named, versioned game state and stable game slots; `internal/app` shares live/demo rendering and navigation. `internal/storage` writes native saves with temporary-file replacement and a validated backup, and browser saves through local storage. Only actual game start/resume contributes to local play counts. Gallery ordering stays stable while browsing; switching tabs recomputes the ranking. Back/settings/lifecycle interruptions pause real-time games, and reopening saved games never starts motion automatically. See `reports/five-games/checks.json`. Settings preferences remain deferred.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
