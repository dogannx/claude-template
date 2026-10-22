#!/usr/bin/env bash
# Optional Stop hook: a GENTLE nudge when a session of real work ends without the
# bookkeeping done. Off by default (Stop fires every turn → noisy). Enable in
# settings.json only if you want it, and keep it cheap + non-blocking.
#
# It only speaks up when there is uncommitted code BUT memory/ and notes/ are untouched
# — i.e. work happened but wasn't recorded. Otherwise it stays silent.
set -uo pipefail
cd "$(git rev-parse --show-toplevel 2>/dev/null || echo .)" || exit 0
command -v git >/dev/null && git rev-parse --git-dir >/dev/null 2>&1 || exit 0

code_changed=$(git diff --name-only -- backend/ frontend/ 2>/dev/null | grep -v '/docs/' | head -1)
record_changed=$(git diff --name-only -- memory/ notes/ todo/ 2>/dev/null | head -1)

if [ -n "$code_changed" ] && [ -z "$record_changed" ]; then
  echo "🟡 You changed code but didn't update memory/, notes/, or todo/."
  echo "   Say "notlarını al" (/kaydet) to record decisions and progress."
fi
exit 0
