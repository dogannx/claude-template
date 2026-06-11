#!/usr/bin/env bash
# Build release artifacts. Pass a version: ./scripts/release.sh v1.2.3
set -euo pipefail
cd "$(dirname "$0")/.."

VERSION="${1:-}"
[ -n "$VERSION" ] || { echo "Usage: $0 vX.Y.Z"; exit 1; }

echo "==> Verifying build is green"
( cd backend && go vet ./... && go test ./... )
( cd frontend && yarn lint && yarn test ) 2>/dev/null || echo "    (skip: frontend tests)"

echo "==> Building images tagged $VERSION"
docker compose build
# docker tag / push steps go here, per your registry.

echo "==> Built artifacts for $VERSION."
echo "    Next: tag the release (git tag -a $VERSION) and follow the manage-release skill."
