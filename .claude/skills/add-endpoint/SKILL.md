---
name: add-endpoint
description: Scaffold a new Fiber API endpoint end-to-end — model, repository, service, handler, route, DTO, and test — following project conventions.
---

# Add an API endpoint (Go + Fiber + bun)

Use when adding a new HTTP endpoint to the backend. Follow the layered architecture
(handler → service → repository) from `backend/CLAUDE.md` and `docs/coding-standards.md`.

## Before you start

- Confirm the route, method, request shape, response shape, and auth requirement.
- Check `memory/project-facts.md` for auth enforcement location and invariants.
- If the schema changes, do the `create-migration` skill first.

## Steps

1. **Model** (`internal/model/`): add/extend the bun model with correct tags. Respect
   invariants in `memory/project-facts.md` (e.g. money in cents, UTC timestamps).
2. **Repository** (`internal/repository/`): add data-access method(s). Use
   `context.Context`, parameterized queries, and a transaction for multi-step writes.
3. **Service** (`internal/service/`): put business logic here. Validate domain rules,
   return wrapped errors (`fmt.Errorf("...: %w", err)`).
4. **DTO** (`internal/http/`): request/response structs. Add validation; never trust input.
5. **Handler** (`internal/http/`): thin — bind+validate DTO → call service → map to the
   shared response/error shape.
6. **Route**: register in the central router, in the correct domain group, behind auth
   middleware if required.
7. **Test**: at minimum a handler/service test covering the happy path and one error path.

## Definition of done

- `go build ./... && go vet ./... && go test ./...` pass.
- Consistent error responses; input validated.
- If behavior is notable, note it; if architectural, propose `/new-adr`.
