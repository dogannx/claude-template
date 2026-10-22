# Database Notes

PostgreSQL + bun specifics for this project. Schema lives in migrations; this file
captures the *reasoning* and operational notes that code doesn't show.

## Schema overview

- Key tables and relationships: TODO (or link to an ERD).

## Conventions

- `snake_case` tables (plural) and columns.
- All timestamps `timestamptz`, stored UTC.
- Soft-delete? TODO (yes/no + column).
- Money: TODO (e.g. integer cents).

## Migrations

- Tooling: SQL migrations in `backend/migrations/` (embedded, bun migrator), run via
  `make migrate` / `make migrate-down` / `make migrate-status`. `down` rolls back the last
  **group** — everything one `up` applied.
- Never edit a shipped migration; add a new one. See the `create-migration` skill.

## Indexing & performance

- Document non-obvious indexes and why they exist.
- Known hot queries and their access patterns: TODO.

## Operational

- Backup: `scripts/backup.sh`. Reset (dev): `scripts/reset-db.sh`.
