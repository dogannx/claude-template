# CLAUDE.md — Project Root Context

> This file is the **primary context entry point** for Claude and other AI assistants.
> Keep it short, high-signal, and always up to date. Link out to detailed docs instead
> of duplicating them here. If something is stale, fix it — Claude trusts this file.

## 1. What is this project?

<!-- One paragraph: what the product does, who it is for, current stage. -->
TODO: Describe the project in 3–5 sentences.

## 2. Repository layout

| Path          | Purpose                                                      |
|---------------|-------------------------------------------------------------|
| `memory/`     | **Always-true project facts** — read this every session.    |
| `backend/`    | Go + Fiber API server. See `backend/CLAUDE.md`.             |
| `frontend/`   | Vue/React SPA (yarn). See `frontend/CLAUDE.md`.             |
| `docs/`       | Long-lived knowledge: architecture, standards, ADRs.        |
| `notes/`      | Session notes, decisions log, progress timeline.            |
| `todo/`       | Living task lists (TODO / backlog / done).                  |
| `scripts/`    | Operational automation (setup, reset-db, backup, release).  |
| `references/` | Optional: project-specific snippets/cheatsheets.            |
| `.claude/`    | Reusable commands, skills, templates, shared settings.      |

### The three knowledge layers (keep them distinct)

- `memory/` → "what is true **right now**" — tiny, always loaded, rarely changes.
- `docs/`   → "**how** it works" — large, read on demand, changes occasionally.
- `notes/`  → "**what happened**" — append-only history.

## 2.5 Where to write what (decision routing)

When something happens during a session, route it to **one** home — never both. If you
hesitate between two rows, prefer the lighter-weight one (quick log over ADR, single
memory line over a new doc).

| What happened?                                          | Where it goes                                            |
|---------------------------------------------------------|----------------------------------------------------------|
| Big architectural decision (schema, auth, contract)     | `docs/architecture/decisions/NNNN-…md` (ADR)             |
| Small technical decision (library, port, file format)   | `notes/decisions/README.md` (one dated line)             |
| Task completed                                          | `todo/TODO.md` `[x]` **then move to** `todo/done.md`     |
| New always-true fact (auth method, prod URL, owner)     | `memory/project-facts.md`                                |
| New hard constraint (no PII in logs, region locked)     | `memory/constraints.md`                                  |
| Conscious shortcut / tradeoff to revisit                | `memory/tech-debt.md`                                    |
| Bug / gotcha / footgun to avoid                         | `memory/known-issues.md`                                 |
| Project-specific snippet or cheatsheet                  | `references/` (only if real content exists)              |
| Session story (what we did + why)                       | `notes/sessions/YYYY-MM-DD-<slug>.md`                    |
| One-line dated timeline entry                           | `notes/progress.md`                                      |

The `/kaydet` command walks this table automatically at session end. ADR vs quick-log
threshold lives in [`notes/decisions/README.md`](notes/decisions/README.md).

## 3. Tech stack (summary)

- **Backend:** Go, Fiber, bun ORM, PostgreSQL
- **Frontend:** Vue, yarn
- **Infra:** Docker / docker-compose, GitHub Actions (CI)
- **Full detail:** [`docs/tech-stack.md`](docs/tech-stack.md)

## 3.5 Information precedence (resolve conflicts in THIS order)

When sources disagree, trust the higher-numbered-priority source and flag the conflict:

1. **`memory/`** — current project truth (auth, constraints, known issues). Wins over all.
2. **`todo/TODO.md` + active session note** — what we're doing right now.
3. **`docs/architecture/decisions/` (Accepted ADRs)** — binding decisions in force.
4. **`docs/`** — standards, conventions, architecture reference.
5. **Sub-project `CLAUDE.md`** (backend/frontend) — local specifics within the above.
6. **The code itself** — ground truth for behavior; if it contradicts `memory/`, the
   memory is stale → tell the user, don't silently follow either.
7. **Your own assumptions** — last resort. State them explicitly; prefer asking.

If `memory/` and the code disagree, that's a decay signal — surface it, don't guess.

## 4. Golden rules for AI assistants

1. **Read before you write.** Always read `memory/` first, then `docs/coding-standards.md`
   and the relevant `CLAUDE.md` (backend/frontend) before generating code.
2. **Respect existing conventions** over your defaults. Match the surrounding code.
3. **Small, reviewable changes.** Prefer focused diffs; explain non-obvious choices.
4. **No secrets in code.** Use `.env` (see `.env.example`). Never commit credentials.
5. **Record decisions.** Architectural choices → an ADR in `docs/architecture/decisions/`.
6. **Update the todo + session notes** at the end of meaningful work (see workflow below).
7. **Tests and migrations matter.** New behavior → tests. Schema change → a migration.
8. **Ask when ambiguous** instead of guessing on irreversible or outward-facing actions.

## 4.5 Natural-language commands (trigger phrases)

The user works in Turkish and speaks naturally. **When the user says any phrase below —
even without the leading slash — execute the mapped command's full workflow immediately.**
Match on intent, not exact wording. This is how the user records everything without manual
file edits. Full catalog: [`docs/commands.md`](docs/commands.md).

| If the user says (examples)                                            | Run        | Effect                                              |
|-----------------------------------------------------------------------|------------|-----------------------------------------------------|
| "notlarını al", "kaydet", "oturumu kapat", "projeden çıkıyorum", "işim bitti" | `/kaydet`  | Session note + TODO + progress + memory + health check → "safe to leave" summary |
| "başlayalım", "neredeyiz", "durum ne", "kaldığımız yer"               | `/basla`   | Read state, brief where we are + today's focus       |
| "bunu not et: …", "şunu hatırla", "yapılacak: …", "bug: …", "borç: …" | `/not`     | Smart-capture into the right file (todo/memory/notes) |
| "karar: …", "şunu karar olarak yaz", "buna karar verdik"              | `/karar`   | Record as ADR (big) or quick decision log (small)    |
| "özetle", "ne yaptık", "nereye geldik"                                | `/ozet`    | Summarize session (no writing)                       |

Rule: if a phrase clearly maps here, **do the action**, then tell the user what you saved
and where. If unsure which file a note belongs in, ask in one line, then write it.

## 5. Daily workflow

- **Session start:** run `/session-start` (or read `notes/progress.md` + `todo/TODO.md`).
- **During work:** keep diffs scoped; create ADRs for big decisions.
- **Session end:** run `/session-end` to write a session note and update todos/progress.

Full guide: [`docs/onboarding.md`](docs/onboarding.md).

## 6. Common commands

```bash
make up         # start the full stack (docker-compose)
make backend    # run backend locally
make frontend   # run frontend dev server
make test       # run all tests
make migrate    # apply DB migrations
```

(See the `Makefile`. Adjust per project.)

## 7. Where to look next

- Architecture overview → [`docs/architecture/overview.md`](docs/architecture/overview.md)
- Coding standards → [`docs/coding-standards.md`](docs/coding-standards.md)
- Conventions (naming, git, PRs) → [`docs/conventions.md`](docs/conventions.md)
- Onboarding for new people → [`docs/onboarding.md`](docs/onboarding.md)
- How this OS works + maintenance → [`docs/system-guide.md`](docs/system-guide.md)

## 8. Automation (so the system doesn't rely on memory/discipline)

- **SessionStart hook** runs `scripts/check-system.sh` → surfaces stale `memory/`,
  uncommitted facts, missing session notes at the start of every session.
- **Git hooks** (`.githooks/`, install via `make hooks-install`) → block unformatted Go,
  committed secrets, and non-conventional commit messages locally.
- **CI** (`.github/workflows/ci.yml`) → enforces format/lint/test/commit rules as a backstop.
- Run the health check anytime: `make check`.
