# STMPCK next work

[ROADMAP.md](ROADMAP.md) is the canonical checklist; its requirement entries own the linked plans. This file is a short pointer, not a competing backlog.

The five games and three-column gallery are implemented and run on the Android test phone. The gear controls, including Games wide (1–8), are mockups.

Remaining work includes:

- Implement settings overrides and persist preferences, including gallery columns.
- Implement the shared audio suite and per-game music.
- Complete accessibility, physical-controller, screen-size, lifecycle, performance, and longer balance playtests.
- Revisit Merge grid-size difficulty and suitable PvP modes at their scheduled milestones.
- Select release platforms, signing, credentials, packaging, and publishing destinations before store deployment.
- Review inherited optional scaffolding for later simplification; it is not required by the game or ordinary build workflow.

Current implementation and verification checkpoint: [2026-11-07 journal](journal/2026-11-07.md). Native Linux compilation still requires X11 development headers; Android and WebAssembly builds pass. No inference/governance service was used for this checkpoint.
