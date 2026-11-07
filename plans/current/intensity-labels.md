# Intensity facial-expression controls

For the [foundation deliverable](../../ROADMAP.md#5-establish-presentation-and-accessibility-rules), address **two edge-aligned facial-expression controls that move one intensity level at a time, without a center control**. Direct endpoint selection was confusing and would skip intermediate levels as more presentations are added.

## Work and validation

Keep ← 🙂 at the left and 🤩 → at the right. Both face and arrow select the adjacent level in their direction, using the same bounded step as swipes and keyboard actions. There is no center control or direct endpoint shortcut. Remove the center face and selection markers. Dim an arrow when its direction has no next level. Bundle emoji graphics locally with attribution. The original four presentations are followed by four breathing/sparkle levels; new levels must preserve this navigation behavior.

## Status

Implemented. Menu state tests and Android/WebAssembly builds pass; the updated APK is installed. Linux compilation still lacks X11 headers. Updated phone interaction verification remains pending because the phone locked during the check. Earlier phone evidence in `reports/menu-tabs/` describes the superseded direct-selection layout. Broader accessibility validation remains pending in the parent deliverable.

## Intensity-driven autoplay revision

The mapped roadmap requirement now orders still monochrome wireframe, still default-style view, default animated view, and higher-intensity animated view. All levels left of default are still. The first left step preserves the default rendering and frozen preview state. Remove the play/pause button and its keyboard/controller shortcuts. Autoplay is derived from intensity, not stored independently. Validate both left steps and return to default.

Revision validation: state tests cover both still levels, both animated levels, and bounded navigation. Android/WebAssembly builds and device install/launch pass (`.tools/intensity-autoplay-build.log`). Physical visual/touch checks of this revision remain pending; older screenshots show the previous controls.

## Requested refinement

Move the existing less/more controls to the sorting-tab row, with clear gaps and outer-edge alignment. Reduce unused header height and update rendering and hit targets together.

Implementation and focused regression checks are active. Previous verification describes the earlier revision.
