# AI Project OS — a Claude Code project template

[![CI](https://github.com/dogannx/claude-template/actions/workflows/ci.yml/badge.svg)](https://github.com/dogannx/claude-template/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A full-stack project template (Go + Fiber + bun / Vue + Vite / PostgreSQL) built around
**Claude Code**. It gives the AI durable project memory, a place for every kind of
knowledge, and a handful of commands, so context survives between sessions without
any manual bookkeeping.

> Built from real day-to-day use: start a session with "başlayalım", work, say
> "notlarını al", and everything that happened is recorded where it belongs.

## Why

AI assistants forget everything between sessions. The usual fix — piling more rules into
a prompt file — turns into a wall of `MUST`/`ALWAYS` that duplicates itself and drifts.
This template takes a different approach:

- **Knowledge has layers.** `memory/` holds what is true *now*, `docs/` explains *how*,
  `notes/` records *what happened*, `todo/` tracks *what's next*. Each fact has one home.
- **Memory is loaded, not requested.** `CLAUDE.md` imports `memory/` with `@` imports,
  so current project truth is in context every session.
- **Bookkeeping is a command, not a discipline.** Saving a session, capturing a note, or
  recording a decision is one phrase.
- **Hard rules are enforced by code, not prose.** Git hooks and CI block unformatted Go,
  committed `.env` files and non-conventional commits; markdown only carries context.
- **Decay is surfaced.** Memory files carry a `last-verified` date; a SessionStart hook
  and CI warn when they go stale.

## Natural-language commands (Turkish)

Commands can be typed as slash commands or just said in plain Turkish:

| Say…                                        | Command   | What happens                                                  |
|---------------------------------------------|-----------|---------------------------------------------------------------|
| "başlayalım", "neredeyiz"                   | `/basla`  | Reads progress, TODOs and the last session note; briefs you   |
| "bunu not et: …", "bug: …", "borç: …"       | `/not`    | Files the note in the right place (todo / memory / notes)     |
| "karar: …"                                  | `/karar`  | ADR for big decisions, one-line log for small ones            |
| "özetle", "ne yaptık"                       | `/ozet`   | Summarizes the session without writing anything               |
| "notlarını al", "işim bitti"                | `/kaydet` | Session note + TODO + progress + memory + health check        |

Plus `/new-feature` (plan before coding) and `/new-adr`, and skills for `add-endpoint`,
`create-migration`, `generate-tests`, `review-pr`, `debug-production`, `manage-release`.
Adding your own command in any language is one markdown file — see
[`docs/commands.md`](docs/commands.md).

## A typical session

```
"başlayalım"                              → /basla   brief + suggested focus
… development …
"bunu not et: refresh token süresi 7 gün" → /not     → memory/project-facts.md
"bug: login sonrası boş sayfa"            → /not     → memory/known-issues.md
"karar: auth için JWT kullanıyoruz"       → /karar   → ADR + memory update
"notlarını al, çıkıyorum"                 → /kaydet  → everything recorded, safe to leave
```

## Structure

```
.
├── CLAUDE.md          # AI entry point: routing table, project rules, @memory imports
├── memory/            # current truth: facts, constraints, known issues, tech debt
├── docs/              # architecture, ADRs, standards, conventions, system guide
├── notes/             # session notes, quick decisions log, progress timeline
├── todo/              # TODO / backlog / done
├── .claude/
│   ├── commands/      # /basla /kaydet /not /karar /ozet /new-feature /new-adr
│   ├── skills/        # repeatable procedures (endpoint, migration, tests, review…)
│   ├── hooks/         # SessionStart health check
│   └── templates/     # session note, ADR, feature plan
├── backend/           # Go + Fiber + bun API (own CLAUDE.md)
├── frontend/          # Vue 3 + Vite SPA (own CLAUDE.md)
├── scripts/           # setup, reset-db, backup, release, check-system
├── .githooks/         # pre-commit (gofmt, .env guard), commit-msg (Conventional Commits)
└── .github/workflows/ # CI: OS health, commit lint, Go, frontend
```

## Using this template

1. Click **Use this template** on GitHub (or clone it).
2. Set up locally:
   ```bash
   cp .env.example .env
   make hooks-install   # git hooks
   make up              # postgres + backend + frontend via docker compose
   ```
3. Open the project in Claude Code and say **"başlayalım"**.
4. Replace the `TODO`s in `CLAUDE.md` and `memory/project-facts.md` with your project's
   details (until then `make check` warns that `last-verified` is unset), rename the Go
   module in `backend/go.mod`, and replace this README.

Backend runs on `http://localhost:8080`, frontend on `http://localhost:5173`.
`make help` lists all tasks; `make check` runs the memory health check.

## Status

The workflow layer (commands, memory, hooks, CI) is in daily use. The app scaffold is
intentionally minimal: a health endpoint with tests, a Vue shell, and a placeholder
`cmd/migrate` you wire to bun's migrator for your project.

How the system fits together and how to keep it from rotting:
[`docs/system-guide.md`](docs/system-guide.md).

## License

[MIT](LICENSE)
