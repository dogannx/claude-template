# Quick Decisions Log

A lightweight log for **small** decisions that don't warrant a full ADR but are worth
remembering (e.g. "use X library for Y", "name this field Z because...").

## When does it belong here vs. in an ADR?

**Write a full ADR** (`docs/architecture/decisions/NNNN-…md`) if **any** is true:

- Changes the schema or an external API contract
- Shifts data ownership or a responsibility boundary between components
- Is costly to undo (would need a "supersede" entry to reverse)
- Expected to be load-bearing for 6+ months

**Log it here** (one dated line) if **all** are true:

- Library / version / tool choice (e.g. Mailhog vs MailDev, lib version pin)
- Port, file name, config value, naming convention call
- Local-only dev choice (testcontainers, fixtures, etc.)
- Easily revisitable — "this for now, may change later"

If in doubt → log here first. Promote to an ADR later if it turns out to be load-bearing.

## Log

- 2026-06-11 — Keep template docs in English for portability; chat can be in any language.

<!--
Format:
- YYYY-MM-DD — decision — one-line rationale.
-->
