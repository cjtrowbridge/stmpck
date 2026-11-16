# Menu Tabs

Implement the corresponding [roadmap requirement](../../ROADMAP.md): Implement ⭐ personal most-played, 👥 descending bundled placeholder popularity, and ✨ only never-played games in descending global popularity order. Scores are explicitly synthetic familiarity estimates, not measured market share or live telemetry. Refresh discovery when returning from play and handle an empty catalog safely.

## Verification

Validate sorting/filtering and empty states, build Android and WebAssembly through the canonical pipeline, and inspect the updated layouts, tile travel, and digit rendering on the phone. Implemented. Sorting, filtering, empty-state, and responsive-scroll tests pass with the race detector. Android and WebAssembly build successfully. Phone screenshots in reports/menu-refinements/ confirm full-width boards, empty discovery, complete Territory digits, and an intermediate Merge tile position. Physical-controller and wider accessibility checks remain release work.

## Placeholder scores and rationale

Snake 100, Merge (2048) 90, Bounce (Pong) 80, Linear (Connect Four) 70, Territory (Risk-inspired) 60. These are developer-assigned ordinal familiarity weights, not percentages, sales totals, or measured comparative market share. Snake's historical bundled-phone reach and 2048's mass web reach motivate the first two positions. Pong's broad historical recognition motivates third; the two board-game inspirations follow as a provisional editorial choice. All values ship locally; no network or telemetry is added.

Context sources: [Nokia's historical Snake reach](https://blogs.windows.com/devices/2009/01/20/history-of-nokia-part-2-snake/), [2048 creator's account](https://medium.com/@gabrielecirulli/2048-success-and-me-7dc664f7a9bd), [The Strong's Pong history](https://www.museumofplay.org/games/pong/). These sources describe different metrics and do not establish this exact cross-game ranking.

The six-game expansion appends Brick 85, Echo 75, Bloom 65, Planets 58, Flow 55, and Caliban 50 as provisional editorial familiarity weights. Existing scores and identifiers stay unchanged. These extra scores are not new market research. Ordering, discovery filtering, mosaic bounds and old-save compatibility are tested with all eleven games. The existing roadmap citation remains valid; implementation coverage expands with the catalog.

The seven-game expansion appends Stack 98, Deduction 92, Wave 82, Drift 78, Crossing 76, Cluster 72, and Shift 68, again as synthetic editorial weights. Earlier scores remain unchanged. Sorting, filtering, mosaic hit targets and save compatibility now cover eighteen games. The original roadmap mapping is unchanged.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
