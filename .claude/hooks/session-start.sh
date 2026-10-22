#!/usr/bin/env bash
# SessionStart hook: prints OS health into Claude's context at the start of a session,
# so decay (stale memory, uncommitted facts) is visible without anyone having to
# remember to check. Wired in .claude/settings.json.
set -uo pipefail
cd "${CLAUDE_PROJECT_DIR:-$(dirname "$0")/../..}" || exit 0

echo "=== AI Project OS — session start health ==="
if [ -f scripts/check-system.sh ]; then
  bash scripts/check-system.sh
else
  echo "(scripts/check-system.sh not found)"
fi
