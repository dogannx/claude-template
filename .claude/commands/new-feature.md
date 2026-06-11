---
description: Scaffold a new feature with a plan before writing code.
argument-hint: <feature name>
---

We want to build: **$ARGUMENTS**

Before writing any code, produce a short plan using
`.claude/templates/feature-template.md` as a guide:

1. **Goal & scope** — what's in, what's explicitly out.
2. **Affected areas** — backend endpoints/models, frontend views, DB migrations.
3. **Data model changes** — tables/columns, bun models, migration needed?
4. **API contract** — routes, request/response shapes.
5. **Tests** — what to cover.
6. **Risks / open questions.**

Wait for my approval on the plan, then implement in small, reviewable steps.
Follow `docs/coding-standards.md` and the relevant `CLAUDE.md`.
