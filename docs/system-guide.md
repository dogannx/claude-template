# System Guide — How this AI Project OS works

This repo is a lightweight **operating system for AI-assisted development**: a folder
layout plus a few commands and hooks that keep project knowledge alive with minimal
manual upkeep. Read once; it rarely changes.

## Mental model: four kinds of knowledge

| Layer       | Question            | Lifespan   | Loaded                        |
|-------------|---------------------|------------|-------------------------------|
| `memory/`   | true *right now*    | rare churn | every session (CLAUDE.md `@`) |
| `docs/`     | *how* it works      | slow churn | on demand                     |
| `notes/`    | *what happened*     | append     | for history                   |
| `todo/`     | *what's next*       | constant   | at `/basla`                   |

Putting a fact in the wrong layer is the main failure mode. The routing table in root
`CLAUDE.md` §3 is the single source of truth for "where does this go?".

## The daily loop

```
"başlayalım"  (/basla)   → read progress + todo + last note → brief → pick focus
   ↓
develop (/new-feature, skills, "bunu not et: …", "karar: …")
   ↓
"notlarını al" (/kaydet) → session note + todo + progress + memory → safe to leave
```

If you do nothing but `/basla` and `/kaydet`, the system stays alive.

## Who maintains what

| Layer              | Updated by                  | If it rots →                      | Guard                                    |
|--------------------|-----------------------------|-----------------------------------|------------------------------------------|
| `memory/`          | AI via `/kaydet`, `/not`    | AI acts on wrong facts (worst)    | `last-verified` + SessionStart hook + CI |
| `todo/`            | AI via `/kaydet`, `/not`    | lose track of work                | read at `/basla`                         |
| `notes/`           | AI via `/kaydet`, `/karar`  | lose the "why"                    | written by `/kaydet`                     |
| ADRs               | dev + AI via `/karar`       | decisions re-litigated            | "Memory impact" section                  |
| `docs/` standards  | dev, rarely                 | guidance drifts from reality      | CI enforces what they describe           |
| skills / scripts   | dev, when a task repeats 3× | manual repetition                 | additive, low decay                      |

## Maintenance tiers (anti-bloat)

- **Tier 0 — heartbeat (handled by commands):** `memory/project-facts.md`, `todo/TODO.md`,
  `notes/progress.md`, root `CLAUDE.md`.
- **Tier 1 — set up once:** `docs/` standards, architecture + ADRs, sub-project
  `CLAUDE.md`s, `.claude/`.
- **Tier 2 — only when there's real content:** other `memory/` files, `references/`,
  `backend/docs/`, `frontend/docs/`, `todo/backlog.md`, new skills and scripts.

## Principles

1. **One fact, one home.** Link instead of duplicating — duplicates drift apart.
2. **Stale is worse than empty.** A wrong `memory/` fact actively misleads; fix or delete.
3. **Instructions are context, not guardrails.** CLAUDE.md and commands describe what is
   specific to this project. Don't restate generic good practice or stack "MUST / ALWAYS /
   double-check" rules — modern models already follow them, and over-emphatic
   instructions make the model over-apply them.
4. **Enforce with code, not prose.** Anything that truly must not happen (unformatted Go,
   committed `.env`, bad commit messages) is blocked by git hooks / CI, not by a sentence
   in a markdown file.
5. **Surface, don't nag.** The SessionStart hook shows decay; it never blocks work.
6. **Prune yearly.** Archive old notes by year; delete dead references.

## Solo vs. small team

- **Solo:** Tier 0 + Tier 1 is plenty. `memory/` and ADRs are your future self's lifeline.
- **Team:** add branch protection + required CI checks; ADRs and `memory/` become shared
  contracts; `onboarding.md` and `glossary.md` start paying off. Consider `CODEOWNERS`.
