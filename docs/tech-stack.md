# Tech Stack

The canonical record of technologies used and **why**. When you swap or add a tool,
update this file (and consider an ADR for significant changes).

## Backend

| Concern        | Choice            | Notes                                          |
|----------------|-------------------|------------------------------------------------|
| Language       | Go                | Primary expertise; fast, simple, great tooling |
| Web framework  | Fiber             | Express-like ergonomics on fasthttp            |
| ORM            | bun               | SQL-first ORM with migrations & query builder  |
| Database       | PostgreSQL        | Primary datastore                              |
| Migrations     | bun migrate       | Versioned, run via `make migrate`              |
| Config         | env vars (.env)   | 12-factor; see `.env.example`                  |

## Frontend

| Concern        | Choice            | Notes                                          |
|----------------|-------------------|------------------------------------------------|
| Framework      | Vue               | SPA                                            |
| Package mgr    | yarn 4 (corepack) | Pinned via `packageManager`; lockfile committed |
| Build / dev    | Vite              | Fast dev server, HMR                           |

## Infrastructure & tooling

| Concern        | Choice            | Notes                                          |
|----------------|-------------------|------------------------------------------------|
| Containers     | Docker / Compose  | Local parity with prod                         |
| VCS            | Git + GitHub      | PR-based workflow                              |
| CI             | GitHub Actions    | Lint, test, build on PR                        |

## Versions

Pin and record key versions here (Go, Node, Postgres, Vue) so onboarding is reproducible.

- Go: `1.2x`
- Node: `2x LTS`
- PostgreSQL: `16`
- Vue: `3.x`
