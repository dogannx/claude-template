# Coding Standards

Rules that keep the codebase consistent. AI assistants **must** follow these; humans too.
Prefer matching existing code over personal preference.

## General

- Clarity over cleverness. Code is read far more than written.
- Small, focused functions and PRs.
- No dead code, no commented-out blocks left behind.
- Handle errors explicitly; never silently swallow them.
- No secrets in source. Use env/config.

## Go (backend)

- Format with `gofmt`/`go fmt`; pass `go vet`. CI enforces this.
- Package layout: `cmd/` for entrypoints, `internal/` for private packages.
- Errors: wrap with context (`fmt.Errorf("doing X: %w", err)`); return, don't panic in
  request paths.
- Handlers thin, logic in services; DB access via a repository layer using bun.
- Context-aware: pass `context.Context` down to DB/HTTP calls.
- Tests: table-driven where it fits; name `TestXxx`; cover error paths.

### Fiber

- Group routes by domain; register in one place.
- Validate and bind request bodies into DTOs; never trust input.
- Consistent error responses (shared error handler / problem shape).

### bun + PostgreSQL

- One model per table; tag fields explicitly.
- All schema changes go through migrations — never edit the DB by hand.
- Use transactions for multi-step writes.
- Avoid N+1; use relations/joins deliberately.

## Vue (frontend)

- Composition API + `<script setup>`.
- Components: PascalCase files; one component per file.
- Keep API calls in a dedicated client/service layer, not in components.
- State management for shared state; local state otherwise.
- Lint/format via project config (`yarn lint`, `yarn format`).

## Comments

- Explain **why**, not **what**. Match the surrounding comment density.
- Public Go identifiers get doc comments.
