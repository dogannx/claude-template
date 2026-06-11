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
go run ./cmd/migrate up
```

## Tests

```bash
go test ./...
```

## Structure

See `CLAUDE.md` in this directory for the layered architecture and conventions.
