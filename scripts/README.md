# scripts/ — Project Automation

Idempotent, documented shell scripts for repeatable operational tasks. If you type a
multi-step command more than twice, it belongs here. Keep them small and safe.

## Rules

- Start every script with `set -euo pipefail`.
- Read config from `.env`; never hardcode secrets.
- **Destructive scripts must confirm** before acting (reset, restore).
- One job per script; print what they do.

## Scripts

| Script         | Purpose                                            | Destructive? |
|----------------|----------------------------------------------------|--------------|
| `setup.sh`     | First-time setup: env, deps, db, migrate, seed     | no           |
| `reset-db.sh`  | Drop & recreate the dev database, then migrate     | **yes**      |
| `backup.sh`    | Dump the database to a timestamped file            | no           |
| `release.sh`   | Build release artifacts (images)                   | no           |

Make them executable: `chmod +x scripts/*.sh`.
