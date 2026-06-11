# Architecture Overview

High-level picture of the system. Update when the shape changes. For specific decisions,
see [`decisions/`](decisions/).

## System diagram (textual)

```
            ┌─────────────┐        HTTP/JSON        ┌──────────────┐
  Browser → │  Frontend   │  ───────────────────►   │   Backend    │
            │  (Vue/Vite) │                          │ (Go + Fiber) │
            └─────────────┘                          └──────┬───────┘
                                                            │ bun
                                                            ▼
                                                     ┌──────────────┐
                                                     │  PostgreSQL  │
                                                     └──────────────┘
```

## Backend layering

```
cmd/api            → entrypoint, wiring
internal/
  ├── http         → Fiber routes, handlers, middleware, DTOs
  ├── service      → business logic
  ├── repository   → data access (bun)
  ├── model        → bun models / domain types
  └── config       → env loading
```

- **Request flow:** route → handler (validate/bind) → service (logic) → repository (DB).
- **Dependencies point inward.** Handlers depend on services; services on repositories.

## Frontend structure

```
src/
  ├── views        → pages
  ├── components   → reusable UI
  ├── stores       → shared state
  ├── api          → backend client/service layer
  └── router       → routes
```

## Cross-cutting concerns

- **Config:** env vars via `.env` (12-factor).
- **Auth:** TODO — describe strategy (sessions/JWT) and link the ADR.
- **Errors:** consistent API error shape; centralized handler.
- **Observability:** TODO — logging/metrics approach.

## Deployment

- Containerized with Docker; orchestrated via docker-compose locally.
- CI builds and tests on each PR (GitHub Actions).
- TODO: describe production target (host, registry, deploy method).
