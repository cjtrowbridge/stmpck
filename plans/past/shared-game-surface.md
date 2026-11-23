# Shared game surface

This plan owns the shared game surface requirement in [Shared experience](../../ROADMAP.md#7-define-the-smallest-shared-experience). Move the shipped Apparat integration into STMPCK so both consumers build the same game and save logic. Keep standalone update controls and storage, expose the embedded viewport and catalog, and verify previews, standalone builds and host acceptance.

Verified the Linux amd64, WebAssembly and signed Android arm64 builds and the Go race suite. All 90 regenerated previews matched their previous image hashes. Firefox checked gallery modes, scrolling, Settings, setup before save, start and return. The standalone phone build was verified at packaging; no standalone phone install was performed.
