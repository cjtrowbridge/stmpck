# Menu browsing

Implement the [gallery requirement](../../ROADMAP.md): a vertically browsable three-column mosaic, with previews filling their tiles while retaining board proportions. Rendering, hit testing, focus, and scroll bounds must use the same geometry, including incomplete final rows and empty discovery.

The live mosaic stays fixed at three columns. The separately mapped [settings mockup](settings-mockup.md) previews a future layout preference; applying and persisting it remains unstarted roadmap work.

Validate grid selection, row-based scrolling, empty cells, Android/WebAssembly builds, and the phone layout.

Verified: grid hit targets (including gutters and the empty last cell), row-based scrolling, and empty discovery pass navigation tests with the race detector. Android/WebAssembly builds pass; the mosaic and settings mockup were visually checked on the authorized phone. Screenshots: reports/mosaic/.
