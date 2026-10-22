# CLAUDE.md — Project Root Context

> Primary context entry point for Claude. Keep it short and current; link out instead of
> duplicating. Only write down what is specific to *this* project — generic good practice
> (read before editing, match conventions, ask before destructive actions) is the
> assistant's default and doesn't need restating here.

## 1. What is this project?

<!-- One paragraph: what the product does, who it is for, current stage. -->
TODO: Describe the project in 3–5 sentences.

## 2. Repository layout

| Path          | Purpose                                                      |
|---------------|-------------------------------------------------------------|
| `memory/`     | Current project truth — auto-loaded below.                  |
| `backend/`    | Go + Fiber API server. See `backend/CLAUDE.md`.             |
| `frontend/`   | Vue SPA (yarn). See `frontend/CLAUDE.md`.                   |
| `docs/`       | Long-lived knowledge: architecture, standards, ADRs.        |
| `notes/`      | Session notes, decisions log, progress timeline.            |
| `todo/`       | Living task lists (TODO / backlog / done).                  |
| `scripts/`    | Operational automation (setup, reset-db, backup, release).  |
| `references/` | Optional: project-specific snippets/cheatsheets.            |
| `.claude/`    | Commands, skills, templates, hooks, shared settings.        |

Knowledge layers: `memory/` = true **right now** · `docs/` = **how** it works ·
`notes/` = **what happened** · `todo/` = **what's next**.

## 3. Where to write what

The single source of truth for routing — commands (`/kaydet`, `/not`, `/karar`) use this
table. One fact, one home; when torn between two rows, pick the lighter one.

| What happened?                                          | Where it goes                                            |
|---------------------------------------------------------|----------------------------------------------------------|
| Big architectural decision (schema, auth, contract)     | `docs/architecture/decisions/NNNN-…md` (ADR)             |
| Small technical decision (library, port, file format)   | `notes/decisions/README.md` (one dated line)             |
| Task completed                                          | move from `todo/TODO.md` to `todo/done.md` (dated)       |
| New always-true fact (auth method, prod URL, owner)     | `memory/project-facts.md`                                |
| New hard constraint (no PII in logs, region locked)     | `memory/constraints.md`                                  |
| Conscious shortcut / tradeoff to revisit                | `memory/tech-debt.md`                                    |
| Bug / gotcha / footgun to avoid                         | `memory/known-issues.md`                                 |
| Project-specific snippet or cheatsheet                  | `references/` (only if real content exists)              |
| Session story (what we did + why)                       | `notes/sessions/YYYY-MM-DD-<slug>.md`                    |
| One-line dated timeline entry                           | `notes/progress.md`                                      |

ADR vs quick-log threshold: [`notes/decisions/README.md`](notes/decisions/README.md).
When you edit a `memory/` file, set its `last-verified` to today.

## 4. Project rules

- Schema changes go through migrations in `backend/migrations/`; never edit a shipped one.
- Accepted ADRs are binding. `memory/` reflects only Accepted ADRs.
- If `memory/` contradicts the code, the memory is probably stale — say so rather than
  silently following either.
- Commits follow Conventional Commits (enforced by `.githooks/` and CI).

## 5. Natural-language commands

The user works in Turkish and speaks naturally. When a message clearly matches one of
these intents — with or without the slash — run that command, then say what was saved
where. Catalog: [`docs/commands.md`](docs/commands.md).

| User says (examples)                                         | Run      |
|--------------------------------------------------------------|----------|
| "notlarını al", "kaydet", "oturumu kapat", "işim bitti"      | `/kaydet` |
| "başlayalım", "neredeyiz", "durum ne", "kaldığımız yer"      | `/basla`  |
| "bunu not et: …", "yapılacak: …", "bug: …", "borç: …"        | `/not`    |
| "karar: …", "buna karar verdik"                              | `/karar`  |
| "özetle", "ne yaptık", "nereye geldik"                       | `/ozet`   |

## 6. Common commands

```bash
make up         # start the full stack (docker-compose)
make backend    # run backend locally
make frontend   # run frontend dev server
make test       # run all tests
make migrate    # apply DB migrations
make check      # AI Project OS health check (memory freshness)
```

## 7. Where to look next

- Architecture → [`docs/architecture/overview.md`](docs/architecture/overview.md)
- Coding standards → [`docs/coding-standards.md`](docs/coding-standards.md)
- Conventions (naming, git, PRs) → [`docs/conventions.md`](docs/conventions.md)
- How this OS works + maintenance → [`docs/system-guide.md`](docs/system-guide.md)

## 8. Current project truth (auto-loaded)

@memory/project-facts.md
@memory/constraints.md
@memory/known-issues.md
@memory/tech-debt.md
