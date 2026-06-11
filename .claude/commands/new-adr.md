---
description: Create a new Architecture Decision Record (ADR).
argument-hint: <short decision title>
---

Create a new ADR for: **$ARGUMENTS**

1. Find the highest existing number in `docs/architecture/decisions/` and use the next
   one, zero-padded to 4 digits (e.g. `0007-`).
2. Copy `.claude/templates/adr-template.md` to
   `docs/architecture/decisions/NNNN-<kebab-title>.md`.
3. Fill in context, the decision, considered alternatives, and consequences based on
   our discussion. Set status to `Accepted` (or `Proposed` if not yet decided).
4. **Complete the "Memory impact" section** and actually apply it: if the decision adds
   or changes a project fact or constraint, update `memory/project-facts.md` /
   `memory/constraints.md` now so the decision is in force, not just archived.
5. If this ADR **supersedes** an earlier one: set the old ADR's status to
   `Superseded by NNNN`, update the index, and revert any now-obsolete memory entries.
6. Add a line to `docs/architecture/decisions/README.md` index.

Show me the final ADR plus the memory changes for review.
