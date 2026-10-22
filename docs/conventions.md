# Conventions

Naming, Git, and process conventions. Keep these boring and consistent.

## Git

- **Branches:** `type/short-description` — e.g. `feat/user-auth`, `fix/login-redirect`.
- **Default branch:** `main`. Never commit directly; use PRs.
- **Commits:** Conventional Commits style:
  - `feat:`, `fix:`, `refactor:`, `docs:`, `test:`, `chore:`, `perf:`, `build:`, `ci:`
  - Imperative mood, short subject; body explains why if non-obvious.

## Pull requests

- One logical change per PR. Keep them reviewable.
- Use the PR template (`.github/pull_request_template.md`, auto-loaded by GitHub).
- CI must pass (lint + test). Include migrations and tests with the change.

## Naming

- **Files (Go):** `snake_case.go`. **Vue components:** `PascalCase.vue`.
- **DB:** `snake_case` tables (plural) and columns.
- **Env vars:** `UPPER_SNAKE_CASE`.
- **API routes:** `kebab` or plural nouns, e.g. `/api/v1/users`.

## Versioning & releases

- Semantic versioning for tagged releases.
- Changelog entry per release (or generated from Conventional Commits).

## Documentation rule of thumb

- A **decision** → ADR in `docs/architecture/decisions/`.
- A **session's work** → note in `notes/sessions/`.
- A **task** → entry in `todo/`.
- **How something works long-term** → `docs/`.
