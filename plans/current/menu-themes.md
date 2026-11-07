# Menu sensory presentations

This plan addresses the [foundation deliverable](../../ROADMAP.md#4-define-the-smallest-shared-experience) requirement: **Move through intensity levels in order, one adjacent level per navigation action. Left means less intense; right means more intense. Stop at the least and most intense levels without wrapping. Preserve this behavior as more levels are added.** The intended outcome and work are described below.

## Work and validation

Provide restrained, colorful, and exuberant palettes with horizontal swipes and visible previous/next face controls that move one adjacent level. Derive autoplay from intensity: wireframe and still-default are still, default and higher levels animate. No audio or haptics are introduced. Verify clamped endpoints and that vertical browsing does not change the theme.

## Status

Eight presentations, clickable faces, keyboard/gamepad actions, and horizontal gestures are implemented. Unit tests pass for endpoint clamping and intensity-derived autoplay. Android and WebAssembly builds pass; the updated APK is installed on the Pixel. Earlier phone checks passed for direct face selection (now replaced by adjacent-step controls), both clamped swipe endpoints, and restart defaults. Evidence: `reports/menu-tabs/checks.json`.

## Requested refinement

In-game intensity changes only through explicit more/less button clicks/taps. Swipes and directional keys/controllers operate gameplay only.

Implementation and focused regression checks are active. Previous verification describes the earlier revision.
