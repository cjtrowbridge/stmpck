# STMPCK Roadmap

Build one small, complete app containing five games: Merge, Snake, Territory, Linear, and Bounce.

The [README's foundational principles and promises](README.md#foundational-principles-and-promises) govern every milestone. [AGENTS.md](AGENTS.md) translates them into development rules. Demoscene first: keep the codebase and the app as simple as possible, and make every milestone runnable.

**Current state:** Merge, Snake, Territory, Linear, and Bounce are playable, with a three-column mosaic of game-driven previews, placeholder global popularity and never-played discovery, eight intensity levels, local session saves, pause/resume, deliberate new-game controls, and Android Back navigation. Android and WebAssembly builds pass; all five were exercised on the Pixel 10 Pro XL. Territory is real-time with immutable fixed-speed orders, one starting territory per player, and neutral territory expansion. Merge supports immediate grid selection and quick tile slides; Linear is computer-only. The gear screen previews 1–8 columns but remains non-functional. Model tests cover rules, movement, opponents, collisions, sorting/filtering, mosaic hit targets, save round trips, and corrupt-save recovery. Physical controller/screen-reader verification, broader accessibility, long playtests, the audio suite, working settings overrides, and store packaging remain open. Native Linux compilation still needs X11 headers.

**Checklist meaning:** checked items have their stated implementation scope completed with recorded checks. This does not imply every device/input/accessibility check has passed; remaining verification has explicit unchecked entries. Requirements that previously combined completed menu work with future gameplay have been split, preserving both scopes. Historical phone checks describe the revision tested, not automatic verification of later changes.

## Foundation

Define these requirements before expanding the collection. Implement the shared experience, presentation, and persistence alongside Merge so each piece serves a real game. Share additional code as concrete needs emerge in later games.

### 1. Purpose and constraints

- [ ] Deliver the first release as one app containing the five games specified below.
- [ ] Keep every milestone runnable through the established build-and-run workflow.
- [ ] Apply offline operation, no ads or surveillance, accessibility, minimal visible text, sensory control, and continuity of existing games throughout development.

Subscriptions, accounts, global ranking services, and additional games remain outside the initial implementation until their requirements are decided. The 👥 tab uses bundled placeholder popularity estimates until a suitable ranking source is selected; ✨ shows only never-played games in that order. Pricing, permanent ownership, cross-device purchases, and guaranteed release cadence remain open; no proposed business model overrides the foundational promises.

### 2. Choose the first release target

Plan: [plans/current/release-target.md](plans/current/release-target.md)

- [ ] Select the first mobile platform and a desktop platform for development.
- [x] Set up ADB, connect and authorize the test phone, deploy the debug APK, and verify the first successful on-device launch. Plan: [plans/past/adb-setup.md](plans/past/adb-setup.md)
- [x] Build a locally signed Android debug APK of the opening menu for the first phone test. Plan: [plans/past/android-debug-build.md](plans/past/android-debug-build.md)
- [x] Make VS Code `launch.json` build, install, and launch the current Android debug app on the connected test phone as the canonical development launch action. Plan: [plans/current/test-phone-launch.md](plans/current/test-phone-launch.md)
- [ ] Verify that the app builds and runs on each selected release target.
- [ ] Complete broader touch, screen sizing, suspension, and native accessibility verification on the selected platform. Initial portrait touch, Back, and restart-save checks pass for all five games; these do not establish accessibility or every-device support.
- [ ] Identify packaging requirements early; complete store submission later, once the release destination, credentials mechanism, and signing policy are selected.

The first release platform is still an open decision. Build feasibility does not establish a tested store release.

### 3. Reduce the inherited boilerplate

- [ ] Review imported files and identify what STMPCK actually needs.
- [ ] Remove unused scaffolding and dependencies.
- [ ] Preserve a straightforward build-and-run workflow.
- [ ] Measure starting executable size and startup time so growth can be tracked against a baseline.

### 4. Define the smallest shared experience

Plan: [plans/current/shared-experience.md](plans/current/shared-experience.md)

- [x] Implement ⭐ personal most-played, 👥 descending bundled placeholder popularity, and ✨ only never-played games in descending global popularity order. Scores are explicitly synthetic familiarity estimates, not measured market share or live telemetry. Refresh discovery when returning from play and handle an empty catalog safely. Plan: [plans/current/menu-tabs.md](plans/current/menu-tabs.md)
- [x] Connect personal ordering to actual gameplay counts when games are playable. Keep counts on-device. Plan: [plans/current/play-counts.md](plans/current/play-counts.md)
- [ ] Select an honest, privacy-compatible source for worldwide popularity and bundle rankings for offline use; keep the clearly labeled bundled placeholder scores until data exists.
- [x] Open each new session at the default animated intensity, autoplay on, and sound off. Plan: [plans/current/menu-defaults.md](plans/current/menu-defaults.md)
- [x] Provide a vertically browsable three-column mosaic of game previews, preserving board proportions within each tile. Use the same grid geometry for rendering, hit testing, scrolling, and focus. Keep names off cards; do not add search. Plan: [plans/current/menu-browsing.md](plans/current/menu-browsing.md)
- [x] Replace procedural illustrations with previews driven by actual game code, and connect selection to playable games with return to the same catalog position. Plan: [plans/current/playable-gallery.md](plans/current/playable-gallery.md)
- [x] Move through intensity levels in order, one adjacent level per navigation action. Left means less intense; right means more intense. Stop at the least and most intense levels without wrapping. Preserve this behavior as more levels are added. During gameplay, intensity changes only by tapping/clicking the more/less buttons; game swipes, keys, and controller directions never change intensity. Plan: [plans/current/menu-themes.md](plans/current/menu-themes.md)
- [x] Provide consistent gameplay pause, resume, restart, and back actions. Preview autoplay follows intensity and has no separate play/pause control. Plan: [plans/current/game-controls.md](plans/current/game-controls.md)
- [x] Implement menu touch/mouse browsing and selection, horizontal intensity navigation, keyboard equivalents, and standard controller directional/confirmation/back mappings. Plan: [plans/current/menu-input.md](plans/current/menu-input.md)
- [x] Provide visible intensity arrows/faces, sorting tabs, settings, back, and focus cues with optional control explanations; no preview play/pause button or center intensity control. Plan: [plans/current/menu-discoverability.md](plans/current/menu-discoverability.md)
- [x] Replace the question-mark help screen with a gear-emoji settings mockup containing clearly labeled, non-functional override switches and a Games wide selector preview (1–8, with 3 selected). Plan: [plans/current/settings-mockup.md](plans/current/settings-mockup.md)
- [x] Handle Android system Back one level at a time: close settings, return from a game to the home menu while preserving catalog position, and stay at home if already there. Share navigation with Escape/controller Back. Plan: [plans/current/android-back.md](plans/current/android-back.md)
- [ ] Verify current touch, keyboard, and physical controller behavior, Android Back button/gesture behavior, and startup/lock/suspend transitions on supported devices. Earlier successful touch checks do not cover all later revisions.
- [x] Extend the same upward Back behavior to playable games when implemented, preserving unfinished game state. Plan: [plans/current/game-back.md](plans/current/game-back.md)

### 5. Establish presentation and accessibility rules

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

### 6. Define local persistence

Plan: [plans/current/local-persistence.md](plans/current/local-persistence.md)

- [x] Remember each game's unfinished session. Sensory-preference persistence is scheduled after the first five games below. Plan: [plans/current/game-sessions.md](plans/current/game-sessions.md)
- [x] Give games stable internal identifiers. Plan: [plans/current/game-identifiers.md](plans/current/game-identifiers.md)
- [x] Save safely when interrupted and recover gracefully from missing or damaged saves. Plan: [plans/current/save-recovery.md](plans/current/save-recovery.md)
- [x] Keep everything needed for ordinary play on the device. Plan: [plans/current/offline-games.md](plans/current/offline-games.md)

### 7. Define when a game is complete

These are recurring acceptance criteria for each game, not claims that the foundation menu already implements gameplay. Apply this completion checklist to each of the five games:

- [ ] Rules work correctly and controls feel good on the target device.
- [ ] The preview demonstrates actual gameplay using the game's own rendering code.
- [ ] Every presentation preserves the same rules and understandable game state.
- [ ] Pause, restart, save, resume, and returning to the catalog work reliably.
- [ ] The game works fully offline and respects all sensory preferences.
- [ ] Accessibility and supported input methods have been checked on the target platform.
- [ ] Adding the game preserves the behavior of previously completed games.
- [ ] Executable size and performance costs have been checked, including startup and resource use on the target device.

### 8. Build the shared audio output suite

Plan: [plans/future/audio-suite.md](plans/future/audio-suite.md)

Audio is planned, not implemented. Build one small, offline audio system shared by the menu and all five games. Sound stays off by default; enabling it is independent of intensity. The requirements below remain unstarted and receive individual plans when detailed planning or implementation begins.

- [ ] **Shared sound vocabulary:** Define recognizable core riffs for related actions across the app: navigation/selection, accepted actions, unavailable or invalid actions, progress, success, and failure. Similar events in different games should sound related, with success and failure clearly distinguishable. Failure feedback should be informative rather than punitive.
- [ ] **Intensity-dependent effects:** Give each core riff a consistent identity across intensity levels. Lower levels use simpler, sparser arrangements; higher levels progressively add rhythmic detail, harmony, timbral richness, and ornamentation. Cover all eight current levels and preserve the progression when new levels are added. Increased intensity must not simply mean increased volume.
- [ ] **Per-game background music:** Give each of the five games its own recognizable, looping musical theme. Preserve its core motif while lower intensity reduces pace and ornamentation and higher intensity adds pace, layers, and musical detail. Music tempo must not change gameplay speed or rules. Choose each theme's tempo range during its audio planning; do not automatically equate music tempo with the border-breathing rate.
- [ ] **Coordinated playback:** Make action riffs and background music work together so feedback remains recognizable. Transition smoothly between intensity arrangements and game themes, without clicks, abrupt volume jumps, restarting the game, or resetting its state. Prevent rapid repeated actions from creating an uncontrolled pileup of sounds.
- [ ] **Player control and lifecycle:** Provide independent controls for sound effects and music, with a clear way to silence all audio. Changing intensity must never unmute either. Stop or suspend audio appropriately when the app loses focus or is interrupted. Audio preferences remain session-only until the already scheduled post-five-games persistence work.
- [ ] **Demoscene implementation:** Prefer compact procedural synthesis, reusable riffs, and shared sequencing over large recordings or a general-purpose audio engine. Keep all sounds and music local, with no downloads, tracking, or runtime network dependency. Bound simultaneous voices and measure CPU, memory, battery, and artifact-size costs.
- [ ] **Verification:** Check every event category and game theme across intensity levels, including rapid input, intensity changes, game switching, mute, interruptions, and resume. Verify intelligibility and comfortable levels on phone speakers and headphones. Keep visual equivalents for meaningful audio cues so every game remains fully playable in silence.

### 9. Implement settings overrides

Plan: [plans/future/settings-overrides.md](plans/future/settings-overrides.md)

The gear screen is initially a mockup: its switches demonstrate the layout but do not change app behavior or save preferences. Implement these controls before treating them as usable sensory overrides. Individual requirements remain unstarted; create their plans when detailed work begins.

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

Plan: [plans/current/merge.md](plans/current/merge.md)

- [x] **Board:** Let the player select a grid size before starting a new game; start the selected board with two tiles. Offer stacked 4×4, 5×5, and 6×6 choices. Selecting a size starts immediately; omit setup Start/Back buttons while retaining system/header Back. Plan: [plans/current/merge-board.md](plans/current/merge-board.md)
- [x] **Rules:** Slide every tile in one direction. Matching tiles combine into the next value; each tile can merge only once per move. Plan: [plans/current/merge-rules.md](plans/current/merge-rules.md)
- [x] **New tiles:** After a move that changes the board, spawn one tile in an empty cell. Spawn 2 with 90% probability and 4 with 10% probability. Plan: [plans/current/merge-new-tiles.md](plans/current/merge-new-tiles.md)
- [x] **Controls:** Swipe or use directional controls. Support touch, mouse interaction, and keyboard/controller equivalents. Plan: [plans/current/merge-controls.md](plans/current/merge-controls.md)
- [x] **Goal:** Reach a tile worth 2,048 on every grid size, then allow continued play. The board ends when no legal moves remain. Plan: [plans/current/merge-goal.md](plans/current/merge-goal.md)
- [x] **Undo:** Undo the previous move, including its randomly spawned tile. Repeating that move produces the same result. Plan: [plans/current/merge-undo.md](plans/current/merge-undo.md)
- [x] **Feedback:** Clearly show movement, merging, and new tiles. Input remains predictable even when decorative effects are disabled. Plan: [plans/current/merge-feedback.md](plans/current/merge-feedback.md)
- [x] **Presentation:** Use distinct continuous exponent colors and show quick 200 ms tile travel on every player move, including low-intensity views. Queue rapid moves so transitions do not jump. Keep numbers/symbols legible. Plan: [plans/current/merge-presentation.md](plans/current/merge-presentation.md)
- [x] **Persistence:** Save the board, score, random state, and undo state. Resume exactly where the player stopped. Plan: [plans/current/merge-persistence.md](plans/current/merge-persistence.md)
- [x] **Preview:** Demonstrate actual sliding and merging with quick interpolated movement, using the same tile travel data as gameplay. Still menu intensity levels freeze the preview. Plan: [plans/current/merge-preview.md](plans/current/merge-preview.md)
- [x] **Completion controls:** Continue when possible, restart deliberately, or return to the collection. Plan: [plans/current/merge-completion-controls.md](plans/current/merge-completion-controls.md)

**Deferred:** difficulty balancing between grid sizes, including size-dependent targets and any relationship between cell count and target exponent. Grid-size selection itself is in the initial scope. No scaling formula has been selected.

### 2. Snake

Plan: [plans/current/snake.md](plans/current/snake.md)

- [x] **Board:** A fixed 16×12 grid with clearly visible boundaries. Plan: [plans/current/snake-board.md](plans/current/snake-board.md)
- [x] **Rules:** The snake moves continuously, eats food, grows, and ends the round when it hits a wall or its own body. Plan: [plans/current/snake-rules.md](plans/current/snake-rules.md)
- [x] **Food placement:** Spawn food only in an unoccupied cell. Filling the board completes the round. Plan: [plans/current/snake-food-placement.md](plans/current/snake-food-placement.md)
- [x] **Controls:** Use screen swipes for touch steering, with keys/controller equivalents and no on-screen directional arrow buttons. Prevent direct reversal into the snake's neck and handle rapid direction changes consistently. Plan: [plans/current/snake-controls.md](plans/current/snake-controls.md)
- [x] **Pace:** Offer 5, 7.5, or 12 moves per second, selected before a round. Keep the chosen speed constant throughout that round. Plan: [plans/current/snake-pace.md](plans/current/snake-pace.md)
- [x] **Starting and resuming:** Wait for deliberate input before movement begins, including after returning from a pause. Plan: [plans/current/snake-starting-and-resuming.md](plans/current/snake-starting-and-resuming.md)
- [x] **Feedback:** Make the head, body, food, movement direction, and collision understandable with minimal decoration. Plan: [plans/current/snake-feedback.md](plans/current/snake-feedback.md)
- [x] **Presentation:** Themes can change the snake and food's appearance while preserving the grid, hitboxes, and movement timing. Plan: [plans/current/snake-presentation.md](plans/current/snake-presentation.md)
- [x] **Persistence:** Save the snake, direction, food, speed, and score. Pause automatically when the app loses focus. Plan: [plans/current/snake-persistence.md](plans/current/snake-persistence.md)
- [x] **Preview:** Show a short sequence of moving, eating, and growing. Plan: [plans/current/snake-preview.md](plans/current/snake-preview.md)
- [x] **Completion controls:** Show the result unobtrusively and provide immediate restart or return to the collection. Plan: [plans/current/snake-completion-controls.md](plans/current/snake-completion-controls.md)

### 3. Territory

Plan: [plans/current/territory.md](plans/current/territory.md)

- [x] **Map:** Offer several live map previews generated from non-grid geography, with irregular coastlines and truly varied territory shapes, sizes, and adjacency. Let the player scramble map geometry independently of starting positions. Selecting a map begins the match; omit setup Start/Back buttons. Plan: [plans/current/territory-map.md](plans/current/territory-map.md)
- [x] **Players:** Choose one to three computer opponents; each player starts with exactly one territory and ten units, with all remaining territories neutral and empty. Show ownership immediately in previews and scramble starting positions independently. Plan: [plans/current/territory-players.md](plans/current/territory-players.md)
- [x] **Territory state:** Show high-contrast ownership dots with smaller live unit counts. Deduct dispatched units from the source immediately; represent units in transit separately. Plan: [plans/current/territory-territory-state.md](plans/current/territory-territory-state.md)
- [x] **Real-time simulation:** Run all players simultaneously in real time, with no turns or end-turn control. Pause simulation on pause, navigation away, or backgrounding. Movement and production rates are independent of intensity. Plan: [plans/current/territory-real-time.md](plans/current/territory-real-time.md)
- [x] **Reinforcement:** Each owned territory produces one unit every three seconds of active simulation; neutral territories do not produce units. Initial rate is a balancing default. Plan: [plans/current/territory-reinforcement.md](plans/current/territory-reinforcement.md)
- [x] **Attack selection:** Select an owned source, then an adjacent territory to issue an immediate order. Allow friendly transfers and movement into neutral or hostile territory. Plan: [plans/current/territory-attack-selection.md](plans/current/territory-attack-selection.md)
- [x] **Committed forces:** Orders send all but one available unit and subtract that force immediately. Routes, destinations, and committed forces cannot be changed or recalled. Move at a fixed map-space speed of 0.12 units per second for every player. Plan: [plans/current/territory-committed-forces.md](plans/current/territory-committed-forces.md)
- [x] **Combat:** Resolve combat only on arrival: each side rolls one die per round, ties favor the defender, and one unit is lost per round until one force is exhausted. Friendly arrivals reinforce; empty neutral territory is occupied. Plan: [plans/current/territory-combat.md](plans/current/territory-combat.md)
- [x] **Conquest:** Surviving arrivals occupy a defeated territory. Units already in transit retain their original owner even if their source changes hands. Plan: [plans/current/territory-conquest.md](plans/current/territory-conquest.md)
- [x] **Further actions:** Allow new legal orders while previous orders travel, including friendly transfers. No end-turn or manual reinforcement action. Plan: [plans/current/territory-further-actions.md](plans/current/territory-further-actions.md)
- [x] **Computer opponent:** Computer players act concurrently using the same production, movement, commitment, and combat rules. Choose orders at a fixed cadence without hidden advantages. Plan: [plans/current/territory-computer-opponent.md](plans/current/territory-computer-opponent.md)
- [x] **Goal:** Capture every territory and defeat any remaining enemy forces in transit. Show human elimination when neither owned territory nor a force in transit remains. Plan: [plans/current/territory-goal.md](plans/current/territory-goal.md)
- [x] **Presentation:** Draw all borders clearly with relationship palettes, high-contrast dots, and small complete unit counts. Render digits without cropped or partially visible glyphs across previews and gameplay; preserve redundant ownership shapes. Plan: [plans/current/territory-presentation.md](plans/current/territory-presentation.md)
- [x] **Persistence:** Save ownership, stationed units, immutable in-transit orders and progress, production/AI clocks, and random state. Preserve existing maps and armies when migrating old turn-based saves; do not advance while closed. Plan: [plans/current/territory-persistence.md](plans/current/territory-persistence.md)
- [x] **Preview:** Show real-time production, committed movement, and conquest using the actual game rules. Plan: [plans/current/territory-preview.md](plans/current/territory-preview.md)

**Initial scope:** twelve territories per map and one to three computer opponents; multiple generated map previews and independent map/start-position scrambling. Cards, continent bonuses, and diplomacy remain later possibilities. New matches use the real-time rules; earlier saved armies and maps migrate in place.

### 4. Linear

Plan: [plans/current/linear.md](plans/current/linear.md)

- [x] **Board:** A fixed seven-column, six-row board. Plan: [plans/current/linear-board.md](plans/current/linear-board.md)
- [x] **Rules:** Players alternate dropping one piece into a column. Pieces settle in the lowest available position. Plan: [plans/current/linear-rules.md](plans/current/linear-rules.md)
- [x] **Legal moves:** Full columns reject placement without consuming a turn. Plan: [plans/current/linear-legal-moves.md](plans/current/linear-legal-moves.md)
- [x] **Win condition:** Four connected pieces horizontally, vertically, or diagonally. A full board without a winner is a draw. Plan: [plans/current/linear-win-condition.md](plans/current/linear-win-condition.md)
- [x] **Modes:** Play against the computer only for now. Move local two-player mode to the future PvP roadmap; existing local two-player saves retain their board and continue against the computer. Plan: [plans/current/linear-modes.md](plans/current/linear-modes.md)
- [x] **Controls:** Tap/click a column, or select a column with directional controls and confirm. Plan: [plans/current/linear-controls.md](plans/current/linear-controls.md)
- [x] **Turn feedback:** Clearly indicate whose turn it is and where a piece will land. Prevent accidental extra moves while the computer acts or a piece is settling. Plan: [plans/current/linear-turn-feedback.md](plans/current/linear-turn-feedback.md)
- [x] **Computer opponent:** Take an immediate winning move, block an immediate loss when possible, and otherwise use a small, understandable position preference. Plan: [plans/current/linear-computer-opponent.md](plans/current/linear-computer-opponent.md)
- [x] **Presentation:** Distinguish the two sides through shape or pattern as well as color. Highlight a winning line without requiring animation. Plan: [plans/current/linear-presentation.md](plans/current/linear-presentation.md)
- [x] **Persistence:** Save the board, mode, side to move, and player-side assignment. Plan: [plans/current/linear-persistence.md](plans/current/linear-persistence.md)
- [x] **Preview:** Demonstrate alternating placements and a completed line. Plan: [plans/current/linear-preview.md](plans/current/linear-preview.md)
- [x] **Completion controls:** Offer replay and return to the collection, with starting sides alternating between matches. Plan: [plans/current/linear-completion-controls.md](plans/current/linear-completion-controls.md)

### 5. Bounce

Plan: [plans/current/bounce.md](plans/current/bounce.md)

- [x] **Arena:** One rectangular playfield with two paddles, a ball, and clear scoring boundaries. Plan: [plans/current/bounce-arena.md](plans/current/bounce-arena.md)
- [x] **Players:** One human against a computer paddle. Plan: [plans/current/bounce-players.md](plans/current/bounce-players.md)
- [x] **Rules:** The ball rebounds from the top and bottom walls and paddles. Missing the ball awards a point to the opponent. Plan: [plans/current/bounce-rules.md](plans/current/bounce-rules.md)
- [x] **Controls:** Move the paddle with relative up/down swipes or drags anywhere on the app screen, including outside the arena, so the finger need not obscure play. Keep keyboard/controller equivalents; a tap on a control still performs that control action. Plan: [plans/current/bounce-controls.md](plans/current/bounce-controls.md)
- [x] **Physics:** Use consistent collision handling that prevents the ball passing through paddles or becoming trapped in nearly vertical rallies. Plan: [plans/current/bounce-physics.md](plans/current/bounce-physics.md)
- [x] **Pace:** Offer slow, medium, and fast settings before a match. Keep gameplay speed independent of decorative animation and effects. Plan: [plans/current/bounce-pace.md](plans/current/bounce-pace.md)
- [x] **Computer opponent:** Give it bounded movement and reaction ability so it is beatable. It uses the same paddle and collision rules. Plan: [plans/current/bounce-computer-opponent.md](plans/current/bounce-computer-opponent.md)
- [x] **Serving:** Wait for deliberate player input before the opening serve and after each point. Plan: [plans/current/bounce-serving.md](plans/current/bounce-serving.md)
- [x] **Scoring:** Use a short fixed match—first to five points. Represent progress with simple markers where practical. Plan: [plans/current/bounce-scoring.md](plans/current/bounce-scoring.md)
- [x] **Feedback:** Clearly communicate contact, direction, points, and match completion. Sound and haptics are optional. Plan: [plans/current/bounce-feedback.md](plans/current/bounce-feedback.md)
- [x] **Presentation:** Trails, particles, and themes must preserve visibility of the ball, paddles, and arena boundaries. Plan: [plans/current/bounce-presentation.md](plans/current/bounce-presentation.md)
- [x] **Persistence:** Save scores, positions, velocity, and speed setting. Resume paused so the player is ready before movement continues. Plan: [plans/current/bounce-persistence.md](plans/current/bounce-persistence.md)
- [x] **Preview:** Show a short rally and a scored point. Plan: [plans/current/bounce-preview.md](plans/current/bounce-preview.md)
- [x] **Completion controls:** Offer an immediate rematch or return to the collection. Plan: [plans/current/bounce-completion-controls.md](plans/current/bounce-completion-controls.md)

## After the first five games

- [ ] Persist menu-animation, master-audio, haptics, and decorative-effect overrides alongside sound-effects, music, and intensity settings (with autoplay derived from intensity unless overridden) across closing and reopening the app. Until this is implemented, each new session starts at the default animated intensity with autoplay on and sound off. Plan: [plans/future/menu-preferences.md](plans/future/menu-preferences.md)

## Future PvP modes

Plan: [plans/future/pvp.md](plans/future/pvp.md)

- [ ] Identify games that benefit from player-versus-player play, beginning with Linear; consider Territory and Bounce where their input and turn structure fit.
- [ ] Design simple offline local play first, with clear player identity, turn ownership, setup, and save/resume rules.
- [ ] Reintroduce Linear’s local two-player mode through the shared PvP design rather than keeping a separate initial-game mode.
- [ ] Evaluate other-device or network play separately; do not introduce accounts, telemetry, or an internet requirement for ordinary play.

These requirements are unstarted and will receive individual plans when activated.
