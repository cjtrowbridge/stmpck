---
title: "Generic Game Terminology"
source: "https://chatgpt.com/c/6aeed6cc-6ec8-83e8-914e-e4de3ebc6939"
model: "gpt-6-thinking"
model_slug: "gpt-6-thinking"
create_time: "2026-11-07T07:44:12.809Z"
update_time: "2026-11-07T09:16:39.344Z"
exported_at: "2026-11-07T17:07:25.785Z"
author: "ChatGPT"
word_count: 39715
line_count: 6119
message_count: 79
---
# Generic Game Terminology

---

## User

I want to create demoscene-style extremely simple, ad-free versions of popular games that are already ubiquitous with countless versions as obnoxious ad-driven slop. Some of the ones I want to do are the concepts like 2048, risk, etc. I'm curious about existing language and terminology for these examples that avoid trademarks, etc.

---

## ChatGPT

I’ll look at the established genre names for games like 2048 and Risk, along with the naming conventions used by open-source clones. I’ll also distinguish terminology that describes the mechanics from names that might create trademark or other IP issues.

---

## ChatGPT

I think there's a useful distinction between **reimplementing a game** and **reimplementing a game mechanic**. What you're describing is closer to the second: extracting the essential rules of familiar games, stripping away the commercial apparatus, and implementing them as extremely small, self-contained programs.

There are already several established terms for this, including **game clones, minimalist games, demakes, microgames, and open-source reimplementations**. But there's also a more specific vocabulary for describing the underlying mechanics without referring to the commercial games by name.

For your project, I'd distinguish three things:

1. **Genre:** What kind of game is this? For example, *territory-control strategy*.
2. **Mechanic:** What actually happens? For example, *adjacency-based conquest with dice combat*.
3. **Implementation philosophy:** What are you making? For example, *a minimalist, ad-free, standalone implementation*.

That separation would let you develop an entire collection without borrowing commercial branding.

## 1. Generic terminology for your first two examples

| Familiar game | Established generic terminology | More precise mechanical description |
|---|---|---|
| **2048** | Sliding-tile puzzle; number-merging puzzle | Directional slide-and-merge puzzle |
| **Risk** | Area-control strategy; territory-control game; conquest wargame | Turn-based, dice-based territorial conquest |
 
### 2048 → Slide-and-merge puzzle

The technical terminology is particularly straightforward.

The Puzzle Genome Atlas categorizes 2048 as a *slide-and-merge* puzzle, identifying directional movement, collision-based merging, and random tile spawning as its defining mechanics. ([Puzzle Genome Atlas](https://puzzlegenome.org/en/games/2048/?utm_source=chatgpt.com))

Other descriptive terms include *number-merging puzzle*, *grid-merging puzzle*, and *sliding-number puzzle*.

For your implementation, I would favor **Merge** or **Slide & Merge** as a simple descriptive title, subject to a trademark search.

There's an interesting historical detail here: 2048 was itself a minimalist, open-source implementation of ideas developed in earlier games. Gabriele Cirulli released it in 2014, based on 1024, which was influenced by Threes. The original 2048 was distributed under the MIT license. ([Cloned Games](https://clonedgames.com/2048?utm_source=chatgpt.com))

In a sense, you'd be continuing the original project's philosophy rather than inventing a new one.

### Risk → Territory-control strategy

Risk belongs to a much older family of tabletop wargames.

Its defining mechanics include controlling territories, deploying reinforcements, attacking adjacent territories, resolving combat probabilistically, and conquering a map. ([The Strong National Museum of Play](https://www.museumofplay.org/toys/risk/?utm_source=chatgpt.com), [BoardGameGeek](https://boardgamegeek.com/boardgame/181/risk?utm_source=chatgpt.com))

The best generic terminology depends on which aspects you're preserving:

| Term | Meaning |
|---|---|
| **Area control** | Players compete to dominate regions of a board. |
| **Area majority** | Players compete to have the greatest presence in regions, without necessarily controlling them exclusively. |
| **Territory control** | Players capture and hold discrete territories. |
| **Territorial conquest** | Players expand by taking territories from opponents. |
| **Dice-based wargame** | Combat outcomes are determined using dice. |
| **Turn-based strategy** | Players take sequential turns to make strategic decisions. |
| **Map-conquest game** | Players progressively conquer a geographic or abstract map. |

For a close mechanical equivalent of Risk, **territorial conquest** is probably the most useful designation.

There's even an existing open-source implementation called *Domination*, which uses its own name while explicitly describing its relationship to Risk. ([AlternativeTo](https://alternativeto.net/software/risk/?utm_source=chatgpt.com))

## 2. Other popular games and their generic names

There are dozens of games that would fit this concept. Many have become so ubiquitous that their mechanics are effectively part of gaming culture, even when their names remain trademarks.

Here is a potential starting catalog.

| Recognizable game | Generic genre or mechanic | Possible minimalist title |
|---|---|---|
| 2048 | Slide-and-merge puzzle | Merge |
| Risk | Territorial conquest / area control | Conquest |
| Tetris | Falling-block puzzle | Stack |
| Minesweeper | Mine-detection / logic deduction | Mines |
| Snake | Grid-based movement / growth | Snake |
| Solitaire | Single-player card game / patience | Patience |
| FreeCell | Open-cell solitaire | FreeCell |
| Sudoku | Number-placement logic puzzle | Sudoku |
| Wordle | Word deduction / feedback puzzle | Five Letters |
| Bejeweled | Match-three puzzle | Match |
| Candy Crush | Match-three / tile-matching puzzle | Match Three |
| Angry Birds | Physics-based projectile puzzle | Trajectory |
| Flappy Bird | One-button obstacle avoidance | Flap |
| Crossy Road / Frogger | Lane-crossing obstacle avoidance | Crossing |
| Connect Four | Four-in-a-row connection game | Four in a Row |
| Battleship | Hidden-fleet deduction game | Fleet |
| Mastermind | Code-breaking deduction game | Codebreaker |
| Yahtzee | Dice-combination / push-your-luck game | Dice |
| Monopoly | Property-trading / economic board game | Landlord |
| Chess | Abstract strategy | Chess |
| Checkers | Draughts | Draughts |
| Reversi / Othello | Disk-flipping abstract strategy | Reversi |
| Peggle | Ball-drop / physics puzzle | Pegboard |
| Bubble Shooter | Projectile-based color matching | Bubble Match |
| Space Invaders | Fixed-screen shooter | Invaders |
| Breakout | Paddle-and-ball brick breaker | Brick Breaker |
| Pac-Man | Maze chase | Maze Chase |
| Agar.io | Multiplayer growth / consumption game | Cells |
| Territorial.io | Real-time territory conquest | Territories |

These are **descriptive suggestions, not trademark-cleared titles**. Even a generic-sounding word can already be registered for games.

It's also worth distinguishing exact reimplementations from games that share a genre. For instance, a basic match-three game is not necessarily a mechanical reproduction of Candy Crush.

## 3. The language of minimalist game development

Your demoscene reference is particularly relevant.

There are several existing development traditions that overlap with what you're describing.

| Term | What it emphasizes | How well it fits |
|---|---|---|
| **Game clone** | Recreating an existing game's gameplay | Excellent |
| **Clean-room reimplementation** | Independently implementing behavior without copying original code | Excellent, if developed that way |
| **Demake** | Recreating a game using simpler or older technology | Very good |
| **Microgame** | Very small, simple game | Good, although often implies short play sessions |
| **Minimalist game** | Removing nonessential presentation and features | Excellent |
| **Sizecoding** | Making programs extremely small | Excellent if executable size matters |
| **Code golf** | Minimizing source-code length | Related, but not necessarily the same goal |
| **One-file game** | Packaging an entire game in a single file | Excellent |
| **Zero-dependency game** | Eliminating external software dependencies | Excellent |

Sizecoding is an especially interesting tradition because the demoscene has established categories for productions as small as 256 bytes, 4 KB, and 64 KB. The focus is on creating impressive software within severe size constraints. ([Nano Gems](https://nanogems.demozoo.org/?utm_source=chatgpt.com))

I think **minimalist game reimplementations** is the clearest description of your project.

But **demakes** might be the more culturally interesting identity, particularly if you intentionally use the graphical and computational constraints of earlier computers.

There's also an opportunity to create a recognizable format: games that have no accounts, no advertising, no tracking, no networking requirements, and no assets loaded from external services.

For web distribution, a single self-contained HTML file would be particularly compelling.

---

## 4. The intellectual property distinction

The United States Copyright Office explicitly states that game ideas and methods of play are not protected by copyright. Original artwork, written material, and other expressive elements may be protected. ([U.S. Copyright Office](https://www.copyright.gov/register/tx-games.html?utm_source=chatgpt.com))

However, there are three distinct issues to consider.

**Copyright:** You can generally implement game mechanics independently, but should avoid copying original artwork, music, code, text, or sufficiently distinctive audiovisual expression.

**Trademark:** You should avoid using another company's brand as your game's title. Trademark analysis depends on likelihood of consumer confusion, and simply changing a word or two doesn't necessarily resolve that risk. ([USPTO](https://www.uspto.gov/trademarks/search/likelihood-confusion?utm_source=chatgpt.com))

**Trade dress and overall presentation:** A game can infringe copyright even when its source code is independently written and its name is different. In *Tetris Holding v. Xio Interactive*, a court found infringement involving the substantial reproduction of Tetris's protected visual expression. ([Willamette Law Online](https://wlo.willamette.edu/ip/2012/05/tetris-holding-llc-v-xio-interactive-inc.html?utm_source=chatgpt.com))

So your instinct to use generic terminology is useful, but adopting independent artwork and visual presentation is equally important.

For 2048, you have an additional advantage: the original browser implementation is available under an MIT license. You can potentially reuse its code under the license terms rather than having to develop everything independently. ([Cloned Games](https://clonedgames.com/2048?utm_source=chatgpt.com))

---

## 5. A possible organizing principle: the mechanics are the games

I think there's a particularly appealing design philosophy hiding in this idea.

Instead of treating *2048*, *Risk*, *Minesweeper*, and *Tetris* as products, you treat them as **instances of more fundamental interactive systems**.

For example:

| Game family | Fundamental system |
|---|---|
| Merge | Spatial compression and combination |
| Conquest | Graph-based territorial competition |
| Mines | Constraint satisfaction and deduction |
| Stack | Geometric packing |
| Match | Pattern recognition and elimination |
| Snake | Path planning with an expanding constraint |
| Patience | Sequencing under incomplete information |
| Fleet | Spatial deduction with hidden information |
| Breaker | Deterministic physics and interception |

This would give the collection a consistent identity based on mechanics rather than intellectual property.

It also opens the possibility of making the rules configurable.

For instance, your **Conquest** implementation could let users choose different maps, reinforcement rules, combat probabilities, and victory conditions without losing the underlying concept.

Similarly, **Merge** could support arbitrary board dimensions, merging rules, and target values.

That gets at something interesting about the demoscene: the demonstration isn't merely the game, but how much interesting behavior can emerge from a remarkably small amount of code.

## 6. My suggested direction

I'd organize the collection around **short, generic, mechanic-based names**.

A unified collection could simply contain:

**Merge · Conquest · Mines · Stack · Match · Snake · Patience · Fleet · Crossing · Breaker**

Each game could have its own tiny executable or standalone HTML file, with the same interface conventions and a consistent minimalist visual language.

There is one important distinction I'd make in your design philosophy:

**Minimalism shouldn't mean removing gameplay complexity. It should mean removing everything that isn't gameplay.**

Risk is an especially good example. The strategic depth comes from the interaction of simple rules, not from having a complicated user interface.

You could preserve the entire strategic system while eliminating the advertising, progression systems, currencies, unlocks, accounts, animations, and other layers that have accumulated around many commercial implementations.

That could make the games feel much closer to the original tabletop, arcade, and early home-computer experiences.

And I particularly like the idea of making each game extremely small, because it provides an objective demonstration of how little code is actually needed to deliver the experience.

**One question that would help narrow the terminology:** Are you primarily interested in recreating the *exact rules* of familiar games under generic names, or in creating *canonical implementations of entire game-mechanic families*?

Those approaches overlap, but the second could become something substantially more interesting than a collection of clones.

---

## User

with my existing ebe boilerplate, give me your top five easiest ones to implement

---

## ChatGPT

Given your existing **EBE (Ebitengine) boilerplate in Go**, I'd prioritize games that require minimal assets, simple input handling, and very little game-state management.

Your existing rendering, navigation, keyboard/controller input, and UI components should eliminate much of the setup work.

## Top five easiest implementations

| Rank | Game | Complexity | Estimated code | Estimated time |
|---|---|---|---|---|
| 🥇 1 | **Snake** | Very low | 100–200 lines | 30–60 minutes |
| 🥈 2 | **2048 (Merge)** | Very low | 150–250 lines | 1–2 hours |
| 🥉 3 | **Pong** | Very low | 100–200 lines | 1–2 hours |
| 4 | **Minesweeper (Mines)** | Low | 200–350 lines | 2–3 hours |
| 5 | **Connect Four** | Low | 150–300 lines | 2–3 hours |

These are rough estimates for playable MVPs built on your existing boilerplate, not tested implementation times.

### 1. Snake — easiest overall

**Why:** Almost no complicated game logic.

You need a grid, a snake represented by a list of coordinates, a direction, randomly placed food, and collision detection.

Ebitengine already provides everything needed for rendering and polling input.

**Particularly interesting:** You can implement the entire game using rectangles without loading a single image asset.

### 2. Merge (2048) — best first puzzle

**Why:** Entirely deterministic turn-based logic, except for spawning new tiles.

You need a 4×4 integer matrix, four directional movement operations, a merging algorithm, and score tracking.

The four movement directions can use the same algorithm by transforming rows and columns.

**Particularly interesting:** The game requires no continuous simulation, collision system, or animation.

### 3. Pong — simplest arcade game

**Why:** Just a few moving rectangles.

You need two paddles, one ball, boundary collisions, paddle collisions, and scoring.

A simple AI opponent can follow the ball's vertical position.

**Particularly interesting:** Unlike Snake or Merge, Pong demonstrates Ebitengine's continuous animation and real-time input handling.

### 4. Mines — easiest deduction game

**Why:** The rules are simple, and the board is generated procedurally.

You need a 2D grid, randomized mine placement, adjacent-mine counts, tile revealing, flood-fill expansion, and flagging.

A basic mouse interface is sufficient, although your controller support could make this more interesting.

**Particularly interesting:** Nearly all its complexity exists in the underlying logic. The UI can be exceptionally simple.

### 5. Four in a Row — easiest multiplayer strategy game

**Why:** The game has a tiny state space compared with Risk or chess.

You need a 7×6 board, alternating turns, gravity-based piece placement, and horizontal/vertical/diagonal win detection.

Local two-player support requires almost no additional complexity.

A basic computer opponent can be added later using minimax.

**Particularly interesting:** This is probably the easiest way to demonstrate a turn-based strategy-game architecture that could eventually support Conquest.

---

## My recommended implementation order

I'd actually build them in this order:

**Merge → Snake → Pong → Mines → Four in a Row**

Merge is a particularly useful first implementation because it establishes a reusable grid-based architecture without introducing real-time simulation.

Your existing EBE patterns also suggest a natural separation:

```text
games/
├── merge/
│   ├── game.go
│   ├── rules.go
│   └── render.go
├── snake/
├── pong/
├── mines/
└── four/
```

Each game could implement a common interface:

```go
type Game interface {
    Update() error
    Draw(screen *ebiten.Image)
    Layout(outsideWidth, outsideHeight int) (int, int)
}
```

That is Ebitengine's native game interface, so you wouldn't need to invent another abstraction.

Your existing named-action input system could provide standardized movement, confirmation, restart, and pause controls.

**I'd start with Merge**, even though Snake is marginally easier. It would give you a working proof of concept for the entire collection, including keyboard/controller interaction, grid rendering, procedural state, and a complete game loop.

All five could share your existing HUD and UI components while keeping their actual game implementations exceptionally small.

One additional architectural recommendation: **keep game rules independent of Ebitengine.** That would make them deterministic, unit-testable, and potentially reusable in future terminal, browser, or multiplayer implementations.

---

## User

i'm curious about the feasibility of making one single app that has a bunch of these games, all offline, all ad-free, all surveillance-free, and having regular updates that add more games without changing old ones

---

## ChatGPT

I think this is an excellent fit for your existing EBE/Ebitengine boilerplate. In fact, **one application containing dozens or hundreds of tiny, self-contained games** could be substantially easier to develop, maintain, and distribute than dozens of independent applications.

The concept would be:

> **One app. Hundreds of games. No ads, accounts, tracking, subscriptions, or internet connection required. New games added regularly. Existing games left alone.**

The important architectural decision is to treat the application as a **game collection with a stable runtime**, rather than one enormous game whose codebase gets progressively more complicated.

That would let you keep adding games without continually modifying the existing ones.

## 1. How feasible is it?

**Very feasible.** The core architecture is straightforward, especially for the kinds of games you're proposing.

| Goal | Feasibility | Main consideration |
|---|---|---|
| One application containing many games | Excellent | Common game interface |
| Completely offline gameplay | Excellent | Bundle everything locally |
| No advertising | Trivial | Don't integrate advertising systems |
| No analytics or surveillance | Excellent | Avoid telemetry and third-party SDKs |
| No user accounts | Trivial | Save locally |
| Regular updates adding games | Excellent | Modular game packages |
| Existing games remain unchanged | Excellent | Immutable game modules and regression testing |
| Extremely small application | Very good | Share the engine, UI, fonts, and assets |
| Desktop, mobile, and browser support | Good | Packaging and platform-specific integration |
| Updates without internet access | Excellent | Support importing offline update packages |

An additional advantage is that games like Merge, Snake, Pong, and Mines don't need complex graphics or large asset libraries.

You could potentially distribute a collection of dozens of games in an executable that's smaller than the advertising and analytics dependencies in a single modern mobile game.

The exact size would depend on your Ebitengine build, assets, and target platform.

## 2. The architecture I'd use

I'd build a single Ebitengine application that serves as a launcher, with each game implemented as an independent Go package.

Something like:

```text
games/
├── main.go
├── engine/
│   ├── app.go
│   ├── input.go
│   ├── audio.go
│   ├── storage.go
│   └── registry.go
├── ui/
│   ├── launcher.go
│   ├── settings.go
│   └── components.go
├── games/
│   ├── merge/
│   ├── snake/
│   ├── pong/
│   ├── mines/
│   ├── four/
│   ├── conquest/
│   └── ...
└── assets/
    └── fonts/
```

Each game would implement a standard interface, extending whatever conventions your existing EBE boilerplate already provides.

For example:

```go
type MiniGame interface {
    ID() string
    Name() string

    Update() error
    Draw(screen *ebiten.Image)

    Reset()
}
```

The launcher would handle the window, screen transitions, input routing, audio settings, and persistent storage.

Individual games would know almost nothing about the launcher.

You could maintain a static registry:

```go
var Games = []GameEntry{
    {ID: "merge", Name: "Merge", New: merge.New},
    {ID: "snake", Name: "Snake", New: snake.New},
    {ID: "pong", Name: "Pong", New: pong.New},
}
```

Adding a game would involve creating a new package and registering it.

Importantly, **none of the other games would need to change**.

### The Go-specific consideration

Go normally produces statically linked executables. This is actually an advantage for your concept.

Rather than implementing dynamically loaded plugins, I'd compile all the games into one application.

That gives you:

- One executable to distribute.
- No runtime dependencies on separately installed game modules.
- No network downloads to play individual games.
- Easy cross-platform compilation.
- No need for a plugin installation or permission system.

It also means every update would ordinarily be a new build of the whole application, even when the only source-code change is the addition of one game.

For small games, that's perfectly reasonable.

---

## 3. How to guarantee that old games don't change

This is the part of your concept I'd take most seriously.

You could establish an explicit development rule:

**Once a game is published, its rules and default behavior are immutable.**

For example:

| App release | New game | Existing games |
|---|---|---|
| 1.0 | Merge | - |
| 1.1 | Snake | Merge unchanged |
| 1.2 | Pong | Merge, Snake unchanged |
| 1.3 | Mines | Previous three unchanged |
| 1.4 | Four | Previous four unchanged |
| 1.5 | Conquest | Previous five unchanged |

There's an important distinction between *not modifying the source code* and *guaranteeing unchanged behavior*.

Changes to shared dependencies, input handling, rendering, or save formats can affect old games even without modifying their packages.

I'd address that with three mechanisms.

**First: isolate each game's rules.** Keep the rules separate from the rendering system.

**Second: version each game's rules and save format.** For example, `merge/rules/v1` remains available even if you eventually introduce a different merging game.

**Third: create automated regression tests.** Every release replays recorded inputs against every previously published game and verifies the resulting game state.

Ebitengine 2.10, released September 8, 2026, actually makes this more interesting: it introduced application virtualization, input injection, and opt-in headless testing for automated game testing. That could help verify both gameplay and presentation across releases. ([Ebitengine](https://ebitengine.org/en/documents/2.10.html?utm_source=chatgpt.com), [Ebitengine](https://ebitengine.org/en/blog/v2.10.0.html?utm_source=chatgpt.com))

In other words, the bigger the collection becomes, the more valuable an automated test suite becomes.

I would still allow bug fixes and compatibility patches, but distinguish those from changing game rules or features.

---

## 4. Distribution: this could be one app on almost everything

Ebitengine already supports Windows, Linux, macOS, browsers through WebAssembly, Android, and iOS. ([Ebitengine](https://ebitengine.org/?utm_source=chatgpt.com))

That gives you several possible distribution strategies.

| Platform | Distribution | Offline capability |
|---|---|---|
| Linux | Standalone executable / package | Fully offline |
| Windows | Standalone executable / installer | Fully offline |
| macOS | Application bundle | Fully offline |
| Android | APK / app store | Fully offline |
| iOS | App Store | Fully offline |
| Browser | WebAssembly / installable PWA | Offline after initial installation and caching |

Mobile requires some platform-specific packaging. Ebitengine provides bindings for Android and iOS, but doesn't eliminate their native application build processes. ([Ebitengine](https://ebitengine.org/en/documents/mobile.html?utm_source=chatgpt.com))

### Android is particularly attractive

Imagine distributing a single APK containing 50 or 100 tiny games.

No ads. No login. No internet permission.

You could even make the lack of network permission part of the product identity.

On Android, an app that never declares the `INTERNET` permission has a strong, inspectable basis for claiming that it cannot make ordinary network connections.

On desktop, you could publish the source code and reproducible build instructions so users can verify the absence of telemetry.

For iOS, the equivalent claim would depend more heavily on auditing the application and its dependencies because there isn't a directly equivalent general internet permission.

### There's one distribution decision worth making early

Do you want the app to **automatically update itself**, or simply receive updates through its distribution channel?

I'd choose the second.

This avoids needing an application-level update service.

For example, users could install updated versions through their app store, package manager, or manually downloaded release.

They would only need internet access to obtain updates, never to play.

You could also distribute complete update packages through USB drives, local networks, or other offline methods.

---

## 5. How many games could fit in one app?

Potentially hundreds, without particularly large storage requirements.

Most of the games you're discussing require very little unique data.

Consider the actual contents:

| Game | Main data |
|---|---|
| Merge | 4×4 integer array |
| Snake | List of grid coordinates |
| Pong | A few positions and velocities |
| Mines | Grid of integers and flags |
| Four | 7×6 board |
| Conquest | Territory graph, armies, ownership |
| Chess | Board state and move history |
| Sudoku | 81 cells and puzzle generator |
| Patience | Deck state and card positions |

In many cases, the gameplay rules could occupy only a few kilobytes of compiled code.

The shared Go runtime, Ebitengine, UI, fonts, and platform libraries would likely account for much more of the application's size than each additional game.

**The incremental cost of adding a game could be extremely small.**

A 100-game collection isn't necessarily 100 times the size of a single-game application.

It might only be modestly larger, depending on how aggressively you share assets and minimize dependencies.

I would avoid making a specific size promise before measuring an actual release build.

---

## 6. An even more interesting idea: append-only game development

This is where I think your concept becomes distinctive.

Instead of a conventional software development model, you could adopt an **append-only release philosophy**.

Every new release adds something to the collection, but existing games remain available in their original form.

You might call this a *cumulative game anthology*.

The release history becomes:

```text
Version 1.0
  + Merge

Version 1.1
  + Snake

Version 1.2
  + Pong

Version 1.3
  + Mines

Version 1.4
  + Four

Version 1.5
  + Conquest

Version 1.6
  + Patience
```

And so on.

The promise to users would be that every update expands their library without taking anything away.

That's a meaningful contrast with modern mobile games that routinely change their progression systems, monetization mechanics, interfaces, or availability.

You could even preserve previous game implementations as tagged, immutable Go packages.

For instance:

```text
games/
├── merge/
│   └── v1/
├── snake/
│   └── v1/
├── pong/
│   └── v1/
└── conquest/
    └── v1/
```

There's no obligation to delete old versions when you develop better ones.

And since each implementation would be small, preserving them would be inexpensive.

---

## 7. What I'd build first

I would start with a collection containing exactly three games:

**Merge, Snake, and Pong.**

Not because these are necessarily the three most compelling games, but because they establish three different kinds of interactions.

| Game | Architectural capability demonstrated |
|---|---|
| Merge | Turn-based input and deterministic state transitions |
| Snake | Fixed-timestep simulation and grid-based movement |
| Pong | Continuous movement, physics, and real-time input |

Once those three work within the same launcher, you've demonstrated most of the basic capabilities needed for a substantial collection.

I would also establish a handful of conventions from the beginning: local-only save files, shared controller navigation, consistent pause and restart controls, per-game high scores, and a common exit-to-library action.

Then every subsequent game could be almost entirely self-contained.

For an initial development milestone, I'd aim for:

**One executable, three games, zero external assets beyond shared fonts, zero networking, and a test suite proving that adding a fourth game doesn't change the first three.**

That would be a strong foundation.

---

## 8. The biggest potential advantage: the release cadence

Suppose the architecture is established and you're adding relatively simple games.

A plausible development rhythm might look like this:

| Development stage | Games in collection | Estimated effort |
|---|---:|---|
| Initial application architecture | 3 | Several days |
| First release | 5 | Another few days |
| First month | 10-20 | Depends on polish and testing |
| Three months | 25-50 | With a repeatable implementation process |
| One year | 100+ | Possible with sustained development or contributors |

Those are illustrative targets, not reliable forecasts. More complex games such as Conquest, chess, or multiplayer games could take substantially longer.

But with the modular architecture, development effort would increasingly be devoted to creating new games instead of maintaining existing ones.

And because each game would be independently testable, this is a natural fit for agentic coding workflows.

You could have separate coding agents implementing different games in parallel, with each operating inside its own package and conforming to the same interface.

Your automated tests could reject contributions that modify existing games.

That would allow the collection to grow without requiring increasingly complicated coordination.

---

## 9. One thing I would deliberately avoid

**I would not build a general-purpose scripting engine or dynamic plugin system at the beginning.**

It might seem attractive to support downloadable games without updating the application, but that introduces considerable complexity:

- A runtime interpreter or sandbox.
- Compatibility and API versioning.
- Code signing or package verification.
- A distribution and update mechanism.
- Potential security vulnerabilities.
- More complicated cross-platform deployment.

You don't need any of that to achieve your stated objective.

A statically compiled Go application containing all the games is simpler, more portable, easier to audit, and more consistent with the demoscene philosophy.

If the collection eventually reaches hundreds of games and rebuild times become problematic, you can reconsider.

For now, static linking is a feature, not a limitation.

---

## 10. The product identity

I think the strongest selling point isn't actually that the games are free.

It's that the application is **finished software**.

Once you've installed a game, it works. It doesn't ask for anything. It doesn't require an account. It doesn't phone home. It doesn't interrupt you. It doesn't monetize your attention.

And the next update simply gives you another game.

That's a remarkably clear product promise.

I'd consider an explicit set of project guarantees:

> **The Collection**
>
> - Every game works completely offline.
> - No advertising, telemetry, analytics, or user tracking.
> - No accounts, subscriptions, or in-app purchases.
> - Games are never removed or redesigned for monetization.
> - Updates add new games without changing existing gameplay.
> - All game state is stored locally.
> - The application is open source and independently auditable.

The privacy guarantees could be made technically enforceable through the application's architecture and distribution permissions, not just asserted in a privacy policy.

### My recommendation

**Build one Ebitengine application containing a growing, append-only library of games.** Give each game its own Go package, keep game logic independent of the shared UI, and use automated regression testing to enforce backward compatibility.

I'd also seriously consider publishing the entire collection under a permissive open-source license.

The combination of an append-only library, reproducible builds, local-only storage, and no network access could make this feel like a modern successor to the old shareware game collections-except without any of the commercial baggage.

And unlike most ambitious game projects, the first release could be genuinely complete with only three games. Every subsequent game would simply make the collection better.

---

## User

is go and ebe particularly well suited to this kind of cross-platform deployment on various app stores?

---

## ChatGPT

I’ll check Ebitengine’s current platform support and the packaging requirements for the major app stores, since the main question is less whether the games run cross-platform and more how smoothly they can be shipped and updated.

---

## ChatGPT

**Yes. Go + Ebitengine is particularly well suited to your project, especially with Ebitengine 2.10's improvements to cross-platform compilation.** However, there's an important distinction:

**Ebitengine makes it easy to run the same game on many platforms. It doesn't make publishing to every app store equally easy.**

For your proposed collection of lightweight, offline, ad-free games, I'd rate the stack as follows.

## 1. Platform-by-platform assessment

| Platform | Go + Ebitengine suitability | Publishing difficulty | Main limitation |
|---|---|---|---|
| Windows | ⭐⭐⭐⭐⭐ | Low-medium | Packaging and signing |
| Linux | ⭐⭐⭐⭐⭐ | Low | Distribution formats |
| macOS | ⭐⭐⭐⭐⭐ | Medium | Signing and notarization |
| Android / Google Play | ⭐⭐⭐⭐ | Medium | Native Android wrapper |
| iOS / Apple App Store | ⭐⭐⭐⭐ | Medium-high | Xcode, signing, review |
| Web / WebAssembly | ⭐⭐⭐⭐⭐ | Low | Browser-specific packaging |
| Steam | ⭐⭐⭐⭐⭐ | Medium | Store integration and submission |
| itch.io | ⭐⭐⭐⭐⭐ | Low | Few platform-specific requirements |
| F-Droid | ⭐⭐⭐⭐ | Medium | Reproducible, compliant Android builds |

These are qualitative assessments rather than measured engineering scores.

### The biggest improvement: Ebitengine 2.10

Released September 8, 2026, Ebitengine 2.10 introduced pure-Go desktop builds on Windows, macOS, Linux, and FreeBSD.

Previously, some desktop platforms required Cgo and platform-specific C development dependencies.

Now, desktop cross-compilation is considerably simpler, and Ebitengine requires Go 1.25 or later. Mobile builds still require Cgo and platform-specific tooling. ([Ebitengine](https://ebitengine.org/en/documents/2.10.html?utm_source=chatgpt.com))

For example, you can build your application for several desktop targets using Go's standard cross-compilation capabilities:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o games-linux .

# Windows
GOOS=windows GOARCH=amd64 go build -o games.exe .

# macOS ARM64
GOOS=darwin GOARCH=arm64 go build -o games-macos .
```

Assuming your project and dependencies are compatible with these targets, you can produce all three binaries from a Linux development machine.

That is a significant advantage for your existing development environment.

---

## 2. Android: particularly attractive for your concept

Android is probably your most important distribution target.

The Google Play Store is full of the kinds of advertising-heavy puzzle and casual games you're proposing to replace.

And Ebitengine supports Android directly.

However, the deployment architecture looks like this:

```text
Your Go code
     │
     ▼
Ebitengine
     │
     ▼
ebitenmobile bind
     │
     ▼
Android library
     │
     ▼
Android application wrapper
     │
     ▼
Android App Bundle (.aab)
     │
     ▼
Google Play Store
```

Ebitengine's `ebitenmobile bind` tool generates platform bindings that can be incorporated into a native Android application.

The maintainers explicitly explain that this approach is necessary because building complete mobile applications entirely in Go is not currently practical. ([Ebitengine](https://ebitengine.org/en/documents/mobile.html?utm_source=chatgpt.com))

**The important point is that you wouldn't need to rewrite any of your games in Kotlin or Java.**

You would maintain a small Android-specific project responsible for the application lifecycle, packaging, and native integration.

Your games themselves would remain Go.

### One current Google Play requirement

As of August 31, 2026, new Google Play submissions and updates must generally target Android 16, API level 36 or higher. ([Google Help](https://support.google.com/googleplay/android-developer/answer/11926878?hl=en-en&utm_source=chatgpt.com))

That means you need a sufficiently current Android SDK, Gradle configuration, and compatible mobile build pipeline.

This isn't a problem inherent to Go, but it does mean you can't treat Android publishing as a simple `go build` operation.

---

## 3. iOS: fully supported, but Apple adds friction

Ebitengine also supports iOS, including its Metal graphics backend.

As with Android, the underlying game stays in Go, but you need a small native application wrapper and Apple's build tools.

The process is approximately:

```text
Go + Ebitengine
       │
       ▼
ebitenmobile bind
       │
       ▼
iOS framework
       │
       ▼
Xcode project
       │
       ▼
Signed iOS application
       │
       ▼
Apple App Store
```

You'll need access to macOS and Xcode to complete the usual iOS build, signing, and submission process.

Apple's Developer Program costs $99 annually in the United States. ([Apple Developer](https://developer.apple.com/help/account/membership/program-enrollment/?utm_source=chatgpt.com))

### One concern: Apple's copycat policies

This is potentially more important to your project than the programming language.

Apple's App Review Guidelines explicitly prohibit apps that simply copy popular existing applications with minor cosmetic changes.

They also discourage publishing multiple substantially similar apps. ([Apple Developer](https://developer.apple.com/app-store/review/guidelines/?utm_source=chatgpt.com))

Your proposed collection has two advantages:

1. You're publishing one unified application rather than dozens of individual clones.
2. You're creating an independently designed collection with a distinct minimalist interface and philosophy.

Neither guarantees approval, but both help differentiate the product.

I would present it as an **original collection of classic game mechanics**, not as a collection of clones of popular commercial applications.

### Frequent updates are supported

You can add new games to the existing iOS application by publishing new versions. Each version goes through the normal submission and review process. ([Apple Developer](https://developer.apple.com/help/app-store-connect/update-your-app/create-a-new-version?utm_source=chatgpt.com))

A monthly release containing several new games might be more practical than submitting a new version every day.

---

## 4. WebAssembly: an unusually good fit

This might actually be Ebitengine's biggest advantage for your project.

Ebitengine supports compilation to WebAssembly directly from Go:

```bash
GOOS=js GOARCH=wasm go build -o games.wasm .
```

The resulting application runs in a browser using the Go WebAssembly runtime. ([Ebitengine](https://ebitengine.org/en/documents/webassembly.html?utm_source=chatgpt.com))

You could offer the same collection as:

- A website where every game is immediately playable.
- An installable Progressive Web App.
- A native Android application.
- A native iOS application.
- Standalone desktop executables.

All using essentially the same game logic and rendering code.

A PWA could cache its files locally so the games remain playable without an internet connection after installation.

One caveat is that Go's WebAssembly output can be relatively large compared with equivalent JavaScript implementations. Browser downloads and startup performance are worth benchmarking.

Nevertheless, I think **WebAssembly should be a first-class target from the beginning**, especially since you're building a collection rather than one complicated game.

---

## 5. How does Go + Ebitengine compare with alternatives?

For your specific project, I'd compare four stacks.

| Criterion | Go + Ebitengine | Godot | Flutter + Flame | JavaScript + Canvas |
|---|---|---|---|---|
| Desktop deployment | Excellent | Excellent | Good | Browser / wrappers |
| Android | Good | Excellent | Excellent | PWA / wrappers |
| iOS | Good | Excellent | Excellent | PWA / wrappers |
| Browser | Good | Good | Good | Excellent |
| Minimal dependencies | Excellent | Good | Good | Excellent |
| Simple 2D games | Excellent | Excellent | Excellent | Excellent |
| Small game-specific code | Excellent | Good | Good | Excellent |
| Native mobile UI integration | Limited | Good | Excellent | Limited |
| Existing EBE compatibility | Native | Requires rewrite | Requires rewrite | Requires rewrite |
| Single shared codebase | Yes | Yes | Yes | Yes |
| Automated headless game testing | Excellent | Good | Good | Good |

### Go + Ebitengine versus Godot

Godot is probably the strongest alternative if your primary objective is publishing games to Android and iOS.

It has an established editor, mobile export tooling, and a larger ecosystem of mobile game development resources.

However, Godot also brings a more elaborate engine architecture.

For a collection of minimalist games that might consist largely of rectangles, text, and simple state machines, most of that functionality isn't necessary.

Ebitengine's lightweight approach is arguably better aligned with your concept.

### Go + Ebitengine versus Flutter + Flame

Flutter might be the better choice if you wanted a conventional mobile application with lots of native-feeling interface components and some embedded games.

For example, if your application had extensive settings, social features, profiles, achievements, and cloud synchronization, Flutter's mobile app infrastructure could become attractive.

But those are precisely the kinds of features you're trying to avoid.

For a full-screen game collection with a custom interface, Ebitengine makes more sense.

### Go + Ebitengine versus JavaScript

A JavaScript implementation would offer extremely easy browser deployment and straightforward PWA support.

However, desktop and mobile distribution would generally involve additional packaging frameworks.

Since you already have a functioning Go boilerplate, switching to JavaScript doesn't seem worthwhile.

---

## 6. One architectural decision could make cross-platform support much easier

I would strongly recommend separating your game mechanics from Ebitengine itself.

For example:

```text
collection/
│
├── games/
│   ├── merge/
│   │   ├── logic.go
│   │   └── render.go
│   ├── snake/
│   ├── pong/
│   └── ...
│
├── platform/
│   ├── desktop/
│   ├── android/
│   ├── ios/
│   └── web/
│
├── ui/
│   ├── launcher.go
│   └── controls.go
│
└── main.go
```

All platforms would use the same Ebitengine-based rendering and game logic.

The platform-specific code would only handle things such as storage locations, lifecycle events, window behavior, and native packaging.

### Input abstraction is particularly important

Your existing EBE input abstraction could be a major advantage.

The same game might need to support:

| Input action | Desktop | Mobile | Controller |
|---|---|---|---|
| Move up | Arrow key | Swipe | D-pad |
| Move down | Arrow key | Swipe | D-pad |
| Move left | Arrow key | Swipe | D-pad |
| Move right | Arrow key | Swipe | D-pad |
| Confirm | Enter | Tap | A |
| Back | Escape | Back button | B |
| Pause | Escape | Pause icon | Start |

Ebitengine already supports keyboard, mouse, gamepad, and touch input, though not every input type is available on every platform. ([Ebitengine](https://ebitengine.org/en/documents/features.html?utm_source=chatgpt.com))

Your games should consume abstract actions rather than raw device inputs.

That would let you write the game logic once and adapt the controls for phones, tablets, computers, and controllers.

I'd also make the launcher itself controller-navigable.

---

## 7. Automating releases

This is where the combination of Go and Ebitengine becomes particularly powerful.

You could configure GitHub Actions to build all your targets whenever you create a release tag.

For example:

```text
git tag v1.12.0
       │
       ▼
GitHub Actions
       │
       ├── Run tests
       │
       ├── Build Linux AMD64
       │
       ├── Build Linux ARM64
       │
       ├── Build Windows AMD64
       │
       ├── Build macOS ARM64
       │
       ├── Build macOS AMD64
       │
       ├── Build WebAssembly
       │
       ├── Build Android AAB
       │
       └── Build iOS IPA
                 │
                 ▼
        Distribution channels
```

The desktop and WebAssembly builds would be relatively straightforward.

Android would require its SDK and native build environment. iOS would require a macOS runner and the appropriate signing credentials.

Store submissions can also be automated to varying degrees, although they still must satisfy each platform's review requirements.

You could have a release pipeline that:

1. Verifies every previously published game's regression tests.
2. Builds the application for all supported targets.
3. Produces checksums and release artifacts.
4. Publishes the desktop binaries and web version.
5. Submits mobile builds for store review.

That would make regular releases substantially less burdensome.

---

## 8. My recommendation: stay with Go + Ebitengine

For your particular goals, I would not switch frameworks.

The existing EBE boilerplate is already valuable, and Ebitengine is unusually well matched to the games you're proposing.

The biggest limitations are not related to Go's ability to implement the games.

They're related to:

- Native mobile packaging.
- App store submission and review.
- Touchscreen layouts and resolution scaling.
- Platform-specific filesystem behavior.
- Mobile lifecycle handling.

Those are manageable, and they're largely one-time infrastructure investments.

### The one thing I'd validate before committing

**Build your existing EBE application for Android and iOS before implementing the collection.**

Not a complete game-just the existing launcher rendering a screen and responding to touch input.

If both builds succeed, the most significant deployment risk will have been addressed.

Then I'd build the collection around that validated cross-platform foundation.

### Bottom line

I would rate **Go + Ebitengine approximately 9/10 for your specific project**, with one caveat: it's closer to 10/10 for desktop and browser development, and perhaps 8/10 for mobile app-store deployment.

Godot and Flutter provide more integrated mobile development workflows, but neither offers a compelling enough advantage to justify abandoning your existing Go infrastructure for a collection of simple, self-contained 2D games.

**The most compelling architecture is one Go codebase, one Ebitengine runtime, multiple distribution targets, and a continually expanding library of games.**

---

## User

do you think people would be willing to pay a subscription fee if there was a guarantee of a certain number of new games on some regular basis?

---

## ChatGPT

I’ll compare this with existing paid game libraries and look at what subscribers would actually be paying for: the games themselves, the release cadence, or the guarantee that the collection stays ad-free and privacy-respecting.

---

## ChatGPT

**Yes. I think there is a viable market for this, particularly at $1-3 per month, provided the subscription is framed as paying for a continually expanding collection rather than renting access to games.**

The distinction is important because your concept has an unusual advantage: you're proposing to remove almost everything people dislike about mobile games while continuously adding things they do like.

I'd be particularly interested in testing this proposition:

> **$1.99/month. Four new games every month. No ads, no tracking, no accounts, no internet required to play. Every game you receive is yours to keep.**

I suspect that would be substantially more appealing to your target audience than a conventional subscription that disables all the games when someone stops paying.

There are precedents for both models, and one is remarkably close to what you're describing.

## 1. The strongest precedent: Playdate

Panic's Playdate is probably the closest existing example of your proposed release model.

Its original offering included 24 games delivered over 12 weeks, with two new games arriving each week.

Players keep the games permanently.

As of November 2026, Playdate has launched three seasons, including a third season released in October 2026. Later seasons are sold separately. ([Playdate Help](https://help.play.date/games/season/?utm_source=chatgpt.com), [Playdate](https://play.date/?utm_source=chatgpt.com))

This is significant because Playdate demonstrates a model in which **the anticipation of receiving new games is itself part of the product**.

Players aren't necessarily purchasing a collection because they want every single game. They're purchasing a recurring experience of discovery.

There is also evidence of an actual paying audience: by April 2025, Panic reported selling more than 70,000 Playdate handhelds. However, that doesn't establish demand for your particular subscription, because Playdate customers also purchase distinctive hardware and access to original games. ([Engadget](https://www.engadget.com/gaming/panic-will-showcase-the-second-season-of-playdate-games-on-april-17-181556782.html?utm_source=chatgpt.com))

Your proposition would be much cheaper and work on devices people already own.

## 2. What people are already paying

Two major competitors establish that consumers will pay for ad-free game collections.

| Service | US price | Offering | Key difference |
|---|---:|---|---|
| Apple Arcade | $6.99/month | 200+ games, no ads or in-app purchases | Access lasts while subscribed |
| Google Play Pass | $4.99/month | 1,000+ games and apps, no ads or in-app purchases in the catalog | Access lasts while subscribed |
| Your collection | Proposed $1.99/month | Four new minimalist games monthly | Games permanently retained |

Apple confirmed its $6.99 price in August 2026. Google Play Pass currently advertises more than 1,000 games and apps; its $4.99 US price is also listed by Verizon. ([Apple](https://www.apple.com/newsroom/2026/08/apple-arcade-levels-up-with-block-blast-plus-and-art-of-fauna-cozy-puzzles-plus/?utm_source=chatgpt.com), [Google Help](https://support.google.com/googleplay/answer/9473027?hl=en&utm_source=chatgpt.com), [Verizon](https://www.verizon.com/solutions-and-services/add-ons/entertainment/google-play-pass/?utm_source=chatgpt.com))

These aren't direct competitors in terms of game production budgets, but they provide useful pricing anchors.

Your advantage wouldn't be having a bigger catalog.

It would be offering an unusually simple, trustworthy, and inexpensive alternative.

### What I'd expect people to pay

These are hypotheses I'd test, not estimates derived from consumer survey data.

| Monthly price | My assessment | Main challenge |
|---|---|---|
| $0.99 | Very approachable | Transaction fees and low revenue |
| **$1.99** | **Strongest starting candidate** | Requires consistently useful releases |
| $2.99 | Plausible | Expectations for quality increase |
| $3.99 | More difficult | Approaches Google Play Pass pricing |
| $4.99+ | Unlikely for an initial catalog | Competes with much larger services |

The price is less important than the perceived value of the promise.

Someone might not pay $1.99 for another implementation of Snake.

But they might pay $1.99 every month for a collection that already contains 30 familiar games, adds four more every month, works on their phone and computer, and never introduces advertising.

Those are meaningfully different propositions.

## 3. The unusual opportunity: a subscription that builds ownership

I think the strongest model would be closer to a **game-of-the-month club** than Netflix.

Consider three possible arrangements.

| Model | What subscribers receive | What happens after cancellation |
|---|---|---|
| Traditional subscription | Access to all games | Games become inaccessible |
| Membership | Access to the collection and new releases | Retain a basic collection |
| **Ownership subscription** | **New games permanently unlocked each month** | **Keep everything previously unlocked** |

I favor the third.

Suppose someone subscribes for six months at $1.99/month.

They pay $11.94 and permanently acquire 24 games, assuming four monthly releases.

They can cancel, stop paying, and continue playing those games indefinitely.

If they return after a year, they can subscribe again and start acquiring whatever new games are being released.

This has several appealing consequences.

First, it reduces the sense of financial commitment.

Second, it aligns with your broader premise of rejecting predatory game monetization.

Third, it gives subscribers something tangible for every payment.

Fourth, you could make the entire payment and licensing model relatively simple.

**There's a catch:** ownership entitlements and the subscription itself introduce account, billing, and restoration complexity. You could avoid collecting personal information yourself by delegating purchases to Apple and Google, but you would still need to integrate their purchase systems.

Also, Google Play explicitly requires subscriptions to provide sustained or recurring value rather than merely one-time purchases. Apple explicitly recognizes subscriptions that provide new game levels, episodic content, and substantive recurring updates. Your recurring game releases might qualify, but I'd verify the specific ownership-unlock arrangement before building around it. ([Apple Developer](https://developer.apple.com/app-store/review/guidelines/?utm_source=chatgpt.com), [Google Help](https://support.google.com/googleplay/android-developer/answer/18265346?hl=en&utm_source=chatgpt.com))

An alternative would be prepaid, nonrenewing seasonal bundles, which more closely resemble Playdate's approach.

## 4. What would the economics look like?

At $1.99/month, the numbers become interesting even with a relatively small audience.

| Paying subscribers | Monthly gross revenue | Annual gross revenue |
|---:|---:|---:|
| 100 | $199 | $2,388 |
| 500 | $995 | $11,940 |
| 1,000 | $1,990 | $23,880 |
| 2,500 | $4,975 | $59,700 |
| 5,000 | $9,950 | $119,400 |
| 10,000 | $19,900 | $238,800 |
| 25,000 | $49,750 | $597,000 |

These figures assume the subscriber count remains constant throughout the year. They exclude store commissions, taxes, refunds, customer acquisition, and development costs.

For a project with a highly automated production pipeline and minimal infrastructure expenses, even a few thousand subscribers could be financially meaningful.

### The development cost is the interesting variable

Let's suppose you promise four new games per month.

That's 48 games annually.

If each game takes an average of eight hours to implement, test, polish, and release, you're looking at:

**384 hours annually**, or approximately 7.4 hours per week.

At 1,000 subscribers, the $23,880 in annual gross revenue would work out to approximately $62 per development hour before expenses.

At 5,000 subscribers, it would be approximately $311 per development hour.

Those numbers are attractive, but the eight-hour assumption is optimistic. Some games will take substantially longer, and platform maintenance, marketing, customer support, and release management add overhead.

The important point is that **your marginal cost per additional subscriber could be extremely low**, because the application doesn't require servers, multiplayer infrastructure, or ongoing content delivery.

Your main cost would be producing the games and acquiring customers.

---

## 5. I would not promise weekly releases immediately

I would start with a monthly guarantee.

Something like:

**Four new games every month, guaranteed.**

Then release them together or on a predictable weekly schedule.

A weekly release is exciting, but it also creates a very public deadline every seven days.

Monthly batches would give you more flexibility to maintain quality.

You could build up a backlog before launching and maintain several months of completed games in reserve.

I'd also consider explicitly defining a minimum standard for what qualifies as a new game. Otherwise, subscribers might become frustrated if four nearly identical variants count as four separate releases.

For example, a new game could have to introduce an independently playable ruleset, rather than a new board size, skin, or difficulty mode.

That would make the guarantee much more meaningful.

---

## 6. There's one important tension with the privacy promise

Your original goal was an application that is entirely offline, ad-free, and surveillance-free.

A subscription makes that harder.

Not impossible, but harder.

Conventional subscription systems need some way to determine whether a user has paid.

That generally involves store-managed purchase records, periodic entitlement validation, or both.

You can make the actual games completely offline, but payments and restoration of purchases will sometimes require connectivity.

If your priority is absolute simplicity, there's an attractive alternative:

### Sell annual collections instead of subscriptions

For example:

**2027 Collection - $9.99**

Includes 48 games, delivered in twelve monthly updates.

At the end of 2027, the buyer permanently owns all 48 games.

Then:

**2028 Collection - $9.99**

Includes another 48 games, delivered throughout the year.

This creates recurring revenue without requiring automatic subscription renewals.

It also offers a more credible ownership promise.

Users can buy a season, receive all its scheduled updates, and keep the resulting games forever.

You'd still need to handle purchases and updates, but not continuously establish subscription status merely to let users play.

This is particularly close to the Playdate model.

---

## 7. What I'd test before choosing a business model

I wouldn't start by asking people whether they'd pay $1.99/month.

People often say yes to hypothetical low-priced subscriptions, but that doesn't necessarily translate into actual purchases or renewals.

Instead, I'd launch a polished free collection of approximately 10-15 games and present three real purchase options:

| Option | Offer | Test price |
|---|---|---:|
| Monthly | Four new games each month | $1.99/month |
| Annual | 48 new games over 12 months | $14.99/year |
| Season | 12 new games delivered over three months, yours forever | $4.99 |

I'd initially make these alternative purchase plans rather than offering overlapping products simultaneously.

Then measure actual purchases, cancellations, and repeat purchases.

I would be particularly interested in whether people are more attracted to **the number of games** or **the promise of ownership**.

Those two value propositions might appeal to different audiences.

There is also an important challenge: these are mostly familiar game mechanics that people can already play for free. Your customers aren't necessarily paying for access to unique gameplay.

They're paying for the experience of having a consistently good, trustworthy, distraction-free collection.

That means presentation, reliability, and the reputation of the collection will matter at least as much as the number of games.

---

## 8. My preferred business model

I would lean toward a hybrid:

**Free permanent starter collection + paid annual seasons.**

The starter collection would always be free and would never contain advertisements, tracking, or artificial restrictions.

A paid season would guarantee a fixed number of additional games over a defined period.

Users would own each season permanently.

That model has several advantages over conventional subscriptions:

- No recurring billing required.
- No need to revoke access after cancellation.
- A clear and finite obligation to deliver games.
- Minimal billing infrastructure.
- A straightforward offline ownership story.
- A new product to sell each year.

And you could still offer an optional automatically renewing annual subscription for people who want to receive every future season.

### My overall assessment

I think there's a credible niche market here, but I wouldn't assume that millions of people want yet another game subscription.

The strongest opportunity is probably among people who are specifically frustrated with the modern mobile gaming experience: intrusive advertising, tracking, manipulative engagement systems, mandatory accounts, and games that depend on remote servers.

**You wouldn't be selling 2048 or Snake. You'd be selling a guarantee that people can enjoy familiar games without being exploited.**

That is a meaningfully differentiated product.

And the recurring delivery model provides a reason for satisfied users to keep supporting the project.

**The question I'd investigate next is whether an annual $9.99-$14.99 season with permanent ownership would outperform a $1.99 monthly subscription.** My hypothesis is that the ownership model would be particularly attractive to the audience you're trying to reach, while reducing the technical and administrative complexity of maintaining a truly offline game collection.

---

## User

what about something like $1/mo or $10/year and make it cross-platform so they can play all the games on all their devices

---

## ChatGPT

That pricing feels more aligned with the simplicity of the product. I’ll check the main complication: whether Apple, Google, and desktop purchases can unlock the same library across devices without forcing users into a conventional tracking-heavy account system.

---

## ChatGPT

I actually like **$1/month or $10/year with access across all devices** considerably more than the $1.99/month model we discussed earlier.

At that price, you're getting close to something that feels like a public utility for casual gaming rather than another commercial gaming service.

The proposition becomes extremely simple:

> **Every game. Every device. One dollar a month.**
>
> A growing collection of classic games, with new games added every month.
>
> No ads. No tracking. No microtransactions. No internet required to play.
>
> $1/month or $10/year. One membership works everywhere.

I think the cross-platform aspect is especially important. A person who pays through Google Play should also be able to open the collection on their Windows laptop, iPad, Linux desktop, or web browser without paying again.

And at $10/year, the purchase is inexpensive enough that someone might subscribe simply because they want the project to exist.

## 1. The pricing is compelling

| Plan | Price | Effective monthly cost |
|---|---:|---:|
| Monthly | $1 | $1.00 |
| Annual | $10 | $0.83 |
| Free | $0 | Starter collection |

The annual plan provides a 16.7% discount compared with paying monthly for a year.

I would make both paid plans identical in functionality. The only difference would be the billing interval.

If you can consistently deliver four new games monthly, an annual subscriber would receive 48 additional games for $10.

That's approximately **21 cents per new game**, in addition to access to the existing collection.

There is an obvious caveat: these are familiar game mechanics rather than 48 professionally produced original games. Their perceived value will depend on the quality of the implementations and the collection itself.

But that very low price makes the proposition easier to understand.

### Revenue potential

Assuming a mix of monthly and annual subscribers, the average revenue per subscriber would fall between $10 and $12 annually.

| Paying subscribers | Annual gross revenue |
|---:|---:|
| 500 | $5,000-$6,000 |
| 1,000 | $10,000-$12,000 |
| 2,500 | $25,000-$30,000 |
| 5,000 | $50,000-$60,000 |
| 10,000 | $100,000-$120,000 |
| 25,000 | $250,000-$300,000 |
| 50,000 | $500,000-$600,000 |

These figures assume stable subscriptions for a full year and exclude commissions, taxes, refunds, and operating expenses.

I wouldn't assume that reaching 10,000 subscribers is easy. Distribution and customer acquisition would likely be harder than writing the games.

But the pricing gives you a straightforward way to test whether there's a meaningful market.

---

## 2. Cross-platform purchases are possible under current app-store rules

This is the most encouraging finding.

**Apple explicitly allows multiplatform services to recognize subscriptions purchased on other platforms**, provided the corresponding content or subscription is also offered through in-app purchases where required.

That's in section 3.1.3(b) of Apple's App Review Guidelines. ([Apple Developer](https://developer.apple.com/app-store/review/guidelines/?utm_source=chatgpt.com))

Google Play also supports applications that recognize subscriptions purchased elsewhere. Its policies allow consumption-only applications where users sign in and access content they previously purchased through another channel. ([Google Help](https://support.google.com/googleplay/android-developer/answer/10281818?hl=en&utm_source=chatgpt.com))

So the basic model is feasible:

| Where someone subscribes | Where they can play |
|---|---|
| Apple App Store | iOS, Android, Windows, macOS, Linux, web |
| Google Play | Android, iOS, Windows, macOS, Linux, web |
| Your website | Web, desktop, Android, iOS |
| Desktop distribution | Desktop, web, Android, iOS |

The important qualification is that **the app stores don't automatically share their customer identities or purchase records**.

You would need a way to associate the purchase with the same user across platforms.

That's where the architecture gets interesting.

## 3. I'd use a single, privacy-preserving membership system

Rather than using Apple, Google, or Steam accounts as your primary identity system, I'd create a very small independent membership service.

The architecture could look like this:

```text
             MEMBERSHIP
             $1 / month
             $10 / year
                  │
                  ▼
         Purchase verification
          Apple / Google / Web
                  │
                  ▼
          Membership service
                  │
                  ▼
         Signed access license
                  │
       ┌──────────┼──────────┐
       │          │          │
       ▼          ▼          ▼
     Phone      Desktop     Browser
       │          │          │
       └──────────┼──────────┘
                  ▼
          Local game library
                  │
                  ▼
           Offline gameplay
```

Your servers would only need to handle membership status and device authorization.

They would not need to receive gameplay data.

### My preferred authentication design

Users could create a membership with a passkey, or use a recovery code that they can enter on another device.

You wouldn't need to require an email address, although optional recovery email would be useful.

On each new device, the user could authorize their copy of the application and receive a cryptographically signed license.

The application stores that license locally.

The device can then verify the signature without contacting your servers.

**This allows the games to operate entirely offline.**

For example, a license might contain:

```json
{
  "license_version": 1,
  "membership": "annual",
  "valid_until": "2027-11-06",
  "library": "complete"
}
```

The actual license would be cryptographically signed and could be bound to a particular installation.

You could make authorization valid through the end of the subscription plus a short grace period.

That means someone who goes camping, travels internationally, or simply doesn't connect a particular computer to the internet can still play.

The tradeoff is that a device must eventually reconnect to receive renewed access, unless you're offering permanent ownership of already acquired games.

### One account, unlimited devices?

I'd seriously consider allowing unlimited personal devices.

At $10/year, the cost of policing device limits could exceed the financial benefit.

However, I'd distinguish between unlimited devices and unlimited independent users.

A reasonable starting policy would be one membership per person, usable across their own devices.

You could avoid invasive device fingerprinting entirely.

---

## 4. The actual payment processing cost matters at $1

There's an interesting difference between app-store subscriptions and direct website subscriptions.

At your proposed price, fixed payment-processing fees become important.

Apple's Small Business Program provides qualifying developers with a 15% commission rate. Google's current US subscription fee structure generally totals 15% when Play Billing is used, including its billing fee. ([Apple Developer](https://developer.apple.com/app-store/small-business-program/?utm_source=chatgpt.com), [Google Help](https://support.google.com/googleplay/android-developer/answer/112622?hl=en-en&utm_source=chatgpt.com))

Stripe's standard US online card rate is 2.9% plus $0.30 per successful transaction, before any additional applicable billing fees. ([Stripe](https://stripe.com/pricing?locale=en-US&utm_source=chatgpt.com))

Here's what that means:

| Payment method | $1 monthly payment | $10 annual payment |
|---|---:|---:|
| Gross revenue | $1.00 | $10.00 |
| Apple / Google at 15% | $0.85 | $8.50 |
| Stripe at 2.9% + $0.30 | $0.67 | $9.41 |

These are simplified examples before taxes, refunds, or additional fees.

**The annual subscription is considerably more attractive for direct payment processing.**

At $1/month, Stripe's fixed transaction fee alone consumes 30% of the payment.

At $10/year, that fixed fee consumes only 3%.

I'd consider allowing monthly subscriptions through the app stores, while emphasizing annual subscriptions for direct website purchases.

Store-specific billing and linking rules would still apply.

---

## 5. I'd separate membership from game ownership

There are two possible ways to implement the $10 annual subscription.

### Model A: Subscription access

Users receive access to the complete collection while subscribed.

When they cancel, they retain access until their paid period expires.

Then the paid games become unavailable.

This is the conventional subscription model and is straightforward to implement.

### Model B: Permanent game ownership

Users receive new games throughout their membership.

Once a game has been delivered to them, they can continue playing it permanently-even after canceling.

They need to renew only to receive additional games.

This is closer to the Playdate model we discussed.

I think **Model B is much more consistent with the philosophy of your project**.

For example:

| Customer | Payment history | Games permanently available |
|---|---|---|
| New user | Free | Starter collection |
| One-month subscriber | $1 | Starter collection + games acquired during that month |
| One-year subscriber | $10 | Starter collection + games acquired during the year |
| Returning subscriber | Renews after cancellation | Previously acquired games + newly acquired games |

There's a technical complication: you'll need to decide how people acquire games added while their membership is active, especially if they never open the application.

I'd award games based on release dates and subscription dates, not whether the person actually downloaded them.

That way, the user doesn't miss anything simply because they didn't open the application for a few months.

For offline ownership, signed license records could preserve the specific games someone has acquired.

You would need a way to restore those entitlements on new devices and account for refunds.

**I wouldn't necessarily commit to permanent ownership before validating store billing requirements**, but it would be my preferred design.

---

## 6. The architecture could be surprisingly inexpensive to operate

This is particularly appealing.

You don't need to host the actual games on your servers during gameplay.

The games are embedded directly into the application.

Your backend would only need to maintain subscription entitlements and issue licenses.

Something like:

| Component | Required? | Purpose |
|---|---|---|
| Game server | No | All games run locally |
| User analytics | No | No gameplay tracking |
| Advertising infrastructure | No | No advertising |
| Cloud saves | Optional | Could be local-only |
| Authentication service | Yes, for cross-platform membership | Associate purchases across devices |
| Entitlement database | Yes | Record membership and game ownership |
| License signing service | Yes | Enable offline verification |
| Payment integrations | Yes | Apple, Google, web purchases |
| Game-content CDN | Not necessarily | Games can ship inside app updates |

This could support a substantial user base with a very small backend.

The important scaling factor would be authentication and subscription activity, rather than gameplay.

For an initial implementation, a small Go service with SQLite or PostgreSQL would be entirely reasonable.

You could also allow users to export their game saves without ever connecting to a server.

---

## 7. I'd make cross-platform access the headline feature

Consider how many competing casual game services are platform-specific.

Apple Arcade is fundamentally an Apple ecosystem service.

Google Play Pass is oriented around Android.

Your collection could deliberately be independent of those ecosystems.

Imagine someone starting a game of Mines on their Android phone and later opening the same game on their Linux desktop.

Their membership works on both.

I'd initially support independent local saves on each device, then add optional save transfer.

Automatic cross-device save synchronization would require additional infrastructure and could complicate your privacy promise.

A simple encrypted export/import system might be enough for the first release.

And if you eventually implement synchronization, it could be end-to-end encrypted.

---

## 8. The major uncertainty isn't technical

The architecture is feasible, but the business question remains unproven:

**Will enough people pay $10 annually for minimalist versions of familiar games?**

Your pricing has two advantages.

First, it's below the price of most commercial game subscriptions.

Second, it's low enough to appeal to people who simply want to support an alternative to ad-driven mobile gaming.

But it also has a weakness: the games themselves are widely available for free.

That means your brand promise needs to be unusually clear.

I'd position the offering around four qualities:

**Simplicity. Privacy. Permanence. Portability.**

A user should be able to understand the offering in seconds.

They shouldn't need to compare tiers, understand virtual currencies, or wonder whether a game will disappear.

---

## 9. What I'd actually launch

I would start with this product structure:

| Feature | Free | $1/month | $10/year |
|---|---|---|---|
| Starter games | ✓ | ✓ | ✓ |
| New games | - | ✓ | ✓ |
| Offline gameplay | ✓ | ✓ | ✓ |
| No advertising | ✓ | ✓ | ✓ |
| No tracking | ✓ | ✓ | ✓ |
| Cross-platform | ✓ | ✓ | ✓ |
| Local saves | ✓ | ✓ | ✓ |
| Permanent access to acquired games | - | ✓ | ✓ |

I'd offer approximately 10 free games initially, with a commitment to four new games monthly for subscribers.

I'd also make every free game a fully functional implementation. No artificial restrictions, timers, or locked features.

The paid membership would simply fund the ongoing creation of additional games.

### One important implementation detail

I would build the app so that **every game is included in the binary, regardless of subscription status**.

The membership controls which games are unlocked, not which game files are downloaded.

That has several advantages: it enables completely offline gameplay, simplifies updates, and avoids separate content downloads.

It also creates a less obvious advantage: the application is genuinely the same across platforms.

The caveat is that Go binaries can be inspected and modified, so local unlock checks are not strong DRM. I would deliberately accept that tradeoff rather than introduce intrusive licensing controls.

---

## My recommendation

I think **$1/month or $10/year, one membership for every device, with permanent access to acquired games**, is the strongest version of the concept you've proposed so far.

It has a straightforward value proposition, a potentially low operating cost, and a meaningful distinction from conventional subscriptions.

I would build it as two mostly independent systems:

**The game collection:** A fully offline, cross-platform Go/Ebitengine application with no tracking or networking requirements during gameplay.

**The membership service:** A tiny, privacy-preserving system that verifies purchases and issues portable, signed licenses.

Those systems could evolve independently.

You could publish dozens of new games without modifying the membership infrastructure, and you could improve the membership infrastructure without changing any game.

**One final strategic consideration:** At $10/year, I wouldn't optimize for maximizing revenue per user. I'd optimize for making the product so simple, reliable, and unobjectionable that people recommend it to their friends and family.

If the collection became the default recommendation whenever someone asked for an ad-free version of a familiar game, that could be a very effective distribution strategy.

---

## User

so my friend future infinitive has an anarchist notaflof ai school in the mission and their website is basically like clickable rows of vibrant animations that take you to more detailed content. I am thinking game selection should work like that , with animations of gameplay for each game playing on repeat, so the user scrolls through a list of vibrant video loops of each game and selects exactly the type of gameplay they actually want. they click on it, and then they are in the game. 

so the free one could still include the entire list of games, but only the first four or five are accessible, and the rest go to a screen that explains the subscription options. 

i also think it's very important to have the list sortable by how often the user has played each game by default, or by popularity, etc. and the user should be able to favorite games and have those stick to the top of the list, etc.

---

## ChatGPT

I think this is a much stronger interface than a conventional game launcher. You're describing something closer to an **interactive catalog of gameplay**, where users choose what to play by seeing the mechanics in action rather than recognizing a title or interpreting a thumbnail.

That makes particular sense for your project because many of the games are deliberately generic. Someone might not know what "territorial conquest" means, but they can immediately recognize the experience they want when they see armies expanding across a map.

And because your games are minimalist and procedurally rendered, you have an unusual advantage: **you could generate the animated previews using the actual game code**, rather than producing and maintaining separate promotional videos.

I'd build the entire application around that idea.

## 1. The core interface: a scrolling wall of gameplay

Imagine opening the app and immediately seeing:

```text
┌─────────────────────────────────────────┐
│                                         │
│  ALL GAMES                 ⚙  Search    │
│  Sort: Most Played ▾                    │
│                                         │
│  ★ FAVORITES                            │
│                                         │
│  ┌───────────────────────────────────┐  │
│  │                                   │  │
│  │     2  4  8  16                  │  │
│  │        ↘  MERGE!                  │  │
│  │                                   │  │
│  │  MERGE                       ▶    │  │
│  └───────────────────────────────────┘  │
│                                         │
│  ┌───────────────────────────────────┐  │
│  │                                   │  │
│  │      ▪ ▪ ▪ ● ● ●                  │  │
│  │       SNAKE                       │  │
│  │                                   │  │
│  │  SNAKE                       ▶    │  │
│  └───────────────────────────────────┘  │
│                                         │
│  ┌───────────────────────────────────┐  │
│  │       TERRITORY CAPTURE           │  │
│  │    ▧▧▧▧ → ▩▩▩▩▩                   │  │
│  │                                   │  │
│  │  CONQUEST                🔒       │  │
│  └───────────────────────────────────┘  │
│                                         │
└─────────────────────────────────────────┘
```

Each row is a wide, vibrant, animated demonstration of its game.

Tapping anywhere on an unlocked row launches the game immediately.

Tapping a locked row displays a simple membership screen.

No intermediate description page, installation step, loading a game from the internet, or conventional storefront.

**The preview is the description.**

This also means people can browse the complete collection without paying, which I think is important for communicating the value of the subscription.

## 2. The sorting and favorites system is essential

I particularly agree with making **Most Played** the default ordering.

A growing library of 100 games could become frustrating to navigate if users have to scroll past 90 games they rarely play to reach their favorites.

I'd support these sorting options from the beginning:

| Sort mode | Behavior | Requires internet? |
|---|---|---|
| **Most Played** | Sort by local play count, descending | No |
| Recently Played | Most recently opened first | No |
| Most Popular | Global popularity ranking | Only to refresh ranking |
| Newest | Most recently released first | No |
| Alphabetical | Sort by game name | No |
| Random | Shuffle the catalog | No |
| Least Played | Surface neglected games | No |

I'd make favorites work independently of sorting.

**Favorites should always be pinned above the rest of the collection**, regardless of the selected sorting mode.

Within favorites, users could reorder games manually or apply the selected sort.

A few additional details would make this feel particularly good:

- A single star icon on each row adds or removes a favorite.
- Favorites persist locally without requiring an account.
- The application remembers the last selected sort mode.
- Opening a game increments its local play count.
- Returning to the library preserves the previous scroll position.

One subtle issue: counting game launches isn't always the same as measuring how frequently someone actually plays.

For the initial release, I'd count a play after someone has interacted with a game, rather than merely opening and immediately closing it.

You can refine the definition later without collecting any information externally.

### A better default ordering for first-time users

I'd use:

1. Favorites, if any.
2. Most-played games.
3. All remaining games in a curated order.

Otherwise, when a new user has played nothing, every game has the same play count.

I'd avoid breaking ties alphabetically because it makes the initial experience feel arbitrary.

And importantly, **I wouldn't automatically push every newly released game to the top**. That would disrupt the user's carefully established ordering.

You could instead have a small, dismissible *New Games* section or badge.

---

## 3. Don't use actual video files unless necessary

This is the technical decision that interests me most.

You described looping animations of gameplay, and I think the ideal implementation is **deterministic gameplay previews rendered directly in Ebitengine**.

Each game could implement a preview interface:

```go
type GamePreview interface {
    Update()
    Draw(screen *ebiten.Image)
    Reset()
}
```

The main launcher could render those previews into offscreen images and then draw them inside the scrolling rows.

There are two approaches worth considering.

### Option A: Actual live gameplay simulations

Each preview runs a simplified AI demonstration of its game.

Examples:

| Game | Preview behavior |
|---|---|
| Merge | Tiles slide and combine into larger numbers |
| Snake | Snake automatically navigates toward food |
| Pong | Two computer-controlled paddles exchange shots |
| Mines | Tiles reveal themselves in expanding patterns |
| Conquest | Territories change ownership across a map |
| Chess | Pieces move through a recorded opening |
| Solitaire | Cards move automatically into valid positions |

This has some appealing properties.

The animation always reflects real gameplay, and you wouldn't need to maintain video assets.

However, running multiple simulations and rendering them simultaneously could become unnecessarily expensive on phones.

### Option B: Recorded deterministic gameplay

Instead of simulating AI in real time, each game supplies a short sequence of recorded inputs.

For example, Merge might contain:

```go
var PreviewMoves = []Direction{
    Left,
    Down,
    Right,
    Up,
    Left,
    Down,
}
```

The preview initializes from a known random seed and replays those actions in a loop.

This produces a predictable, repeatable demonstration of the game.

You could even ensure that the sequence includes something visually satisfying, like merging two 128 tiles.

**I favor this approach**, with one modification: don't necessarily update every preview at 60 FPS.

For example, you could render only the currently visible previews, at 12-24 FPS, and pause everything outside the viewport.

That would be substantially more efficient.

For particularly complicated games, pre-rendered animation frames might eventually be preferable.

### Important: separate the preview from the game

I'd extend your EBE game interface with optional preview support:

```go
type GameDefinition struct {
    ID          string
    Name        string
    Description string
    ReleaseDate string
    Free        bool

    NewGame    func() Game
    NewPreview func() Preview
}
```

That would let each new game arrive with its own animated catalog entry.

And because the preview is part of the game's package, it remains unchanged when you add more games.

The launcher doesn't need to understand how Snake, Merge, or Conquest work. It only knows how to display their previews.

---

## 4. Making the free version feel generous

I agree with showing the entire catalog to everyone.

But I'd distinguish **browseable** from **playable**.

Every user can see every animated preview.

Everyone gets the first four or five games completely free.

Paid games display a small lock icon.

A locked game's animation keeps playing, so users can actually see what they're missing.

When someone taps one, they get a straightforward membership screen.

Something like:

> **Unlock the complete collection**
>
> Play every game, including four new games every month.
>
> **$1/month** or **$10/year**
>
> Works on all your devices.
>
> No ads. No tracking. Play offline.
>
> [Subscribe] [Not Now]

I would avoid making the free user feel like they're constantly encountering obstacles.

There's a distinction between showing a paid catalog and aggressively pushing subscriptions.

You can display locked games prominently without using timers, repeated pop-ups, artificial scarcity, or other engagement tactics.

### One change I'd make to your proposal

I wouldn't automatically define the free games as the first five entries in the sortable list.

I'd give them explicit permanent free status.

For example:

**Merge, Snake, Pong, Mines, and Four in a Row** could always be free.

Then sorting, filtering, and favoriting would never affect which games are unlocked.

A new user could also filter the catalog to show only playable games.

I would also avoid putting all paid games at the bottom by default, since part of your idea is allowing people to discover the entire collection visually.

Apple requires clear disclosure when showcased content needs additional purchases, and paid functionality on both Apple and Google Play is subject to their in-app billing policies. Your proposed locked-game previews are compatible in principle with a normal in-app purchase flow, but the payment screen needs platform-appropriate integration. ([Apple Developer](https://developer.apple.com/app-store/review/guidelines/?utm_source=chatgpt.com), [Google Help](https://support.google.com/googleplay/android-developer/answer/9858738?hl=en&utm_source=chatgpt.com))

---

## 5. Popularity without surveillance

This is probably the most interesting tension in your proposal.

You want users to sort by popularity.

But you also want an application that does not monitor its users.

Those goals aren't necessarily incompatible.

There are several ways to approach this.

### Local popularity

This is easiest.

Each game records how often the user plays it.

Everything stays on the device.

No accounts or network calls required.

### Global popularity

To rank games by how much they're played across the entire user base, you'd need some source of aggregate information.

You could collect anonymous, opt-in usage counts, but that still constitutes telemetry.

I think there's a better default for your project.

**Publish a global popularity ranking as part of each application update.**

For example:

```json
{
  "ranking_version": "2026-11",
  "games": [
    {"id": "merge", "rank": 1},
    {"id": "mines", "rank": 2},
    {"id": "snake", "rank": 3},
    {"id": "conquest", "rank": 4}
  ]
}
```

Those rankings could come from voluntary community polls, public website voting, or other explicitly disclosed sources.

No data needs to be transmitted from the application.

A user can sort by global popularity even while completely offline.

The limitation is that the rankings wouldn't necessarily represent actual usage. I'd label community-voted rankings accordingly rather than presenting them as measured player counts.

If you eventually decide to offer opt-in aggregate gameplay statistics, you can do so without changing the rest of the architecture.

---

## 6. What data should the application save?

Very little.

I would define a local preferences file containing:

```json
{
  "sort": "most_played",
  "favorites": [
    "merge",
    "snake",
    "conquest"
  ],
  "play_counts": {
    "merge": 42,
    "snake": 17,
    "pong": 8
  },
  "last_played": {
    "merge": "2026-11-07T08:30:00Z"
  }
}
```

You would also store individual game save files and, where needed, a local membership entitlement.

No gameplay history needs to leave the device.

One architectural point I'd decide early: **should favorites and play counts synchronize across devices?**

My preference would be local-only by default, with optional transfer or end-to-end encrypted synchronization later.

This avoids requiring a network connection every time the launcher opens.

And it means that the entire application behaves identically whether the device is online or offline.

---

## 7. One subtle design improvement: don't autoplay everything

I love the visual direction, but the implementation needs to respect the capabilities of different devices.

On a phone, perhaps only three game rows are visible at once.

On a large desktop display, there might be eight.

I'd render animations only for visible rows, plus perhaps one row above and below the viewport.

Everything else remains paused.

That makes the system scale to hundreds of games without hundreds of active simulations.

I'd also support reduced motion. If a user enables the operating system's reduced-motion setting, the catalog could display representative still frames instead of continuous animation.

That preserves the visual identity while making the interface more accessible.

An especially nice detail would be to pause the preview when a user opens the game, then resume from the same position when they return to the library.

The entire interface could feel continuous rather than like a conventional application with separate screens.

---

## 8. The experience becomes self-explanatory

This is what I like most about your concept.

Imagine someone browsing the collection who has never played 2048.

They don't need to know the title, read a description, or watch a tutorial.

They see numbers moving around a grid and combining.

They immediately understand the basic interaction.

Likewise, someone who doesn't recognize Risk can watch territories change color and armies move around a map.

They understand that it's a conquest game.

**You're replacing the conventional game-discovery interface with a visual demonstration of gameplay itself.**

That's particularly appropriate for a collection built around familiar mechanics and minimal presentation.

It also distinguishes your application from a typical retro game launcher, which relies heavily on box art, screenshots, and recognizable brands.

Your games wouldn't need famous names or recognizable artwork to communicate what they offer.

The gameplay would sell itself.

---

## 9. How I'd prioritize implementation

For the first release, I'd establish the launcher before building a large catalog.

| Feature | Priority |
|---|---|
| Animated, clickable game rows | Essential |
| Favorites pinned to top | Essential |
| Local play-count tracking | Essential |
| Sort by most played | Essential |
| Sort by newest and alphabetically | Essential |
| Free versus locked games | Essential |
| Subscription screen | Essential |
| Preserve scroll position | Essential |
| Search | High |
| Filter to playable games | High |
| Global popularity | Later |
| Cross-device preference sync | Later |
| Community ratings | Probably unnecessary |

The key is that the launcher should remain simple even as the number of games grows.

I'd want the interface to work beautifully with five games and still feel manageable with 500.

### One thing I'd consider adding

A **Recently Added** view that highlights games introduced in the latest update.

This would reinforce your ongoing release promise without requiring notifications or engagement tracking.

People could open the application, see what was added, and immediately play.

It would be a natural way to make the monthly subscription feel valuable.

---

## My overall recommendation

I'd make the animated catalog the defining feature of the entire application.

Not just a launcher, but a distinctive experience that makes browsing simple games enjoyable in itself.

The architecture would have four clean layers:

**Game modules → deterministic animated previews → sortable/favoritable catalog → membership access control.**

The most important technical decisions are to make previews part of each game's package, keep catalog preferences local, and render only visible animations.

And the most important product decision is to make the interface equally beautiful for free and paid users.

Everyone should be able to discover the complete collection. Paying should unlock more games, not remove deliberately introduced frustrations.

**I think the strongest version of this product is an animated, scrollable library of games where every row shows you exactly what you're getting-and one tap takes you straight into playing it.**

---

## User

i think i would like to have it as close to zero text as possible. definitely no search. maybe at the top there is a star emoji for the user's most played games, a group of people emoji for the global most popular ones, and a sparkle emoji for a discovery menu. and not having any text on each game in the list, so instead of trying to brand each one with a non-trademarked word, it can just be the video of the gameplay.

---

## ChatGPT

I think that's a substantially better direction. **The games don't need names in the interface at all.** If the gameplay is sufficiently recognizable, the animation becomes both the identity of the game and the means of selecting it.

It also makes the collection feel less like a commercial storefront and more like an interactive exhibition.

## The interface

I'd reduce the entire launcher to three controls:

| Icon | Meaning | Behavior |
|---|---|---|
| ⭐ | Yours | Games ordered by personal play frequency |
| 👥 | Popular | Games ordered by global popularity |
| ✨ | Discover | Games arranged to encourage exploration |

The default would be ⭐, with the user's favorite and most-played games appearing first.

Below those controls: **nothing but an edge-to-edge scrolling collection of animated gameplay.**

No game titles, descriptions, search box, ratings, categories, or promotional banners.

Each animation is a button. Tap it, and you're playing.

### A possible layout

```text
┌──────────────────────────────┐
│                              │
│       ⭐     👥     ✨        │
│                              │
├──────────────────────────────┤
│                              │
│     [ANIMATED GAMEPLAY]      │
│                              │
├──────────────────────────────┤
│                              │
│     [ANIMATED GAMEPLAY]      │
│                              │
├──────────────────────────────┤
│                              │
│     [ANIMATED GAMEPLAY]      │
│                              │
├──────────────────────────────┤
│                              │
│     [ANIMATED GAMEPLAY]      │
│                              │
└──────────────────────────────┘
```

I'd even eliminate visible borders between rows if each game's animation has a distinct visual identity.

That would produce a continuous vertical gallery of gameplay.

## One distinction: most-played versus favorites

I wouldn't combine these concepts entirely.

**Most-played should be automatic. Favorites should be intentional.**

For example, someone might play Conquest 100 times but want Merge pinned at the very top.

You could handle both without introducing another navigation control:

- Tap ⭐ to display games ordered by personal play frequency.
- Long-press a game to toggle its favorite status.
- Favorited games always appear before other games in the ⭐ view.

A small star could briefly appear over the animation when a game is favorited, then disappear.

That keeps the interface almost completely free of text.

On desktop, right-click could perform the same action. For keyboard and controller accessibility, a dedicated favorite action could be available when a game is focused.

## The discovery mode is especially interesting

I wouldn't make ✨ simply another random sorting option.

I'd make it a genuinely different way to explore the collection.

For example, the discovery ordering could favor games that the user hasn't played much, periodically introduce newly released games, and avoid repeatedly displaying games they've just played.

A simple weighted ranking might consider:

$
D(g)=w_1N(g)+w_2R(g)+w_3U(g)
$

Where:

- $N(g)$ represents how new the game is.
- $R(g)$ introduces a random component.
- $U(g)$ favors games the user hasn't tried.

The weights would determine how strongly each factor affects discovery.

This could work entirely offline.

You wouldn't need machine learning, behavioral profiling, recommendation servers, or any other infrastructure.

Just a deterministic local algorithm with a little randomness.

I'd also make the ordering stable during a browsing session. Otherwise, returning to the launcher could unexpectedly rearrange everything.

## What about locked games?

I'd actually avoid putting permanent padlock icons over the previews.

They would interfere with the aesthetic you're describing.

Instead, **every game preview would look identical regardless of subscription status**.

When someone selects an unlocked game, it launches immediately.

When they select a locked game, they encounter the subscription screen.

You could make that screen nearly textless too, although the actual purchase options and legally required disclosures should remain clear.

For instance:

**🔓 All games · $1/month · $10/year**

With two large purchase buttons and a close control.

The normal interface would never need to display prices, advertisements, or promotional language.

And once subscribed, the user would never encounter that screen again while their membership remains valid.

---

## The most compelling technical implication

This is where the idea becomes especially compatible with Ebitengine.

**You don't actually need traditional game thumbnails, names, or catalog descriptions.**

Every game could expose two things:

```go
type Game interface {
    Update() error
    Draw(screen *ebiten.Image)
}

type Preview interface {
    Update()
    Draw(screen *ebiten.Image)
}
```

The launcher simply composes previews into a vertically scrolling canvas.

It doesn't need to know whether a particular preview depicts Snake, Merge, Pong, or a territorial conquest game.

All it needs is a unique internal identifier, a preview renderer, and an entry point.

The user-facing identity is entirely visual.

This also reduces the amount of work involved in adding games. You won't need to invent names, write promotional copy, or create thumbnail assets for each release.

You'd still maintain internal names for persistence, tests, accessibility, and development.

### Accessibility without visual text

It's worth separating *minimal visible text* from *lack of accessible labels*.

Each animated row should have a screen-reader label, such as "Sliding number puzzle" or "Territory conquest strategy."

That allows the interface to be nearly textless without making it inaccessible to people who use assistive technology.

Likewise, the three emoji controls need accessible names, and reduced-motion users should see representative still images.

---

## One design decision I'd change

I would **not put the user's most-played games exclusively under ⭐ and the globally popular games exclusively under 👥**.

I would make those three icons represent *ways to arrange the same complete collection*.

So:

- ⭐ displays the entire collection, ordered by personal preference.
- 👥 displays the entire collection, ordered by global popularity.
- ✨ displays the entire collection in discovery order.

Every game remains accessible from every view.

That preserves the elegant interaction model while preventing games from becoming difficult to find.

For global popularity, you could bundle periodically updated rankings with application releases. That would preserve the zero-telemetry architecture, although the rankings would need an independent source such as voluntary community voting.

---

## The one interaction I'd be particularly careful about

**Long-pressing to favorite could conflict with the natural expectation that tapping immediately launches a game.**

I'd use a small interaction distinction:

- **Tap:** Play.
- **Long press:** Favorite or unfavorite.
- **Vertical swipe:** Browse.
- **Back:** Return to the exact catalog position.

On desktop, clicking launches, scrolling browses, and right-clicking favorites.

Those interactions would make the application feel extremely immediate.

And because you don't have text overlays, almost the entire screen is devoted to the actual gameplay demonstrations.

---

## The central idea

I think this is approaching a particularly coherent design:

**The application doesn't tell you what games are called. It shows you what games are.**

You don't browse brands, descriptions, screenshots, ratings, or marketing materials.

You browse gameplay.

You see something interesting, touch it, and play.

The only persistent navigation is:

# ⭐　👥　✨

Everything else is the game collection itself.

And because the underlying games are extremely simple, you could make their previews unusually vivid, smooth, and expressive without requiring large video files or expensive rendering.

The resulting application could have a recognizable visual identity without depending on the visual identity or trademarks of any individual game.

**I'd preserve that radical simplicity as a design constraint, rather than gradually adding conventional app-store features as the catalog grows.**

---

## User

maybe instead of favorite, long-licking opens a menu with pin, hide, share?

---

## ChatGPT

Yes. I think **long-press → 📌 Pin, 🙈 Hide, ↗️ Share** is much better than using long-press exclusively for favorites.

It gives you a complete game-management interface without adding any permanent UI clutter.

The interaction model becomes:

| Gesture | Action |
|---|---|
| **Tap** | Launch game immediately |
| **Long-press** | Open contextual menu |
| **Swipe vertically** | Browse games |
| **Back** | Return to the same position in the catalog |

### Context menu

Long-pressing a gameplay preview could pause its animation and display a small floating menu:

```text
┌───────────────────────────────┐
│                               │
│      [GAMEPLAY PREVIEW]       │
│                               │
│         ┌───────────┐         │
│         │ 📌  Pin   │         │
│         │ 🙈  Hide  │         │
│         │ ↗️  Share │         │
│         └───────────┘         │
│                               │
└───────────────────────────────┘
```

I'd keep the menu labels. Three short words appearing only when requested wouldn't compromise your nearly textless interface.

### 📌 Pin

Pinned games always appear at the top of the ⭐ view, above games ordered by play frequency.

I'd probably make pinned games appear first in all three views, so their position remains predictable.

Long-pressing a pinned game would show **Unpin**.

### 🙈 Hide

Hidden games disappear from the regular catalog.

This is particularly useful as the collection grows. Someone might enjoy only 15 games out of 200 and want to remove everything else from their browsing experience.

However, there needs to be a way to recover hidden games.

I'd suggest allowing a **long-press on the ✨ icon** to open a small menu containing a hidden-games view.

That avoids adding a fourth persistent navigation icon.

I would also make hiding reversible and local to the device.

### ↗️ Share

This could become your most effective acquisition mechanism.

Sharing a game should produce a link directly to that game's entry in the collection.

For example:

`https://example.org/g/merge`

Someone opening that link could immediately see the animated gameplay preview.

If they have the application installed, the link could launch the game directly.

Otherwise, it could open a minimal web page where they can try the game or install the collection.

**I'd make shared links to the five free games immediately playable in the browser**, with no signup or installation required.

For paid games, the link could still open the animated preview and offer membership options.

That creates organic discovery without social networks, referral tracking, or advertising infrastructure.

---

## One additional option: make the menu extensible

I would initially ship only three actions:

**📌 Pin · 🙈 Hide · ↗️ Share**

But I'd implement the menu as a generic component so you could eventually add actions such as resetting a game or viewing its controls without changing the browsing interface.

For now, three is enough.

### A small UX detail I'd recommend

When someone hides a game, display a brief, nonintrusive **Undo** action.

Accidentally hiding something should never require navigating into a settings menu to recover it.

And the menu should support a secondary click on desktop and an equivalent keyboard/controller action, not only long-press.

## The resulting design

You now have a surprisingly complete interface with only three persistent controls:

**⭐ Most Played · 👥 Popular · ✨ Discover**

Every game is represented by an animated, clickable preview.

The user can tap to play, long-press to organize, and swipe to browse.

Everything else—pinning, hiding, sharing, subscriptions, and preferences—appears only when relevant.

I particularly like that **the interface doesn't require users to learn the names of games**. Even sharing can be based on a visual preview rather than a title.

The one architectural detail I'd establish immediately is to give every game a permanent internal ID. That ID should remain stable across app updates so pins, hidden status, play counts, saves, and shared links continue working even if the game itself never has a visible name.

---

## User

ok lets talk about names for htis

---

## ChatGPT

I think the name should reflect the most distinctive thing about the project: **an ever-growing collection of games, stripped of everything except play.**

The combination is unusual enough to suggest several directions:

- **Demoscene:** Small, elegant programs; playful, experimental, independent.
- **Anti-platform:** No ads, surveillance, accounts, algorithms, or engagement traps.
- **Play as a commons:** Familiar games treated as shared cultural experiences rather than commercial products.
- **Visual discovery:** No titles, descriptions, or search. Just animated gameplay you can immediately enter.
- **Permanence:** Games don't disappear or change. The collection just grows.

I'd avoid names like *GamePass*, *Arcade+*, or *Retro Games*. They sound like conventional commercial products, and your concept is more interesting than that.

## My favorite candidates

| Name | Why I like it | Potential weakness |
|---|---|---|
| **Playground** | Immediately understandable; a place to play without commercial pressure | Common name, likely crowded |
| **Toybox** | A collection of small, self-contained things to play with | Could feel childish |
| **Playthings** | Charming, unpretentious, emphasizes the games as objects | Slightly old-fashioned |
| **Commons** | Games as shared culture, outside commercial enclosures | Doesn't immediately communicate gaming |
| **Playfield** | Evokes the actual interactive space rather than a storefront | Somewhat technical |
| **No Quarter** | Arcade reference, anti-monetization double meaning, punk sensibility | Existing uses; also sounds combative |
| **Freeplay** | Arcade terminology for machines that don't require coins | Could imply everything is free |
| **Bits** | Demoscene heritage, short, memorable | Extremely generic |
| **Pastime** | Relaxed, familiar games without engagement mechanics | Less distinctive |
| **Little Games** | Says exactly what it is, with no marketing pretense | Descriptive rather than ownable |
| **Play Forever** | Longevity, permanence, games that don't disappear | A little promotional |
| **Just Play** | Captures the zero-friction philosophy | Very common phrase |

These are creative candidates, not trademark-cleared names.

## Three directions I find particularly compelling

### 1. **Playground**

This is probably the most immediately intuitive name.

A playground is a place where people go to play. You don't need an account, advertising, progression system, or even instructions.

It also works with your interface concept: the user enters a space full of things happening, sees something appealing, and joins in.

**Playground**  
*Nothing but games.*

The main drawback is that it's so generic that distinguishing it in app stores and trademark searches could be difficult.

### 2. **Freeplay**

This has a nice historical connection to arcades.

A machine in free-play mode lets you play without inserting coins.

Your project effectively takes the modern mobile-game economy and flips the switch to free play.

It also connects nicely to the demoscene, where software is often created for the pleasure of making and sharing it.

**Freeplay**  
*Games without the games.*

I like the double meaning of that tagline, although the $1/month subscription could confuse people who interpret *Freeplay* as promising that everything costs nothing.

### 3. **Commons**

This is my favorite philosophically.

The games you're describing are familiar cultural forms, many of which have been implemented thousands of times.

Instead of inventing new intellectual property around each one, you're creating a shared library of playable mechanics.

**Commons**  
*A place to play.*

It fits the independent, non-extractive philosophy and could develop into an open-source community project.

The downside is that it doesn't immediately communicate what the application does.

---

## Some more unconventional possibilities

Given the visual, demoscene-inspired interface, I also like the idea of a name that's short, distinctive, and slightly playful.

| Name | Association |
|---|---|
| **Blip** | Tiny digital experiences |
| **Bloop** | Playful retro-computing sound |
| **Wiggle** | Motion, interaction, animated previews |
| **Tinker** | Exploration and experimentation |
| **Fidget** | Little things to play with |
| **Dabble** | Casual, low-commitment play |
| **Loop** | Repeating gameplay animations |
| **Sprites** | Classic 2D games and demoscene graphics |
| **Pixels** | Retro digital aesthetic |
| **Playroom** | A welcoming collection of games |
| **Bits & Pieces** | Small programs collected together |
| **Oddments** | A collection of miscellaneous little things |
| **Diversions** | Short, enjoyable interruptions |
| **Arcadia** | Arcade heritage and a playful ideal world |
| **Frolic** | Play for its own sake |

Of these, **Blip, Frolic, Dabble, and Oddments** have the strongest personality.

## One idea I particularly like: **Diversions**

There's something appealing about using an older word for games and amusements.

A diversion is something you do for enjoyment rather than productivity.

It doesn't imply addiction, achievement, competition, or endless engagement.

And it fits a collection that could eventually contain hundreds of completely different games.

**Diversions** sounds like a collection rather than a platform.

It also leaves room for your minimalist interface to do the work of communicating what the product actually is.

---

## My shortlist

If I were narrowing this to five names worth developing, I'd choose:

| Rank | Name | Primary strength |
|---|---|---|
| 1 | **Freeplay** | Strongest connection to arcade history and anti-monetization |
| 2 | **Diversions** | Most distinctive editorial personality |
| 3 | **Playground** | Most immediately understandable |
| 4 | **Blip** | Best fit for a tiny, vibrant, demoscene-inspired application |
| 5 | **Commons** | Strongest connection to the project's philosophy |

There's also a more radical direction worth considering.

Because you're designing an interface with almost no text, **the application name doesn't necessarily need to be a conventional word**.

It could be a symbol, a short invented word, or even a sequence of characters that evokes play and interaction. The name could function more like a demoscene group's signature than a commercial brand.

That might be especially effective if the application is intended to have a distinctive, recognizable visual identity.

### One question

Do you imagine the name having more of a **punk/demoscene/anti-commercial identity** (something like *No Quarter* or *Blip*), or a **friendly, universal identity** (*Playground*, *Toybox*, *Diversions*)?

I think that distinction will determine whether the name feels like a political statement about the gaming industry or simply a delightful place to play games.

---

## User

i think it should be clear that this is targetted at people with autism/adhd. i like the word arcade or puns on that

---

## ChatGPT

I think **Arcade** is a great anchor, especially if you're explicitly designing for autistic and ADHD users.

The concept already has several qualities that could appeal to neurodivergent players: immediate visual communication, predictable gameplay, no advertising interruptions, no manipulative reward systems, and control over what appears in the interface.

I'd explore names that combine **arcade culture, stimming, hyperfocus, sensory enjoyment, and neurodivergent identity**.

There's also a particularly interesting opportunity for wordplay around *arcade*, *arcadia*, and *neurodivergence*.

## 1. My favorite names

| Name | Concept | My assessment |
|---|---|---|
| **Stimcade** | Stim + Arcade | ⭐ Best overall |
| **Neurocade** | Neuro + Arcade | ⭐ Clearest audience association |
| **ArcADHD** | Arcade + ADHD | Clever, explicitly ADHD |
| **Hypercade** | Hyperfocus + Arcade | Strong, energetic, memorable |
| **Auticade** | Autistic + Arcade | Explicitly autistic |
| **Divercade** | Divergent + Arcade | Clever, more subtle |
| **Fidgetcade** | Fidget + Arcade | Playful, approachable |
| **Spincade** | Spinning / stimming + Arcade | Short, animated |
| **Stim City** | Stimming + arcade culture | Fun, but less distinctive |
| **Neuroplay** | Neurodivergence + Play | Clear but conventional |
| **Arcadia** | Utopian place of enjoyment | Beautiful, but less explicitly neurodivergent |
| **InfiniCade** | Infinite + Arcade | Emphasizes the growing collection |

## 2. Stimcade

This might be my strongest recommendation.

# **STIMCADE**
*An arcade for your brain.*

It has several advantages.

**First, it's immediately recognizable as a gaming product.** The *-cade* ending evokes arcade games without requiring the complete word.

**Second, "stim" has a meaningful association with autistic and other neurodivergent communities.** It suggests sensory enjoyment, repetition, and self-directed engagement.

**Third, it's short and memorable.**

**Fourth, it fits the visual language.** Your scrolling wall of vibrant, looping gameplay animations could itself be part of the sensory appeal.

And importantly, Stimcade doesn't necessarily imply that the games are educational, therapeutic, or intended to treat ADHD or autism.

They're just enjoyable games.

The main consideration is that stimming is not synonymous with all autistic or ADHD experiences, and some players may not identify with the term.

## 3. Neurocade

# **NEUROCADE**
*Play your way.*

This is probably the most straightforward name if you want to communicate neurodivergent identity.

It sounds like a real arcade, but the prefix suggests that it's designed with different neurological experiences in mind.

I particularly like it because it doesn't limit the audience to a particular diagnosis.

It could appeal to autistic people, people with ADHD, and others who simply prefer games without sensory overload, interruptions, or manipulative engagement systems.

It also sounds slightly futuristic, which fits your demoscene aesthetic.

The drawback is that "neuro" can sound medical or clinical, depending on the branding.

You could counteract that with a vibrant, playful visual identity.

## 4. Hypercade

# **HYPERCADE**
*Nothing between you and play.*

This is my favorite from a purely branding perspective.

It combines *hyperfocus* with *arcade*, while also evoking hyperactivity and energetic gameplay.

It's punchy, sounds like an established gaming brand, and doesn't explicitly reference a diagnosis.

That last point is both an advantage and a disadvantage.

It could attract a wider audience, but it wouldn't immediately communicate that the project was designed for autistic and ADHD users.

There's also a potential mismatch: *Hypercade* sounds fast, intense, and stimulating, whereas some of your games may be slow, contemplative, or relaxing.

Still, I think it's worth considering.

## 5. ArcADHD

This is the most direct pun.

# **ArcADHD**

It has a certain irreverent quality that I like.

And visually, you could emphasize the embedded ADHD:

**Arc·ADHD**

But there are two drawbacks.

First, it's harder to pronounce naturally.

Second, it centers ADHD rather than autism and the broader neurodivergent community.

It could make an excellent name for an ADHD-specific game collection, but perhaps not the entire project.

---

## 6. Some more experimental names

I'd also explore names that lean into neurodivergent culture without necessarily using clinical terminology.

| Name | Wordplay |
|---|---|
| **AuDHDcade** | AuDHD + Arcade |
| **Stimtopia** | Stim + Utopia |
| **Arcadivergent** | Arcade + Neurodivergent |
| **Neurotopia** | Neuro + Utopia |
| **Dopacade** | Dopamine + Arcade |
| **Dopamine Arcade** | Direct reference to reward and motivation |
| **Patterncade** | Pattern recognition + Arcade |
| **Loopcade** | Loops + Arcade |
| **Infocade** | Information + Arcade |
| **Autopia Arcade** | Autism + Utopia |
| **Stimmersion** | Stim + Immersion |
| **Focuscade** | Focus + Arcade |
| **Special Arcade** | Special interests + Arcade |
| **Parallel Play** | A familiar concept in discussions of autistic social interaction |
| **Sidequest Arcade** | ADHD associations with pursuing tangents and diversions |

**Loopcade** is particularly interesting given your interface.

The app is literally a scrolling collection of looping gameplay animations.

The name would describe both the browsing experience and the appeal of repetitive, predictable mechanics.

It could also be very visually distinctive.

---

## 7. What would make it feel genuinely neurodivergent-friendly?

I think this is an important distinction for the brand.

Rather than just labeling the application as designed for autistic and ADHD users, you could make its design principles obvious.

| Design choice | Potential benefit |
|---|---|
| No advertising | No involuntary interruptions |
| No tracking | Privacy and autonomy |
| No autoplaying audio | Avoids unexpected sensory stimulation |
| Adjustable animation intensity | Accommodates different sensory preferences |
| Persistent favorites and hidden games | Gives users control over their environment |
| Instant gameplay | Reduces friction between intention and action |
| No streaks or daily rewards | Avoids pressure to return |
| No forced tutorials | Lets users explore independently |
| Offline gameplay | Predictable availability |
| Games never unexpectedly change | Preserves familiar experiences |

I would especially emphasize the last one.

**Predictability is a product feature.**

Your append-only development philosophy, in which new games are added without changing existing ones, could be particularly valuable to people who prefer familiar and stable experiences.

And letting users hide or pin games means the application adapts to their preferences without trying to shape their behavior.

That's a meaningful alternative to conventional engagement optimization.

---

## 8. My final shortlist

Taking the audience, visual identity, subscription model, and game-selection interface together:

| Rank | Name | Why |
|---|---|---|
| 🥇 | **Stimcade** | Most distinctive combination of arcade culture and neurodivergent identity |
| 🥈 | **Neurocade** | Most straightforward broad neurodivergent association |
| 🥉 | **Hypercade** | Strongest conventional brand potential |
| 4 | **Loopcade** | Best connection to the animated interface |
| 5 | **Arcadivergent** | Most explicitly references neurodivergence |

These are naming concepts, not trademark or domain-availability checks.

### One more possibility I really like

**STIMCADE — The Infinite Arcade**

That gives you a short primary brand name and a secondary phrase that explains the continually growing collection.

The brand wouldn't need to say "autism" or "ADHD" in every context, but its identity could be explicitly rooted in neurodivergent community values.

And the product itself would reinforce that identity through its design.

**My vote is Stimcade**, with Neurocade as the strongest alternative.

The question I'd explore next is whether you want the name to feel like it's **made by and for neurodivergent people**, or more broadly like **an arcade designed around neurodivergent-friendly principles that everyone can enjoy**. Those are compatible goals, but they suggest slightly different branding.

---

## User

research if stimcade is available on app stores

---

## ChatGPT

I’ll check the exact name “Stimcade” on Apple’s App Store and Google Play, then look for existing uses of the name that could still create a branding conflict even if no app currently uses it.

---

## ChatGPT

## Stimcade - App Store availability research
**Checked: November 7, 2026**

**Good news: I found no indexed mobile game or application named “Stimcade” on either Apple's App Store or Google Play.**

However, I discovered an existing software product using exactly that name, which creates a potential trademark concern.

| Platform | Finding | Assessment |
|---|---|---|
| Apple App Store | No exact-name listing found in indexed search results | 🟢 Promising |
| Google Play | No exact-name listing found in indexed search results | 🟢 Promising |
| Existing software | **StimCADE**, by Schlumberger | 🟡 Potential conflict |
| Trademark history | Documented trademark use and a Mexican registration record | 🟡 Requires investigation |

**Important:** A missing search result doesn't guarantee that an app name is available for registration. Apple and Google can have unpublished or reserved names, and trademark rights are separate from store listing availability.

### The existing StimCADE

Schlumberger developed an engineering application called **StimCADE**, used to design and simulate oil-well stimulation treatments.

The name appears in technical documentation dating to at least 2002, and software listings were still being updated in 2026. ([Software Informer](https://stimcade.software.informer.com/?utm_source=chatgpt.com), [Scribd](https://www.scribd.com/document/598485302/UG-V4-Intl-tr?utm_source=chatgpt.com))

I also found a Mexican trademark record identifying Schlumberger Offshore Services as the owner, with a filing date of September 4, 2002, and a registration date of October 31, 2002. That source does not establish whether the registration remains active today. ([empremexico.com](https://empremexico.com/tecnologia-telecomunicaciones/stimcade/?utm_source=chatgpt.com))

The good news is that petroleum engineering software and casual arcade games are substantially different markets.

That makes consumer confusion less likely than if another gaming company already used the name.

Nevertheless, I wouldn't treat the name as legally cleared.

## What this means for your project

**Stimcade looks promising for app-store availability, but it isn't a completely unused name.**

The existing product is sufficiently different from an arcade game collection that I wouldn't automatically abandon the name. However, using the exact spelling of an established software trademark introduces avoidable uncertainty.

### How to confirm availability

**Apple App Store:** [App Store Connect](https://appstoreconnect.apple.com/?utm_source=chatgpt.com)

Apple allows an app name to be used once per localization. You can test whether **Stimcade** is available by creating an app record. This is more authoritative than searching published listings because names can be reserved before an app launches. ([Apple Developer](https://developer.apple.com/help/app-store-connect/create-an-app-record/add-a-new-app?utm_source=chatgpt.com))

**Google Play:** [Google Play Console](https://play.google.com/console/?utm_source=chatgpt.com)

Google Play identifies applications by their package IDs, so name duplication works differently. An existing app with the same title would not necessarily prevent you from using it, although impersonation and trademark rules still apply.

**US trademarks:** [USPTO Trademark Search](https://tmsearch.uspto.gov/?utm_source=chatgpt.com)

I found evidence of historical trademark use, but I couldn't establish the current US registration status from the accessible records. The USPTO recommends checking potentially confusingly similar marks, not just exact matches. ([USPTO](https://www.uspto.gov/trademarks/search/federal-trademark-searching?utm_source=chatgpt.com))

### My assessment

| Question | Finding |
|---|---|
| Is there an existing arcade game called Stimcade? | None found |
| Is there an indexed iOS app called Stimcade? | None found |
| Is there an indexed Android app called Stimcade? | None found |
| Has the name been used commercially? | **Yes, by Schlumberger** |
| Is it definitely available for registration? | Not established |
| Would I continue considering it? | **Yes, with caution** |

**My recommendation:** Keep **Stimcade** on your shortlist, but verify the US trademark status and attempt to reserve the name in App Store Connect before investing in its visual identity.

The strongest positive finding is that I found no apparent gaming competitor using the name. The strongest negative finding is that you're not starting with an entirely original word.

---

## User

what about hypercade

---

## ChatGPT

I’ll check Hypercade the same way: exact-name app listings, existing gaming brands, and trademark signals. The gaming overlap matters more here than it did for Stimcade.

---

## ChatGPT

## Hypercade - Name Availability Research
**Checked: November 7, 2026**

**I would avoid Hypercade.** Unlike Stimcade, it has already been used for an actual arcade game collection, and there's an active company using the exact name in the gaming industry.

### 1. Existing uses of Hypercade

| Existing use | Description | Concern |
|---|---|---|
| **Hypercade (2016)** | VR arcade collection developed by Hidden Path Entertainment | 🔴 Very high |
| **Hypercade.io** | Active gaming-industry software company | 🔴 High |
| **HYPER ARCADE** | Registered US trademark covering video games | 🟠 High |
| **Hypercade.com** | Domain listed for sale | 🟡 Moderate |
| Apple App Store | No exact-name listing found in indexed searches | 🟢 Promising |
| Google Play | No current exact-name listing found in indexed searches | 🟢 Promising |

### 2. The biggest problem: Hypercade already exists as a game collection

In 2016, Hidden Path Entertainment released **Hypercade**, a collection of four arcade-style games.

It launched on Android/Gear VR and subsequently appeared on Windows and Oculus Go.

Its premise was remarkably similar to yours: several simple arcade games collected into one application.

The original collection included games inspired by classic arcade mechanics, including shooters and action games. ([MobyGames](https://www.mobygames.com/game/134992/hypercade/?utm_source=chatgpt.com), [MobyGames](https://www.mobygames.com/game/134992/hypercade/releases/?utm_source=chatgpt.com))

This is the most significant conflict because it isn't merely an unrelated company using the same name.

**It's an existing product in almost exactly the same category.**

Even though I couldn't establish that it's still commercially available, historical use creates meaningful branding and potential legal concerns.

### 3. There's also an active Hypercade company

The website [hypercade.io](https://hypercade.io/?utm_source=chatgpt.com) currently advertises a platform providing development infrastructure and analytics for independent game studios.

Its services include game development tools, automated testing, release infrastructure, and audience intelligence.

Although it isn't directly selling a consumer arcade collection, the company operates within the gaming industry.

That makes confusion more plausible than with Stimcade's petroleum-engineering software.

### 4. Trademark concerns

I found a particularly relevant registered US trademark:

**HYPER ARCADE**
- Owner: Metaverse Magna Inc.
- USPTO serial number: **98556255**
- Registration number: **8336498**
- Registered: July 7, 2026.
- Covered products include downloadable computer games and mobile gaming software.

The mark is registered for a stylized design incorporating those words, rather than necessarily providing exclusive rights to every use of either word independently. Nevertheless, it is relevant to a similarity analysis. ([Furm](https://furm.com/trademarks/hyper-arcade-98556255?utm_source=chatgpt.com), [Trademarkia](https://www.trademarkia.com/hyper-arcade-98556255?utm_source=chatgpt.com))

You can investigate the registration through [USPTO Trademark Search](https://tmsearch.uspto.gov/?utm_source=chatgpt.com).

### 5. Domain availability

The exact domain **hypercade.com** is advertised for sale through BrandBucket for **$3,995**.

That is an advertised price, not a verified completed sale or guarantee of current availability. ([BrandBucket](https://www.brandbucket.com/names/hypercade?utm_source=chatgpt.com))

Given the existing uses, I wouldn't pay that amount without obtaining trademark advice.

---

## Stimcade vs. Hypercade

| Criterion | Stimcade | Hypercade |
|---|---|---|
| Existing game with exact name | None found | **Yes** |
| Active gaming-industry company | None found | **Yes** |
| Existing trademark concerns | Historical industrial software | Registered similar gaming mark |
| Communicates neurodivergent focus | Strong | Moderate |
| Fits your animated arcade concept | Excellent | Excellent |
| Branding risk | Moderate | **High** |
| My recommendation | Worth investigating | Avoid |

### My recommendation

**Stimcade is substantially more promising than Hypercade.**

The strongest reason to reject Hypercade isn't even the trademark registration. It's that someone has already used the exact name for a collection of simplified arcade games.

That makes it harder to establish an independent identity.

I'd continue exploring names built around **stim**, **arcade**, and **neurodivergent play**, preferably invented words that don't already have a commercial gaming history.

Of the two names researched so far, **Stimcade remains my clear favorite**.

---

## User

ok what else

---

## ChatGPT

I think we should keep exploring **short, invented names that combine arcade culture with autistic/ADHD experiences**, especially stimming, repetition, pattern recognition, hyperfocus, and sensory enjoyment.

Given that **Stimcade** looks promising but has an existing industrial-software use, while **Hypercade** has direct gaming-industry conflicts, I'd prioritize names that are both distinctive and easy to pronounce.

## My strongest new candidates

| Name | Wordplay | Why it works |
|---|---|---|
| **Stimmarcade** | Stimming + Arcade | Explicitly neurodivergent, immediately understandable |
| **Stimmcade** | Stimming + Arcade | A variation on Stimcade with a clearer reference to stimming |
| **Neurocadian** | Neuro + Arcadian | Playful, slightly utopian |
| **Dopacade** | Dopamine + Arcade | Short, catchy, ADHD-adjacent |
| **Fidgetcade** | Fidget + Arcade | Communicates small, satisfying interactions |
| **Loopcade** | Loop + Arcade | Perfect for the looping gameplay previews |
| **Patterncade** | Pattern + Arcade | Emphasizes puzzles and pattern recognition |
| **Stimtropolis** | Stim + Metropolis | A whole world of things to play with |
| **Divercade** | Divergent + Arcade | Subtle neurodivergent identity |
| **Neuroarcade** | Neuro + Arcade | Clear, simple, approachable |
| **Auticade** | Autistic + Arcade | Direct autistic identity |
| **Arcadivergent** | Arcade + Neurodivergent | Explicit community association |
| **Fidgetopia** | Fidget + Utopia | Playful, friendly, suggests a collection |
| **Stimverse** | Stim + Universe | An expanding collection of experiences |
| **Focuscade** | Focus + Arcade | Short, but less explicitly neurodivergent |

Of these, I especially like **Divercade, Stimmcade, Loopcade, and Fidgetcade**.

But I think there's another direction that might be more distinctive.

## Another direction: make the name feel like neurodivergent slang

Instead of combining a clinical-sounding word with *arcade*, we could draw on language familiar within autistic and ADHD communities.

| Name | Inspiration | Personality |
|---|---|---|
| **Stimmies** | Stimming | Cute, friendly, informal |
| **Stimmiverse** | Stimming + universe | Expansive, colorful |
| **Stimnasium** | Stim + gymnasium | A place to play and move |
| **Stimmatic** | Stim + automatic | Retro-tech, arcade-machine feel |
| **Stimbox** | Stim + toybox | Minimalist and approachable |
| **AuDHD Arcade** | Combined autism/ADHD identity | Explicitly community-oriented |
| **Parallel Arcade** | Parallel play | Subtle, thoughtful |
| **Sidequest** | ADHD-associated tangents | Gaming-native, playful |
| **Special Interest Arcade** | Special interests | Direct and affectionate |

I particularly like **Stimnasium**. It suggests a place full of things to do, without the clinical associations of *Neurocade*.

And **Stimmatic** has a wonderful 1980s arcade-machine quality.

## Preliminary name-search findings

I ran exact-name web searches for several of the new candidates.

| Candidate | Initial search result | My interpretation |
|---|---|---|
| **Divercade** | No obvious exact-name gaming product surfaced | Promising |
| **Stimmcade** | No obvious exact-name gaming product surfaced | Promising |
| **Fidgetcade** | No obvious exact-name gaming product surfaced | Promising |
| **Dopacade** | No obvious exact-name gaming product surfaced | Promising |
| **Loopcade** | Existing use of the name in online content | Less distinctive |

These are preliminary web-search results, **not verified availability in Apple App Store Connect, Google Play Console, or trademark databases**.

## My updated top five

| Rank | Name | Why I'd consider it |
|---|---|---|
| 🥇 | **Stimmcade** | Closest to Stimcade, more clearly rooted in stimming |
| 🥈 | **Divercade** | Neurodivergent identity without sounding clinical |
| 🥉 | **Fidgetcade** | Immediately communicates playful, tactile satisfaction |
| 4 | **Stimnasium** | Most original personality; feels like a place |
| 5 | **Dopacade** | Short, memorable, and energetic |

### One interesting distinction

I think **Stimmcade** and **Divercade** communicate two slightly different things.

**Stimmcade** says: *This is an arcade for stimming.*

**Divercade** says: *This is an arcade for neurodivergent people.*

Since your collection will include everything from repetitive puzzles to complicated strategy games, **Divercade might actually be the better umbrella brand**. It doesn't suggest that every game needs to be a sensory or repetitive experience.

On the other hand, **Stimmcade** has a stronger connection to the immediacy and sensory appeal of your animated interface.

I'd narrow the next round to **Stimmcade, Divercade, and Fidgetcade**. Those three seem to best combine an identifiable audience, an arcade association, and a plausible path toward a distinctive brand.

---

## User

what about some play on the word hyperdrive

---

## ChatGPT

I like this direction. **Hyperdrive** has a retro-futurist, demoscene quality, and the *hyper* prefix evokes hyperfocus and ADHD without making the name sound clinical.

The interesting opportunity is combining **hyperdrive + arcade + neurodivergent play**.

### Names built around Hyperdrive

| Name | Wordplay | Impression |
|---|---|---|
| **Hyperdive** | Hyperdrive − r | Diving into hyperfocus; immersive play |
| **Hyperdiver** | Hyperdrive + neurodivergent | Especially interesting for your audience |
| **Hyperdiverse** | Hyper + neurodiverse | Explicitly inclusive, futuristic |
| **Hyperdrome** | Hyper + aerodrome/velodrome | A futuristic arena for games |
| **Hyperloop** | Hyperfocus + looping gameplay | Strong concept, but heavily associated with transportation |
| **Hyperplay** | Hyperfocus + play | Simple and approachable |
| **Hyperverse** | Hyper + universe | Expanding collection of games |
| **Hypercade** | Hyper + arcade | Great sound, but we've already identified naming conflicts |
| **Hyperarcadia** | Hyperdrive + Arcadia | A futuristic world of play |
| **Hyperstim** | Hyperfocus + stimming | Direct neurodivergent association |
| **Stimdrive** | Stimming + hyperdrive | Retro-futuristic and playful |
| **Arcadrive** | Arcade + hyperdrive | A particularly natural arcade pun |
| **Neurodrive** | Neurodivergence + hyperdrive | Futuristic, somewhat technical |
| **Diverdrive** | Neurodivergent + hyperdrive | Distinctive, community-oriented |
| **Hyperdrift** | Hyperdrive + drifting attention | Evokes playful movement and exploration |

## My three favorites

### 1. Hyperdive
**Hyperfocus meets play.**

This is probably the most elegant pun.

It changes just one letter from *Hyperdrive*, turning forward motion into immersion.

It could imply diving into a game, diving into a special interest, or entering a state of hyperfocus.

It also sounds natural when spoken aloud.

The weakness is that it doesn't immediately identify the product as an arcade.

### 2. Hyperdiver
**A play on hyperdrive and neurodivergence.**

This one is particularly interesting because *diver* can suggest both someone who dives deeply into interests and someone who is neurodivergent.

It feels more like an identity than a product category.

The branding could even use a stylized **HYPERDIVER** wordmark with vibrant animated typography.

I think this might be the most promising direction if you want the neurodivergent association to be clever rather than explicit.

### 3. Arcadrive
**Arcade meets hyperdrive.**

This is my favorite if you want people to immediately recognize that it's a gaming product.

It sounds like a fictional arcade machine, a retro console, or a piece of demoscene software.

It also fits an application that's constantly expanding.

The name doesn't explicitly signal autism or ADHD, but the visual identity and messaging could establish that association.

---

## A different possibility: **Hyperdivergent**

This isn't really an arcade pun, but it has a stronger conceptual identity.

It combines hyperfocus, hyperactivity, and neurodivergence into something that sounds almost like a movement.

**HYPERDIVERGENT**

*An arcade for different minds.*

It has a slightly rebellious quality that fits your ad-free, surveillance-free philosophy.

The downside is that it's longer and less immediately playful.

---

## My ranking

| Rank | Name | Strength |
|---|---|---|
| 🥇 | **Hyperdive** | Best wordplay and strongest sound |
| 🥈 | **Arcadrive** | Best connection to arcade gaming |
| 🥉 | **Hyperdiver** | Best subtle neurodivergent association |
| 4 | **Stimdrive** | Best combination of stimming and retro-futurism |
| 5 | **Hyperdivergent** | Strongest identity and mission |

**I would explore Hyperdive and Arcadrive first.**

There's also a visual opportunity here: a stylized arcade cabinet or hyperspace tunnel as the application icon, with the actual interface remaining nothing but scrolling, colorful gameplay loops.

One caveat: these are creative suggestions, not availability-checked names. In particular, anything beginning with *Hyper-* deserves a trademark search given how heavily that prefix is used in gaming and technology.

---

## User

what about something esoteric from neuromancer which is about to be super famous

like what about some play on deck

---

## ChatGPT

“Deck” could be a stronger direction: it connects Neuromancer’s cyberdeck imagery to a portable collection of games. I’ll check the adaptation context and explore names that evoke that world without borrowing protected character or product names.

---

## ChatGPT

Oh, **deck** is a much more interesting direction. It connects several things about your project simultaneously:

- **Neuromancer:** The cyberdeck as a personal interface to cyberspace.
- **Neurodivergence:** The association between *neuro*, computers, and alternative ways of interacting with information.
- **Arcade culture:** The idea of a portable collection of games.
- **Your interface:** A scrolling deck of animated experiences that you select and enter.
- **Demoscene culture:** Small programs, hardware constraints, and independent software.

And the timing is particularly good. Apple's *Neuromancer* adaptation premieres **January 22, 2027**, so the cyberdeck aesthetic may become much more culturally recognizable. ([Apple](https://www.apple.com/tv-pr/news/2026/07/apple-tvs-highly-anticipated-cyberpunk-series-neuromancer-hacks-into-san-diego-comic-con-with-first-ever-teaser/?utm_source=chatgpt.com))

What I particularly like is that *deck* doesn't need to be an explicit reference to Neuromancer. It's already a generic computing and gaming term.

## 1. My favorite deck-based names

| Name | Wordplay | My assessment |
|---|---|---|
| **Stimdeck** | Stimming + Cyberdeck | ⭐ Best connection to your original concept |
| **Diverdeck** | Neurodivergent + Cyberdeck | ⭐ Distinctive and community-oriented |
| **Arcadeck** | Arcade + Deck | ⭐ Best pure gaming pun |
| **Deckade** | Deck + Arcade | ⭐ Clever, retro, memorable |
| **Deckdrive** | Cyberdeck + Hyperdrive | Futuristic, energetic |
| **Stimdrive** | Stimming + Hyperdrive | Strong alternative without "deck" |
| **Hyperdeck** | Hyperfocus + Cyberdeck | Excellent sound, but already used commercially |
| **Neurodeck** | Neurodivergence + Cyberdeck | Already a video game |
| **Playdeck** | Play + Deck | Already used for gaming products |
| **Loopdeck** | Looping animations + Cyberdeck | Particularly suited to your interface |
| **Deckspace** | Cyberdeck + Cyberspace | Immersive, science-fiction aesthetic |
| **Deckwave** | Cyberdeck + New Wave | Retro-futuristic, demoscene feel |
| **Deckshift** | Cyberdeck + Shifting attention | Playful, technological |
| **Stimmatrix** | Stimming + Matrix | Direct Gibson-era cyberpunk association |
| **Neurodrive** | Neurodivergence + Hyperdrive | Strong science-fiction feeling |

These are naming candidates, not trademark-cleared options.

## 2. Stimdeck

# STIMDECK

This is probably my favorite so far across all our naming discussions.

It takes the strongest feature of *Stimcade*-its explicit association with stimming-and combines it with the cyberdeck concept.

It's also interesting because it sounds like a device rather than a conventional gaming service.

That fits your project.

You're not really building a game storefront. You're creating a personal interface containing a collection of small, interactive experiences.

A user opens Stimdeck, browses animated programs, and enters whichever one appeals to them.

The name feels natural for a Go/Ebitengine application that runs on everything from Linux computers to Android phones.

**Stimdeck also avoids the implication that every game must be an arcade-style action game.** It could contain puzzles, strategy games, simulations, and other interactive experiences.

## 3. Diverdeck

# DIVERDECK

This is a close second.

It's a contraction of *divergent* and *deck*, but it also suggests diving into experiences.

I like that it subtly communicates the neurodivergent audience without requiring someone to recognize the word *stim*.

It also has a distinctive cyberpunk sound.

Imagine:

**DIVERDECK**

And beneath that, a dense wall of brightly colored gameplay animations.

The wordmark could be quite minimal.

The downside is that the neurodivergent reference isn't immediately obvious to everyone.

But that might be acceptable if your audience discovers the application's identity through its community and design.

## 4. Arcadeck

# ARCADECK

This might be the cleverest pure wordplay.

It's simply:

**ARCADE + DECK → ARCADECK**

It sounds like the name of a fictional 1980s portable game console.

And because the application is a collection of games, the meaning is immediately relevant.

It doesn't explicitly reference autism or ADHD, but I don't think that necessarily disqualifies it.

The product's identity can be neurodivergent even if the name isn't.

In fact, Arcadeck may be more accessible to people outside the community, while still providing the desired cyberdeck association.

## 5. Deckade

# DECKADE

I find this one particularly interesting.

It combines *deck* with the sound of *arcade*, while also suggesting *decade*.

That gives it a retro-computing association without explicitly describing the games as retro.

It's short, unusual, and easy to pronounce.

The main drawback is that people might initially read it as a stylization of "decade" rather than an arcade reference.

But that could also make it more distinctive as a brand.

---

## 6. A few names I'd rule out immediately

I checked some of the most obvious options against existing products.

| Name | Existing use | Assessment |
|---|---|---|
| **HyperDeck** | Blackmagic Design's professional video recorders | ❌ Avoid |
| **Neurodeck** | Existing psychological deckbuilding video game | ❌ Avoid |
| **PlayDeck** | Gaming platform and an App Store party-game collection | ❌ Avoid |

Blackmagic has an established HyperDeck hardware product line. ([Blackmagic Design](https://www.blackmagicdesign.com/products/hyperdeckstudio?utm_source=chatgpt.com))

*Neurodeck* has been sold on Steam since 2021, making it particularly problematic for another gaming product. ([Steam Store](https://store.steampowered.com/app/1253860/Neurodeck/?utm_source=chatgpt.com))

*PlayDeck* is also already used in gaming, including an application that bundles multiple casual games. ([PlayDeck](https://www.playdeck.io/?utm_source=chatgpt.com), [App Store](https://apps.apple.com/us/app/playdeck-party-games/id6810917226?utm_source=chatgpt.com))

I haven't established trademark or app-store availability for Stimdeck, Diverdeck, Arcadeck, or Deckade.

---

## 7. There's a more esoteric Neuromancer reference worth considering

In *Neuromancer*, a cyberdeck isn't just a computer.

It's an interface through which a console cowboy enters another environment.

That's remarkably similar to what you're proposing.

Your launcher is a collection of animated windows into different interactive worlds. Selecting one takes you inside.

In the novel, this is associated with *jacking in*, *punching deck*, and entering cyberspace. ([SuperSummary](https://www.supersummary.com/neuromancer/index-of-terms/?utm_source=chatgpt.com))

That suggests some additional possibilities:

| Name | Inspiration |
|---|---|
| **Jackdeck** | Jacking into a cyberdeck |
| **Deckjack** | A connection between operator and machine |
| **Joyjack** | Joystick + neural jack |
| **Stimjack** | Stimming + jacking in |
| **Deckrunner** | Console cowboys and cyberspace runs |
| **Softdeck** | Software + cyberdeck |
| **Simdeck** | Simulation + cyberdeck |
| **Simstim** | The fictional sensory-experience technology from Neuromancer |

I would be more cautious with **Simstim**, because it's a distinctive fictional term associated with Gibson's work.

But **Simdeck** is appealing. It sounds like a personal simulation device, and it could encompass your entire growing collection of games.

---

## My updated ranking

Considering your audience, visual interface, subscription model, and cyberpunk inspiration:

| Rank | Name | Best quality |
|---|---|---|
| 🥇 | **Stimdeck** | Combines neurodivergent identity and cyberdeck culture |
| 🥈 | **Arcadeck** | Most elegant gaming/cyberdeck pun |
| 🥉 | **Diverdeck** | Strong neurodivergent identity |
| 4 | **Deckade** | Most distinctive and playful |
| 5 | **Simdeck** | Broadest potential beyond conventional games |

### One thought about Stimdeck versus Stimcade

I think **Stimdeck is a better name than Stimcade** for what you've actually designed.

*Stimcade* sounds like a place containing arcade machines.

*Stimdeck* sounds like a personal device that contains interactive experiences.

Your application is meant to work across phones, desktops, tablets, and browsers, and the interface is essentially a personal, customizable collection of animated programs.

That's much closer to a cyberdeck than a traditional arcade.

It also lets you retain a subtle association with *Neuromancer* without using a recognizable character name or directly branding the product as part of that fictional universe.

**My strongest candidates are Stimdeck and Arcadeck.** Stimdeck communicates who it's designed for; Arcadeck communicates what it does.

And given that you want the entire interface to be almost completely textless, I think either could support a distinctive visual identity without needing a subtitle or explanation.

---

## User

think more like Automatic Jack

---

## ChatGPT

Ah, **Automatic Jack**. That's a much better reference point for the aesthetic you're after.

Not *Neurodeck* or *Stimcade*, where two words are mashed together to describe the product. You're looking for something that sounds like it came out of the Sprawl: **a little mysterious, casually technological, slightly dangerous, and full of personality.**

Automatic Jack is actually from Gibson's *Burning Chrome*, rather than *Neuromancer* itself. He's Bobby Quine's hardware specialist, with a myoelectric prosthetic arm. It's particularly apt because he's a maker and tinkerer who exists outside corporate systems. ([MITH Archive](https://archive.mith.umd.edu/digitalstorytelling/wp-content/uploads/GibsonW_Burning_Chrome.pdf?utm_source=chatgpt.com))

And the *Automatic Jack* naming pattern is excellent: an ordinary word paired with a technological or mechanical one, producing something that sounds like underground software, a pirate radio station, or a forgotten arcade.

## Names in the spirit of Automatic Jack

| Name | Association | Character |
|---|---|---|
| **Automatic Joy** | Automatic Jack + immediate gratification | Playful, effortlessly cyberpunk |
| **Automatic Play** | Games without friction | Minimalist, confident |
| **Automatic Arcade** | Automatic Jack + arcade | Direct, memorable |
| **Electric Jack** | Cybernetic hardware + jacking in | Classic cyberpunk |
| **Soft Jack** | Software + neural jack | Subtle, slightly strange |
| **Jolly Roger** | Pirate culture + playful rebellion | Underground, irreverent |
| **Lucky Jack** | Chance, games, a cyberpunk character | Friendly, retro |
| **Electric Playground** | Cyberpunk + play | Expressive, accessible |
| **Soft Machine** | Software, machines, 1960s counterculture | Literary, experimental |
| **Idle Circuit** | Low-pressure play + computing | Demoscene aesthetic |
| **Automatic Dream** | Cyberpunk + self-directed experience | Atmospheric |
| **Chrome Playground** | Burning Chrome + games | Strong Gibson association |
| **Electric Diversion** | Old-fashioned amusement + technology | Unusual, elegant |
| **Little Matrix** | Small interactive worlds | Cyberpunk, but more recognizable |
| **Jackpot Circuit** | Jacking in + arcade games | Playful and technological |

These are creative directions, not availability-checked names.

## More esoteric, less obviously gaming-related

These are closer to the spirit of Gibson's fictional brands, nightclub names, and hacker aliases.

| Name | Impression |
|---|---|
| **Gentleman Loser** | Direct reference to the bar in *Burning Chrome* |
| **Blue Ant** | Mysterious corporate name from Gibson's later fiction |
| **Memory Palace** | Information, cognition, exploration |
| **Soft Animal** | Organic meets digital |
| **Chrome Garden** | Synthetic nature, interactive environments |
| **Little Waldo** | Remote manipulation, cybernetics |
| **Electric Ghost** | Digital experiences without permanence or friction |
| **Lucky Circuit** | Arcade chance + hardware |
| **Neon Habit** | Repetition, sensory pleasure, cyberpunk |
| **Soft Arcade** | Gentle interaction rather than overstimulation |
| **Ghost Circuit** | Invisible machinery and electronic play |
| **Automatic Garden** | A growing collection of little experiences |
| **Pocket Sprawl** | A whole city of experiences in your pocket |
| **Parallel Jack** | Parallel play + neural interface |

I particularly like **Parallel Jack** for your audience. It quietly combines the autistic-community association of parallel play with cyberpunk's neural-jack terminology.

It doesn't explain itself, which is part of its appeal.

## A few that feel especially right

### 1. Automatic Joy

This might be the strongest.

It sounds like a lost piece of 1980s cyberpunk software.

There's also a nice conceptual contrast: the app doesn't manipulate users into pursuing rewards. It simply provides things they enjoy.

**AUTOMATIC JOY**

No subtitle required.

### 2. Parallel Jack

This has more personality than Stimdeck.

It's simultaneously a reference to parallel play, computing, and jacking into cyberspace.

**PARALLEL JACK**

It also sounds like a character who could exist in *Burning Chrome*.

### 3. Soft Arcade

I really like how this sounds.

It's simple, but it suggests a completely different kind of arcade experience: one without pressure, exploitation, or aggressive stimulation.

**SOFT ARCADE**

The contrast between *soft* and the frenetic reputation of arcades could be an appealing part of the identity.

### 4. Automatic Garden

This is less obviously a game collection, but it captures the append-only aspect beautifully.

A garden grows. New things appear. Existing things remain.

It also feels like the name of an obscure experimental computer program from 1987.

### 5. Neon Habit

This is probably the most visually evocative.

**NEON HABIT**

It connects naturally to your interface full of brightly colored, looping animations.

However, *habit* could imply compulsive engagement, which might conflict with your rejection of attention-extraction mechanics.

---

## One more direction: fictional hardware names

Gibson's **Ono-Sendai Cyberspace 7** is a good example of the naming style: specific, slightly inscrutable, and evocative of a consumer electronics product from an alternate future.

You could borrow that *style* without borrowing the fictional brand.

Some possibilities:

- **Jack Seven**
- **Soft Circuit**
- **Automatic Zero**
- **Pocket Terminal**
- **Parallel Machine**
- **Little Console**
- **Neon Terminal**
- **Electric Deck**
- **Joy Terminal**
- **Soft Interface**

Of these, **Automatic Zero** and **Joy Terminal** stand out to me.

They sound like real artifacts from a fictional technological culture rather than names invented by a marketing department.

## My top five

| Rank | Name | Why |
|---|---|---|
| 1 | **Automatic Joy** | Closest to the feel of Automatic Jack |
| 2 | **Parallel Jack** | Clever neurodivergent association, cyberpunk personality |
| 3 | **Soft Arcade** | Beautiful contrast with exploitative gaming |
| 4 | **Automatic Zero** | Most mysterious and demoscene-like |
| 5 | **Joy Terminal** | Sounds like a fictional retro-futurist device |

**I think Automatic Joy and Parallel Jack are much closer to what you're looking for than any of the earlier portmanteaus.**

They don't try to explain the product. They give it an identity.

And since the application itself is deliberately almost textless, a name that feels like an alias or mysterious piece of software may be more appropriate than one that literally describes an arcade.

---

## User

what about a nonbinary gentleman loser

---

## ChatGPT

Oh, I like this direction considerably more. **The Gentleman Loser** from *Neuromancer* has exactly the kind of underground, cyberpunk, anti-establishment personality you're looking for. Making it nonbinary adds another layer of queer irreverence.

The trick is finding something that preserves the wonderfully contradictory combination of *gentleman* and *loser* without sounding like an awkward gender-neutral substitution.

Here are my favorites:

| Name | Character |
|---|---|
| **Gentle Loser** | Elegant, genderless, understated. My favorite. |
| **Gentlethem Loser** | Explicitly nonbinary, irreverent, a little absurd |
| **The Gentle Loser** | Sounds like an underground arcade or cyberpunk bar |
| **Gentlefolk Loser** | Queer, communal, slightly whimsical |
| **Gentlethem's Arcade** | A playful reinterpretation of the original |
| **Gentle Losers** | Sounds like a queer hacker collective |
| **Them Losers** | Punk, funny, unapologetic |
| **Gentle Them** | Short, clever, unmistakably queer-coded |
| **Loser Arcade** | More immediately recognizable as a gaming product |
| **Beautiful Loser** | Countercultural, affectionate, and memorable |

### My three favorites

**1. Gentle Loser**

This is particularly compelling because dropping *man* doesn't require replacing it with anything.

It preserves the apparent contradiction between gentleness and losing.

And for a collection of games that explicitly rejects exploitative engagement mechanics, there's something appealing about embracing *loser* as an identity.

You don't have to win. You don't have to compete. You don't have to optimize your behavior.

You can just play.

**2. Gentlethem Loser**

This is much more overtly queer.

It's almost a joke about the awkwardness of making traditionally masculine terms gender-neutral, but that's part of its appeal.

It sounds like the name of a queer cyberpunk collective or an obscure demoscene group.

**3. Gentle Losers**

This might actually work best as an organization or project name.

It sounds like a community that produces little games for people who don't particularly care about winning.

You could have:

**GENTLE LOSERS**  
*An arcade for the rest of us.*

That feels like a subcultural identity rather than a commercial gaming brand.

---

### One interesting possibility

What if **Gentle Losers** is the name of the collective, and **Automatic Joy** is the name of its arcade?

That would give you something reminiscent of the demoscene's distinction between the group producing software and the productions themselves.

It also lets you embrace the Gibson reference without making the arcade's name too derivative.

But if we're choosing a single name, **Gentle Loser** is my favorite so far.

It feels queer, cyberpunk, affectionate, and anti-competitive without needing to explain any of those things.

And unlike *Stimcade*, it doesn't sound like you're marketing a specialized product to autistic people. It sounds like something made **by people who share a particular sensibility, for anyone who wants to participate.**

One caveat: these are creative candidates, not trademark-cleared names.

---

## User

i feel like you're just making a lot of portmanteaus instead of really understanding the semiotics of these terms and concepts

---

## ChatGPT

You're right. I've been treating the names as combinations of keywords rather than understanding why names like **Automatic Jack** and **Gentleman Loser** work in the first place.

The important thing about Gibson's naming isn't the cyberpunk vocabulary. It's the **social world implied by the names**.

*Automatic Jack* sounds like a person who acquired a nickname through some combination of reputation, technology, and personal history. The name doesn't describe his function to a consumer. It suggests that he exists independently of the reader, in a world with its own culture.

*Gentleman Loser* is even more interesting. It's the name of a bar, but it carries an entire social philosophy. *Gentleman* invokes respectability, class, manners, and masculine social status. *Loser* negates the conventional measures of success on which that status depends. Together, they suggest a place where people who have rejected—or been rejected by—the dominant status hierarchy can nevertheless possess dignity, sophistication, and community.

And because it's a bar, the name also identifies a **third place**. It's somewhere people gather outside institutions, workplaces, and domestic life. It doesn't need to explain what those people do there.

That's much closer to what you're describing than *Neurocade*.

## The semiotics of your arcade

I think there are several tensions that make the project culturally interesting.

**Play versus productivity.** Games are activities undertaken for their own sake. Yet modern gaming platforms increasingly measure play in terms of retention, engagement, conversion, and monetization.

**Neurodivergence versus normalization.** You're not necessarily proposing games that accommodate people who deviate from a presumed normal player. You're proposing an environment built around different assumptions about what people want from software: repetition, predictability, autonomy, intensity on their own terms, and the freedom to disengage.

**The arcade versus the platform.** An arcade is a place containing machines that do things. A platform is an institution mediating relationships between users, developers, advertisers, and owners. Your project uses contemporary distribution infrastructure to recreate something closer to the former.

**The loser versus the gamer.** The commercial gaming industry is obsessed with winning, leveling, achievement, performance, and competitive status. There's something genuinely subversive about an arcade that doesn't care whether you're good at its games.

**The cyberdeck versus the smartphone.** Gibson's deck represents an instrument through which an individual exercises agency within a networked world. The contemporary smartphone, despite being vastly more powerful, is frequently designed as a mechanism through which institutions exercise agency over its owner.

Your project is, in a small way, about reclaiming the computer as a personal instrument.

And this is why *Automatic Jack* is such a useful reference. It belongs to a world of people who know how their machines work and use them for their own purposes.

## The question of gender in Gentleman Loser

I also misunderstood your suggestion of a nonbinary Gentleman Loser.

I treated *gentleman* as a word requiring a gender-neutral replacement.

But the masculine coding is part of the original name's meaning. Merely replacing *man* with *them* doesn't preserve the relationship between gender, class, respectability, and exclusion.

A nonbinary reinterpretation might instead ask:

**What does the Gentleman Loser become when neither conventional gender nor conventional success confers social legitimacy?**

That isn't necessarily answered by a gender-neutral synonym for gentleman.

It might be answered by an entirely different social archetype.

Consider these—not as a shortlist of product names, but as distinct cultural figures:

| Figure | Social meaning |
|---|---|
| **The Amateur** | Someone who does something for love rather than professional compensation |
| **The Dilettante** | Someone who pursues interests without submitting to professional specialization |
| **The Fool** | Someone whose apparent foolishness permits truths that respectable people cannot express |
| **The Misfit** | Someone whose difference is defined by their relationship to social expectations |
| **The Eccentric** | Someone whose independence appears strange from the perspective of convention |
| **The Idler** | Someone who refuses to make productivity the measure of a worthwhile life |
| **The Flâneur** | Someone who wanders, observes, and experiences without an instrumental destination |
| **The Jester** | Someone who uses play to invert established hierarchies |

These have very different implications.

*Misfit* still defines the person in relation to a norm.

*Eccentric* suggests someone who possesses enough social independence to disregard that norm.

*Amateur* is particularly interesting because its etymology concerns love. The amateur does something because they love doing it.

*Idler* rejects the demand that every moment be economically productive.

And *Fool* has an especially rich relationship to games, social hierarchy, and the subversion of authority.

## What I think we're actually looking for

Not a name that says:

*This is an arcade for neurodivergent people.*

But a name that **expresses a cultural attitude toward play that neurodivergent people might recognize as their own**.

There's an enormous difference.

One approach markets neurodivergence as an identity category. The other creates something from within a particular cultural experience.

Likewise, the name shouldn't necessarily advertise the absence of surveillance, advertising, or subscriptions. Those are characteristics of the product, but they aren't its cultural meaning.

*Gentleman Loser* doesn't advertise cheap drinks or a welcoming atmosphere. It names the kind of world the bar belongs to.

That's the level at which I think we should work.

And there's one additional point about Gibson: his fictional names are effective partly because they aren't all constructed from futuristic terminology. He juxtaposes ordinary, sometimes archaic social language with technological environments. The result is a world where technology is culturally assimilated rather than perpetually novel.

**I'd like to explore a name that sounds like an actual institution, gathering place, or subcultural artifact from that world—not a brand trying to look cyberpunk.**

The most productive question may be whether your arcade is conceptually **a place for outsiders to gather, a machine for personal freedom, or an institution dedicated to the intrinsic value of play**.

Those three ideas could lead to very different names, and all three are more interesting than another variation on *-cade*.

---

## User

ok so what are some ideas

---

## ChatGPT

I think we should look for names that sound like **a bar in the Sprawl, an underground computer club, a defunct amusement company, or a place where people go to do something that doesn't need to be justified**.

The names should have a cultural life beyond the product. They shouldn't explain themselves too readily.

Here are the directions I find most interesting.

## 1. The Gentleman Loser tradition
*Respectability, failure, gender, and affectionate defiance.*

| Name | Semiotics |
|---|---|
| **The Idle Gentleman** | Aristocratic leisure detached from productivity; gendered respectability becomes faintly ridiculous. |
| **The Happy Amateur** | Someone who plays for love rather than competence, prestige, or money. |
| **The Lucky Loser** | Fortune and failure coexist; winning isn't a prerequisite for belonging. |
| **The Velvet Fool** | Softness, theatricality, queerness, and the fool's freedom to ignore convention. |
| **The Elegant Disaster** | Dignity without conformity or competence. |
| **The Amateur Hour** | Reclaims a phrase normally used to dismiss unprofessional work. |
| **The Leisure Class** | A pointed reference to Veblen, reversing the class politics of leisure. |
| **The Beautiful Losers** | Outsider solidarity, though already strongly associated with Leonard Cohen's novel. |
| **The Minor Aristocracy** | An invented social class whose distinctions are deliberately insignificant. |

**The Amateur Hour** is particularly compelling.

It sounds like a late-night radio program, an underground television broadcast, or a recurring gathering at a bar.

It also asserts that expertise, productivity, and commercial success aren't prerequisites for participation.

And it's especially appropriate for a collection of small games developed in the demoscene tradition.

## 2. The cyberdeck tradition
*Personal technology, illicit autonomy, obsolete futures.*

These names evoke the social world of cyberpunk without depending on recognizable Gibson terminology.

| Name | Semiotics |
|---|---|
| **Soft Terminal** | A personal interface rather than an institutional platform. |
| **Pleasure Circuit** | Technology devoted to enjoyment rather than productivity. |
| **The Idle Machine** | A computer explicitly refusing to work. |
| **Automatic Leisure** | The technological automation of something that shouldn't need automating. |
| **Ghost Terminal** | An independent machine existing outside surveillance infrastructure. |
| **Public Access** | Community computing, public-access television, and communications outside corporate gatekeepers. |
| **Pleasure Console** | A machine built for the user's experience rather than someone else's purposes. |
| **The Last Terminal** | A surviving artifact from a different technological future. |
| **Soft Authority** | A contradiction between gentleness and control. |
| **Unauthorized Joy** | Enjoyment occurring without institutional permission. |

I particularly like **Public Access**.

It invokes the world of public-access television, community computer labs, pirate broadcasting, and independent media.

It's also an interesting counterpoint to contemporary digital platforms, where access is increasingly conditional on surrendering personal information.

**Unauthorized Joy** has a different kind of appeal. It's almost a political slogan, but it could also be the name of a fictional nightclub or underground arcade.

## 3. The neurodivergent experience, without clinical language
*Repetition, attention, intensity, divergence, and independent pleasure.*

Here I'd avoid diagnostic references and instead explore cultural experiences that can resonate with autistic and ADHD communities without suggesting they're universal.

| Name | Semiotics |
|---|---|
| **Parallel Play** | Independent activity in shared space; a direct reference to a recognizable social concept. |
| **Happy Accidents** | Unexpected discoveries become valuable rather than errors. |
| **The Same Again** | Repetition as pleasure rather than deficiency. |
| **Special Interest** | Reclaims terminology associated with autistic experience. |
| **Minor Obsessions** | Affectionate treatment of intense interests. |
| **The Pattern Club** | Recognition, repetition, order, and community. |
| **Pleasant Distractions** | Refuses to pathologize wandering attention. |
| **The Quiet Riot** | A contradiction between sensory calm and social rebellion. |
| **Perfectly Strange** | Difference without a demand for correction. |
| **The Endless Particular** | Finding extraordinary meaning in small details. |

**The Same Again** is one of my favorites in this category.

There is something wonderfully defiant about treating repetition as an end in itself.

Modern software is obsessed with novelty and constant change. You're proposing a collection where people can return to familiar experiences indefinitely, knowing they haven't been altered.

It also sounds like something someone might say at a bar when ordering another drink.

**Minor Obsessions** is another strong candidate. It captures the intimate intensity of seemingly trivial interests without pathologizing them.

## 4. The arcade as an institution
*An establishment with its own culture, rituals, and regulars.*

These sound like places that could have existed before the application.

| Name | Semiotics |
|---|---|
| **The Penny Arcade** | Historic amusement culture; inexpensive, accessible entertainment. |
| **The Pleasure Palace** | Old amusement halls, theatrical excess, and playful decadence. |
| **The Little Palace** | A grand institution devoted to modest pleasures. |
| **The Electric Social** | Early electrical entertainment and working-class social clubs. |
| **The Idle Arcade** | A place specifically for unproductive activity. |
| **The Amateur Society** | A collective organized around love rather than professional status. |
| **The Leisure Bureau** | An absurd government office administering enjoyment. |
| **The Department of Play** | Bureaucratic authority repurposed for pleasure. |
| **The Arcade Social** | A gathering place rather than a commercial gaming platform. |
| **The Permanent Vacation** | An institution existing outside ordinary working time. |

I think **The Leisure Bureau** is especially interesting.

It sounds like an institution from a slightly absurd alternate history.

The name borrows the authority of bureaucracy and redirects it toward something bureaucracy normally struggles to accommodate: aimless enjoyment.

It also creates a very different aesthetic possibility from conventional arcade branding.

Imagine a stark institutional logo, with the application itself opening into an explosion of vibrant, looping gameplay.

The contrast could be beautiful.

## 5. Names that sound like fictional people

This is probably closest to *Automatic Jack*.

Names that might belong to a hacker, a bartender, a forgotten arcade operator, or a demoscene personality.

| Name | Character |
|---|---|
| **Lucky Nobody** | Fortune without status. |
| **Automatic Fool** | A technological trickster who refuses conventional intelligence. |
| **Electric Amateur** | Someone who loves machines without needing institutional authority. |
| **Little Trouble** | Affectionate rebellion on a manageable scale. |
| **Soft Criminal** | Illegitimacy without violence; the crime of refusing conventional expectations. |
| **The Last Amateur** | Someone still making things for love in a fully commercialized world. |
| **Idle Jack** | A cyberpunk character who has opted out of the productivity economy. |
| **Lucky Stranger** | An unknown person who belongs wherever they arrive. |
| **The Gentle Outlaw** | Rebellion without domination. |
| **Ordinary Magic** | Everyday technology experienced as something wondrous. |

**Lucky Nobody** really stands out to me.

It's short, memorable, and has the contradictory quality that makes *Gentleman Loser* compelling.

*Nobody* implies a lack of social importance. *Lucky* implies someone who has somehow escaped the consequences.

It sounds like someone who could be sitting at the bar in the Gentleman Loser.

---

## My actual shortlist

Taking all these directions together, I'd narrow it to seven.

| Name | Why I think it works |
|---|---|
| **The Amateur Hour** | Demoscene culture, noncommercial creativity, outsider community |
| **Lucky Nobody** | Strong character, affectionate anti-status identity |
| **The Same Again** | Repetition, predictability, and neurodivergent pleasure |
| **Public Access** | Anti-platform politics, technological autonomy, community |
| **Minor Obsessions** | Neurodivergent cultural resonance without clinical framing |
| **The Leisure Bureau** | Distinctive institutional fiction, ironic authority |
| **Unauthorized Joy** | The project's political philosophy expressed as a place or experience |

### The three I'd explore most seriously

**Lucky Nobody** is closest to the literary quality of *Automatic Jack* and *Gentleman Loser*. It implies a person and a social milieu without explaining either.

**The Amateur Hour** is closest to the culture of the project. It captures the idea of making small, elegant programs for the pleasure of making and using them.

**The Leisure Bureau** is potentially the strongest brand. It implies an entire fictional institution, with room for a distinctive visual identity and recurring releases.

But there's another possibility worth considering: **The Same Again**.

A neurodivergent-oriented arcade built around familiar games, stable rules, predictable experiences, and the freedom to repeat them endlessly could hardly have a more appropriate name.

It doesn't mention games, computers, autism, ADHD, or cyberpunk. Yet it expresses something fundamental about why the project exists.

And that, I think, is closer to the semiotic territory you were pointing toward with *Automatic Jack*.

---

## User

imagine a version of his denisovan embassy that sits between his universes and ours

---

## ChatGPT

Now that's a much more interesting framing.

You're imagining something like **the Denisovan Embassy from William Gibson's *Agency*, except that it exists at the intersection of Gibson's fictional universes and our own reality.**

Not a cyberpunk-themed arcade. An establishment that feels as though it somehow crossed over from one of those universes and opened for business here.

The Denisovan Embassy is a particularly good reference because its name contains a small, deliberate impossibility. Denisovans are an extinct human lineage. An embassy implies an organized political entity, diplomatic recognition, and a relationship between distinct societies.

Together, those words imply a world that the author doesn't have to explain.

The Guardian's review of *Agency* identifies the Denisovan Embassy as an example of Gibson's ability to suggest an entire social reality through apparently incidental details. ([The Guardian](https://www.theguardian.com/books/2020/jan/22/agency-william-gibson-review?utm_source=chatgpt.com))

And there's something especially appropriate about **an embassy** as your model.

An embassy is a little piece of one political world operating inside another. It has its own customs, affiliations, protocols, and relationship to authority.

Your arcade could be imagined as just such an enclave: a place where the prevailing rules of the commercial internet don't apply.

No surveillance. No advertising. No behavioral manipulation. No mandatory social performance.

Not because someone has made these things illegal, but because *that's simply not how this place operates*.

## The naming principle: an institution from an impossible civilization

I'd look for two ordinary concepts that become strange when placed together.

The first establishes an institutional or social identity. The second introduces a contradiction that implies an unseen history.

Consider these:

| Name | The world it implies |
|---|---|
| **The Amateur Consulate** | An unofficial diplomatic mission for people who do things for love rather than money. |
| **The Embassy of Elsewhere** | A diplomatic outpost representing somewhere that doesn't officially exist. |
| **The Parallel Society** | A community operating alongside conventional society, under different assumptions. |
| **The Leisure Authority** | An institution claiming jurisdiction over something that should require no permission. |
| **The Permanent Visitor** | Someone who belongs to a place without ever becoming its citizen. |
| **The Minor Embassy** | A diplomatic mission whose sovereign territory may be no bigger than a room. |
| **The Unrecognized Republic** | A country that exists regardless of whether anyone acknowledges it. |
| **The People's Amusement Bureau** | An agency from an alternate history where recreation is public infrastructure. |
| **The Embassy of Nobody** | An institution representing people who have no official standing. |
| **The Temporary Autonomous Zone** | A space of self-directed activity outside ordinary institutional control. |
| **The Other People's Club** | A social institution whose membership consists entirely of outsiders. |
| **The Office of Small Pleasures** | A bureaucracy responsible for things too intimate and trivial to administer. |

I think **The Amateur Consulate**, **The Minor Embassy**, and **The Unrecognized Republic** come closest to the literary mechanism behind the Denisovan Embassy.

But we can push the idea further.

## What if it's an archaeological artifact from a future that never happened?

Gibson's *The Peripheral* and *Agency* introduce the idea of divergent timelines, or *stubs*. That's especially useful here.

The arcade could be understood as a product from a timeline where personal computing developed differently.

A world where computers remained tools belonging to their operators, games remained small and portable, and the commercial internet never became an infrastructure of behavioral extraction.

Its name might be:

**The Last Public Terminal**

A surviving machine from a civilization where public computing was a common good.

**The Pleasure Annex**

An institutional addition whose actual parent institution has disappeared.

**The Bureau of Lost Futures**

An organization preserving technologies and cultural practices from futures that never materialized.

**The Parallel Arcade**

An arcade operating in a neighboring reality, accessible through ordinary computers.

**The Recreational Embassy**

An outpost representing a society that treats play as a legitimate purpose in itself.

**The Amateur Republic**

A fictional polity organized around voluntary activity rather than economic productivity.

These names are less explicitly strange than *Denisovan Embassy*, but they establish a similar relationship between the mundane and the impossible.

## I think the neurodivergent aspect could be more subtle and more interesting

Rather than defining the institution as belonging to autistic or ADHD people, imagine that its inhabitants have an entirely different understanding of attention, repetition, and play.

In their society, spending six hours arranging numbers in a grid isn't considered unusual.

Neither is returning to the same activity every day for twenty years.

Nobody thinks a computer should interrupt you to encourage a purchase. Nobody measures the value of recreation by how effectively it increases engagement.

These aren't accommodations. They're the prevailing cultural assumptions.

**The arcade is an embassy from a society in which neurodivergent experiences aren't departures from the norm.**

That premise could be expressed through something like **The Parallel Society** or **The Amateur Republic**, without using diagnostic language.

And it creates a compelling relationship to Gibson's fiction: the visitor isn't entering a fantasy world to escape reality. They're briefly crossing into another possible reality, one with different social arrangements.

## My strongest candidates

I would narrow this direction to five:

| Name | Why it works |
|---|---|
| **The Minor Embassy** | Beautifully understated; suggests an entire alternate political order. |
| **The Pleasure Annex** | Sounds like an actual location in the Sprawl; mildly illicit and bureaucratically absurd. |
| **The Amateur Republic** | Gives the project an ideology without making it a slogan. |
| **The Parallel Arcade** | Connects the fictional premise directly to the actual experience. |
| **The Embassy of Nobody** | Combines the outsider dignity of Gentleman Loser with the impossible diplomacy of Denisovan Embassy. |

### My favorite is The Pleasure Annex.

It has the quality I think you've been searching for.

*Pleasure* can mean sensory enjoyment, amusement, indulgence, or something slightly illicit.

*Annex* suggests an extension of some larger institution, but we're never told what that institution is.

It sounds like a place that was built for another purpose, subsequently occupied by an eccentric community, and now operates according to customs no one remembers establishing.

It's also plausible as the name of a nightclub, an experimental computer laboratory, a strange little arcade, or a place in Gibson's London.

And it doesn't try to explain the product.

**The Pleasure Annex** could simply be the name that appears when someone launches the application, immediately before encountering the wall of animated games.

The name suggests a world. The application lets them enter it.

One caveat worth preserving: this is an original naming direction inspired by Gibson's worldbuilding, not a suggestion to reproduce his fictional locations or imply an official connection to his work. None of these names has been checked for commercial availability.

**The most interesting question now isn't necessarily what to call the arcade, but what sort of civilization would have built it.** Once that civilization feels real, I suspect its name will become almost obvious.

---

## User

None of these are very good. I'm basically giving you the peripheral universe which is full of related terminology

---

## ChatGPT

You're right—I kept inventing vaguely Gibsonian names instead of drawing from the actual vocabulary and social institutions of *The Peripheral* and *Agency*. I'll focus on the terms Gibson uses, what they mean in that world, and which could plausibly name your arcade.

---

## ChatGPT

You're right. You've been pointing me toward **the actual cultural vocabulary of Gibson's Jackpot trilogy**, and I've been responding with generic cyberpunk branding.

*The Peripheral* and *Agency* already contain a rich taxonomy of people, places, technologies, and social institutions. And several of those terms are directly relevant to what you're building.

Most importantly, your idea of an arcade that exists between Gibson's universes and ours suggests something very specific: **an interface through which people in one reality can enter and interact with another.**

That's almost literally what a peripheral does.

But the distinction between *peripheral*, *polt*, *stub*, and *continuum* matters. They're not interchangeable science-fiction words. Each describes a different relationship between identity, embodiment, technology, and reality.

## The actual vocabulary worth exploring

| Term from the books | Meaning in Gibson's world | Connection to your project |
|---|---|---|
| **Peripheral** | An artificial body remotely inhabited by a human operator | Entering another experience through an interface |
| **Polt** | Someone acting remotely in another world, like a poltergeist | The player as a visitor rather than a consumer |
| **Stub** | A newly branching alternate timeline | Each game as a small, self-contained world |
| **Continua** | Multiple timelines or realities | Your growing collection of independently accessible worlds |
| **Continua enthusiast** | Someone who explores alternate timelines as a hobby | An extraordinarily apt description of your players |
| **Haptics** | Technology connecting nervous systems and machines | Sensory interaction and embodied play |
| **Sigil** | A visual identifier or symbol | Your almost entirely visual, textless interface |
| **Fab** | Fabrication, often through advanced automated manufacturing | Small programs created and distributed as artifacts |
| **Assemblers** | Advanced automated manufacturing technology | Procedurally constructed experiences |
| **Aunties** | Powerful automated intelligence systems | Computational assistance without conventional human interfaces |
| **The Jackpot** | The cascading catastrophes that transform society | The exploitative technological world your project rejects |

These aren't speculative associations: *polt*, *peripheral*, *stub*, *continua*, *haptics*, and *sigil* are part of the books' established vocabulary. ([PATRICK D. JOYCE](https://patrickdjoyce.com/blogs/blog/a-peripheral-glossary?utm_source=chatgpt.com), [Amazing Stories](https://amazingstories.com/2014/11/days-future-now-peripheral-william-gibson/?utm_source=chatgpt.com), [University of Adelaide Digital Library](https://digital.library.adelaide.edu.au/bitstreams/837d8332-4ac4-4a02-a0a2-d0bcc2f5c974/download?utm_source=chatgpt.com))

And the term **continua enthusiast** is particularly wonderful.

In *The Peripheral*, Lev Zubov treats alternate timelines as a hobby. They are places to explore, observe, and sometimes interfere with. The books also complicate this through the imperial relationship between the people who initiate stubs and those who inhabit them. ([Reactor](https://reactormag.com/the-peripheral-spoiler-review//?utm_source=chatgpt.com))

Your app reverses that relationship somewhat: every user independently chooses which world to enter, and the worlds don't extract anything from their visitors.

## Names I'd actually consider

Now that we're working with the source material rather than its aesthetic, these are much closer to what I think you're asking for.

| Candidate | Connection | Why it works |
|---|---|---|
| **Polt** | A remote visitor inhabiting another reality | Short, strange, playful; almost a nickname |
| **Polt Club** | The social identity of people visiting other worlds | Sounds like an actual subculture |
| **The Continua** | The collection of alternate realities | Elegant name for a collection of playable worlds |
| **Stub Club** | People exploring divergent timelines | Irreverent, memorable, almost childlike |
| **The Thylacinery** | Lev's thylacines and Gibson's extinct-animal imagery | Sounds like an impossible establishment |
| **Neoprimitive** | A cultural movement in the post-Jackpot world | Excellent connection to minimalist software and technological refusal |
| **The Cosplay Zone** | A designated area associated with historical reenactment | A place where another reality is deliberately performed |
| **Easy Ice** | Flynne Fisher's gaming identity | A real connection to play, skill, and virtual presence |
| **Remembrancer** | A term in the novel's vocabulary | Suggests preservation of experiences and cultural memory |
| **Autonomic Cutout** | A technological term in *The Peripheral* | Suggests interruption of involuntary systems; a strange, clinical-sounding name |

Several of these are actual distinctive names or phrases from the novels, so using them commercially would warrant particular copyright and trademark caution. ([PATRICK D. JOYCE](https://patrickdjoyce.com/blogs/blog/a-peripheral-glossary?utm_source=chatgpt.com), [University of Adelaide Digital Library](https://digital.library.adelaide.edu.au/bitstreams/837d8332-4ac4-4a02-a0a2-d0bcc2f5c974/download?utm_source=chatgpt.com))

### 1. Polt

**POLT** is probably my favorite.

In *The Peripheral*, it's a somewhat dismissive term for a person remotely participating in another time or reality.

But consider its other associations:

- Poltergeist: a presence that manifests by manipulating physical objects.
- An invisible participant whose agency becomes evident through what they do.
- A consciousness acting through an artificial interface.

That's a remarkably good description of someone playing a game.

And *Polt* has the phonetic quality of an obscure 1980s computer program. It doesn't sound like a conventional software brand.

It also offers a nice reversal: a term originally used to diminish people from less technologically advanced realities becomes a self-chosen identity.

### 2. Stub Club

This might be the most naturally playful.

**STUB CLUB**

It's funny without being a joke.

It's a club for people who like exploring little alternate worlds.

There's also an unintentional but useful association with tiny programs: small, self-contained pieces of code.

The strongest connection is architectural. Every game in your collection is a separate, persistent little universe governed by its own rules.

And because the collection grows without replacing older games, the worlds accumulate.

The caveat is that *stub* in Gibson's books has imperial connotations. It's a name imposed on alternate realities from the perspective of people who consider their own reality more important.

That tension could be part of the appeal.

### 3. The Thylacinery

This is probably the closest to your **Denisovan Embassy** example.

A thylacine is an extinct marsupial carnivore. Gibson's future includes recreated thylacines, including the pair belonging to Lev Zubov.

A *thylacinery* sounds like a facility devoted to keeping them.

But what kind of facility?

A zoo? A private club? A laboratory? A sanctuary? An eccentric aristocrat's estate?

The name implies an institution whose existence requires an entire alternate history.

It's delightfully esoteric.

However, it also has the weakest relationship to the actual functionality of your app.

### 4. Neoprimitive

I find this one conceptually fascinating.

**NEOPRIMITIVE**

It's already a recognizable cultural category within Gibson's post-Jackpot world.

But outside the fiction, consider what it communicates about your software:

You're using modern hardware and cross-platform technology to recreate radically simple computing experiences.

No cloud. No engagement algorithms. No sprawling dependencies.

Just small programs that do things.

That's almost a technological neoprimitivism.

It also recalls the demoscene's fascination with achieving complex results using deliberately restricted resources.

Unlike *Stimcade*, the name doesn't advertise a diagnosis or an audience. It expresses a philosophy of technology.

---

## There's one especially interesting possibility: Continua

**CONTINUA**

I think this deserves serious consideration.

It's already a real word, not a fictional invention, and it conveys multiplicity without fragmentation.

Your application would contain dozens or hundreds of little interactive worlds.

They exist independently but are accessible through the same interface.

The user scrolls through their visual representations, selects one, and enters it.

That's almost a literal depiction of browsing continua.

It also has a connection to neurodiversity through the idea of variation across continua rather than a binary distinction between normal and abnormal.

The word isn't inherently associated with a particular diagnosis, though.

Commercially, its existing use as an ordinary word would also create branding challenges.

---

## My ranking has changed substantially

| Rank | Name | Reason |
|---|---|---|
| **1** | **Polt** | Strongest combination of literary meaning, identity, and playful brevity |
| **2** | **Stub Club** | Best social identity; closest to an underground arcade |
| **3** | **The Thylacinery** | Closest to the impossible-institution quality of Denisovan Embassy |
| **4** | **Continua** | Best description of the actual interaction model |
| **5** | **Neoprimitive** | Strongest connection to minimalist, anti-extractive computing |

These haven't been checked for app-store or trademark availability.

### But I think there's a deeper connection we haven't addressed

In *The Peripheral*, Flynne's early experience of the future is mediated by something she initially understands as **a game**.

That's central to the novel, not incidental.

She believes she's performing a familiar kind of paid gaming work, but her actions are actually taking place in another reality.

Your interface presents the reverse proposition.

The user sees a continuous display of apparently autonomous little worlds. They select one, and suddenly they're inside it, able to intervene.

**The animated preview is the window into the stub. The game is the peripheral. The player is the polt.**

That is a much stronger conceptual foundation for the product than simply decorating an arcade with cyberpunk references.

And it makes **Polt** and **Stub Club** my strongest recommendations-not because they sound futuristic, but because they describe the relationship between your users and the worlds you're creating.

If you wanted the *Denisovan Embassy* quality specifically, however, I would explore **The Thylacinery** further. It has the distinctive property of sounding like an establishment that exists for reasons the audience doesn't yet understand.

---

## User

you're just pulling random words out of these stories. you're note really thinking or synthesizing any new ideas

---

## ChatGPT

I think the mistake I've been making is treating Gibson's fiction as a vocabulary to mine rather than a **system of ideas from which to derive something new**.

The Denisovan Embassy is interesting not because Denisovans are extinct or because embassies are unusual places. It's interesting because the name implies **an impossible political relationship that nevertheless seems entirely ordinary to the people inhabiting that world**.

Automatic Jack works differently. It's a nickname that implies a biography. Gentleman Loser suggests an entire social milieu. Neither explains itself. Both make you infer the world that produced them.

And *The Peripheral* adds another dimension: people in one reality casually interacting with another, often without fully understanding the nature of that relationship.

Your arcade sits at a fascinating intersection of those ideas.

### The fictional premise I'd start from

Imagine that somewhere in the post-Jackpot world, an eccentric institution discovers our timeline.

They find our internet rather appalling.

Not because our technology is primitive, but because nearly every trivial pleasure has been enclosed within an apparatus of surveillance, behavioral manipulation, and rent extraction.

They decide to establish a tiny cultural outpost.

It doesn't attempt to communicate with governments or influence history. Instead, it distributes little recreational programs through an interface that runs on almost any computer.

The programs are familiar because the underlying games are familiar across multiple timelines. Chess, territory conquest, number puzzles, and other mathematical structures recur independently.

The institution doesn't regard itself as a gaming company. It considers itself a kind of **cultural exchange**.

The people operating it are perhaps neurodivergent by our standards, but their world doesn't organize cognition into the same categories. Their particular relationship to attention, repetition, and pattern is simply part of their culture.

They don't understand why someone would put advertisements inside a game.

They might charge a nominal membership fee because maintaining the exchange requires resources, but selling attention would strike them as an extraordinary breach of hospitality.

Now the naming question becomes much more interesting:

**What would these people call their little outpost, and what would we call it?**

Those might not even be the same thing.

## Names derived from that premise

### 1. The Other Embassy

This isn't an embassy *of* somewhere. It's the other embassy.

The definite article implies that its existence is already common knowledge, while *other* leaves the relationship undefined.

It sounds like a place that might be mentioned casually in conversation without anyone explaining which nation it represents.

I like the social ambiguity, but it doesn't quite have the idiosyncrasy of Gentleman Loser.

### 2. The Friendly Occupation

This is more interesting politically.

An occupation is the seizure of territory by an outside power. *Friendly* introduces an unsettling contradiction.

Your application occupies a tiny portion of a device that would otherwise be dominated by commercial applications. But unlike those applications, it makes no demands on its host.

It's also suggestive of the post-Jackpot world's relationship to stubs: intervention framed as benevolence, with all the questions about power that implies.

There's something appropriately mischievous about an arcade calling itself **The Friendly Occupation**.

### 3. The Unscheduled Visitor

This draws on the temporal instability of the Jackpot trilogy without invoking its terminology.

The visitor is someone who arrives from elsewhere, but *unscheduled* implies a bureaucratic system responsible for regulating such arrivals.

It's also a quiet reference to ADHD, spontaneity, and nonconformity without reducing those experiences to a joke.

As a name, though, it sounds more like a character or an event than a place.

### 4. The Recreational Incident

This is closer to the tonal contradiction I think we want.

An *incident* is something authorities investigate, document, or contain.

*Recreational* suggests something undertaken purely for enjoyment.

Together they imply a society where spontaneous pleasure is sufficiently unusual to merit an official report.

It could be a nightclub, an unauthorized gathering, or a piece of mysterious software.

And the name has an implicit joke: the existence of a simple, non-extractive game collection has somehow become an incident.

### 5. The Minor Incursion

An incursion is a crossing into territory controlled by someone else.

*Minor* makes it sound bureaucratically insignificant.

Within the fictional premise, the application is literally an incursion from another timeline.

Within our world, it's a tiny intrusion into the commercial software ecosystem.

It has a slightly dangerous, understated quality that I associate with Gibson's naming.

But it doesn't communicate the pleasure of playing games.

### 6. The Voluntary Annex

This might be the strongest conceptual contradiction.

An annex is territory incorporated into something larger, often through an exercise of power.

*Voluntary* reverses the expected relationship.

The application becomes a little piece of another world that you choose to let into yours.

Unlike most digital platforms, it doesn't attempt to expand its jurisdiction over your behavior, data, or relationships.

It has territory only because you give it territory.

The name also sounds like a location whose original purpose has been forgotten.

### 7. The Unnecessary Machine

This connects more directly to the cultural philosophy.

Most software is justified by what it accomplishes.

A game is different. Its purpose is the experience of interacting with it.

An unnecessary machine is therefore not a failed machine. It's a machine that exists entirely because someone wanted it to.

There's also an affectionate rejection of instrumental rationality here, which feels relevant to both neurodivergent experiences and demoscene culture.

I can imagine *The Unnecessary Machine* existing in Gibson's world as a small, eccentric establishment that nobody quite knows how to categorize.

---

## The one I keep returning to: The Recreational Incident

Not necessarily because it's the best final name, but because it operates on several levels simultaneously.

It sounds like the title of a classified report concerning an inexplicable event.

It implies that someone somewhere considers recreation a matter of institutional concern.

It's simultaneously slightly threatening and completely unserious.

And it has the possibility of a fictional history.

Imagine the application displaying an introductory line only on first launch:

*INCIDENT CLASSIFICATION: RECREATIONAL*

Then the user enters the scrolling gallery of games.

No lore dump. No explanation. The name implies the rest.

That is much closer to the literary mechanism you've been pointing toward.

## But I think there's another possibility

We've been assuming that the name should belong to the **institution**.

What if it instead names the **relationship between our world and theirs**?

The remarkable thing about the Denisovan Embassy is that an embassy implies reciprocal recognition. It isn't merely an artifact from somewhere else. It's evidence that two societies acknowledge one another.

Your application could be framed as a small diplomatic privilege extended to people in our reality.

Not a product.

Not a platform.

**A recreational agreement between worlds.**

That suggests a different family of names: *mutual recognition, informal diplomatic relations, reciprocal hospitality, cultural exchange, extraterritorial privileges*.

I think there's a genuinely interesting name somewhere in that territory.

And it should probably sound less like a science-fiction title and more like something an eccentric institution would have named itself for reasons that were perfectly obvious to its founders.

**Of the directions we've explored, I'd pursue the diplomatic relationship rather than the cyberpunk technology.** It's the part of your Denisovan Embassy comparison that has the most conceptual depth—and the part I hadn't really engaged with before.

---

## User

this project is more solarpunk than cyberpunk

---

## ChatGPT

That changes the framing substantially. I was interpreting Gibson's worlds primarily through their cyberpunk aesthetics—intrusion, underground networks, illicit technology—when what you're building is much closer to **a solarpunk response to the conditions Gibson describes**.

And *The Peripheral* is a particularly interesting reference because its future isn't simply a dystopia. It's a world where astonishing technological capabilities coexist with the consequences of catastrophic political and economic failures.

Your project imagines taking some of those capabilities—small, sophisticated, accessible computational experiences—and organizing them according to an entirely different set of values.

### The distinction I was missing

Cyberpunk traditionally asks: *How do people exercise agency when technology is controlled by powerful institutions?*

Solarpunk asks: *What would technology look like if we organized it around mutual benefit, autonomy, accessibility, and ecological and social sustainability?*

Your arcade answers that second question in a very concrete way.

It's not an underground resistance network fighting the advertising industry. It's simply a working demonstration that **the advertising industry is unnecessary to making and distributing enjoyable games**.

And that changes the naming possibilities.

Rather than imagining a mysterious institution operating illegally between timelines, imagine **an ordinary public amenity from a more desirable timeline**.

An institution whose name implies a society where the right to play, explore, and enjoy technology is taken for granted.

The interesting contrast is that something entirely ordinary in that world would seem almost utopian in ours.

## Names I'd explore from that perspective

| Name | What it implies |
|---|---|
| **Common Pleasure** | Enjoyment is a shared good rather than a commodity. |
| **The Public Arcade** | An arcade understood as public infrastructure, not a commercial venue. |
| **Amateur Electric** | Technology produced for the love of it rather than profit. |
| **The Pleasure Cooperative** | Recreation organized through mutual ownership and participation. |
| **Open Leisure** | Unrestricted time and enjoyment; an alternative to engagement extraction. |
| **The People's Arcade** | Games as a cultural commons. |
| **Common Play** | Play as something shared, ordinary, and universally available. |
| **The Leisure Commons** | An institution devoted to preserving access to uncommodified recreation. |
| **Mutual Amusement** | People supporting one another's enjoyment rather than monetizing attention. |
| **The Public Garden** | A cultivated, growing collection that people enter to enjoy without being exploited. |

But I think the stronger insight is about **the kind of name**, not the specific words.

### Imagine a solarpunk version of the Denisovan Embassy

An embassy represents a political entity.

What would a solarpunk civilization establish in another timeline?

Perhaps not an embassy at all.

It might establish a **library, garden, cooperative, conservatory, or public arcade**.

Those institutions are interesting because they have recognizable social purposes that don't depend on extracting maximum economic value from their visitors.

A public garden doesn't need to know how many minutes you spent admiring a particular flower.

A library doesn't need to interrupt you with advertisements to justify the existence of books.

And a cooperative doesn't necessarily need to maximize returns to outside investors.

Your game collection has elements of all three.

It's a **library** because the games are preserved and accumulate.

It's a **garden** because it grows incrementally.

And it's an **arcade** because people enter it to play.

The solarpunk quality comes from how those functions are organized, rather than from visual symbols of greenery, sunlight, or solar panels.

### A name I find newly interesting: Mutual Amusement

**Mutual Amusement** sounds like an institution from a world where entertainment developed through cooperative rather than advertising-driven economics.

It also has a faintly antiquated quality, like a nineteenth-century mutual aid society or an early electrical amusement company.

It isn't saccharine. It doesn't announce itself as virtuous. And it doesn't sound like a wellness application targeted at neurodivergent consumers.

It sounds like an establishment.

I also like **Amateur Electric**, which has more of the sensibility of Automatic Jack.

*Amateur* originally denotes someone who does something out of love. *Electric* invokes technological modernity without implying corporate ownership.

Together, they suggest a culture of people making things simply because they enjoy making them.

Neither name is availability-checked.

---

## The direction I'd pursue

I think the essential idea is:

**A little piece of a better technological civilization, distributed into ours.**

Not an escape from the present, but a functioning example of an alternative.

That is both solarpunk and compatible with Gibson's interest in how technologies acquire meaning through the social arrangements surrounding them.

For naming, I'd now favor something that sounds like a **social institution, a communal amenity, or an eccentric maker collective**, with just enough ambiguity to suggest a larger world.

And I would stop trying to encode autism, ADHD, gaming, cyberpunk, and solarpunk into a single clever compound word.

Those things should inform the institution's character. The name only needs to make that character believable.

**The closest candidates so far, to me, are Amateur Electric and Mutual Amusement.** They evoke a world where joyful, accessible software is ordinary rather than remarkable.

---

## User

what about a play on autoerotica

---

## ChatGPT

**Autoerotica** is much closer to the kind of name you've been reaching for.

It has several overlapping meanings that actually relate to the project, rather than merely referencing its aesthetics.

There's **auto** as self-directed, autonomous, automatic, and computational. There's **erotica** in the broader sense of pleasure, desire, and sensory experience. And the whole word carries a slightly transgressive, queer, countercultural quality.

Most importantly, it suggests **pleasure that doesn't require anyone else's participation or approval**. That's an interesting proposition for an offline, surveillance-free arcade designed around neurodivergent autonomy.

There's also an association with *autoeroticism* that could be funny without necessarily making the product sexual. The joke is that you're spending hours privately pressing buttons because it feels good.

And there's a nice connection to the autoerotic dimensions of stimming: self-directed sensory pleasure, repetition, and regulation without needing to justify the activity through productivity.

## Names I'd explore

| Name | Semiotic possibilities |
|---|---|
| **Autoerotica** | The strongest original word; technological autonomy and unapologetic pleasure |
| **Autotelica** | From *autotelic*: an activity whose purpose is intrinsic to itself |
| **Autotelic Arcade** | An arcade where playing is its own reward |
| **Autotelic** | Philosophically precise, elegant, and understated |
| **Autonomica** | Autonomy, automation, and the suggestion of an alternative technological culture |
| **Autopoetica** | Self-creation, generative systems, and poetic technology |
| **Autoludica** | Self-directed play; from *ludic*, relating to play |
| **Auterotica** | A deliberately strange mutation that retains the erotic association |
| **Erotomatic** | Erotic + automatic; sounds like a forgotten coin-operated amusement machine |
| **Pleasure Principle** | Freud, self-directed enjoyment, and the rejection of instrumental rationality |

### The particularly interesting one: Autotelica

I think **Autotelica** may be the strongest synthesis of everything you've described.

*Autotelic* comes from Greek roots meaning self and end or purpose. An autotelic activity is undertaken for its own sake, rather than as a means to some external reward.

That is almost an exact description of your design philosophy.

Your games don't exist to generate advertising impressions, collect behavioral information, or manipulate people into purchasing something else.

**The purpose of playing is playing.**

The connection to ADHD and autism is also interesting without being reductive: immersion in an activity because the activity itself is satisfying, whether that's repetition, pattern recognition, experimentation, or mastering a system.

And *Autotelica* retains much of the phonetic character of *Autoerotica* without sounding explicitly sexual.

It also feels like the name of an institution or technology from an alternate future.

### But Autoerotica has something Autotelica doesn't

**Autoerotica is funny.**

And I think we've been underestimating the importance of humor.

*Gentleman Loser* is funny. *Automatic Jack* is funny in a much more understated way. The *Denisovan Embassy* is funny because its impossibility is presented matter-of-factly.

*Autoerotica* has that quality.

It sounds like an establishment from an alternate history where self-directed pleasure is a perfectly respectable technological discipline.

It also has the potential to be a little embarrassing to say aloud, which might be either a feature or a problem.

The practical drawback is that Apple and Google's policies around sexual content, as well as app-store search and discovery, could make an explicitly erotic-sounding name unnecessarily difficult to market—even if the application contains nothing sexual.

## My strongest three

**Autoerotica** — Most provocative, funny, and culturally distinctive.

**Autotelica** — Most philosophically coherent with the actual experience you're creating.

**Erotomatic** — Most evocative of a fictional arcade machine or strange piece of retro-futurist technology.

Of those, I'd lean toward **Autotelica**, but with an important qualification: *Autoerotica* has a personality that the more intellectually respectable name lacks.

And that tension might be exactly what you're looking for.

I also think **Autotelic Arcade** is worth considering, because it makes the philosophical argument almost explicitly: an arcade devoted to experiences that are ends in themselves.

These are creative suggestions, not trademark-cleared names.

**The most promising insight is the connection between autoeroticism and autotelic play: enjoyment that is self-directed and intrinsically meaningful, rather than contingent on external validation or reward.**

---

## User

it still feels like you dont really understand hat im asking for

---

## ChatGPT

I think you're right. I've been repeatedly translating your references into abstract concepts and then generating names that literally describe those concepts. That's almost the opposite of what you're asking for.

You gave me three particularly revealing references:

- **Automatic Jack** — A name with rhythm, personality, humor, and an implied social history. It belongs to a particular cultural world.
- **Gentleman Loser** — Queer-coded in its relationship to masculinity, respectability, failure, and outsider culture, without needing to announce any of those things.
- **Denisovan Embassy** — An absurdly specific institution that implies an entire alternate social reality. It sounds like a place that genuinely exists, not a name invented to communicate a theme.

And you've been describing a **solarpunk, neurodivergent, queer, technologically sophisticated arcade** that might exist somewhere between Gibson's fictional futures and our own.

But you don't want a name that *describes* those attributes. You want something that could plausibly **emerge from the culture that produced such a place**.

That's why my response to *Autoerotica* missed the point.

I immediately converted it into *autotelic play*, explained the philosophy, and recommended *Autotelica*. I effectively removed the sex, humor, ambiguity, and irreverence that made your suggestion interesting.

*Autoerotica* has a completely different register. It could be the name of a dubious electronics shop, an experimental art collective, a queer nightclub, or an obscure piece of software. It plays with autonomy, technology, sensuality, and self-directed pleasure without settling into a single explanation.

And it doesn't sound like it was named by someone trying to communicate four product benefits.

**I think you're looking for a name that feels discovered rather than designed.** Something with the cultural density of a good fictional proper noun, where the associations aren't all deliberate or even entirely consistent.

The other thing I've been missing is that **solarpunk doesn't mean wholesome**. A queer, anarchist, neurodivergent solarpunk culture could be vulgar, funny, sexually liberated, eccentric, and technically sophisticated. It doesn't need to sound like a nonprofit promoting community gardens.

So I want to correct my approach rather than give you another twenty candidates.

The question I should be asking when evaluating a name isn't *"How many aspects of the product does it communicate?"*

It's **"What kind of culture would produce this name, and would I want to visit the place it names?"**

That's a much more demanding creative problem, and it's the one you've actually been asking me to engage with.

---

## User

this thing is more like a medpack or battlestim that buffs the user than like a tool the user needs to accomplish something

---

## ChatGPT

That changes the metaphor considerably. **The arcade isn't a destination, a tool, or even a collection of games. It's a consumable experience that temporarily improves the user's state.**

More *stimpack* than *cyberdeck*.

And the distinction matters. A cyberdeck extends your capabilities so you can accomplish something. A stimpack is something you use because of what it does **to you**. The experience is the point, not the task.

That also makes your earlier suggestion of *Autoerotica* much clearer. You were circling a vocabulary of self-administered pleasure, stimulation, and bodily autonomy, not just making a provocative pun.

There are some interesting tensions here:

- **Medicine versus recreation:** The product borrows the language of treatment and enhancement without claiming that play is a medical intervention.
- **Enhancement versus normalization:** The user isn't being corrected or made more productive. They decide which experience feels good.
- **Military technology versus mutual care:** A battlestim is normally administered to keep someone fighting. Your version exists to help someone enjoy being alive, without any subsequent demand for performance.
- **Intoxication versus control:** The user can seek stimulation, repetition, or absorption without surrendering control to an engagement algorithm.

That last distinction is particularly important. Conventional mobile games are often engineered to keep people playing. Your application would be engineered to let people play **exactly as much as they want**.

## This suggests a different naming vocabulary

I'd look at the intersection of fictional pharmaceuticals, recreational compounds, emergency supplies, and underground consumer products.

Not literal clinical terminology, but names that sound like things someone might keep in their pocket in a more humane future.

| Name | What it evokes |
|---|---|
| **Recreational Dose** | Pleasure treated with the matter-of-fact language of medicine |
| **Soft Ration** | A small, portable allotment of comfort or stimulation |
| **Pleasure Reserve** | Something kept available for when it's wanted |
| **Civilian Stim** | A battlestim repurposed for ordinary life |
| **Personal Effects** | Both belongings and the effects of something on a person |
| **Little Remedy** | Affectionate, intimate, slightly dubious |
| **Over the Counter** | Accessible without permission; also a place of informal exchange |
| **Common Tonic** | An everyday restorative available to everyone |
| **Side Effects** | Both pharmacological consequences and the unintended pleasures of technology |
| **Maintenance Dose** | Something used routinely because it feels right |
| **Recreational Aid** | Bureaucratic language for voluntary pleasure |
| **Field Comfort** | A portable provision for difficult environments |

Of these, I find **Civilian Stim**, **Personal Effects**, and **Side Effects** especially interesting.

### Civilian Stim

This gets closest to your battlestim analogy.

It implies that the technology originated in an environment of coercion or conflict, but has been appropriated for civilian life.

A battlestim helps someone continue fighting. A civilian stim has no military objective. It belongs to the person using it.

That transition from military technology to everyday mutual care has a genuinely solarpunk quality.

It also has a certain Gibsonian casualness: *"You got a civilian stim?"* sounds like something a character could say without explanation.

### Personal Effects

This is subtler and perhaps more sophisticated.

*Personal effects* ordinarily means someone's belongings.

But *effects* also describes what a substance or experience does to a person.

The phrase therefore operates simultaneously as a collection of personal possessions and a collection of self-administered experiences.

It also sounds like the name of a strange little shop or software publisher.

### Side Effects

This is the most playful.

It takes the language of unwanted consequences and makes those consequences the entire point.

In a conventional medical context, side effects are something you tolerate in exchange for the intended therapeutic benefit.

Here, there's no external therapeutic objective. The sensory experience itself is what you came for.

It also sounds like a demoscene group.

---

But I think the central naming question is now much more precise:

**What would a society call a recreational stimpack when it had stopped regarding self-directed stimulation as suspicious, pathological, or unproductive?**

Probably not something explicitly therapeutic.

It might have a slang name, a manufacturer's name, or a nickname derived from an effect people enjoy.

Something that sounds as ordinary as *gum*, *coffee*, *a smoke*, or *a pick-me-up*.

That's the territory I'd explore next: **a little dose of voluntary, self-directed pleasure, packaged like a piece of everyday technology.**

Not a place to go. Not a machine to operate. **Something you take.**

---

## User

no these are all terrible

---

## ChatGPT

Fair. I've been overexplaining the references and then producing names that sound like marketing exercises.

You're looking for something with the wit, irreverence, and cultural specificity of *Automatic Jack*—something that could exist in a queer, neurodivergent, solarpunk future, with the playful pharmacology of a battlestim and the sensuality of *Autoerotica*.

I keep making the names too literal, too wholesome, or too self-consciously meaningful.

I don't think another list of twenty names would help. **I need to stop trying to make every name explain the product.** The name should have enough personality to stand on its own, even if its relationship to the arcade is oblique.

---

## User

ok like hat

---

## ChatGPT

I'd try names that sound like **recreational pharmaceuticals, illicit consumer electronics, queer underground culture, or fictional products from a future where pleasure is a public good**.

Not necessarily wholesome. Not necessarily respectable. A little funny, a little suspect.

| Name | Character |
|---|---|
| **Softcore** | Sexual innuendo, software, gentleness, pleasure without intensity being mandatory |
| **Joyride** | Unauthorized use of a machine purely for pleasure |
| **Afterglow** | The pleasurable residue of an experience |
| **Good Chemistry** | Neurochemistry, interpersonal affinity, recreational substances |
| **Pleasure Unit** | A measurement of enjoyment, a machine, or a person |
| **Happy Accident** | Something delightful that wasn't planned |
| **Recreational Chemistry** | An underground laboratory or a respectable-sounding vice |
| **Special Delivery** | Something you've been looking forward to receiving |
| **Sweet Relief** | A pleasurable interruption of an unpleasant condition |
| **Electric Candy** | A little piece of synthetic sensory pleasure |
| **Minor Vice** | Something indulgent, harmless, and slightly disreputable |
| **Candyflip** | Psychedelic slang, altered states, and playful experimentation |
| **Soft Landing** | Comfort after intensity; a place to come down |
| **Pleasure Model** | An artificial construct designed for enjoyment, with a suggestive double meaning |
| **Good Medicine** | Something enjoyable that happens to make life better |

Three strike me as closer to the register you're describing.

**Softcore** has an interesting ambiguity. It could refer to software, erotic media, an aesthetic, or a deliberately gentler alternative to hardcore gaming. It doesn't require an explanation. It sounds like a demoscene production or an underground label.

**Minor Vice** has some of the social wit of *Gentleman Loser*. It makes a small, harmless pleasure sound like a moral transgression. There's an implicit refusal to justify enjoyment.

**Joyride** connects particularly well to the experience. You temporarily take control of a little machine, go somewhere for no practical reason, enjoy yourself, and leave. No achievement required. No destination.

Of these, **Minor Vice** is the one I'd most want to find on the sign of a little arcade in Gibson's future.

But **Softcore** might be the better name for the actual software: short, memorable, suggestive, technologically resonant, and not trying to explain itself.

These are creative candidates, not availability-cleared names.

---

## User

what if there was a mechanic of changing the color and icon schemes from foods to gems to biologicals, etc

---

## ChatGPT

I think that's a genuinely interesting mechanic, especially if it applies **across the entire arcade rather than just to individual games**.

You could separate each game's underlying mechanics from its sensory presentation. The same game could be experienced through fruit, gemstones, microorganisms, geometric shapes, flowers, or something completely abstract.

That makes the collection more like a **configurable sensory environment** than a conventional arcade.

And it gives users control over what kind of stimulation they want.

## 1. One game, many sensory interpretations

Take your Merge (2048) implementation.

The rules remain identical, but the visual representation changes.

| Theme | Visual progression | Aesthetic |
|---|---|---|
| 🍒 Food | Cherry → Strawberry → Peach → Watermelon | Bright, organic, playful |
| 💎 Gems | Quartz → Amethyst → Sapphire → Diamond | Saturated colors, crystalline shapes |
| 🦠 Biological | Cell → Colony → Organism → Ecosystem | Organic, evolving, microscopic |
| 🌱 Botanical | Seed → Sprout → Flower → Tree | Solarpunk, naturalistic |
| 🌌 Cosmic | Dust → Asteroid → Planet → Star | Expansive, astronomical |
| ⚙️ Mechanical | Gear → Engine → Machine → Factory | Industrial, kinetic |
| 🔷 Geometric | Triangle → Square → Pentagon → Hexagon | Clean, abstract, mathematical |
| 🔢 Numeric | 2 → 4 → 8 → 16 | Original number-based gameplay |

The important point is that the theme should not change the rules or probabilities.

A 128 tile is still a 128 tile, even when represented as an orchid or a diamond.

This makes it possible to offer different sensory experiences without introducing different game variants that require independent maintenance.

## 2. Make themes a global feature

Instead of asking players to configure every game separately, I'd make the selected theme apply to the entire collection wherever it makes sense.

For example:

**💎 Gem mode**

Merge uses gemstones.

Snake becomes a string of crystals collecting smaller gems.

Mines uses mineral deposits and excavation.

Four in a Row uses colored stones.

Conquest uses crystalline territories.

The themes would provide shared colors, shapes, icons, textures, and potentially sound effects.

Your existing EBE framework could expose a theme interface:

```go
type Theme interface {
    Color(role ColorRole) color.Color
    Icon(role IconRole) *ebiten.Image
    Sound(event SoundEvent) AudioClip
}
```

I would distinguish semantic roles from the actual artwork.

For instance, games could request `PlayerOne`, `PlayerTwo`, `Hazard`, `Reward`, or `Background` rather than explicitly requesting red, blue, or green.

That would allow the entire interface to change coherently.

## 3. This creates an opportunity for sensory profiles

I'd avoid assuming that autistic or ADHD users all prefer the same kinds of stimulation.

Instead, let people choose the presentation they enjoy.

| Setting | Possible options |
|---|---|
| Palette | Pastel, vivid, monochrome, high contrast |
| Symbol system | Food, gems, biological, botanical, geometric |
| Motion | Static, gentle, animated, intense |
| Effects | Off, subtle, expressive |
| Audio | Silent, minimal, arcade, ambient |
| Background | Solid, gradient, animated |

One person might want bright crystals with explosive merging animations.

Someone else might want muted botanical graphics with no sound.

Both should be able to play exactly the same game.

And importantly, neither should be treated as the default or the more appropriate experience for neurodivergent users.

## 4. This could make your animated catalog much more distinctive

Imagine the player changing from 🍓 to 💎.

The entire scrolling gallery transforms.

Every game preview changes to the gem theme, while the underlying gameplay remains identical.

That's an unusually compelling visual experience.

It also reinforces the idea that your arcade isn't a collection of unrelated games. It's one coherent environment containing many kinds of play.

I'd make the theme selector another contextual interaction rather than adding more permanent navigation.

You could potentially access it by long-pressing the ✨ discovery icon.

The selector itself could be almost entirely visual: a scrolling row of theme swatches, each showing a short animation.

No descriptions required.

## 5. I'd support mixing themes

This is where the idea gets particularly fun.

Instead of forcing a single global theme, users could choose a global default and override individual games.

For example:

- Global default: 🌱 Botanical
- Merge: 💎 Gems
- Snake: 🍒 Food
- Mines: 🦠 Biological

Those preferences would be saved locally.

You could also allow themes to vary by game, so the player doesn't have to sacrifice the particular visual interpretation they enjoy most.

One technical consideration: not every theme will map naturally to every game. A theme system should define common visual roles but allow each game to provide its own interpretation.

---

## 6. It could also help with the naming problem

This makes me think the concept is drifting away from a conventional game collection and toward something closer to **a sensory toybox or recreational stim system**.

The games are the underlying structures.

The skins are the sensory materials.

The user chooses the combination that produces the experience they want.

That connects more directly to your earlier references to autoerotica, battlestims, and Gibson's technology than any of the names I proposed.

The central design principle could be:

**Mechanics are stable. Sensation is customizable.**

That is especially compelling alongside your append-only development philosophy.

You promise that the underlying games won't change unexpectedly, while allowing users to change their sensory experience as much as they like.

### One caution

I'd keep numeric and shape-based accessibility cues available even in themed modes. If different values or players are distinguishable only by color, the game becomes harder to use for people with color-vision differences.

And the theme picker should have a way to return to a known, simple default.

---

**I think this is worth making a foundational part of the architecture**, not adding it later as cosmetic customization.

With Go/Ebitengine, you can build game logic once, expose semantic visual roles, and have a shared theme system supply artwork, colors, animations, and sounds.

The resulting product becomes considerably more distinctive: not merely hundreds of ad-free games, but hundreds of familiar game mechanics that users can experience through whichever sensory vocabulary they prefer.

---

## User

the games become a way for people to solve a solvable problem in the way they want to with the kind of dressing and color they feel like

---

## ChatGPT

Yes. And I think this gets much closer to the actual identity of the project.

**You're offering people a collection of bounded, solvable problems, with control over both how they engage with those problems and how solving them feels.**

The game is the structure. The theme is the sensory experience. The player chooses the combination.

That creates three distinct dimensions of agency:

1. **Cognitive:** What kind of problem do I feel like solving? Spatial reasoning, pattern recognition, optimization, deduction, strategy, timing?
2. **Sensory:** What do I want that problem to look, sound, and feel like? Gemstones, fruit, cells, flowers, numbers, saturated colors, muted colors, silence, animation?
3. **Temporal:** How long do I want to engage? One move, one round, twenty minutes, or the same familiar game for hours?

None of these choices needs to be justified by an external objective.

### The distinction between solving and winning

I think *solvable* is an especially important word here.

Not every game needs to be easy, and not every problem needs to have a guaranteed solution. But the player should encounter **an intelligible system in which their actions produce understandable consequences**.

That's different from the contemporary mobile-game experience where the apparent puzzle is frequently entangled with deliberately engineered friction, artificial scarcity, and monetization incentives.

In your collection, a difficult puzzle is genuinely difficult because of its rules—not because an algorithm wants the player to buy another attempt.

That distinction could be especially meaningful to people seeking predictable, self-directed experiences, including some autistic and ADHD players, without assuming everyone in those communities wants the same thing.

### The themes are more than skins

I would take your idea a little further.

Imagine that the player selects a problem involving combining equivalent objects.

They might experience it as:

- **Gemstones:** Crystals fuse into increasingly elaborate formations.
- **Botany:** Seeds develop into plants, which combine into larger ecological systems.
- **Microbiology:** Cells divide, merge, and form increasingly complex structures.
- **Numbers:** Pure mathematical relationships, without metaphor or decoration.

The underlying state transitions remain identical.

But the interpretation changes what the player notices, anticipates, and enjoys.

A numerical progression might satisfy someone who enjoys arithmetic regularity. A botanical progression might satisfy someone who enjoys watching things grow.

These aren't necessarily interchangeable experiences psychologically, even though they're computationally equivalent.

And that's important: **the player isn't merely selecting a cosmetic preference. They're choosing how they want to encounter the problem.**

### This suggests a different product architecture

I would separate the application into three layers:

| Layer | Responsibility | Example |
|---|---|---|
| **Problem** | Rules, states, constraints, solutions | Slide-and-merge |
| **Expression** | Symbols, palettes, animation, sound | Gemstones |
| **Interaction** | Controls, pace, feedback, accessibility | Swipe, keyboard, controller |

A given problem could support several expressions and interaction styles without requiring separate implementations.

You might even eventually let people create and share their own expressions without modifying the games.

That could become a particularly attractive open-source contribution model: someone who doesn't write Go could contribute a botanical theme, a new visual progression, or a set of sensory effects.

### And I think it changes how we should think about the name

Earlier, I kept trying to name a fictional arcade, cyberdeck, or recreational stimulant.

But the thing you're actually describing is closer to **a collection of small, voluntary encounters with order**.

The pleasure comes partly from having control over the conditions of the encounter.

You choose the problem. You choose its appearance. You choose how to approach it. You choose when you've had enough.

There's something gently radical about making a computer behave that way.

It doesn't try to understand you, classify you, optimize you, or decide what you should want next.

It just offers possibilities and responds to your choices.

**The computer isn't trying to solve the user. The user gets to solve the problem.**

I think that's the clearest articulation of the project's philosophy we've reached so far.

---

## User

maybe the list scrolls left/right which is the same set of games but with a different theme, from less (left) to more (right) on a spectrum from black and white wireframes to bubblegum orgasm dragonforce happyhardcore vibes

---

## ChatGPT

**Yes. This is the first interface concept that really unifies the mechanics, sensory customization, and philosophy of the project.**

You're describing a **two-dimensional arcade** in which the two axes represent fundamentally different choices:

- **↕ Vertical:** What kind of problem do I want to solve?
- **↔ Horizontal:** How do I want that problem to feel?

And the horizontal dimension isn't just a collection of unrelated skins. It's a **spectrum of sensory intensity**, from almost entirely abstract mathematical representation to maximalist, ecstatic audiovisual excess.

The same games exist at every point along that spectrum.

## 1. A two-dimensional arcade

Imagine the application opening onto a grid:

| | ◀ Minimal | | | Maximal ▶ |
|---|---|---|---|---|
| **Merge** | Monochrome wireframes | Gemstones | Bioluminescent organisms | Exploding rainbow crystals |
| **Snake** | White line on black | Colored geometric segments | Botanical growth | Psychedelic neon dragon |
| **Mines** | Simple black-and-white grid | Mineral formations | Cellular landscapes | Chromatic chain reactions |
| **Conquest** | Abstract territory outlines | Color-coded regions | Living ecosystems | Galactic-scale fireworks |
| **Pong** | Two lines and a dot | Soft glowing shapes | Fluid organic motion | Laser trails and particle explosions |

The rows are different computational experiences.

The columns are different sensory interpretations.

But here's the important part: **the user doesn't necessarily need to see a grid.**

They could experience it as one continuous navigable surface.

Scroll vertically to explore games.

Swipe horizontally to transform the entire collection.

The user's current horizontal position determines how every game looks and sounds.

### The controls could remain almost completely textless

At the top:

**⭐　👥　✨**

Underneath: the animated game previews.

Swipe up/down to find a game.

Swipe left/right to transform its sensory presentation.

Tap to play.

Long-press for 📌 Pin, 🙈 Hide, ↗️ Share.

That's nearly the whole interface.

---

## 2. I'd make the sensory spectrum continuous

This is where I'd depart slightly from a conventional theme-selection system.

Rather than jumping directly between eight fixed themes, imagine that the interface interpolates between progressively more elaborate presentations.

For example:

| Position | Visual character | Motion and sound |
|---|---|---|
| **0%** | Black-and-white wireframes | No effects, silent |
| **20%** | Solid shapes, limited palette | Minimal transitions |
| **40%** | Saturated colors, familiar objects | Smooth animation, subtle sounds |
| **60%** | Organic shapes, rich textures | Particles, responsive motion |
| **80%** | Bioluminescence, iridescence | Elaborate effects, layered audio |
| **100%** | Bubblegum orgasm dragonforce happy-hardcore | Maximal audiovisual spectacle |

The interaction could feel like turning a knob on an instrument.

At the far left, Snake might be a white line moving across a black screen.

As you swipe right, the line gains color, texture, a body, expressive movement, and increasingly elaborate reactions.

At the far right, eating a piece of food might trigger a ridiculously extravagant visual celebration.

**The puzzle remains exactly the same. Only the experience changes.**

That creates an interesting possibility: a player might find the exact point on the spectrum where a particular game feels most satisfying.

Not everyone wants minimalism. Not everyone wants maximalism. And the same person might want both at different times.

## 3. The distinction between theme and intensity

There's one architectural subtlety I'd preserve.

**Food, gemstones, and biological organisms are themes. Wireframe-to-happy-hardcore is intensity.**

Those are related, but not identical.

A gemstone theme could be monochromatic and minimal.

A biological theme could be bright, chaotic, and full of movement.

So I wouldn't permanently bind *gemstones* to 40% and *biological organisms* to 60%.

Instead, I'd give the system two independent properties:

**Representation:** What are the game objects?

**Intensity:** How elaborately are those objects presented?

That allows combinations such as:

- Minimalist monochrome gemstones.
- Pastel botanical organisms with gentle motion.
- Extremely colorful geometric wireframes.
- Hyperkinetic fruit with exaggerated animations.

The horizontal swipe could control intensity, while a separate contextual selector changes representation.

Alternatively, you could make the horizontal axis a curated progression through complete audiovisual presets, combining both dimensions. That would produce a stronger artistic experience at the expense of independent customization.

For a first release, I'd actually favor **curated presets**. They're easier to make coherent and much easier for users to understand through visual previews.

---

## 4. The transition itself could be the experience

This is the most exciting part to me.

Instead of the whole screen abruptly changing when someone swipes horizontally, the current game preview could **morph into the adjacent interpretation**.

Imagine Merge:

At the far left, it's a plain grid containing numbers.

Swipe right.

The numbers become simple colored blocks.

Swipe again.

The blocks become gemstones.

Again.

The gemstones become glowing crystalline structures.

Again.

The screen erupts into iridescent particles as matching objects fuse.

The underlying game state remains identical throughout.

The user isn't switching games or even necessarily switching themes.

They're exploring alternate presentations of the same mathematical system.

That's a much more interesting experience than a conventional settings menu.

And because all the game previews are already animated, you could synchronize their transitions.

A horizontal swipe transforms the entire visible gallery.

## 5. The engineering is feasible with Ebitengine

I'd model it as two separate coordinates:

```go
type ArcadeState struct {
    GameIndex      int
    SensoryLevel   float64
}
```

The launcher would use `GameIndex` to determine which games are visible and `SensoryLevel` to determine how they're rendered.

However, I wouldn't make every game implement its own bespoke sensory scale.

I'd define shared presentation parameters:

```go
type SensoryProfile struct {
    Saturation       float64
    Contrast         float64
    MotionIntensity  float64
    ParticleDensity  float64
    AudioComplexity  float64
    HapticIntensity  float64
}
```

Each game could interpret those values through its own renderer.

A change in sensory intensity could affect the global palette, animation timing, particle systems, audio, and haptics.

For games with more distinctive presentation systems, you could supplement the shared parameters with game-specific mappings.

### One important limitation

You can't necessarily interpolate meaningfully between arbitrary icons.

A mathematical number doesn't gradually become a strawberry in the same way that a color can gradually become more saturated.

You'd probably want discrete representation changes combined with continuous transitions in color, effects, and animation.

That's a technical distinction, not a fundamental obstacle.

---

## 6. The accessibility implications are unusually good

The same architecture could support dramatically different sensory preferences without separating users into diagnostic categories.

Someone who wants the quietest possible experience moves left.

Someone who wants intense audiovisual stimulation moves right.

Someone who wants vivid colors but no sound could override audio independently.

I'd preserve independent controls for flashing, motion, sound, and haptics, especially because maximum intensity shouldn't imply potentially unsafe flashing effects.

Also, I'd make sure the application remembers the last selected position.

A person who prefers monochrome shouldn't have to navigate through an animated rainbow every time they open the app.

And someone who loves maximalist effects shouldn't have to rebuild their preferred environment every session.

**The interface should adapt to the user without needing to infer anything about them.**

That's a particularly elegant privacy principle.

---

## 7. There's an interesting relationship to your earlier naming ideas

Now I understand *Autoerotica* differently.

It's not merely a pun about self-directed stimulation.

The application is essentially a **user-controlled sensory instrument**.

The player chooses a bounded problem and then adjusts the sensory presentation until the experience feels the way they want.

The relationship is closer to a synthesizer than a conventional game menu.

You aren't simply browsing games.

You're tuning the experience of play.

That brings together your demoscene references, the solarpunk philosophy, and the metaphor of a self-administered recreational stimulant.

The user isn't being stimulated according to an engagement algorithm.

**They're controlling the stimulation themselves.**

And the difference between those two arrangements is arguably the entire philosophy of the project.

---

## One design choice I'd make differently

I wouldn't make the horizontal swipe directly change the sensory settings of a game that's already in progress without some indication that this is intentional.

In the launcher, horizontal gestures can navigate through presentations.

Inside the game, those gestures might be gameplay controls.

I'd keep the sensory navigation in the launcher and let the selected presentation carry into the game.

A pause menu could allow the user to change it mid-game, but it shouldn't interfere with playing.

### My strongest recommendation

**Make horizontal navigation feel like moving through parallel versions of the same arcade, not adjusting a conventional settings slider.**

At each horizontal position, the entire collection exists in a different sensory reality.

At the far left, it's an austere mathematical instrument.

At the far right, it's an ecstatic audiovisual toy.

And everything between those extremes remains available.

The resulting interface could be almost entirely textless, yet communicate two powerful questions immediately:

**What do you want to play?**

**How do you want it to feel?**

That's a genuinely distinctive interaction model.

---

## User

ok now think of hat the name needs to accomplish

---

## ChatGPT

I think the name needs to accomplish something more subtle than identifying an arcade or communicating that it's designed for neurodivergent people.

**It needs to name a pleasurable experience of voluntary cognitive and sensory stimulation, while implying a culture in which that experience is ordinary, legitimate, and available on the user's own terms.**

And it needs to do that with the wit and ambiguity of *Automatic Jack*, *Gentleman Loser*, or *Denisovan Embassy*, rather than sounding like a wellness product.

There are several requirements emerging from everything we've discussed.

## 1. It should name the experience, not the software

You're not really selling games in the conventional sense.

You're offering people the opportunity to choose a bounded problem, enter it, and experience the satisfaction of interacting with a comprehensible system.

The player controls the problem, its sensory interpretation, its intensity, and when to stop.

That's why your comparison to a **battlestim or medpack** matters.

The software is something a person reaches for because they want to feel a certain way.

But unlike a pharmaceutical or a therapeutic intervention, it doesn't promise to cure or correct anything.

The name should evoke the pleasure of using it, not suggest that the user needs treatment.

## 2. It should be culturally neurodivergent, not diagnostically branded

I think we've established that you don't want the name to say *ADHD Arcade* or *Autism Games*.

Those names frame neurodivergence as a market segment.

Your project is more interesting because it begins with a particular relationship to attention, repetition, pattern, stimulation, and autonomy, and makes that relationship foundational to the software.

The name can be strange, funny, intense, repetitive, sensual, or irreverent.

It doesn't have to explain why.

**The identity should feel endogenous to the culture, rather than designed to appeal to it from outside.**

That's an important distinction.

## 3. It needs room for both minimalism and maximalism

This is perhaps the biggest implication of your newest interface idea.

At one end, the application is black-and-white wireframes.

At the other, it's bubblegum, dragons, euphoric music, and maximal audiovisual spectacle.

Neither end is the correct experience.

The user chooses.

That makes names emphasizing only calm, focus, simplicity, or intensity misleading.

A name like *Soft Arcade* privileges one end of the spectrum.

A name like *Hypercade* privileges the other.

**The product is about control over the spectrum, not a particular point on it.**

The name should be comfortable appearing on both a monochrome terminal and an absurdly colorful animated screen.

## 4. It should have humor, pleasure, and a little transgression

This is what I missed about *Autoerotica*.

You weren't necessarily suggesting that exact word. You were demonstrating a register.

The word brings together automation, autonomy, erotic pleasure, privacy, and something faintly scandalous.

It refuses the antiseptic language usually used to market technology to people with disabilities or neurological differences.

And it is funny.

I think that matters enormously.

Your arcade should feel like something people enjoy, not something they're prescribed.

The name should have some capacity for mischief.

## 5. It should imply a world, not explain a feature

This is the Gibson influence.

*Denisovan Embassy* doesn't explain its purpose.

It gives us two familiar concepts whose juxtaposition implies an entire political and historical reality.

*Gentleman Loser* suggests a social institution and its relationship to status.

*Automatic Jack* suggests a character with a history.

The names have what I'd call **narrative surplus**: they imply more than they explicitly communicate.

That's what the name of your project needs.

It should sound like something from a slightly better future that has somehow found its way into ours.

Not a utopian manifesto. Not a dystopian warning.

An ordinary artifact of another social arrangement.

## 6. It should communicate autonomy without demanding productivity

The user's agency is central to the entire design.

They choose the game.

They choose the representation.

They choose the intensity.

They choose what appears in their catalog.

They choose when to play and when to leave.

Nothing monitors them or attempts to maximize their time in the application.

This is where the solarpunk philosophy becomes concrete.

**The computer serves the person. The person doesn't serve the computer.**

But I wouldn't put words like *autonomy*, *freedom*, *privacy*, or *commons* directly in the name simply to communicate those principles.

The name should belong to a culture that takes those things for granted.

## 7. It needs to work as an actual product name

There are practical constraints too.

| Requirement | Importance |
|---|---|
| Memorable after hearing it once | Very high |
| Easy to pronounce aloud | High |
| Distinctive enough for app-store discovery | Very high |
| Plausibly trademarkable | Very high |
| Not already used by another game collection | Essential |
| Works internationally | Desirable |
| Doesn't imply a medical treatment | Important |
| Doesn't create unnecessary adult-content classification problems | Important |
| Works as a small app icon and wordmark | High |
| Doesn't rely on knowing Gibson's fiction | Essential |

There's another practical consideration: the name needs to do more work because the interface contains almost no words.

Users will encounter it in app stores, recommendations, shared links, and conversations, but barely at all inside the application.

It needs to be recognizable without constant reinforcement.

---

# I think the central requirement is productive ambiguity

The best name should sit at the intersection of several meanings without resolving into one.

For example, *Autoerotica* operates simultaneously in the registers of technology, sexuality, self-direction, and recreational pleasure.

*Automatic Jack* evokes technology, a person, a nickname, and an implied social world.

*Gentleman Loser* combines status and failure.

And *Denisovan Embassy* combines extinct humanity with diplomatic legitimacy.

These names work because their constituent meanings **complicate one another**.

They're not merely additive.

That's the difference between *Neurocade*, which tells you what two concepts have been combined, and a name that creates an entirely new cultural object.

## The deeper contradiction your name could exploit

I think your project has one especially interesting contradiction:

**Something completely frivolous that is nevertheless treated as essential infrastructure.**

The games don't accomplish anything in the conventional productive sense.

They're recreational. Repetitive. Sometimes absurd. Potentially intensely pleasurable.

And yet you're proposing to make them permanently available, technologically accessible, privacy-preserving, and affordable, with the seriousness normally reserved for important public services.

That's an unusually rich cultural idea.

There's a second contradiction:

**A recreational stimulant whose dosage is controlled entirely by the person using it.**

It can be almost nothing—a few black-and-white lines—or an overwhelming celebration of color, sound, and motion.

But the intensity isn't chosen by an algorithm optimizing retention.

It's chosen by the player.

And a third:

**An advanced technological product that rejects much of what contemporary technology considers advancement.**

It works offline. It doesn't watch you. It doesn't personalize itself by secretly studying you. It doesn't manufacture urgency.

It uses sophisticated software to create an experience of remarkable simplicity and autonomy.

These contradictions are the raw material for the name.

---

## My proposed naming brief

If I were distilling everything into instructions for a naming exercise, I'd write:

> Find a short, memorable name for a fictional recreational technology or establishment from a queer, neurodivergent, solarpunk future.
>
> Its purpose is self-directed cognitive and sensory pleasure. It offers familiar, solvable problems in representations ranging from austere mathematical abstraction to ecstatic audiovisual excess.
>
> The name should feel culturally authentic rather than commercially optimized: witty, slightly transgressive, casually technological, and suggestive of a larger world.
>
> Draw inspiration from the narrative density and social ambiguity of William Gibson's *Automatic Jack*, *Gentleman Loser*, and *Denisovan Embassy*, without reproducing his terminology.
>
> Avoid literal descriptions, clinical language, wellness branding, generic cyberpunk imagery, and obvious portmanteaus.
>
> The result should sound like something that already exists in a more interesting future.

### One last thing

I think **the name should carry a little of the same pleasure as the application itself**.

It should be satisfying to say, slightly surprising, and open to interpretation.

It should be a name people might adopt affectionately, abbreviate, joke about, or turn into a verb.

Not because it describes the product perfectly, but because it has enough personality to become part of people's lives.

**The test I'd use is simple: Could this name plausibly appear in a Gibson novel without explanation, on a queer underground club flyer, and as the title of a tiny demoscene executable?**

If it works in all three settings—and doesn't sound like a marketing department invented it—we're probably getting close.

---

## User

ok no with all that in mind, give me a list of really good name candidates

---

## ChatGPT

I’ll aim for names you might find on the wrapper of a recreational invention—witty, slightly suggestive, culturally specific, and unconcerned with explaining themselves.

---

## ChatGPT

**Please Yourself, Sweet Release, Discrete Pleasures, and Idle Worship** are the strongest directions I see.

| Name | What makes it interesting |
|---|---|
| **Please Yourself** | An invitation, an assertion of independence, and an innuendo. It can be affectionate or gloriously dismissive: nobody else gets to decide what you should enjoy. |
| **Sweet Release** | Software arriving, a problem resolving, tension letting go, an orgasm. It sounds equally plausible on a candy wrapper, a club flyer, or a little executable. |
| **Discrete Pleasures** | Separate, bounded experiences, with the mathematical sense of *discrete* and the spoken suggestion of *discreet*. Private enjoyment made out of little systems. |
| **Idle Worship** | A twist on *idol worship* that treats supposedly unproductive enjoyment with extravagant reverence. Repetition becomes something you cherish rather than something you must justify. |
| **Soft Serve** | Software and confectionery, with a little suggestiveness. The name is tactile and appetizing; its visual world could accommodate both a plain white spiral and an obscene mountain of rainbow toppings. |
| **Pleasure Craft** | A vessel used for enjoyment, an object someone has made, and the practice of making things feel good. It suggests a small, personal means of getting somewhere without needing a destination. |
| **Self-Possessed** | Both being composed and belonging to yourself—with a slightly uncanny suggestion of inhabiting your own body. It has the dignity and ambiguity that the overtly “neuro” names lacked. |
| **Good for Nothing** | An affectionate insult turned into a refusal of instrumental value. No useful output required. The benefit belongs to the person playing, not to an employer, advertiser, or productivity metric. |
| **Second Helping** | A further serving of something enjoyable, but also another kind of help. It feels like a welcome provision rather than a treatment: take the portion you actually want. |
| **Small Mercies** | Little occasions when something becomes manageable. A bounded problem, a satisfying response, a world that behaves intelligibly. More intimate than spectacular, without prescribing a quiet aesthetic. |

### The four I would develop further

**PLEASE YOURSELF**

This has the most attitude. Crucially, it addresses the player without diagnosing, flattering, or instructing them to improve. The words can mean *“do whatever you like”* and *“give yourself pleasure.”* Neither reading cancels the other.

It could be the name of the software, an invitation printed on its packaging, or the unofficial motto of the culture that made it.

**SWEET RELEASE**

This has the strongest connection between the recreational-stim idea and the actual medium. The double meaning isn’t something the branding has to manufacture: a release already belongs to software and bodily experience.

It also accommodates your full visual range. **SWEET RELEASE** in austere monochrome feels intriguingly clinical; the same name surrounded by fluorescent cherries and ecstatic particle effects becomes exuberantly indecent.

**DISCRETE PLEASURES**

This is the most specific to what you’re making without literally describing a game collection. The *discrete* part belongs to the bounded rules and separable little worlds. The *pleasures* belong to the player’s chosen encounter with them.

It sounds like an anthology of slightly disreputable experiences—which is a more interesting identity than “a platform for casual games.”

**IDLE WORSHIP**

This has the strongest subcultural personality. It could belong to a demoscene collective, a queer night, or an eccentric recreational practice. The joke carries a real position: attention doesn’t become valuable only when somebody can put it to work.

**My first choice to put on an actual mockup would be _Please Yourself_.** It preserves the irreverence of *Autoerotica* while making the user’s sovereignty—not a particular sensory intensity—the central invitation.

These are creative candidates; I haven’t checked their commercial or app-store availability.

---

## User

What about stmpck

---

## ChatGPT

**STMPCK** is substantially stronger than most of the names we've explored.

I read it as **STIMPACK with the vowels stripped out**—and that compression is doing something interesting.

It combines several associations without explicitly spelling any of them out\:

- **Stimpack\:** A self-administered boost, somewhere between recreational stimulation and science-fiction consumables.
- **Stim\:** A recognizable reference to stimming within autistic and neurodivergent communities.
- **Pack\:** A collection of small things you carry around and use when you want them.
- **Vowelless typography\:** The visual language of demoscene releases, hacker handles, compressed filenames, and underground software.

Unlike *Stimcade* or *Neurodeck*, it doesn't feel like two marketing keywords glued together.

It feels like the label on a small, slightly illicit piece of software.

And the absence of vowels works particularly well with your nearly textless interface. The name itself participates in the same reduction of unnecessary information.

There's another reason I like it: **STMPCK doesn't commit to either end of your sensory spectrum.** It could appear as six stark white characters on a black screen or as an aggressively colorful, animated logo. Both would feel appropriate.

It also has the informality of an underground software release rather than the earnestness of a wellness brand.

### The main weakness

**Pronunciation.** People will probably say *stimpack*, but the spelling isn't immediately obvious from hearing it. That's a tradeoff between a distinctive visual identity and word-of-mouth discoverability.

I'd establish the intended pronunciation as *stim-pack*, while retaining **STMPCK** as the written name.

### Preliminary availability check

I searched for the exact spelling, including indexed Apple App Store and Google Play listings. I found no apparent gaming product using **STMPCK**\. The search results were largely unrelated documents or incidental letter sequences.

That's encouraging for distinctiveness, but it isn't a trademark clearance or confirmation that the name can be reserved on either store.

### My assessment

Of the names we've considered, **STMPCK is my strongest candidate so far**\.

It has an actual relationship to the product without describing it. It's culturally legible to the people you're targeting, but mysterious enough to acquire its own meaning. And it's short enough to work as an icon, executable name, domain, or demoscene signature.

I would keep **STMPCK**, pronounced *stimpack*, at the top of the shortlist rather than trying to turn it into another clever portmanteau.There's another reason I like it: **STMPCK doesn't commit to either end of your sensory spectrum.** It could appear as six stark white characters on a black screen or as an aggressively colorful, animated logo. Both would feel appropriate.

It also has the informality of an underground software release rather than the earnestness of a wellness brand.

### The main weakness

**Pronunciation.** People will probably say *stimpack*, but the spelling isn't immediately obvious from hearing it. That's a tradeoff between a distinctive visual identity and word-of-mouth discoverability.

I'd establish the intended pronunciation as *stim-pack*, while retaining **STMPCK** as the written name.

### Preliminary availability check

I searched for the exact spelling, including indexed Apple App Store and Google Play listings. I found no apparent gaming product using **STMPCK**. The search results were largely unrelated documents or incidental letter sequences.

That's encouraging for distinctiveness, but it isn't a trademark clearance or confirmation that the name can be reserved on either store.

### My assessment

Of the names we've considered, **STMPCK is my strongest candidate so far**.

It has an actual relationship to the product without describing it. It's culturally legible to the people you're targeting, but mysterious enough to acquire its own meaning. And it's short enough to work as an icon, executable name, domain, or demoscene signature.

I would keep **STMPCK**, pronounced *stimpack*, at the top of the shortlist rather than trying to turn it into another clever portmanteau.
