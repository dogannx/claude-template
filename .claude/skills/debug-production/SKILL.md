---
name: debug-production
description: Systematically analyze a production error or incident — from symptom to root cause to fix and prevention.
---

# Debug a production issue

Use when investigating a production error/incident. Be methodical; don't guess-patch.

## Steps

1. **Capture the symptom:** exact error, stack trace, timestamp, affected env, scope
   (one user vs all). Check `memory/known-issues.md` first — it may be known.
2. **Reproduce:** find the smallest reliable repro (request, input, state). If you can't
   reproduce, gather more data before changing code.
3. **Locate:** trace from the error to the code path (handler → service → repository).
   Check recent deploys/migrations (`git log`, `notes/progress.md`).
4. **Root cause:** explain *why* it happens, not just where. Distinguish trigger from
   underlying cause.
5. **Fix:** smallest correct change. Add a **regression test** that fails before the fix.
6. **Prevent:** if it reveals a class of bug, note it. Update `memory/known-issues.md`
   (or remove it once fixed) and `tech-debt.md` if the cause was a known shortcut.
7. **Document:** record the incident in a session note; propose an ADR if it changes
   architecture.

## Output

Root-cause statement, the fix (with test), and any follow-up prevention items.
