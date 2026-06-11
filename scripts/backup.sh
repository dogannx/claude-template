#!/usr/bin/env bash
# Dump the database to a timestamped file in ./backups (gitignored).
set -euo pipefail
cd "$(dirname "$0")/.."
set -a; [ -f .env ] && . ./.env; set +a

mkdir -p backups
STAMP="$(date +%Y%m%d-%H%M%S)"
OUT="backups/${POSTGRES_DB:-app_dev}-${STAMP}.sql"

echo "==> Dumping database to $OUT"
docker compose exec -T db pg_dump -U "${POSTGRES_USER:-app}" "${POSTGRES_DB:-app_dev}" > "$OUT"
echo "==> Backup written: $OUT"
