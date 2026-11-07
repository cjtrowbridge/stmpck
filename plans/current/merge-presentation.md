# Merge Presentation

Implement the corresponding [roadmap requirement](../../ROADMAP.md): **Presentation:** Use distinct continuous exponent colors and show quick 200 ms tile travel on every player move, including low-intensity views. Queue rapid moves so transitions do not jump. Keep numbers/symbols legible.

## Verification

Validate sorting/filtering and empty states, build Android and WebAssembly through the canonical pipeline, and inspect the updated layouts, tile travel, and digit rendering on the phone. Implemented. Sorting, filtering, empty-state, and responsive-scroll tests pass with the race detector. Android and WebAssembly build successfully. Phone screenshots in reports/menu-refinements/ confirm full-width boards, empty discovery, complete Territory digits, and an intermediate Merge tile position. Physical-controller and wider accessibility checks remain release work.
