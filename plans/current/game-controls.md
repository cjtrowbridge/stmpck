# Game Controls

For the [foundation](../../ROADMAP.md#foundation), implement **Provide consistent gameplay pause, resume, restart, and back actions. Preview autoplay follows intensity and has no separate play/pause control.** This requirement is needed by the first five playable games.

## Work and validation

Implement the shared behavior alongside the pure game models and native/web app integration. Verify deterministic state transitions, save/load round trips, and corruption handling where relevant. Device/accessibility verification remains explicit.

## Status

Implemented. Deterministic model tests, storage recovery tests, Android/WebAssembly builds, and the five-game phone smoke check pass.

`internal/play` owns named, versioned game state and stable game slots; `internal/app` shares live/demo rendering and navigation. `internal/storage` writes native saves with temporary-file replacement and a validated backup, and browser saves through local storage. Only actual game start/resume contributes to local play counts. Gallery ordering stays stable while browsing; switching tabs recomputes the ranking. Back/settings/lifecycle interruptions pause real-time games, and reopening saved games never starts motion automatically. See `reports/five-games/checks.json`. Settings preferences remain deferred.
