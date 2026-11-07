# Merge Preview

Implement the corresponding [roadmap requirement](../../ROADMAP.md): **Preview:** Demonstrate actual sliding and merging with quick interpolated movement, using the same tile travel data as gameplay. Still menu intensity levels freeze the preview.

## Verification

Validate sorting/filtering and empty states, build Android and WebAssembly through the canonical pipeline, and inspect the updated layouts, tile travel, and digit rendering on the phone. Implemented. Sorting, filtering, empty-state, and responsive-scroll tests pass with the race detector. Android and WebAssembly build successfully. Phone screenshots in reports/menu-refinements/ confirm full-width boards, empty discovery, complete Territory digits, and an intermediate Merge tile position. Physical-controller and wider accessibility checks remain release work.
