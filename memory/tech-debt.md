---
last-verified: TODO
review-cadence: 60d
---

# Technical Debt Register

> Intentional shortcuts we took on purpose, and what it costs to repay them.
> This makes debt visible and deliberate instead of forgotten. Not for bugs
> (those go in `known-issues.md`).

| Area      | Shortcut taken                 | Why                | Repay cost | Trigger to fix         |
|-----------|--------------------------------|--------------------|------------|------------------------|
| —         | (none yet)                     | —                  | —          | —                      |

<!--
Example:
| auth      | No refresh-token rotation       | ship MVP faster    | ~1 day     | before public launch   |
| db        | No indexes on reports table     | low data volume    | ~2 hrs     | when >100k rows        |
-->

## Notes

- Review this list during planning. Pay debt when its **trigger** fires, not before.
- A new shortcut taken? Add a row in the same session you take it.
