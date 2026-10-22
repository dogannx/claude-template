# CLAUDE.md — Backend (Go + Fiber)

Backend-specific context; loaded on top of the root `CLAUDE.md`.

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

## Conventions

Handlers thin → services hold logic → repositories hold bun queries. Full list in
`docs/coding-standards.md`.

## Commands

```bash
go run ./cmd/api          # run server
go run ./cmd/migrate up   # apply migrations
go test ./...             # tests
go vet ./... && go fmt ./...
```

## Skills

- New endpoint → `add-endpoint` skill.
- Schema change → `create-migration` skill.
