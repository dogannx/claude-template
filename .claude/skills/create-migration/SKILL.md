---
name: create-migration
description: Create and wire a database migration (bun + PostgreSQL) safely, with an up and a reversible down.
---

# Create a database migration (bun + PostgreSQL)

Use for any schema change. Migrations are SQL files in `backend/migrations/`, embedded into
the binary (`migrations/migrations.go`) and applied by `cmd/migrate` via bun's migrator.
Never edit a migration that has already shipped — add a new one.

## Steps

1. **Create the pair** in `backend/migrations/`, timestamp prefix in UTC:
   - `<YYYYMMDDHHMMSS>_<name>.tx.up.sql` — the forward change
   - `<YYYYMMDDHHMMSS>_<name>.tx.down.sql` — its exact reverse

   `.tx` runs the file in a transaction. Drop it only for statements PostgreSQL can't run
   inside one (e.g. `CREATE INDEX CONCURRENTLY`). Several statements in one file are fine.
2. **Write `up`** without breaking existing rows: a new non-null column on a populated
   table → add nullable, backfill, then set not-null (or give it a default). Add indexes for
   new query patterns (check `memory/constraints.md` perf budgets).
3. **Write `down`** so the change is fully reversible.
4. **Update bun models** in `internal/model/` to match.
5. **Run locally:** `make migrate`, `make migrate-status`, then `make migrate-down` and
   `make migrate` again to prove the round trip. `down` rolls back the last **group**
   (everything one `up` applied). A crashed run can leave a lock: `go run ./cmd/migrate unlock`.
6. If this is the project's first real migration, delete the
   `20261023000000_create_examples` example pair.

## Safety checks

- Large tables: prefer non-blocking operations; call out any locking risk.
- Destructive change (drop column/table): confirm with the user first.

## Definition of done

- `up` and `down` both run cleanly; models match; `go test ./...` passes (CI also runs
  the project's migrations up and fully down against PostgreSQL).
