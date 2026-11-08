# Menu launch defaults

For the [foundation deliverable](../../ROADMAP.md#7-define-the-smallest-shared-experience), address **Open each new session at the default animated intensity, autoplay on, and sound off.**

## Work and validation

Stop loading/saving the previous native sensory-preference file. Initialize the default animated presentation with animation enabled and no audio output. Keep choices through ordinary navigation/suspend within the session, but reset on process restart. Validate new defaults and phone restart without deleting existing user data. Persistence of sound and intensity is deferred until after all five games.

## Status

New-session and intensity-sequence tests pass. Android and WebAssembly builds pass, and the updated APK installs and launches through the canonical device command. The default is now level 2 of four (zero-based), with two still levels to its left. Earlier evidence in `reports/menu-tabs/` tested the superseded independent autoplay control; touch/visual verification of the new four-level arrangement remains pending. No audio implementation is present; the app remains silent.
