# STMPCK Roadmap

Build one small, complete app, starting with Merge, Snake, Territory, Linear, and Bounce and expanding with the selected additions below.

The [README's foundational principles and promises](README.md#foundational-principles-and-promises) govern every milestone. [AGENTS.md](AGENTS.md) translates them into development rules. Demoscene first: keep the codebase and the app as simple as possible, and make every milestone runnable.

**Current state:** Merge, Snake, Territory, Linear, and Bounce are playable, with a three-column mosaic of bundled gameplay sprite previews, placeholder global popularity and never-played discovery, eight intensity levels, local session saves, pause/resume, deliberate new-game controls, and Android Back navigation. Android and WebAssembly builds pass; all five were exercised on the Pixel 10 Pro XL. Territory is real-time with immutable fixed-speed orders, one starting territory per player, and neutral territory expansion. Territory accepts swipe/drag orders on release as well as tap selection. Merge supports immediate grid selection and quick tile slides with swipe/drag or keyboard/controller input; its on-screen directional arrows are removed. Linear is computer-only. The gear screen previews 1–8 columns but remains non-functional. Model tests cover rules, movement, opponents, collisions, sorting/filtering, mosaic hit targets, save round trips, and corrupt-save recovery. Physical controller/screen-reader verification, broader accessibility, long playtests, the audio suite, working settings overrides, and store packaging remain open. Native Linux compilation still needs X11 headers.

**Additional games:** Flow, Caliban, Bloom, Echo, Brick, and Planets are implemented, followed by Drift, Stack, Deduction, Bubble Volley, Wave, Crossing, and Shift, bringing the catalog to eighteen. Rule/save/race tests and Android/WebAssembly builds pass. Browser checks exercise all additions and persistence across reload. Phone verification, initial balancing and wider input/device checks remain open. The Planets start/reaction, Caliban drag-anywhere and Bloom pattern refinements are implemented and browser-checked; phone verification remains pending.

**Checklist meaning:** checked items have their stated implementation scope completed with recorded checks. This does not imply every device/input/accessibility check has passed; remaining verification has explicit unchecked entries. Requirements that previously combined completed menu work with future gameplay have been split, preserving both scopes. Historical phone checks describe the revision tested, not automatic verification of later changes.

**Plan lifecycle:** create plans only for entries actually being planned or executed; unstarted roadmap entries need no plan. Archive each plan when its own scope is complete, preserving evidence and updating references. Previously active work may move to `plans/future/` when explicitly deferred. See [the planning rules](AGENTS.md#roadmap-and-plan-correspondence).

## Foundation

Define these requirements before expanding the collection. Implement the shared experience, presentation, and persistence alongside Merge so each piece serves a real game. Share additional code as concrete needs emerge in later games.

### 1. Purpose and constraints

- [ ] Deliver the first release as one app containing the five games specified below.
- [ ] Keep every milestone runnable through the established build-and-run workflow.
- [ ] Apply offline operation, no ads or surveillance, accessibility, minimal visible text, sensory control, and continuity of existing games throughout development.

Performance telemetry, player profiles, cross-device synchronization, and subscription validation are now requested foundation work, recorded in sections 4–6 below. They are not implemented. This direction requires explicitly reconciling the README/AGENTS commitments to no transmitted gameplay history, device-only preferences/counts, and no periodic online authorization. Pricing, permanent ownership, store purchase/restoration mechanisms, and release cadence remain open. The 👥 tab continues to use bundled placeholder popularity estimates; ✨ shows only never-played games in that order. Performance telemetry is not a source of player-popularity rankings.

### 2. Choose the first release target

Plan: [plans/current/release-target.md](plans/current/release-target.md)

- [ ] Select the first mobile platform and a desktop platform for development.
- [x] **Android development updates:** publish validated APKs to an artifacts-only GitHub repository; provide a Settings update button and one silent startup check using installed-APK hash comparison, same-signer/newer-version validation and user-confirmed Android installation. Ordinary play remains offline. Plan: [plans/past/android-updates.md](plans/past/android-updates.md)
- [x] **Android app icon:** use the root `logo.svg` as the packaged launcher icon. Android packaging and Pixel 4a launcher appearance verified; included in published 0.3-dev. Plan: [plans/past/android-app-icon.md](plans/past/android-app-icon.md)
- [x] Produce a native Windows executable automatically when the canonical build runs on Windows. Verified on Windows/amd64; native interactive and accessibility checks remain separate. Plan: [plans/past/windows-build.md](plans/past/windows-build.md)
- [x] Set up ADB, connect and authorize the test phone, deploy the debug APK, and verify the first successful on-device launch. Plan: [plans/past/adb-setup.md](plans/past/adb-setup.md)
- [x] Build a locally signed Android debug APK of the opening menu for the first phone test. Plan: [plans/past/android-debug-build.md](plans/past/android-debug-build.md)
- [x] **Canonicalize the open-source build toolchain:** require Debian's `openjdk-21-jdk-headless` package and Debian's `adb` package, validated consistently by setup and packaging with an explicit apt remedy; shared selection/mismatch regressions, canonical Android packaging and Pixel 4a launch pass; Debian OpenJDK 21 is unpacked locally through the documented fallback. Plan: [plans/past/open-source-build-toolchain.md](plans/past/open-source-build-toolchain.md)
- [x] Make VS Code `launch.json` build, install, and launch the current Android debug app on the connected test phone as the canonical development launch action. Plan: [plans/current/test-phone-launch.md](plans/current/test-phone-launch.md)
- [ ] Verify that the app builds and runs on each selected release target.
- [ ] Complete broader touch, screen sizing, suspension, and native accessibility verification on the selected platform. Initial portrait touch, Back, and restart-save checks pass for all five games; these do not establish accessibility or every-device support.
- [ ] Identify packaging requirements early; complete store submission later, once the release destination, credentials mechanism, and signing policy are selected.

The first release platform is still an open decision. Build feasibility does not establish a tested store release.

### 3. Reduce the inherited boilerplate

Plan: [plans/future/reduce-boilerplate.md](plans/future/reduce-boilerplate.md)

- [x] **Restore the upstream framework submodule:** Replace the copied framework with a real Git submodule, publish the approved framework changes upstream, preserve unrelated local work and pin a retrievable revision. Plan: [plans/past/framework-submodule.md](plans/past/framework-submodule.md)

- [ ] Review imported files and identify what STMPCK actually needs.
- [ ] Remove unused scaffolding and dependencies.
- [ ] Preserve a straightforward build-and-run workflow.
- [ ] Measure starting executable size and startup time so growth can be tracked against a baseline.

### 4. Simplification and demoscenization

Plan: [plans/current/simplification.md](plans/current/simplification.md)

Follow boilerplate reduction with a continuing cycle of simplifying the app, measuring its real costs, improving quality, and reducing resource consumption. This includes a performance-data submission endpoint and a separate process that analyzes the collected data and turns findings into improvements.

- [x] **Canonize deliverable-driven engineering:** Apply the approved upstream planning, implementation and bounded review/measurement guidance locally, preserving host mappings and runtime boundaries. Host execution rules and plan lifecycle are explicit, and completed scopes are archived with references preserved. Guidance has structural/manual scenario review; full prompt lint awaits its dependency and runtime performance improvement remains separate. Plan: [plans/current/deliverable-driven-engineering.md](plans/current/deliverable-driven-engineering.md)
- [ ] **Streamline the implementation:** Review shared and game-specific code, remove duplication and unnecessary abstractions, simplify control/data flows, and reduce dependency and asset costs. Preserve rules, saves, accessible controls, and sensory choices while improving clarity and maintainability. First active scope implemented: serialized snapshot selection/writes and unchanged Echo/Cluster save suppression pass local race/build/preview checks and matched host benchmarks; Android/phone verification remains pending. Plan: [plans/current/streamline-implementation.md](plans/current/streamline-implementation.md)
- [ ] **Performance baselines and budgets:** Record startup/loading time, frame-time distributions and stalls, memory/allocations, CPU/GPU use, storage/I/O, network cost, artifact size, and battery/energy impact where measurable. Define representative devices, games, preview loads, and intensity levels so comparisons are meaningful; identify metrics unavailable on particular platforms rather than fabricating equivalents.
- [ ] **Bounded instrumentation:** Collect versioned performance samples at a modest, configurable cadence with explicit CPU, memory, storage, and battery budgets. Define the minimum fields and useful aggregate measurements before adding instrumentation; keep profile/gameplay records separate from operational performance data.
- [ ] **Telemetry endpoint:** Set up an API endpoint to accept performance batches, validate their schema and size, control abuse, and store enough build/platform context to compare releases. Select hosting, authentication, retention, and access controls during detailed planning; reuse the profile API's infrastructure where that keeps the system simpler.
- [ ] **Regular submission:** Submit bounded batches on a defined regular interval when connectivity and the platform lifecycle allow. Queue locally while offline, cap queue growth, retry with backoff, and resume overdue submissions on a later eligible launch. Uploads must not block play, startup, saving, or suspension; exact sampling and upload intervals remain to be selected.
- [ ] **Separate analysis process:** Establish a recurring analysis job or workflow distinct from ingestion and gameplay. Aggregate measurements, compare builds and baselines, identify regressions and costly paths, and produce actionable findings with supporting evidence. Define its schedule and how findings become prioritized roadmap requirements and plans.
- [ ] **Iterative quality improvement:** Use the repeated cycle of measure → analyze → prioritize → simplify/optimize → verify → compare. Record the measured effect of each change on quality and resource consumption, using correctness/accessibility/save regressions alongside performance checks. Include the overhead of telemetry itself in every comparison.
- [ ] **Privacy and product commitments:** Define collection disclosure, player controls/consent, allowed metadata, retention/deletion, and whether operational samples can be associated with a profile. Update README and AGENTS to accurately describe the selected policy before enabling remote collection. Preserve the no-ads/no-spyware commitment; do not quietly turn this endpoint into behavioral analytics.

### 5. Cross-platform SQLite persistence

SQLite is the requested local database for settings and player data on **Android, Linux, Windows, macOS, and iOS**. These are design and verification targets, not claims that all five platforms currently build or ship.

- [ ] **Portable database integration:** Select a small SQLite integration compatible with the Go/app stack and each target's packaging, application-data directories, and lifecycle. Verify actual builds and database operations on all five targets. Decide how the existing WebAssembly build participates without assuming native SQLite support in the browser.
- [ ] **Settings and player schema:** Store persistent sensory and menu settings, local/profile identity references, games played, per-game scores/results, and the player statistics needed by the app. Keep local settings functional offline and retain stable game identifiers. Coordinate with the existing settings-override and sensory-persistence milestones rather than creating competing implementations.
- [ ] **Player statistics and central tendencies:** Define per-game counts and score summaries, including appropriate means, medians, and modes where meaningful. Specify how different rule versions, variants, grid sizes, and difficulty levels are separated, and how abandoned/restarted games affect counts. Treat these as player-owned records; do not infer diagnoses or combine incomparable scores.
- [ ] **Migration and recovery:** Migrate existing local saves and play counts without losing progress. Use versioned schema migrations, transactional writes, backups/recovery, and interruption-safe behavior. Decide whether detailed game snapshots live in SQLite or remain in the current save store, with a clear ownership boundary.
- [ ] **Offline changes and sync metadata:** Record pending profile changes locally so settings and player data can be synchronized after connectivity returns. Define identifiers, revisions, deduplication, and deletion handling so retries or multiple devices do not multiply play counts or lose results.
- [ ] **Storage verification and resource limits:** Test upgrades, failed migrations, low storage, damaged data, concurrent operations, shutdown/backgrounding, and long-term data growth. Measure database size and I/O overhead, and define retention/aggregation without silently deleting information the player expects to keep.

### 6. Player profile API, device synchronization, and subscription sessions

Build the API needed to associate the player's devices with a profile, synchronize their data, and maintain long-lived authenticated sessions with cached subscription status. Local play and recording progress must tolerate several weeks off grid.

- [ ] **Profile and device API:** Define profile creation/sign-in, linking and unlinking devices, and authorized access to the player's settings and records. Implement the API and select its hosting, credentials, and operational ownership during detailed planning. Keep account setup and sync failures from adding unnecessary steps to normal play.
- [ ] **Cross-device synchronization:** Synchronize the selected settings and player data between each device's SQLite database and the profile API. Specify which preferences are shared versus device-specific, how concurrent/offline edits merge, how scores and counts deduplicate, and how deletions propagate. Make synchronization incremental, resumable, and independent of the gameplay update loop.
- [ ] **Long-lived validated sessions:** Define session lifetime, renewal, secure credential storage on every target, logout/device revocation, and account recovery. Distinguish authentication from subscription entitlement so the app does not treat an expired access token or a network failure as proof that a subscription has ended.
- [ ] **Weekly attempts and monthly validation:** Attempt subscription revalidation at least once per week when the app can run and reach the service, with successful validation at least monthly as the requested freshness target. Retry overdue checks when connectivity returns or the app resumes; exact intervals, calendar-month versus fixed-day interpretation, and refresh scheduling must be specified and tested. Platforms cannot guarantee a background check while a device is offline or the app cannot run.
- [ ] **Off-grid continuity and expiry policy:** Cache the last validated entitlement locally and allow several weeks of offline play without interruption. Explicitly settle what happens beyond a month without successful validation, during service outages, after confirmed expiry/cancellation, and when clocks change. The monthly-validation requirement and the original unlimited-offline promise must be reconciled; do not silently introduce a hard cutoff or claim both guarantees simultaneously.
- [ ] **Subscription verification backend:** Choose the subscription authority and purchase/restoration mechanisms, validate entitlement on the server, and return a verifiable cached result to the device. Define renewal, grace, expiry, cancellation, and revocation behavior without inventing pricing or ownership terms that have not been selected.
- [ ] **End-to-end verification:** Exercise multiple devices on all five targets, offline edits followed by reconnection, duplicate/reordered requests, session renewal and revocation, weekly attempts, month-scale offline periods, and server failures. Verify that local progress survives every failure and that the documented access policy matches actual behavior.

Sections 4–6 record the requested future foundation scope. No telemetry endpoint, SQLite migration, profile service, sync, or subscription gate is implemented by these entries. Create one deliverable plan per activated section and a separate linked plan for each requirement as detailed planning begins. Align the product promises and existing persistence/settings plans with the new direction in that work.

### 7. Define the smallest shared experience

- [x] Share the catalog, setup, and resumable game surface with Apparat. [Implementation plan](plans/past/shared-game-surface.md)

Plan: [plans/current/shared-experience.md](plans/current/shared-experience.md)

- [x] Implement ⭐ personal most-played, 👥 descending bundled placeholder popularity, and ✨ only never-played games in descending global popularity order. Scores are explicitly synthetic familiarity estimates, not measured market share or live telemetry. Refresh discovery when returning from play and handle an empty catalog safely. Plan: [plans/past/menu-tabs.md](plans/past/menu-tabs.md)
- [x] Connect personal ordering to actual gameplay counts when games are playable. Keep counts on-device. Plan: [plans/past/play-counts.md](plans/past/play-counts.md)
- [ ] Select an honest, privacy-compatible source for worldwide popularity and bundle rankings for offline use; keep the clearly labeled bundled placeholder scores until data exists.
- [x] Open each new session at the default animated intensity, autoplay on, and sound off. Plan: [plans/current/menu-defaults.md](plans/current/menu-defaults.md)
- [x] Provide a vertically browsable three-column mosaic of game previews, preserving board proportions within each tile. Use the same grid geometry for rendering, hit testing, scrolling, and focus. Keep names off cards; do not add search. Plan: [plans/past/menu-browsing.md](plans/past/menu-browsing.md)
- [x] Replace procedural illustrations with previews driven by actual game code, and connect selection to playable games with return to the same catalog position. Plan: [plans/past/playable-gallery.md](plans/past/playable-gallery.md)
- [x] **Sprite gallery previews:** Standardize deterministic generation of bundled sprite animations and stills from game code; replace live gallery simulations with bounded, visible-only image playback, preserving intensity choices and navigation. Follow [docs/previews.md](docs/previews.md) for regeneration and future game coverage. Asset freshness, pixel-cache bounds and browser playback checks pass; phone memory/battery/performance checks remain pending. Plan: [plans/past/sprite-previews.md](plans/past/sprite-previews.md)
- [x] Move through intensity levels in order, one adjacent level per navigation action. Left means less intense; right means more intense. Stop at the least and most intense levels without wrapping. Preserve this behavior as more levels are added. During gameplay, intensity changes only by tapping/clicking the more/less buttons; game swipes, keys, and controller directions never change intensity. Plan: [plans/current/menu-themes.md](plans/current/menu-themes.md)
- [x] Provide consistent gameplay pause, resume, restart, and back actions. Preview autoplay follows intensity and has no separate play/pause control. Plan: [plans/past/game-controls.md](plans/past/game-controls.md)
- [x] Implement menu touch/mouse browsing and selection, horizontal intensity navigation, keyboard equivalents, and standard controller directional/confirmation/back mappings. Plan: [plans/past/menu-input.md](plans/past/menu-input.md)
- [x] Provide visible intensity arrows/faces, sorting tabs, settings, back, and focus cues with optional control explanations; no preview play/pause button or center intensity control. Plan: [plans/past/menu-discoverability.md](plans/past/menu-discoverability.md)
- [x] Replace the question-mark help screen with a gear-emoji settings mockup containing clearly labeled, non-functional override switches and a Games wide selector preview (1–8, with 3 selected). Plan: [plans/past/settings-mockup.md](plans/past/settings-mockup.md)
- [x] Handle Android system Back one level at a time: close settings, return from a game to the home menu while preserving catalog position, and stay at home if already there. Share navigation with Escape/controller Back. Plan: [plans/past/android-back.md](plans/past/android-back.md)
- [ ] Verify current touch, keyboard, and physical controller behavior, Android Back button/gesture behavior, and startup/lock/suspend transitions on supported devices. Earlier successful touch checks do not cover all later revisions. Specifically, the latest Territory swipe orders and Merge arrow removal are built and deployed but await live phone verification because the device was locked.
- [x] Extend the same upward Back behavior to playable games when implemented, preserving unfinished game state. Plan: [plans/past/game-back.md](plans/past/game-back.md)

### 8. Establish presentation and accessibility rules

Plan: [plans/current/presentation-rules.md](plans/current/presentation-rules.md)

- [x] Implement the first four ordered presentations: still monochrome wireframe → still default-style view → default animated view → higher-intensity animated view. All levels left of default must not autoplay; the first left step changes only autoplay. Remove the preview play/pause button and motion-toggle shortcuts. Place the two intensity controls on the same row as the centered sorting tabs, separated from them and aligned to the outer left/right edges. Keep two edge-aligned controls: ← 🙂 moves one intensity level lower, and 🤩 → moves one level higher. Tapping either face or its arrow performs the same adjacent step. There is no center control. At either endpoint, further movement in that direction does nothing; adding future levels must preserve adjacent-step navigation. Plan: [plans/current/intensity-labels.md](plans/current/intensity-labels.md)
- [x] Extend the vivid view with palette-colored breathing borders and sparse random sparkles, followed by three levels with progressively more frequent sparkles and stronger rainbow cycling. Sparkle births are 2, 5, 10, then 18 per second; breathing progresses from a resting-style pace to happy-hardcore tempo: 15, 45, 120, then 180 full expansion/contraction cycles per minute (4 seconds, about 1.33 seconds, 0.5 seconds, then about 0.33 seconds per cycle). The first effect level keeps the vivid palette; the next three increasingly blend rainbow accents with 24-, 16-, and 10-second color cycles. Use small, smoothly fading sparkles and stable backgrounds, with at most 29 live sparkle candidates. Preserve earlier levels, bounded adjacent navigation, and intensity-driven autoplay. Plan: [plans/current/breathing-sparkles.md](plans/current/breathing-sparkles.md)
- [ ] Keep sound and haptics independently controllable across presentation changes; preview autoplay follows the intensity level.
- [ ] Define shared visual cues for selection, available actions, success, and failure.
- [ ] Preserve understandable game state without depending solely on color, sound, or animation.
- [ ] Provide accessible labels, optional explanations, and alternatives to gesture-only actions; verify that essential actions are discoverable without hidden gestures or written instructions.
- [ ] Apply reduced-motion preferences to previews and transitions as well as game effects.
- [ ] Validate sustained frame rate, battery/resource cost, readability, and sensory accessibility across all eight intensity levels on the test phone.

The complete left-to-right order is **still wireframe → still default style → animated default → vivid → palette breathing/sparkles → gentle rainbow → richer rainbow → full rainbow**. “Default” is the fixed launch view, not a recalculated midpoint as more levels are added. Sound stays off in the current app. Screen-reader integration, reduced-motion integration, and sustained effect-performance testing remain open.

### 9. Define local persistence

Plan: [plans/past/local-persistence.md](plans/past/local-persistence.md)

- [x] Remember each game's unfinished session. Sensory-preference persistence is scheduled after the first five games below. Plan: [plans/past/game-sessions.md](plans/past/game-sessions.md)
- [x] Give games stable internal identifiers. Plan: [plans/past/game-identifiers.md](plans/past/game-identifiers.md)
- [x] Save safely when interrupted and recover gracefully from missing or damaged saves. Plan: [plans/past/save-recovery.md](plans/past/save-recovery.md)
- [x] Keep everything needed for ordinary play on the device. Plan: [plans/past/offline-games.md](plans/past/offline-games.md)

### 10. Define when a game is complete

**Mandatory for every game addition, update and deployment:** follow the [gallery preview methodology and deployment gate](docs/previews.md#required-game-update-and-deployment-gate), including test-phone installs and eventual store releases. Keep previews accurate to the current game, regenerate changed captures, review affected still/animated intensity variants, pass asset tests and build validation for the intended target, and record evidence in the mapped plan and journal. Include regenerated PNGs and their manifest with the code. Completion and deployment are blocked by missing, stale, corrupt, misleading or unreviewed previews; deploy only the artifact built from the validated revision. This ongoing obligation applies even when a game's original implementation checkbox is already checked.

These are recurring acceptance criteria for each game, including all six additions, not claims of completed release acceptance. Apply this completion checklist to every game:

- [ ] Rules work correctly and controls feel good on the target device.
- [ ] The preview demonstrates actual gameplay captured from the game's own rendering code into bundled sprite animations and stills. The required [game-update and deployment gate](docs/previews.md#required-game-update-and-deployment-gate) passes for this revision: accurate/reviewed captures, current manifest, asset tests, intended-target build validation and recorded evidence. Gallery browsing never runs the game simulation; regenerated assets accompany their source changes.
- [ ] Every presentation preserves the same rules and understandable game state.
- [ ] Pause, restart, save, resume, and returning to the catalog work reliably.
- [ ] The game works fully offline and respects all sensory preferences.
- [ ] Accessibility and supported input methods have been checked on the target platform.
- [ ] Adding the game preserves the behavior of previously completed games.
- [ ] Executable size and performance costs have been checked, including startup and resource use on the target device.

### 11. Build the shared audio output suite

Audio remains roadmap scope, not implemented or actively planned. Build one small, offline audio system shared by the menu and the full game collection. Sound stays off by default; enabling it is independent of intensity. The requirements below remain unstarted and receive individual plans when detailed planning or implementation begins.

- [ ] **Shared sound vocabulary:** Define recognizable core riffs for related actions across the app: navigation/selection, accepted actions, unavailable or invalid actions, progress, success, and failure. Similar events in different games should sound related, with success and failure clearly distinguishable. Failure feedback should be informative rather than punitive.
- [ ] **Intensity-dependent effects:** Give each core riff a consistent identity across intensity levels. Lower levels use simpler, sparser arrangements; higher levels progressively add rhythmic detail, harmony, timbral richness, and ornamentation. Cover all eight current levels and preserve the progression when new levels are added. Increased intensity must not simply mean increased volume.
- [ ] **Per-game background music:** Give each game its own recognizable, looping musical theme, including the six additions. Preserve its core motif while lower intensity reduces pace and ornamentation and higher intensity adds pace, layers, and musical detail. Music tempo must not change gameplay speed or rules. Choose each theme's tempo range during its audio planning; do not automatically equate music tempo with the border-breathing rate.
- [ ] **Coordinated playback:** Make action riffs and background music work together so feedback remains recognizable. Transition smoothly between intensity arrangements and game themes, without clicks, abrupt volume jumps, restarting the game, or resetting its state. Prevent rapid repeated actions from creating an uncontrolled pileup of sounds.
- [ ] **Player control and lifecycle:** Provide independent controls for sound effects and music, with a clear way to silence all audio. Changing intensity must never unmute either. Stop or suspend audio appropriately when the app loses focus or is interrupted. Audio preferences remain session-only until the already scheduled post-five-games persistence work.
- [ ] **Demoscene implementation:** Prefer compact procedural synthesis, reusable riffs, and shared sequencing over large recordings or a general-purpose audio engine. Keep all sounds and music local, with no downloads, tracking, or runtime network dependency. Bound simultaneous voices and measure CPU, memory, battery, and artifact-size costs.
- [ ] **Verification:** Check every event category and game theme across intensity levels, including rapid input, intensity changes, game switching, mute, interruptions, and resume. Verify intelligibility and comfortable levels on phone speakers and headphones. Keep visual equivalents for meaningful audio cues so every game remains fully playable in silence.

### 12. Implement settings overrides

The gear screen is initially a mockup: its switches demonstrate the layout but do not change app behavior or save preferences. Implement these controls before treating them as usable sensory overrides. This deliverable and its individual requirements remain unstarted; create their plans only when detailed work begins.

- [ ] **Games wide (1–8):** Implement the gear-menu layout selector and persist its value locally across app restarts. Reflow previews, hit targets, scrolling, and focus for the selected column count. Default to three columns. The current gear control is a mockup and does not change or save the layout.
- [ ] **Disable menu animations:** Make the setting stop preview playback and decorative animation throughout every menu, regardless of intensity. Without this override, the existing intensity-driven autoplay rules apply. Do not alter gameplay speed or rules.
- [ ] **Disable all audio:** Make one master mute suppress music and effects across every game and menu, regardless of intensity or individual audio controls. No setting or navigation change may silently unmute it.
- [ ] **Disable haptics and reduce decorative effects:** Wire the remaining switches to shared output rules across the app. Reduced decoration removes sparkles, breathing, and rainbow animation without hiding meaningful game state.
- [ ] **Accessible controls and lifecycle:** Provide accessible names, clear on/off state, keyboard/controller operation, and immediate consistent application across navigation, game changes, interruptions, and resume. Remove the mockup notice only once these switches work.
- [ ] **Validate overrides:** Check each switch separately and in combination at every intensity, including changes while muted or still. Confirm controls never affect game rules and that meaningful feedback remains available visually.

Persist these options in the existing **After the first five games** milestone below. Implementation of session-level overrides and their later storage are separate requirements; the mockup satisfies neither.

## First five games

Implement in the order below. All five require offline operation, no ads or telemetry, local saves, minimal visible text, accessible controls, adjustable sensory presentation, and stable rules across updates. Each now has a gameplay preview using its own rules and rendering code. Checked game requirements record implementation plus model/build/device smoke checks; the shared release completion checklist above remains a separate gate, especially for accessibility, physical controllers, and longer playtesting.

Merge includes grid-size selection before starting a new game, with a fixed 2,048 target for every size. Initial choices are 4×4, 5×5, and 6×6, with 4×4 selected initially. Use the documented starting rules: Territory runs in real time with locked movement orders and automatic production; Bounce matches are first to five. Do not turn every rule into a setting; document the chosen defaults when settled.

### 1. Merge

Plan: [plans/past/merge.md](plans/past/merge.md)

- [x] **Board:** Let the player select a grid size before starting a new game; start the selected board with two tiles. Offer stacked 4×4, 5×5, and 6×6 choices. Selecting a size starts immediately; omit setup Start/Back buttons while retaining system/header Back. Plan: [plans/past/merge-board.md](plans/past/merge-board.md)
- [x] **Rules:** Slide every tile in one direction. Matching tiles combine into the next value; each tile can merge only once per move. Plan: [plans/past/merge-rules.md](plans/past/merge-rules.md)
- [x] **New tiles:** After a move that changes the board, spawn one tile in an empty cell. Spawn 2 with 90% probability and 4 with 10% probability. Plan: [plans/past/merge-new-tiles.md](plans/past/merge-new-tiles.md)
- [x] **Controls:** Use touch swipes or mouse drags, with keyboard/controller directional equivalents. Do not show on-screen directional arrow buttons. Plan: [plans/past/merge-controls.md](plans/past/merge-controls.md)
- [x] **Goal:** Reach a tile worth 2,048 on every grid size, then allow continued play. The board ends when no legal moves remain. Plan: [plans/past/merge-goal.md](plans/past/merge-goal.md)
- [x] **Undo:** Undo the previous move, including its randomly spawned tile. Repeating that move produces the same result. Plan: [plans/past/merge-undo.md](plans/past/merge-undo.md)
- [x] **Feedback:** Clearly show movement, merging, and new tiles. Input remains predictable even when decorative effects are disabled. Plan: [plans/past/merge-feedback.md](plans/past/merge-feedback.md)
- [x] **Presentation:** Use distinct continuous exponent colors and show quick 200 ms tile travel on every player move, including low-intensity views. Queue rapid moves so transitions do not jump. Keep numbers/symbols legible. Plan: [plans/past/merge-presentation.md](plans/past/merge-presentation.md)
- [x] **Persistence:** Save the board, score, random state, and undo state. Resume exactly where the player stopped. Plan: [plans/past/merge-persistence.md](plans/past/merge-persistence.md)
- [x] **Preview:** Demonstrate actual sliding and merging with quick interpolated movement, using the same tile travel data as gameplay. Still menu intensity levels freeze the preview. Plan: [plans/past/merge-preview.md](plans/past/merge-preview.md)
- [x] **Completion controls:** Continue when possible, restart deliberately, or return to the collection. Plan: [plans/past/merge-completion-controls.md](plans/past/merge-completion-controls.md)

**Deferred:** difficulty balancing between grid sizes, including size-dependent targets and any relationship between cell count and target exponent. Grid-size selection itself is in the initial scope. No scaling formula has been selected.

### 2. Snake

Plan: [plans/past/snake.md](plans/past/snake.md)

- [x] **Board:** A fixed 16×12 grid with clearly visible boundaries. Plan: [plans/past/snake-board.md](plans/past/snake-board.md)
- [x] **Rules:** The snake moves continuously, eats food, grows, and ends the round when it hits a wall or its own body. Plan: [plans/past/snake-rules.md](plans/past/snake-rules.md)
- [x] **Food placement:** Spawn food only in an unoccupied cell. Filling the board completes the round. Plan: [plans/past/snake-food-placement.md](plans/past/snake-food-placement.md)
- [x] **Controls:** Use screen swipes for touch steering, with keys/controller equivalents and no on-screen directional arrow buttons. Prevent direct reversal into the snake's neck and handle rapid direction changes consistently. Plan: [plans/past/snake-controls.md](plans/past/snake-controls.md)
- [x] **Pace:** Offer 5, 7.5, or 12 moves per second, selected before a round. Keep the chosen speed constant throughout that round. Plan: [plans/past/snake-pace.md](plans/past/snake-pace.md)
- [x] **Starting and resuming:** Wait for deliberate input before movement begins, including after returning from a pause. Plan: [plans/past/snake-starting-and-resuming.md](plans/past/snake-starting-and-resuming.md)
- [x] **Feedback:** Make the head, body, food, movement direction, and collision understandable with minimal decoration. Plan: [plans/past/snake-feedback.md](plans/past/snake-feedback.md)
- [x] **Presentation:** Themes can change the snake and food's appearance while preserving the grid, hitboxes, and movement timing. Plan: [plans/past/snake-presentation.md](plans/past/snake-presentation.md)
- [x] **Persistence:** Save the snake, direction, food, speed, and score. Pause automatically when the app loses focus. Plan: [plans/past/snake-persistence.md](plans/past/snake-persistence.md)
- [x] **Preview:** Show a short sequence of moving, eating, and growing. Plan: [plans/past/snake-preview.md](plans/past/snake-preview.md)
- [x] **Completion controls:** Show the result unobtrusively and provide immediate restart or return to the collection. Plan: [plans/past/snake-completion-controls.md](plans/past/snake-completion-controls.md)

### 3. Territory

Plan: [plans/past/territory.md](plans/past/territory.md)

- [x] **Map:** Offer several live map previews generated from non-grid geography, with irregular coastlines and truly varied territory shapes, sizes, and adjacency. Let the player scramble map geometry independently of starting positions. Selecting a map begins the match; omit setup Start/Back buttons. Plan: [plans/past/territory-map.md](plans/past/territory-map.md)
- [x] **Players:** Choose one to three computer opponents; each player starts with exactly one territory and ten units, with all remaining territories neutral and empty. Show ownership immediately in previews and scramble starting positions independently. Plan: [plans/past/territory-players.md](plans/past/territory-players.md)
- [x] **Territory state:** Show high-contrast ownership dots with smaller live unit counts. Deduct dispatched units from the source immediately; represent units in transit separately. Plan: [plans/past/territory-territory-state.md](plans/past/territory-territory-state.md)
- [x] **Real-time simulation:** Run all players simultaneously in real time, with no turns or end-turn control. Pause simulation on pause, navigation away, or backgrounding. Movement and production rates are independent of intensity. Plan: [plans/past/territory-real-time.md](plans/past/territory-real-time.md)
- [x] **Reinforcement:** Each owned territory produces one unit every three seconds of active simulation; neutral territories do not produce units. Initial rate is a balancing default. Plan: [plans/past/territory-reinforcement.md](plans/past/territory-reinforcement.md)
- [x] **Attack selection:** Swipe or drag from an owned source to an adjacent destination to issue an order on release. Retain source-then-destination taps and keyboard/controller selection. Allow friendly transfers and neutral or hostile destinations; reject invalid endpoints without changing existing orders. Plan: [plans/past/territory-attack-selection.md](plans/past/territory-attack-selection.md)
- [x] **Committed forces:** Orders send all but one available unit and subtract that force immediately. Routes, destinations, and committed forces cannot be changed or recalled. Move at a fixed map-space speed of 0.12 units per second for every player. Plan: [plans/past/territory-committed-forces.md](plans/past/territory-committed-forces.md)
- [x] **Combat:** Resolve combat only on arrival: each side rolls one die per round, ties favor the defender, and one unit is lost per round until one force is exhausted. Friendly arrivals reinforce; empty neutral territory is occupied. Plan: [plans/past/territory-combat.md](plans/past/territory-combat.md)
- [x] **Conquest:** Surviving arrivals occupy a defeated territory. Units already in transit retain their original owner even if their source changes hands. Plan: [plans/past/territory-conquest.md](plans/past/territory-conquest.md)
- [x] **Further actions:** Allow new legal orders while previous orders travel, including friendly transfers. No end-turn or manual reinforcement action. Plan: [plans/past/territory-further-actions.md](plans/past/territory-further-actions.md)
- [x] **Computer opponent:** Computer players act concurrently using the same production, movement, commitment, and combat rules. Choose orders at a fixed cadence without hidden advantages. Plan: [plans/past/territory-computer-opponent.md](plans/past/territory-computer-opponent.md)
- [x] **Goal:** Capture every territory and defeat any remaining enemy forces in transit. Show human elimination when neither owned territory nor a force in transit remains. Plan: [plans/past/territory-goal.md](plans/past/territory-goal.md)
- [x] **Presentation:** Draw all borders clearly with relationship palettes, high-contrast dots, and small complete unit counts. Render digits without cropped or partially visible glyphs across previews and gameplay; preserve redundant ownership shapes. Plan: [plans/past/territory-presentation.md](plans/past/territory-presentation.md)
- [x] **Persistence:** Save ownership, stationed units, immutable in-transit orders and progress, production/AI clocks, and random state. Preserve existing maps and armies when migrating old turn-based saves; do not advance while closed. Plan: [plans/past/territory-persistence.md](plans/past/territory-persistence.md)
- [x] **Preview:** Show real-time production, committed movement, and conquest using the actual game rules. Plan: [plans/past/territory-preview.md](plans/past/territory-preview.md)

**Initial scope:** twelve territories per map and one to three computer opponents; multiple generated map previews and independent map/start-position scrambling. Cards, continent bonuses, and diplomacy remain later possibilities. New matches use the real-time rules; earlier saved armies and maps migrate in place.

### 4. Linear

Plan: [plans/past/linear.md](plans/past/linear.md)

- [x] **Board:** A fixed seven-column, six-row board. Plan: [plans/past/linear-board.md](plans/past/linear-board.md)
- [x] **Rules:** Players alternate dropping one piece into a column. Pieces settle in the lowest available position. Plan: [plans/past/linear-rules.md](plans/past/linear-rules.md)
- [x] **Legal moves:** Full columns reject placement without consuming a turn. Plan: [plans/past/linear-legal-moves.md](plans/past/linear-legal-moves.md)
- [x] **Win condition:** Four connected pieces horizontally, vertically, or diagonally. A full board without a winner is a draw. Plan: [plans/past/linear-win-condition.md](plans/past/linear-win-condition.md)
- [x] **Modes:** Play against the computer only for now. Move local two-player mode to the future PvP roadmap; existing local two-player saves retain their board and continue against the computer. Plan: [plans/past/linear-modes.md](plans/past/linear-modes.md)
- [x] **Controls:** Tap/click a column, or select a column with directional controls and confirm. Plan: [plans/past/linear-controls.md](plans/past/linear-controls.md)
- [x] **Turn feedback:** Clearly indicate whose turn it is and where a piece will land. Prevent accidental extra moves while the computer acts or a piece is settling. Plan: [plans/past/linear-turn-feedback.md](plans/past/linear-turn-feedback.md)
- [x] **Computer opponent:** Take an immediate winning move, block an immediate loss when possible, and otherwise use a small, understandable position preference. Plan: [plans/past/linear-computer-opponent.md](plans/past/linear-computer-opponent.md)
- [x] **Presentation:** Distinguish the two sides through shape or pattern as well as color. Highlight a winning line without requiring animation. Plan: [plans/past/linear-presentation.md](plans/past/linear-presentation.md)
- [x] **Persistence:** Save the board, mode, side to move, and player-side assignment. Plan: [plans/past/linear-persistence.md](plans/past/linear-persistence.md)
- [x] **Preview:** Demonstrate alternating placements and a completed line. Plan: [plans/past/linear-preview.md](plans/past/linear-preview.md)
- [x] **Completion controls:** Offer replay and return to the collection, with starting sides alternating between matches. Plan: [plans/past/linear-completion-controls.md](plans/past/linear-completion-controls.md)

### 5. Bounce

Plan: [plans/past/bounce.md](plans/past/bounce.md)

- [x] **Arena:** One rectangular playfield with two paddles, a ball, and clear scoring boundaries. Plan: [plans/past/bounce-arena.md](plans/past/bounce-arena.md)
- [x] **Players:** One human against a computer paddle. Plan: [plans/past/bounce-players.md](plans/past/bounce-players.md)
- [x] **Rules:** The ball rebounds from the top and bottom walls and paddles. Missing the ball awards a point to the opponent. Plan: [plans/past/bounce-rules.md](plans/past/bounce-rules.md)
- [x] **Controls:** Move the paddle with relative up/down swipes or drags anywhere on the app screen, including outside the arena, so the finger need not obscure play. Keep keyboard/controller equivalents; a tap on a control still performs that control action. Plan: [plans/past/bounce-controls.md](plans/past/bounce-controls.md)
- [x] **Physics:** Use consistent collision handling that prevents the ball passing through paddles or becoming trapped in nearly vertical rallies. Plan: [plans/past/bounce-physics.md](plans/past/bounce-physics.md)
- [x] **Pace:** Offer slow, medium, and fast settings before a match. Keep gameplay speed independent of decorative animation and effects. Plan: [plans/past/bounce-pace.md](plans/past/bounce-pace.md)
- [x] **Computer opponent:** Give it bounded movement and reaction ability so it is beatable. It uses the same paddle and collision rules. Plan: [plans/past/bounce-computer-opponent.md](plans/past/bounce-computer-opponent.md)
- [x] **Serving:** Wait for deliberate player input before the opening serve and after each point. Plan: [plans/past/bounce-serving.md](plans/past/bounce-serving.md)
- [x] **Scoring:** Use a short fixed match—first to five points. Represent progress with simple markers where practical. Plan: [plans/past/bounce-scoring.md](plans/past/bounce-scoring.md)
- [x] **Feedback:** Clearly communicate contact, direction, points, and match completion. Sound and haptics are optional. Plan: [plans/past/bounce-feedback.md](plans/past/bounce-feedback.md)
- [x] **Presentation:** Trails, particles, and themes must preserve visibility of the ball, paddles, and arena boundaries. Plan: [plans/past/bounce-presentation.md](plans/past/bounce-presentation.md)
- [x] **Persistence:** Save scores, positions, velocity, and speed setting. Resume paused so the player is ready before movement continues. Plan: [plans/past/bounce-persistence.md](plans/past/bounce-persistence.md)
- [x] **Preview:** Show a short rally and a scored point. Plan: [plans/past/bounce-preview.md](plans/past/bounce-preview.md)
- [x] **Completion controls:** Offer an immediate rematch or return to the collection. Plan: [plans/past/bounce-completion-controls.md](plans/past/bounce-completion-controls.md)

## After the first five games

- [ ] Persist menu-animation, master-audio, haptics, and decorative-effect overrides alongside sound-effects, music, and intensity settings (with autoplay derived from intensity unless overridden) across closing and reopening the app. Until this is implemented, each new session starts at the default animated intensity with autoplay on and sound off. Plan: [plans/future/menu-preferences.md](plans/future/menu-preferences.md)

## Next game additions

The selected additions **Flow, Caliban, Bloom, Echo, Brick, and Planets** now have initial playable implementations, shared previews, controls, and local saves. Checked entries below mean implemented with deterministic rule/save checks and successful Android/WebAssembly builds. Phone testing of these six games is pending because the device is disconnected; balance, physical-controller, screen-reader, and release acceptance work remain open. Deliverable and requirement plans below record the initial defaults and evidence.

All six inherit offline operation, no ads or telemetry, minimal visible text, accessible input alternatives, intensity-independent rules, silent play, shared navigation, local saves, and previews driven by actual game code. Reuse existing code where it keeps the implementation simple.

### Flow

Plan: [plans/current/flow.md](plans/current/flow.md)

Inspired by rotating-pipe/network puzzles.

- [x] **Connections:** Rotate pieces to connect a complete network, with clear visual feedback for connected and disconnected pieces. Plan: [plans/past/flow-connections.md](plans/past/flow-connections.md).
- [x] **Controls and goal:** Use tap/click rotation with keyboard/controller alternatives and show completion visually. Initial version: a solvable 5×5 spanning tree with closed endpoints. Plan: [plans/past/flow-controls-and-goal.md](plans/past/flow-controls-and-goal.md).
- [x] **Integration:** Save the puzzle and piece orientations, provide deliberate restart and a real gameplay preview, and keep connection cues understandable without color alone. Plan: [plans/past/flow-integration.md](plans/past/flow-integration.md).

### Caliban

Plan: [plans/current/caliban.md](plans/current/caliban.md)

A Tempest-inspired twelve-lane shooter. Tap/drag selects a rim lane; left/right or D-pad moves around it. Firing is automatic, three missed enemies end the round, and Next advances after a cleared wave. Gameplay pace is independent of presentation.

- [x] **Arena and action:** Build a small procedural arena with readable lanes, player movement around the rim, and enemies approaching along the arena. Plan: [plans/past/caliban-arena-and-action.md](plans/past/caliban-arena-and-action.md).
- [x] **Controls and progression:** Define simple touch controls with input alternatives, firing behavior, collisions, and wave completion. Keep gameplay pace independent of sensory intensity. Plan: [plans/past/caliban-controls-and-progression.md](plans/past/caliban-controls-and-progression.md).
- [x] **Integration:** Provide pause, deliberate restart, saved progression, and a real gameplay preview. Preserve clear threats and hitboxes in still and monochrome presentations. Plan: [plans/past/caliban-integration.md](plans/past/caliban-integration.md).
- [x] **Drag anywhere:** Relative horizontal dragging anywhere on the gameplay screen moves around the rim: right moves clockwise, left counterclockwise, and one screen width spans twelve lanes. Arena taps and keyboard/controller alternatives remain available; dragging never changes intensity. Plan: [plans/past/caliban-drag-anywhere.md](plans/past/caliban-drag-anywhere.md).

### Bloom

Plan: [plans/current/bloom.md](plans/current/bloom.md)

A cellular-automaton sandbox based on Conway's Game of Life, with Bloom as the STMPCK name.

- [x] **Simulation:** Implement the familiar live/dead-cell evolution rules in a small deterministic model. Make the simulation's actual rate an explicit gameplay choice independent of decorative intensity. Plan: [plans/past/bloom-simulation.md](plans/past/bloom-simulation.md).
- [x] **Interaction:** Let the player create and erase cells, run/pause the simulation, step one generation, and reset deliberately. Prefer direct interaction over instructions; the initial grid wraps at 24×24. Plan: [plans/past/bloom-interaction.md](plans/past/bloom-interaction.md).
- [x] **Integration:** Save the cell arrangement and simulation state, offer a game-driven preview, and support open-ended exploration without a required win condition or score. Plan: [plans/past/bloom-integration.md](plans/past/bloom-integration.md).
- [x] **Recognizable preview patterns:** Gallery playback cycles Glider, Blinker, Toad, Beacon, Block, and Pulsar every two seconds, evolving each with the actual simulation. Still-preview intensity levels freeze playback. Plan: [plans/past/bloom-preview-patterns.md](plans/past/bloom-preview-patterns.md).
- [x] **Pattern picker and placement:** Patterns, P, or controller right shoulder opens a paused popup with six named visual thumbnails. Choose a pattern and grid position, inspect its placement preview, then Place/confirm or Cancel/Back. Centered stamps wrap around edges and add live cells without erasing existing ones. Keyboard/controller navigation is supported; native screen-reader verification remains pending. Plan: [plans/past/bloom-pattern-picker.md](plans/past/bloom-pattern-picker.md).

### Echo

Plan: [plans/current/echo.md](plans/current/echo.md)

A sequence-memory game inspired by Simon.

- [x] **Sequence:** Present a short sequence of distinguishable visual cues for the player to repeat, extending the sequence after success. Plan: [plans/past/echo-sequence.md](plans/past/echo-sequence.md).
- [x] **Controls and feedback:** Use large tap/click targets with keyboard/controller equivalents. Shape and position cues must make the entire game playable silently; audio can reinforce them when the future audio suite is available. Plan: [plans/past/echo-controls-and-feedback.md](plans/past/echo-controls-and-feedback.md).
- [x] **Integration:** Clearly distinguish demonstration from player input, provide deliberate replay and pause, save the round state, and demonstrate the actual interaction in the gallery. Cues last 0.5 seconds with 0.25-second gaps; failure offers the same sequence again. Plan: [plans/past/echo-integration.md](plans/past/echo-integration.md).

### Brick

Plan: [plans/current/brick.md](plans/current/brick.md)

A Breakout-style paddle-and-ball game, reusing Bounce's physics where useful.

- [x] **Arena and collisions:** Deflect a ball with a paddle to clear brick arrangements. Keep collisions reliable at the chosen speeds and make remaining bricks readable. Plan: [plans/past/brick-arena-and-collisions.md](plans/past/brick-arena-and-collisions.md).
- [x] **Controls and rounds:** Define direct paddle movement, deliberate serving, missed-ball behavior, and board completion. Initial version: three lives, three layouts, four collision substeps per tick, and deliberate serves. Plan: [plans/past/brick-controls-and-rounds.md](plans/past/brick-controls-and-rounds.md).
- [x] **Integration:** Save the remaining bricks and round state, provide pause/restart and real gameplay previews, and preserve visibility of the paddle, ball, and bricks across intensity levels. Plan: [plans/past/brick-integration.md](plans/past/brick-integration.md).

### Planets

Plan: [plans/current/planets.md](plans/current/planets.md)

A planetary conquest game inspired by Risk and the existing real-time Territory game. Planets is a separate addition to the collection.

- [x] **Planetary systems:** Each level is one planetary system. Planets move continuously in their orbits during active play; show their positions, ownership, and stationed units clearly. Plan: [plans/past/planets-planetary-systems.md](plans/past/planets-planetary-systems.md).
- [x] **Unrestricted destinations:** Allow forces to travel from any owned planet to any other planet in the system, including friendly transfers and neutral/enemy destinations. Territory adjacency does not limit movement. Plan: [plans/past/planets-unrestricted-destinations.md](plans/past/planets-unrestricted-destinations.md).
- [x] **Orders and travel:** Use source-to-destination taps and swipe/drag orders. Retain real-time, fixed-rate travel, immediate source deductions, and irreversible issued orders. Lock the destination body's identity while it moves. Fleets pursue the destination's current position at 0.28 normalized board widths per second, faster than orbital motion. Plan: [plans/past/planets-orders-and-travel.md](plans/past/planets-orders-and-travel.md).
- [x] **Conquest:** Define unit production, force selection, arrival combat, computer opponents, and system victory using the simplest rules that fit unrestricted travel. Balance these explicitly rather than assuming Territory's current values transfer unchanged. Plan: [plans/past/planets-conquest.md](plans/past/planets-conquest.md).
- [x] **Moons:** Include moons orbiting some planets. Systems two and three include one and two capturable moons respectively, using the same production/combat rules. Plan: [plans/past/planets-moons.md](plans/past/planets-moons.md).
- [x] **Integration:** Save the system/level, orbital phases, ownership, units, production clocks, and forces in transit. Pause orbital motion and gameplay together on pause/backgrounding; do not advance while closed. Provide game-driven previews and clear non-color ownership cues. Select one of three system previews; victory offers the next system. Plan: [plans/past/planets-integration.md](plans/past/planets-integration.md).
- [x] **Player starts the match:** New matches move orbits but hold production and opponent orders until the first valid player dispatch. Selection or Play alone does not start opponents. Save/resume preserves this state; legacy matches migrate as already started. Plan: [plans/past/planets-player-start.md](plans/past/planets-player-start.md).
- [x] **Opponent reaction progression:** New matches begin with a six-second decision interval, decreasing linearly to two seconds over three minutes of active match time. Each interval is scheduled at the previous decision. Timing persists, ignores paused/closed time, and is independent of intensity. Legacy matches retain their two-second cadence. These initial defaults remain subject to playtesting. Plan: [plans/past/planets-reaction-progression.md](plans/past/planets-reaction-progression.md).

The five refinements above are implemented with rule/save/race tests, browser interaction checks, and Android/WebAssembly builds. Evidence is recorded in their separate requirement plans. Phone, physical-controller, native screen-reader and balance verification remain pending; earlier checked entries describe the initial versions.

## Seven further game additions

The seven further additions are **Drift (provisional name), Stack, Deduction, Bubble Volley, Wave, Crossing, and Shift**. These seven games now have initial playable implementations. Rule/save tests and race checks pass; Android and WebAssembly builds succeed. Browser interaction and reload-persistence checks pass. Phone, physical-controller, accessibility and balance verification remain open. The linked plans record the implemented defaults.

All seven inherit the shared completion criteria: offline and silent play, no ads or telemetry, minimal visible text, accessible input alternatives and non-color cues, intensity-independent rules, local saves, deliberate restart, shared navigation, and previews driven by actual game code.

### Drift — provisional name

Plan: [plans/current/drift.md](plans/current/drift.md)

An Asteroids-inspired space shooter. **Drift** is a proposed replacement for the earlier name Orbit; the final name remains open.

- [x] **Flight and action:** Wrapping space, inertial ship movement, automatic firing, and rocks splitting twice. Drag anywhere to steer/thrust; keyboard/controller steering and a thrust toggle. Plan: [plans/past/drift-flight-and-action.md](plans/past/drift-flight-and-action.md).
- [x] **Rounds and integration:** Three lives and three waves with deliberate Next/restart, saved ship/rocks/shots/RNG, pause without catch-up and genuine model previews. Drift remains a provisional display name. Plan: [plans/past/drift-rounds-and-integration.md](plans/past/drift-rounds-and-integration.md).

### Stack

Plan: [plans/current/stack.md](plans/current/stack.md)

A falling-block puzzle inspired by Tetris.

- [x] **Placement and clearing:** 10×20 board, seven four-cell shapes, shuffled seven-piece bags, clockwise rotation with small horizontal kicks, row clearing and blocked-spawn failure. Plan: [plans/past/stack-placement-and-clearing.md](plans/past/stack-placement-and-clearing.md).
- [x] **Pace and integration:** Explicit slow/medium/fast gravity, deliberate hard drop and a 40-line goal. Persist board, falling/next pieces, bag/RNG and clock. Preview uses actual placements. Plan: [plans/past/stack-pace-and-integration.md](plans/past/stack-pace-and-integration.md).

### Deduction

Plan: [plans/current/deduction.md](plans/current/deduction.md)

A Sudoku game; **Deduction** is the selected STMPCK name.

- [x] **Puzzle rules:** 9×9 Sudoku with 3×3 regions. Shuffle a valid completed grid, then remove clues only while a solver proves uniqueness; aim for 36 clues. Plan: [plans/past/deduction-puzzle-rules.md](plans/past/deduction-puzzle-rules.md).
- [x] **Interaction and integration:** Cell selection and visible digit palette, keyboard digits and controller cycling, erase, one-step undo, immutable clues and conflict/completion feedback. Save clues/entries/solution/undo; actual solving preview. Notes/hints deferred. Plan: [plans/past/deduction-interaction-and-integration.md](plans/past/deduction-interaction-and-integration.md).

<a id="cluster"></a>

### Bubble Volley

Plan: [plans/current/cluster.md](plans/current/cluster.md)

A bubble-shooter-style matching game.

- [x] **Aiming and matching:** Eight-column staggered hex grid, four marked colors, fixed-speed launched pieces with side-wall reflection. Snap to a free neighbor; clear connected matches of at least three and drop pieces disconnected from ceiling. Plan: [plans/past/cluster-aiming-and-matching.md](plans/past/cluster-aiming-and-matching.md).
- [x] **Rounds and integration:** Five starting rows; clear all to win, reaching bottom row loses. Drag to aim and release to fire, directional aim/confirm and Fire button. Save board, shot and queued colors; deliberate replay and real-model preview. Plan: [plans/past/cluster-rounds-and-integration.md](plans/past/cluster-rounds-and-integration.md).

### Wave

Plan: [plans/current/wave.md](plans/current/wave.md)

A Space Invaders-inspired arcade shooter.

- [x] **Movement and action:** Eight columns by four rows of procedural enemies, horizontal formations descending at boundaries, automatic player fire and seeded enemy shots. Drag horizontally anywhere or use keyboard/controller movement. Plan: [plans/past/wave-movement-and-action.md](plans/past/wave-movement-and-action.md).
- [x] **Waves and integration:** Three lives, three waves, deliberate Next/restart and pause; persist enemies/projectiles/clocks/RNG. Difficulty grows with wave, independent of intensity; game-driven gallery. Plan: [plans/past/wave-waves-and-integration.md](plans/past/wave-waves-and-integration.md).

### Crossing

Plan: [plans/current/crossing.md](plans/current/crossing.md)

A Frogger-inspired crossing game; **Crossing** is the selected STMPCK name.

- [x] **Traversal:** Nine-column crossing with three road lanes, five water lanes with moving logs, safe banks and median. Swipe or directional controls hop one cell; ride logs, avoid cars/water and reach three marked destinations. Plan: [plans/past/crossing-traversal.md](plans/past/crossing-traversal.md).
- [x] **Levels and integration:** Three lives and three levels, with pace determined by level. Save positions, lane phases, destinations, lives and clocks; pause hazards together. Deliberate retry/Next and model-driven preview. Plan: [plans/past/crossing-levels-and-integration.md](plans/past/crossing-levels-and-integration.md).

### Shift

Plan: [plans/current/shift.md](plans/current/shift.md)

A sliding-tile puzzle.

- [x] **Sliding and solvability:** 4×4 numbered sliding puzzle; generate each scramble by legal moves from the solved board, avoiding immediate reversals and a solved start. Plan: [plans/past/shift-sliding-and-solvability.md](plans/past/shift-sliding-and-solvability.md).
- [x] **Interaction and integration:** Tap a neighbor of the empty cell or use directional input; visual completion, one-step undo and confirmed new puzzle. Persist arrangement/initial/undo/move count; preview uses legal slides. Plan: [plans/past/shift-interaction-and-integration.md](plans/past/shift-interaction-and-integration.md).

## Future PvP modes

- [ ] Identify games that benefit from player-versus-player play, beginning with Linear; consider Territory and Bounce where their input and turn structure fit.
- [ ] Design simple offline local play first, with clear player identity, turn ownership, setup, and save/resume rules.
- [ ] Reintroduce Linear’s local two-player mode through the shared PvP design rather than keeping a separate initial-game mode.
- [ ] Evaluate other-device or network play separately; do not introduce accounts, telemetry, or an internet requirement for ordinary play.

This deliverable and its requirements are unstarted and receive plans only when actively planned or executed.
