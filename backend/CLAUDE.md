# CLAUDE.md — Backend (Go + Fiber)

Context for AI assistants working in the backend. Read the root `CLAUDE.md` first.

## Overview

Go API server using **Fiber** (HTTP), **bun** (ORM), **PostgreSQL** (DB).

## Layout

```
backend/
├── cmd/
│   ├── api/          # main server entrypoint
│   └── migrate/      # migration runner
├── internal/
│   ├── http/         # routes, handlers, middleware, DTOs
│   ├── service/      # business logic
│   ├── repository/   # data access via bun
│   ├── model/        # bun models / domain types
│   └── config/       # env config loading
├── migrations/       # bun migrations
├── go.mod
└── Dockerfile
```

## Conventions (see also `docs/coding-standards.md`)

- Handlers are thin: validate/bind → call service → map result/error to response.
- Business logic lives in services; DB access in repositories.
- Pass `context.Context` through to all DB/external calls.
- Wrap errors with context (`%w`); never panic in request paths.
- Every schema change is a migration in `migrations/` — never hand-edit the DB.
- New behavior ships with tests (`go test ./...`).

## Commands

```bash
go run ./cmd/api          # run server
go run ./cmd/migrate up   # apply migrations
go test ./...             # tests
go vet ./... && go fmt ./...
```

## When adding an endpoint

1. Define/extend the bun model in `internal/model` (+ migration if schema changes).
2. Add repository method(s) in `internal/repository`.
3. Add/extend a service in `internal/service`.
4. Add handler + route + DTO in `internal/http`.
5. Write tests. Update docs/ADR if it changes architecture.
