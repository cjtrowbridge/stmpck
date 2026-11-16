# Persistent sensory preferences — after the first five games

For [After the first five games](../../ROADMAP.md#after-the-first-five-games), address **Persist menu-animation, master-audio, haptics, and decorative-effect overrides alongside sound-effects, music, and intensity settings (with autoplay derived from intensity unless overridden) across closing and reopening the app.** The intended outcome is to restore the player's selected sensory experience on the next launch, once all five games are implemented.

## Timing and scope

Deferred by explicit user direction. Current builds start each new session at the default animated intensity, autoplay on, and sound off; they do not restore earlier preference files. Within-session choices still apply until the app process closes.

When activated, store sound-effects, music, and intensity settings locally and derive autoplay from intensity, validate saved values, recover safely from missing/damaged data, and verify restart behavior on supported platforms. Game-session saves are separate roadmap requirements.

Persist the working settings overrides from [settings overrides](../../ROADMAP.md#12-implement-settings-overrides) in this same post-five-games milestone. Validate stored values and restore the selected still/muted behavior before output begins. Mockup switch state must never be mistaken for implemented preferences.

## Deferral review — 2026-11-16

Retain this previously active scope as deferred: preference loading/saving was disabled by explicit user direction, as recorded in the local-persistence plan. Resume when sensory-preference persistence is selected for active work alongside the working override requirements. Completion requires validated local restoration and restart/recovery checks; the roadmap’s post-five-games timing does not itself reactivate execution.
