---
description: Create a new Architecture Decision Record (ADR).
argument-hint: <short decision title>
---

Create a new ADR for: **$ARGUMENTS**

1. Next number in `docs/architecture/decisions/` (4 digits) →
   `NNNN-<kebab-title>.md`, from `.claude/templates/adr-template.md`.
2. Fill it from our discussion. Status `Accepted`, or `Proposed` if not yet decided.
3. Apply its "Memory impact" to `memory/` so the decision is actually in force.
4. If it supersedes an earlier ADR, mark that one `Superseded by NNNN` and revert its
   memory entries.
5. Add a row to the index in `docs/architecture/decisions/README.md`.

Show me the ADR and the memory changes.
