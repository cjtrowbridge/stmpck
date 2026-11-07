# STMPCK

STMPCK (pronounced “stimpack”) is a collection of small games for self-directed cognitive and sensory pleasure. Choose a problem you feel like solving, choose how you want it to look and feel, and play on your own terms.

Designed primarily for neurodivergent people and welcoming to everyone, STMPCK brings a demoscene approach to familiar game mechanics: the simplest possible codebase and app, capable of experiences ranging from monochrome wireframes to extravagant audiovisual delight.

This is a private development repository for an app intended for app-store distribution. It contains five playable games—Merge, Snake, Territory, Linear, and Bounce—with game-driven previews, eight bounded intensity levels, emoji sorting tabs, and local session saves. Selecting a card starts setup or resumes its saved game. The principles below are binding design commitments, not a claim that all release requirements are complete.

## Foundational principles and promises

1. **Demoscene first.** Create rich experiences from very little code and data. Favor small, understandable implementations, procedural visuals and sound, few dependencies, and efficient use of the device. Every added system must earn its complexity. Small code can produce extravagant sensory experiences.

2. **The simplest possible path to play.** Open the app, see the games, choose one, play. Keep navigation shallow and controls consistent. The collection must remain simple to use as it grows.

3. **Show instead of explain.** Gameplay previews are the primary interface. Use movement, shapes, and direct interaction wherever possible, with as little visible text as practical. Keep text available where it helps accessibility, understanding, or informed purchase decisions. There is no search box.

4. **Neurodivergent people are the primary audience.** Design around varied relationships to repetition, attention, predictability, and stimulation. Welcome everyone without requiring a diagnosis or assuming a single preferred experience. Accessibility is part of the initial design.

5. **The player controls the sensation.** Support the range from monochrome wireframes to extravagant audiovisual delight, and remember the player's preferences. Preview autoplay follows intensity: every level left of the default is still. Sound and haptics remain independent of intensity. Keep a still version of the default visual style immediately to its left. Changing presentation preserves the game's rules and state.

6. **Pleasure is enough.** Playing, repeating, exploring, and enjoying a familiar pattern need no productive justification. STMPCK makes no demand that players improve themselves and makes no therapeutic claims.

7. **Fully offline operation.** Installed games, previews, settings, and saves work without an internet connection. Opening and playing the app never depends on a server responding. Everything needed for available play is bundled locally; going offline does not interrupt play or introduce a periodic online check.

8. **No ads, spyware, or surveillance.** No advertising SDKs, behavioral analytics, tracking identifiers, or transmitted gameplay history. Preferences and any counts used to organize the library stay on the device. These protections apply equally to free and paid play.

9. **Respect the player's attention.** No streak pressure, manufactured urgency, energy timers, nagging notifications, or spending prompts during play. Starting, stopping, and returning should be easy. The player chooses what to play and when to leave.

10. **Preserve familiar experiences.** Updates can add games and optional presentations while preserving existing rules, saves, and chosen settings. Bug fixes and compatibility work maintain that continuity. Deliberate changes to a familiar game should be offered as an optional variant rather than silently replacing it.

## Decisions still open

Pricing, subscriptions, permanent ownership, accounts, cross-device purchases, and a guaranteed release cadence are not settled commitments. Any purchase or restoration design must uphold offline play and the privacy promises above. Obtaining an app, an update, or a new purchase is distinct from running the installed app.

Global popularity needs an honest source compatible with no gameplay surveillance. It is not a reason to introduce telemetry. Platform availability will be documented as builds are verified; app-store deployment is not configured yet.

[chat.md](chat.md) records the exploration that led here, including ideas that were not adopted. This README is the canonical product direction, and [AGENTS.md](AGENTS.md) translates it into development rules. [ROADMAP.md](ROADMAP.md) defines the foundation work, shared completion criteria, and requirements for the first five games.

## Local development

The application uses Go 1.26.4 and Ebitengine v2.9.9, pinned in `go.mod`. Build tooling requires Python 3.10+ and the native prerequisites for the target platform.

Build and verify every locally feasible target:

```bash
python scripts/build.py --verify
```

On Unix, `make build` is the no-flag convenience entrypoint. The build reports every target, its intended output, whether it can be built locally, and why. Artifacts go to `releases/{goos}/{goarch}/app/latest[.exe]` and are ignored by Git.

Run a built host artifact:

```bash
python scripts/run_artifact.py
```

The shared menu is in `internal/app/`, with navigation in `internal/menu/`, deterministic rules in `internal/play/`, and local saves in `internal/storage/`. `cmd/app/` hosts desktop/web, and `mobile/` plus `android/` host Android. Native desktop and JS/WASM builds depend on local toolchain support; Android arm64 debug packaging is configured. Build support does not imply a tested store release.

## Opening menu status and controls

- The top tabs are ⭐ personal most-played, 👥 worldwide most-played, and ✨ discovery. Personal and worldwide show the complete collection; discovery shows only games never played, ordered by global popularity, with an empty state after all games have been tried. Personal counts come from starts/resumes, never previews. Bundled placeholder popularity scores are Snake 100, Merge 90, Bounce 80, Linear 70, and Territory 60. These are synthetic familiarity weights, not measured market share or live player counts; the global tab is labeled “Popularity estimate.” The rationale and sources are in [the tab plan](plans/current/menu-tabs.md).
- Scroll or swipe vertically to browse a three-column mosaic; previews retain their proportions within each tile. Tap/click a card to play.
- Intensity navigation always moves one adjacent level at a time: the left-aligned ← 🙂 control steps toward less intensity, and the right-aligned 🤩 → control steps toward more intensity. Tapping either the face or its arrow performs the same step. In the gallery, horizontal swipes/scrolls and directional keyboard/controller actions follow the same order. The intensity controls flank the sorting tabs on the same row, separated and aligned to the left and right edges. During play, intensity changes only by clicking/tapping these controls; swipes and directional controls operate the game. There is no center button. At either endpoint, further movement in that direction does nothing. These controls must retain this behavior as more intensity levels are added.
- The intensity order is **still monochrome wireframe → still default-style view → default animated view → vivid → breathing/sparkles → gentle rainbow → richer rainbow → full rainbow**. The first step left of default changes only autoplay; the next step left changes to wireframe. Every level left of default must remain still as new levels are added. There is no menu-autoplay play/pause button or motion-toggle shortcut; gameplay has its own pause/resume control. New sessions start at the default animated view with sound off.
- The four levels after vivid add breathing borders and small, randomly positioned sparkles. The first keeps the vivid palette; the next three progressively strengthen rainbow cycling. Breathing accelerates from a resting-style pace to happy-hardcore tempo: 15, 45, 120, then 180 cycles per minute, with one full expansion/contraction per beat. Sparkle birth rates are 2, 5, 10, and 18 per second, with smooth fades and stable backgrounds. Earlier views and the launch default stay unchanged.
- Android system Back (button or gesture), Escape, and controller Back go up one level: close settings first, then return from the game to the same home catalog position, saving its state. At home, Back stays at home. The in-app back arrow also returns to the catalog.
- Keyboard: up/down selects a card, left/right changes presentation, Enter opens, H/F1 opens settings, and 1/2/3 or Tab selects the sorting tab.
- Standard gamepad: D-pad, bottom face button to open, right face button to go back, Start for settings, and shoulder buttons to select sorting tabs. Physical controller verification is pending.

The gear screen also shows a **Games wide** mockup with choices **1–8** and **3** selected. The gallery is fixed at three columns for now; applying and persisting this preference is future roadmap work.

Settings, sound, and intensity persistence is deferred until after the first five games. Current builds keep settings only within the running session and ignore earlier sensory-preference files. Game saves are implemented; native screen-reader integration is still pending. There is no audio, haptic output, networking, advertising, or gameplay telemetry in the menu.

Verification: pure-Go rules/navigation/storage tests pass, including the race detector. Android and WebAssembly build successfully; all five games were exercised on the Pixel 10 Pro XL, including Merge move/undo, the earlier Territory reinforcement input, Linear human/computer moves, Back/pause, and persistence across process restart. Evidence is in `reports/five-games/`. The subsequent layout/control revision also passes rules, geometry, color, and save-migration tests with the race detector. It is installed on the test phone; the shared header, stacked Merge setup, and Territory opponent/shuffle previews were visually checked in `reports/game-refinements/`. Territory now uses real-time movement; dedicated tests cover immediate dispatch, immutable orders, arrival combat, neutral starts, complete border coverage, and mid-flight save continuity. Android debug APK size is about 4.79 MB. Linux still needs X11 development headers. Physical controllers, native screen readers, all screen sizes, longer playtests, and sustained performance remain release work.

## Playing the five games

- **Merge:** tap a vertically stacked 4×4, 5×5, or 6×6 choice to start immediately, without setup Start/Back buttons. Tile colors follow a continuous exponent spectrum, with distinct colors per value and contrasting numbers; wireframe uses distinct grays. Swipe, use arrows/D-pad, or tap the visible direction controls. Tiles slide over 200 ms; rapid inputs queue rather than cutting off movement. This feedback works at every gameplay intensity, while still gallery views remain frozen. Reach 2,048 at every size; continue afterward. Undo restores the move and random spawn. Switch between numbers and power-of-two dot counts.
- **Snake:** choose a speed, then give a direction or press Play. Turn with screen swipes or keyboard arrows/D-pad; there are no touch arrow buttons. Food grows the snake; walls and the body end the round.
- **Territory:** choose one to three opponents and tap an organic map preview to start. Each player begins with one territory and ten units; the rest are neutral. Select your source, then an adjacent destination to send all but one unit immediately. The source count drops at dispatch; committed forces travel at a fixed speed and cannot be recalled or redirected. Friendly arrivals reinforce and hostile arrivals resolve dice combat. All sides play simultaneously, with one unit produced per owned territory every three seconds. These initial rates do not depend on intensity. High-contrast ownership dots contain smaller live counts; shapes reinforce ownership colors. Friendly borders are teal, hostile borders bold red, neutral borders gold, and coasts white; wireframe preserves monochrome with different line weights. Pause/Back/backgrounding stop simulation, with no offline catch-up. Older saves keep their maps and armies when converted to these real-time rules.

- **Linear:** play against the computer. Local two-player play is deferred to the future PvP roadmap; existing two-player saves keep their board and continue against the computer. Tap a column, or select it with left/right and confirm. The landing outline shows the next drop. Replays alternate the starting side.
- **Bounce:** choose a speed and press Play to serve. Swipe/drag vertically anywhere on the app screen to move the paddle relatively, use up/down, or hold the D-pad. First to five wins; each point waits for a deliberate serve.

Game controls: **Space / controller Y** pauses or resumes, **R** requests a new game with confirmation, **U** undoes Merge, **S** changes Merge symbols. Touch controls are visible below the board. Back and opening settings pause games; Snake, Bounce, and saved Territory matches require input to resume. The Android foreground view keeps the display awake during play.

Game state and local play counts are saved privately on the device: Android app files, the desktop user config directory under `stmpck`, or browser local storage. Native saves use atomic replacement and a validated previous-save backup. Game progress is separate from sensory settings; those settings still start fresh each launch. No audio output or network access is required.

The ⚙️ settings screen is currently a **mockup**: its switches only preview on/off positions and neither apply nor save preferences. It previews disabling menu animations, all audio, haptics, and decorative effects. Working app-wide overrides are tracked in the [settings roadmap](ROADMAP.md#9-implement-settings-overrides); persistence stays scheduled after the first five games.

The planned [shared audio suite](ROADMAP.md#8-build-the-shared-audio-output-suite) will use recognizable action, success, and failure riffs across the app, with more ornamentation at higher intensity. Each game will have its own music whose pace and arrangement follow intensity. Effects and music remain independently controllable and off by default; implementation is still pending.

## Android device setup and debug deployment

The first phone deployment is complete and recorded in the [foundation roadmap](ROADMAP.md#2-choose-the-first-release-target). The development app is **STMPCK (dev)**, package `dev.cjtrowbridge.stmpck.debug`. Its manifest requests no permissions. This is a debug build for testing, not a store release.

Local prerequisites are under ignored `.tools/`: Go, JDK 21, Android SDK API 35, Build-Tools 35.0.0, NDK 27.2.12479018, and ADB 37.0.1. On a fresh Linux/amd64 host, `python scripts/setup_android.py` installs the pinned SDK/NDK; provide a JDK on PATH or at `.tools/jdk/` and install Android SDK Platform-Tools separately. The setup helper preserves the downloaded tools' license files.

Enable USB debugging, connect and unlock the test phone, approve its authorization prompt, and verify `.tools/platform-tools/adb devices -l` reports `device`. On Ubuntu, missing USB permissions may require `sudo apt-get install android-sdk-platform-tools-common` followed by reconnecting the phone. See the [official ADB setup documentation](https://developer.android.com/tools/adb).

**Canonical development launch:** in VS Code Run and Debug, select **STMPCK: Launch on test phone** and press F5 (or Run Without Debugging). `.vscode/launch.json` builds the current code, installs the debug APK on the authorized test phone, and opens the app. This launches the app; it does not attach a Go debugger. Connect and unlock the phone first.

The equivalent terminal command is `python3 scripts/build.py --verify --launch-android` (`python` on Windows). Launch stops if this build fails to produce the Android APK; failures for other targets are reported without blocking a successful Android deployment. With multiple authorized devices, use the explicit serial option in the manual deployment command below.

For separate build and deployment steps:

```bash
python scripts/build.py --verify
python scripts/deploy_android.py
```

The no-flag build includes Android when prerequisites are available. Its artifact is `releases/android/arm64/app/latest.apk`. The full matrix currently reports failure for Linux because this host lacks X11 development headers, even when Android and WebAssembly succeed; inspect each target's result.

The deployment helper requires one authorized device, or an explicit `--serial` when multiple devices are connected. It preserves app data on reinstall and verifies that the app process/activity launches. Retain `.tools/android-debug.keystore` to sign subsequent updates with the same development key. Store publishing and release signing are not configured.

Emoji labels are bundled locally from Twemoji; see the [asset attribution](internal/app/emoji/ATTRIBUTION.md).

## Imported tooling

This repository was initialized from `cjtrowbridge/ebe-boilerplate` at commit `e741e1c22bb329698fc3235c3f78710e26dcaa10`. Its `agentic-pipelines` dependency is copied as regular files from commit `6b272ea4cb1a5a28213545222a66005a73d900b1`.

`agentic-pipelines/`, `pipeline.yaml`, and `api.sample.yaml` are optional development scaffolding. They are not app features or prerequisites for building, testing, or playing STMPCK. Inference endpoints are used only when the operator explicitly requests that operation.
