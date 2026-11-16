# Settings mockup

For the [settings mockup requirement](../../ROADMAP.md), include non-functional override switches and a Games wide selector preview with choices 1–8 and 3 selected. The selector must not change or persist the gallery layout yet; that implementation is explicitly deferred in the settings roadmap.

For [Define the smallest shared experience](../../ROADMAP.md#7-define-the-smallest-shared-experience), address **replace the question-mark help screen with a gear-emoji settings mockup containing clearly labeled, non-functional override switches**. Make the proposed settings visible and reviewable without implying that output controls or saving already work.

## Implementation

Use a bundled gear emoji in place of the question mark. Show switches for disabling menu animations, all audio, haptics, and decorative effects. Switches may change their displayed mock state only; explicitly label them as a preview with no functional or saved effect. Keep optional controls/attribution text below the settings. Back closes settings before leaving a preview. Actual behavior belongs to the future settings-overrides deliverable and persistence to the post-five-games milestone.

## Validation

Build Android/WebAssembly, verify shared Back state tests, and deploy via the canonical phone command. Physical visual/accessibility checks are separate.

## Results

Settings/preview/home Back state tests pass. Android/WebAssembly builds and phone installation/launch succeed (`.tools/settings-mockup-build.log`). Physical appearance verification remains pending because the test phone locked; no claim is made that the mock switches control output or persist.

The Games wide mockup was checked on the test phone with all eight choices visible and 3 selected. It has no layout or persistence handler. Evidence: reports/mosaic/settings.png.

## Closure — 2026-11-16

Archived for the completed implementation scope documented above. This is a documentation lifecycle decision, not new runtime or device verification. Historical evidence and limitations remain intact. Broader outstanding acceptance is retained in the [shared game completion criteria](../../ROADMAP.md#10-define-when-a-game-is-complete), [release-target plan](../current/release-target.md), and [presentation/accessibility plan](../current/presentation-rules.md); game-specific pending acceptance remains in the corresponding deliverable plan where applicable.
