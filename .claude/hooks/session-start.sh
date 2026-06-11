#!/usr/bin/env bash
# SessionStart hook: prints OS health into Claude's context at the start of a session,
# so decay (stale memory, uncommitted facts, missing notes) is visible without anyone
# having to remember to check. Wired in .claude/settings.json.
set -uo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || dirname "$0")/.." 2>/dev/null || true

echo "=== AI Project OS — session start health ==="
if [ -x scripts/check-system.sh ]; then
  bash scripts/check-system.sh
else
  echo "(scripts/check-system.sh not found or not executable)"
fi
echo "Reminder: read memory/ before coding; run /session-end before stopping."
