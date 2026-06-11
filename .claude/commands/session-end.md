---
description: Wrap up a work session — write a session note and update todos/progress.
---

We are ending this work session. Do the following:

1. **Create a session note** at `notes/sessions/YYYY-MM-DD-<slug>.md` using the
   template in `.claude/templates/session-template.md`. Fill in:
   - What we did and why
   - Key decisions made (and whether any deserve a full ADR)
   - What changed (files/areas), tests added
   - Open items / next steps

2. **Update `todo/TODO.md`:** check off completed items, add new ones discovered,
   move finished items to `todo/done.md`.

3. **Update `notes/progress.md`:** add a dated one-line entry summarizing progress.

4. **Update `memory/` if facts changed:** new/changed auth, integrations, or invariants
   → `project-facts.md`; new bug/footgun → `known-issues.md` (remove fixed ones);
   shortcut taken → `tech-debt.md`; new hard limit → `constraints.md`.

5. **Flag follow-ups:** if any decision was architectural, propose creating an ADR
   with `/new-adr`. If anything is half-done or risky, call it out explicitly.

Keep notes concise and high-signal. Then show me a summary of what you wrote.
