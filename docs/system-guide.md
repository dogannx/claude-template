# System Guide — How this AI Project OS works

This repo is not just a folder layout; it's a lightweight **operating system for
AI-assisted development**. This guide explains the model and — critically — how to keep
it sustainable. Read once; it rarely changes.

## Mental model: four kinds of knowledge

| Layer       | Question            | Lifespan   | Loaded        |
|-------------|---------------------|------------|---------------|
| `memory/`   | true *right now*    | rare churn | every session |
| `docs/`     | *how* it works      | slow churn | on demand     |
| `notes/`    | *what happened*     | append     | for history   |
| `todo/`     | *what's next*       | constant   | every session |

Putting a fact in the wrong layer is the main failure mode. Reference → `docs/`.
Current truth → `memory/`. Event → `notes/`. Task → `todo/`.

## Maintenance tiers (the anti-bloat contract)

The goal is **minimum daily maintenance**. Files are grouped by obligation:

### Tier 0 — Always, near-zero effort (the heartbeat)
Updated automatically by `/session-start` and `/session-end`. If you only maintain
these, the system still works.
- `memory/project-facts.md`
- `todo/TODO.md`
- `notes/progress.md`
- root `CLAUDE.md`

### Tier 1 — Every project, slow-changing (set up once, edit rarely)
- `docs/coding-standards.md`, `docs/conventions.md`, `docs/tech-stack.md`
- `docs/architecture/overview.md` + `decisions/` (ADRs)
- `backend/CLAUDE.md`, `frontend/CLAUDE.md`
- `.claude/commands/`, `.claude/settings.json`

### Tier 2 — On demand only (don't create empty)
Add a file here **only when you have real content for it**:
- `memory/known-issues.md`, `tech-debt.md`, `constraints.md`
- `.claude/skills/*` (add when a task repeats 3+ times)
- `scripts/*` (add when a command is typed 3+ times)
- `references/` (project-specific only; personal KB stays external)
- `backend/docs/`, `frontend/docs/`
- `todo/backlog.md`, `todo/done.md`

## The daily loop

```
/session-start → read memory + progress + todo → propose focus → confirm
   ↓
develop (use /new-feature, skills; create ADR for big calls)
   ↓
/session-end → session note + update todo + progress + memory → (PR)
```

## Sustainability rules

1. **One fact, one home.** Never duplicate across layers; link instead.
2. **Stale beats missing? No — stale is worse.** A wrong `memory/` fact misleads the AI.
   Fix or delete on sight.
3. **Empty files are debt.** Don't scaffold Tier-2 files "just in case."
4. **Automate the boring parts.** If you maintain it by hand daily, make it a command.
5. **Prune yearly.** Archive old notes by year; delete dead references.

## Layer sustainability contract (who / when / decay / automation)

The honest answer to "will this rot?" for every layer. Designed for the **real** user
who forgets and takes shortcuts — not an ideal one.

| Layer            | Who updates        | When                        | If not updated → risk            | Automation guard                          |
|------------------|--------------------|-----------------------------|----------------------------------|-------------------------------------------|
| `memory/`        | dev + AI           | when a fact changes; `/session-end` | AI acts on wrong facts (worst) | SessionStart hook + `last-verified` staleness check + CI |
| `todo/TODO.md`   | dev + AI           | session start & end         | lose track of work               | `/session-start` reads it; surfaced in check |
| `todo/done.md`   | AI (`/kaydet`)     | every `[x]` (move from TODO)| "what did we ship" trail is lost | `/kaydet` Step 3 enforces the move        |
| `notes/sessions/`| AI (`/session-end`)| each session                | lose "why/what happened"         | `/session-end` writes it; check warns if missing today |
| `notes/progress.md`| AI               | `/session-end`              | no high-level timeline           | written by `/session-end`                 |
| `notes/decisions/`| AI (`/karar`) + dev| at each small decision     | small calls get re-litigated     | `/karar` routes small decisions here      |
| `docs/` (standards)| dev              | rarely (on convention change)| guidance drifts from reality    | CI enforces the rules they describe       |
| ADRs             | dev + AI           | at decision time            | decisions re-litigated/forgotten | `/new-adr` + Memory-impact requirement    |
| `.claude/skills` | dev                | when a task repeats 3×      | manual repetition                | none needed (additive, low decay)         |
| `scripts/`       | dev                | when a command repeats 3×   | error-prone manual ops           | used by hooks/CI → exercised regularly    |
| `references/`    | dev                | ad hoc                      | junk drawer / staleness          | optional; prune yearly                    |

**Principle — minimum maintenance, maximum value:** the only layers requiring active
daily care (`memory/`, `todo/`, session notes) are exactly the ones the commands and
hooks update *for* you. Everything else is slow-changing or additive. If you do nothing
but run `/session-start` and `/session-end`, the system stays alive.

## Designing for the real (imperfect) human

- **Make the right path the easy path:** bookkeeping happens via two commands, not manual edits.
- **Surface, don't nag:** the SessionStart hook *shows* decay; it never blocks work.
- **Fail loud only where it's cheap and important:** git hooks block secrets and bad
  commits (fast, local), CI blocks broken code — but memory staleness only *warns*.
- **Degrade gracefully:** skip a session note and nothing breaks; you just lose history.
  The system bends instead of snapping.
- **Stale is worse than empty:** a wrong `memory/` fact is the one failure that actively
  misleads — which is why it gets the most automation (hook + staleness + CI).

## Single dev vs. small team

- **Solo:** Tier 0 + Tier 1 is plenty. `todo/TODO.md` alone (skip backlog/done).
  `memory/` and ADRs are your future self's lifeline.
- **Team:** add PR template + CI enforcement of `conventions.md`; ADRs and `memory/`
  become shared contracts; `onboarding.md` and `glossary.md` start paying off.
  Consider `CODEOWNERS` and protected `main`.
