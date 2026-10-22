# Backend

Go + Fiber API with bun ORM and PostgreSQL.

## Run

```bash
cp ../.env.example ../.env   # if not done already
go run ./cmd/api
```

Or from the repo root: `make backend`.

## Migrations

```bash
go run ./cmd/migrate up       # apply pending migrations
go run ./cmd/migrate down     # roll back the last group
go run ./cmd/migrate status   # applied / pending
go run ./cmd/migrate unlock   # clear a stale lock after a crashed run
```

SQL files live in `migrations/` and are embedded into the binary. A migration is marked
applied only after its SQL succeeds. The tests in `cmd/migrate` run when
`TEST_DATABASE_URL` points at a PostgreSQL (CI sets it); each test uses its own throwaway
schema, so pointing it at a dev database is safe.

## Tests

```bash
go test ./...
```

## Structure

See `CLAUDE.md` in this directory for the layered architecture and conventions.
