#!/usr/bin/env bash
# Health check for the AI Project OS itself (not the app).
# Surfaces decay: stale memory and uncommitted memory changes.
# Used by: the SessionStart hook, the git pre-commit hook, and CI.
# Exit code: 0 = ok or warnings only; 1 = hard problems (CI may choose to fail).
set -uo pipefail
cd "$(dirname "$0")/.."

# Thresholds (days) before a memory file is considered stale.
STALE_FACTS=${STALE_FACTS:-30}        # project-facts: review monthly
STALE_CONSTRAINTS=${STALE_CONSTRAINTS:-90}
STALE_DEBT=${STALE_DEBT:-60}
STALE_ISSUES=${STALE_ISSUES:-14}      # known-issues: review often

warn=0
say()  { printf '  %s\n' "$1"; }
flag() { printf '  ⚠️  %s\n' "$1"; warn=1; }

days_since() { # echo days since an ISO date, or empty if unparseable
  python3 -c "import sys,datetime as d;print((d.date.today()-d.date.fromisoformat(sys.argv[1].strip())).days)" "$1" 2>/dev/null
}

echo "── Memory freshness ───────────────────────────────"
check_memory() {
  local file="$1" limit="$2"
  [ -f "$file" ] || { say "$file: missing (ok if unused)"; return; }
  local lv
  lv=$(grep -m1 '^last-verified:' "$file" | sed -e 's/last-verified:[[:space:]]*//' -e 's/#.*//' -e 's/[[:space:]]*$//')
  if [ -z "$lv" ] || [ "$lv" = "TODO" ]; then
    flag "$file: no 'last-verified' date set."
    return
  fi
  local d; d=$(days_since "$lv")
  if [ -z "$d" ]; then say "$file: last-verified '$lv' (unparseable, skipped)"; return; fi
  if [ "$d" -gt "$limit" ]; then
    flag "$file: last verified $d days ago (> $limit). Re-verify and bump last-verified."
  else
    say "$file: verified $d days ago (ok)."
  fi
}
check_memory memory/project-facts.md "$STALE_FACTS"
check_memory memory/constraints.md   "$STALE_CONSTRAINTS"
check_memory memory/tech-debt.md     "$STALE_DEBT"
check_memory memory/known-issues.md  "$STALE_ISSUES"

echo "── Uncommitted project-memory ─────────────────────"
if command -v git >/dev/null && git rev-parse --git-dir >/dev/null 2>&1; then
  if ! git diff --quiet -- memory/ 2>/dev/null; then
    flag "memory/ has uncommitted changes — facts may have moved; commit them."
  else
    say "memory/ clean."
  fi
else
  say "(not a git repo — skipping git checks)"
fi

echo "───────────────────────────────────────────────────"
if [ "$warn" -eq 1 ]; then
  echo "System check: warnings above. Address them at /kaydet."
else
  echo "System check: all good."
fi
exit 0
