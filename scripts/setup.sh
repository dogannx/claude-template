#!/usr/bin/env bash
# First-time project setup. Safe to re-run (idempotent).
set -euo pipefail
cd "$(dirname "$0")/.."

echo "==> Ensuring .env exists"
[ -f .env ] || { cp .env.example .env; echo "    created .env from example — fill in secrets"; }

echo "==> Installing git hooks"
git rev-parse --git-dir >/dev/null 2>&1 && git config core.hooksPath .githooks && echo "    core.hooksPath=.githooks" || echo "    (not a git repo yet — run 'git init' then 'make hooks-install')"

echo "==> Starting database"
docker compose up -d db

echo "==> Installing backend deps"
( cd backend && go mod download ) 2>/dev/null || echo "    (skip: backend not initialized yet)"

echo "==> Installing frontend deps"
( cd frontend && yarn install ) 2>/dev/null || echo "    (skip: frontend not initialized yet)"

echo "==> Applying migrations"
( cd backend && go run ./cmd/migrate up ) 2>/dev/null || echo "    (skip: no migrations yet)"

echo "==> Done. Run 'make up' to start the stack."
