# Restore the upstream framework submodule

This plan owns **Restore the upstream framework submodule** under [Foundation 3](../../ROADMAP.md#3-reduce-the-inherited-boilerplate). The framework was copied into STMPCK as ordinary tracked files despite the requested submodule integration. Restore independent version control and publish the approved simplicity revision to `https://github.com/cjtrowbridge/agentic-pipelines.git`.

## Authorized outcome

The user explicitly requested repair and upstream publication on November 15. Work on existing main branches; a detached preparation worktree is allowed without creating a branch. Preserve unrelated host toolchain edits and upstream Ollama changes. Back up the copied framework before replacing it. Only the approved simplicity changes belong in the upstream commit.

- [x] Publish the scoped framework revision on top of current upstream main with available checks.
- [x] Replace copied framework files with `.gitmodules` and a mode-160000 Git entry pinned to that published commit.
- [x] Verify independent Git status, upstream reachability, preserved host customizations and required links; record limitations.

## Synthesis and rollback

Adopt the published upstream framework as a dependency. Keep host AGENTS, README, roadmap, prompts, pipeline, scripts and VS Code configuration authoritative. New upstream source/build/bootstrap policies do not silently migrate ordinary STMPCK development. Optional pipeline runtime/schema changes arrive with the dependency but the pipeline remains unrun; its current config requires a separate conformance audit before use. No model service or app build/deployment is needed for this repository repair.

Original copied tree backup: `/tmp/stmpck-submodule-repair/vendored-framework.tar.gz`; pre-conversion host diff: `/tmp/stmpck-submodule-repair/host-before.patch`. These are recovery artifacts, not tracked dependencies. The old host commit retains the original tracked tree. Do not roll back by discarding unrelated working changes.

## Evidence — 2026-11-15

Published upstream commit `608b803ffe95e4835260e7a54436dd857eec01d4` to main and verified it with `git ls-remote`. The detached publication worktree passed 16 documentation tests, the canonical staged documentation/source gate and diff whitespace checks. Prior full-catalog limitations (missing jsonschema; unrelated existing playbook-size failure) remain documented in the upstream plan. No source/runtime changes were included.

`.gitmodules` points to the canonical HTTPS repository. `git ls-files --stage agentic-pipelines` reports mode 160000 at the published commit; the nested repository has independent clean status. The original sibling checkout was fast-forwarded to that commit and its unrelated Ollama edits restored. Its complete recovery stash and the temporary archives are retained. No new branch was created. Host app behavior and pre-existing toolchain changes are preserved. The large deletion count is removal of vendored files from the parent index, not deletion from the submodule.
