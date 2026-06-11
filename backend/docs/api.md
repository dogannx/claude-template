# API Documentation

The contract for the backend HTTP API. Keep in sync with the code; this is what
frontend and external consumers rely on.

## Conventions

- Base path: `/api/v1`
- Auth: see `memory/project-facts.md`
- Errors: consistent shape — TODO document it (e.g. `{ "error": { "code", "message" } }`)
- Pagination: TODO (e.g. `?page=&limit=`, response includes `total`)

## Endpoints

> Keep this as a living reference. Consider generating from OpenAPI later if it grows.

### Example

```
GET /api/v1/health
→ 200 { "status": "ok" }
```

| Method | Path                | Auth | Description        |
|--------|---------------------|------|--------------------|
| GET    | /api/v1/health      | no   | liveness check     |
| ...    | ...                 | ...  | ...                |
