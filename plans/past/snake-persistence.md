# Snake — Persistence

For [Snake](../../ROADMAP.md#2-snake), implement the roadmap requirement: **Persistence:** Save the snake, direction, food, speed, and score. Pause automatically when the app loses focus.

## Work and validation

Implement this requirement in the Snake state model and shared app integration as appropriate. Keep presentation independent of game rules, preserve state across navigation, and verify relevant behavior with deterministic model tests and the canonical Android/WebAssembly build. Physical input and accessibility validation are recorded separately, not inferred from compilation.

## Status

Implemented; coordinate through [snake.md](snake.md); preserve this one-to-one roadmap link.

## Implementation evidence

Implemented in `internal/play/snake.go`: 16×12 board, three speeds, two queued direction changes with reversal prevention, safe food placement, wall/body collision, score, full-board completion, and deliberate start/resume. Tests cover queued turns, reversal, walls, departing-tail movement, and board fill. A phone session and Back-to-paused saving pass.

This requirement is implemented within the playable game. Relevant pure-Go model/storage tests and Android/WebAssembly builds pass. Phone smoke checks cover the main touch path, Back, and saved sessions, not every possible rule outcome or assistive technology. Shared release verification remains open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
