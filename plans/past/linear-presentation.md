# Linear — Presentation

For [Linear](../../ROADMAP.md#4-linear), implement the roadmap requirement: **Presentation:** Distinguish the two sides through shape or pattern as well as color. Highlight a winning line without requiring animation.

## Work and validation

Implement this requirement in the Linear state model and shared app integration as appropriate. Keep presentation independent of game rules, preserve state across navigation, and verify relevant behavior with deterministic model tests and the canonical Android/WebAssembly build. Physical input and accessibility validation are recorded separately, not inferred from compilation.

## Status

Implemented; coordinate through [linear.md](linear.md); preserve this one-to-one roadmap link.

## Implementation evidence

Implemented in `internal/play/linear.go`: 7×6 board, gravity, full-column rejection, four-direction wins, draws, computer or local two-player mode, landing preview, alternating starts, and immediate-win/block/center-preference AI. Tests cover legal drops and horizontal/vertical/diagonal wins plus opponent priority. Phone human input and computer response pass.

This requirement is implemented within the playable game. Relevant pure-Go model/storage tests and Android/WebAssembly builds pass. Phone smoke checks cover the main touch path, Back, and saved sessions, not every possible rule outcome or assistive technology. Shared release verification remains open.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
