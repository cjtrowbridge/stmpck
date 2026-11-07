# Smallest shared experience

This plan addresses the [Smallest shared experience](../../ROADMAP.md#4-define-the-smallest-shared-experience) as a whole. The intended outcome and work are described below.

## Work and validation

Build the opening gallery incrementally. Coordinate menu browsing, theme selection, input, and discoverability plans. Cards now open playable games, with a setup screen for a new session and the existing state for a resumed session. Verify compilation and input behavior; keep broader accessibility and device verification explicit.

## Status

Opening gallery implemented with five previews, selection/back navigation, themes, intensity-driven autoplay, a gear settings mockup, emoji sorting tabs, and new-session defaults. Sensory-preference persistence is deferred until after the five games. WebAssembly compiles and was visually checked in Firefox at 480×800. Linux build is blocked by missing X11 headers. Gameplay integration and the five-game phone smoke path pass; broader device/controller checks and screen-reader support remain pending.

## Roadmap reconciliation — 2026-11-07

Completed menu UI/mapping requirements are separated from actual gameplay, real ranking data, assistive technology, and physical input verification. Requirement plans retain one-to-one mappings. Android system Back is owned by [android-back.md](android-back.md). The parent deliverable remains in progress until the unchecked shared-experience scope is complete.

Working settings overrides belong to the new future settings deliverable. The mockup only exposes proposed controls.

## Current checkpoint — 2026-11-07

The gallery is a three-column mosaic. Global ordering uses labeled, bundled placeholder familiarity scores; discovery contains only games with zero plays in that order. The gear screen includes a non-functional 1–8-column selector, with application and persistence deferred. All five games have saves and previews; Territory now runs in real time, Merge has quick tile interpolation, and Linear is computer-only. See the [journal](../../journal/2026-11-07.md) for scoped verification and remaining work.
