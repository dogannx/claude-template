---
last-verified: TODO   # YYYY-MM-DD — bump whenever you confirm these facts are still true
review-cadence: 30d   # how often this file should be re-verified
---

# Project Facts

> Invariant truths the AI must always know. Keep entries short and current.
> Replace every TODO with the real value when starting a project.
> When you confirm everything below is still accurate, set `last-verified` to today.

## Identity

- **Project:** TODO (name + one-line purpose)
- **Stage:** TODO (prototype / MVP / production)
- **Primary owner(s):** TODO

## Authentication & authorization

- **Auth method:** TODO (e.g. JWT access + refresh, session cookies, OAuth provider)
- **Where enforced:** TODO (e.g. Fiber middleware `internal/http/middleware/auth.go`)
- **Roles / permissions model:** TODO

## Environments

| Env        | URL / host        | DB                | Notes                  |
|------------|-------------------|-------------------|------------------------|
| local      | localhost         | docker postgres   | `.env`                 |
| staging    | TODO              | TODO              | TODO                   |
| production | TODO              | TODO              | TODO                   |

## External integrations

- TODO (payment, email, storage, third-party APIs) — what, where configured, owner.

## Critical invariants

- TODO (e.g. "all money stored in integer cents", "all timestamps UTC",
  "user emails are unique and lowercased").

## Currently-in-effect decisions

- See `docs/architecture/decisions/` for full ADRs. Key ones in force:
  - TODO
