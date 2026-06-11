---
name: create-migration
description: Create and wire a database migration (bun + PostgreSQL) safely, with an up and a reversible down.
---

# Create a database migration (bun + PostgreSQL)

Use for any schema change. Per `memory/constraints.md`: schema changes are ALWAYS
migrations — never hand-edit the DB, and never modify a migration that already shipped.

## Steps

1. **New file** in `backend/migrations/` with the next sequential prefix and a clear
   name (e.g. `0007_add_users_email_index`).
2. **Write `up`**: the forward change (create table / add column / index).
   - New non-null column on a populated table → add nullable + backfill + set not-null,
     or provide a default. Don't break existing rows.
   - Add indexes for new query patterns (check `memory/constraints.md` perf budgets).
3. **Write `down`**: the exact reverse, so the migration is reversible.
4. **Update bun models** in `internal/model/` to match.
5. **Run locally**: `go run ./cmd/migrate up` then verify; test `down` too.
6. **Seed/fixtures**: update if the change affects them.

## Safety checks

- Large tables: prefer non-blocking operations; call out any locking risk.
- Destructive change (drop column/table)? Confirm with the user first.
- Never edit a shipped migration — add a new one to fix it.

## Definition of done

- `up` and `down` both run cleanly; models match; tests pass.
