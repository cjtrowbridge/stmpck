# Gallery preview methodology

Gallery cards play bundled PNG images captured from the actual game rules and renderer. Gameplay and automated demonstration inputs run only in the development capture tool. Ordinary browsing must never initialize or tick demonstration game states, search moves, or write player saves/counts.

## Required game-update and deployment gate

This methodology is mandatory for every new game, every update to an existing game, and every deployment containing games, including test-device installs and eventual store releases. Preview maintenance is part of completing the game change, not a later polish task.

Before marking a game update complete or deploying its artifact:

1. Review whether the game's rules, rendering, demonstration inputs or shared capture sources changed. Regenerate with `python scripts/build.py --generate-previews` whenever they did or the freshness check fails. For changes outside capture inputs, verify and record that the existing previews still accurately represent the game; unchanged images do not need gratuitous regeneration.
2. Follow the visual and behavioral review below for affected games: representative action, still/wireframe/vivid variants, looping, intensity behavior and gallery navigation. Check in generated images and their manifest with the source change.
3. Pass the preview asset tests and canonical build validation for the revision being deployed. Record the commands/results and preview review in the mapped game/update plan and journal, including any device-specific checks still pending. An unrelated platform failure does not establish or invalidate success for the intended deployment target.
4. Deploy only the artifact built from that validated revision. Missing, stale, corrupt, misleading or unreviewed previews block completion and deployment until corrected. Do not bypass freshness checks, fall back to live menu simulations, or substitute an older APK to get past a failure.

Ordinary builds automatically check source freshness and manifest integrity. Visual accuracy, input behavior and the correspondence between the reviewed revision and deployed artifact still require the workflow above; the manifest alone does not prove them.

## Adding or changing a game

1. Keep its persistent identifier in `internal/play.IDs` stable. Add deterministic demonstration inputs to `internal/play/demos.go`, `additions_demos.go` or `expansion_demos.go`, reusing the game's real rules. The sequence should clearly show the main interaction and, when practical, an outcome. Never use the player's stored session.
2. Implement its actual rendering in the existing `internal/app/render*.go` paths. Capture calls that same renderer with `active=false`. Keep the board recognizable at thumbnail size and preserve its proportions.
3. Run `python scripts/build.py --generate-previews`. The canonical driver builds a development-only WebAssembly capture program, runs it in local headless Firefox, and writes checked-in images plus `internal/preview/assets/manifest.json`. Firefox is a development prerequisite for regeneration; the distributed app and ordinary builds do not require it. There are no third-party Python package dependencies.
4. Inspect the wireframe/default/vivid posters and animated loops. Check representative action, tiny details, clipping, aspect ratio, palette differences and the loop boundary. Verify still/default switching, scrolling, tab filtering, settings and game entry/return in the actual gallery. Generation alone cannot establish visual quality.
5. Run the `internal/preview` asset tests with existing rule/navigation/storage/presentation tests. Run `python scripts/build.py --verify`. Ordinary builds reject missing, stale or altered assets; regenerate after changing rule/render sources or capture settings. Commit the image files and manifest with their source changes.
6. Record compressed asset size, decoded cache limits, cold loading, warm scrolling and frame-time observations on representative devices. Compare against the previous version under the same conditions. Keep phone, accessibility and battery claims separate from browser/model evidence.

## File and playback contract

- Stable filenames: `<game-id>-<style>-still.png` and `<game-id>-<style>-sheet.png`. Style 0 is monochrome wireframe, 1 default, 2 vivid. Only styles 1 and 2 have sheets.
- Frame canvas: 96×96 pixels, eight columns per sheet, 12 frames per second, eight seconds per game. Bloom gets twelve seconds so its six familiar patterns all appear. Change the shared contract and generator together when measurements justify different resolution/duration; regenerate every affected asset.
- Seeds and inputs come from `play.NewDemonstrations` and `Collection.Demonstrate`. They are deterministic. A two-second simulated lead-in gives still cards established gameplay. Capture does not open player storage. The manifest fingerprints rule/render/capture sources and records PNG sizes and checksums.
- The two levels left of default are still. Default and higher levels animate. Still default holds the current default frame when cached; monochrome uses its poster. Higher levels reuse vivid sheets with the existing live border-breathing and sparkle effects. Rainbow cycling affects those shared effects; separate full recordings for each intensity are unnecessary.
- Only visible gallery cards advance. Playback suspends in settings, gameplay and when unfocused. Offscreen cards retain their frame until visible again. Preview playback never changes play counts or saves.
- Posters are decoded once. One background decode is allowed at a time; the update loop uploads at most one completed sheet per update. A 64 MiB least-recently-used cache bounds sheet pixels, plus approximately 1.9 MiB of poster pixels and one temporary decoded sheet (up to approximately 5.1 MiB). GPU/driver allocation overhead is additional and must be measured on devices. Loading or failed sheet decoding leaves a usable poster; it never starts a live-game fallback.
- PNGs are bundled for fully offline operation. No downloads, analytics or telemetry are part of this workflow.
- Currently visible sheets are protected from eviction. If a future catalog or very large viewport exceeds the cache budget, excess cards keep their posters until capacity becomes available; they must not repeatedly decode and evict one another.

## Evidence

Implementation validation and measured sizes are recorded in [the mapped plan](../plans/current/sprite-previews.md) and the development journal. Browser timing is diagnostic evidence for that browser/host, not an Android battery or performance guarantee.
