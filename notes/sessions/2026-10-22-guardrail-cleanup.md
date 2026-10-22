# Session — 2026-10-22 — Template guardrail cleanup

- **Duration / focus:** template review; remove redundant AI guardrails
- **Participants:** Doğan + Claude

## Goal

The template had "sharp boundary" rules written for older models that needed repeated
reminders. Remove them and review the whole template for other improvements.

## What we did

- Slimmed root `CLAUDE.md` to project-specific context; `memory/` is now auto-loaded via
  `@memory/...` imports instead of "always read memory first" rules.
- Removed forced checklists and MUST/REQUIRED wording (`/kaydet`, session + ADR templates,
  `memory/`, coding standards, sub-project `CLAUDE.md`s).
- Single-sourced the routing table in `CLAUDE.md` §3; `/not`, `/kaydet`, `/karar` point to it.
- Moved trigger phrases into command `description` frontmatter; deleted duplicate
  `/session-start`, `/session-end` and `.claude/templates/pull-request-template.md`.
- Rewrote `docs/system-guide.md` (three overlapping tables → one; added "enforce with
  code, not prose" principle).
- Fixed: SessionStart hook cd'd to the repo's parent → health check never ran.
- Fixed: missing `go.sum` → backend didn't build (CI would fail). Ran `go mod tidy`.
- Rewrote `router_test.go` to test the real `NewRouter` (was testing a fake route).
- Dropped the always-noisy "no session note today" health check.

## Decisions

- Drop redundant AI guardrails; enforce hard rules via hooks/CI. (Quick log, not ADR.)

## Changes

- Files/areas touched: `.claude/`, `CLAUDE.md`, `docs/`, `memory/`, `scripts/check-system.sh`,
  `backend/go.mod`, `backend/go.sum`, `backend/internal/http/router_test.go`.
- Migrations: none.
- Tests added/updated: `TestRouter` (health degraded + 404 error envelope).
- Commit `f12dc29`, pushed to `origin/main`.

## Open items / next steps

- [ ] Check the CI run for `f12dc29` is green on GitHub.
- [ ] `cmd/migrate` is still a placeholder (not wired to bun migrate).

## Notes for next time

Keep instructions lean: don't re-add MUST/ALWAYS rules or "double-check" checklists.
