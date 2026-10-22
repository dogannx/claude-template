# Example: session continuity in practice

What the system looks like after real use, and how the next session picks up from it.
The artifacts below come from the session that cleaned up this template (2026-10-22).
They live here, not in `notes/`, so a fresh copy of the template starts clean. Delete this
folder once you don't need the reference.

## Session 1 — work, then "notlarını al"

During the session the user only talked; at the end they said **"notlarını al"**.
`/kaydet` then wrote, without any manual file editing:

| File                                   | What landed there                                             |
|----------------------------------------|---------------------------------------------------------------|
| `notes/sessions/2026-10-22-…md`        | [the full session note](2026-10-22-guardrail-cleanup.md): goal, what changed, decisions, open items |
| `notes/progress.md`                    | `2026-10-22 — Removed redundant AI guardrails, fixed SessionStart hook + backend build.` |
| `todo/done.md`                         | `[x] 2026-10-22 — Template review: drop redundant AI guardrails, fix hook/build/test.` |
| `notes/decisions/README.md`            | `2026-10-22 — Dropped redundant AI guardrails …; hard rules are enforced by hooks/CI instead.` |

The decision went to the quick log, not an ADR: it is easy to revisit and doesn't touch
schema or contracts (threshold in `notes/decisions/README.md`). Nothing in `memory/`
changed, so `/kaydet` left it alone.

## Session 2 — "başlayalım"

A new Claude Code session has no memory of session 1. On **"başlayalım"**, `/basla`:

1. already has `memory/` in context (imported by `CLAUDE.md`);
2. reads `notes/progress.md`, `todo/TODO.md`, and the latest session note;
3. checks `git status` / `git log`;
4. briefs: where we are, the suggested focus, and open questions. Here that brief would
   surface the two open items from the note: check that CI is green, and that
   `cmd/migrate` was still a placeholder.

Both open items were picked up and closed in the following session (CI fixed, migrator
wired with a round-trip test) — the continuity loop working as intended.

## Reproduce it yourself (≈10 minutes)

1. Copy the template, open it in Claude Code, say **"başlayalım"**.
2. Ask for a tiny persisted feature, e.g. *"add a `notes` table and `GET /api/v1/notes`"*
   (`create-migration` + `add-endpoint` skills).
3. Mid-way, say **"karar: notes are soft-deleted"** and **"yapılacak: pagination"**.
4. Say **"notlarını al"** and look at what changed under `notes/`, `todo/`, `memory/`.
5. Quit, start a new session, say **"başlayalım"** — the brief should mention the
   soft-delete decision and the pagination TODO, and continue from there.
