#!/usr/bin/env bash
# DESTRUCTIVE: drops and recreates the dev database, then re-runs migrations.
set -euo pipefail
cd "$(dirname "$0")/.."
set -a; [ -f .env ] && . ./.env; set +a

DB="${POSTGRES_DB:-app_dev}"

echo "!! This will DROP and recreate database '$DB'. All data will be lost."
read -r -p "Type the database name to confirm: " CONFIRM
[ "$CONFIRM" = "$DB" ] || { echo "Aborted."; exit 1; }

echo "==> Dropping & recreating $DB"
docker compose exec -T db psql -U "${POSTGRES_USER:-app}" -d postgres \
  -c "DROP DATABASE IF EXISTS \"$DB\";" -c "CREATE DATABASE \"$DB\";"

echo "==> Re-running migrations"
( cd backend && go run ./cmd/migrate up )

echo "==> Database reset complete."
