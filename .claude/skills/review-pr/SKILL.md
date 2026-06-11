---
name: review-pr
description: Review a pull request or the current diff for correctness, security, conventions, and tests before merge.
---

# Review a pull request

Use to review the current branch diff (or a given PR). Be specific and cite
`file:line`. Prioritize real bugs over style nits.

## Gather context

- `git diff main...HEAD` (or the PR diff) and `git log --oneline main..HEAD`.
- Re-read `docs/coding-standards.md`, `docs/conventions.md`, and `memory/constraints.md`.

## Review dimensions (in priority order)

1. **Correctness:** logic errors, edge cases, nil/zero handling, error paths, race
   conditions, `context` propagation.
2. **Security:** input validation, authz checks, SQL injection (parameterized queries),
   secrets, PII in logs — cross-check `memory/constraints.md`.
3. **Data:** migrations reversible? N+1 queries? transactions for multi-step writes?
4. **Tests:** new behavior covered? error paths tested? do they actually assert?
5. **Conventions:** layering respected, naming, commit style, no dead code.
6. **Docs:** ADR/memory updated if architecture or facts changed?

## Output

- A short summary verdict (approve / changes requested).
- Findings grouped by severity (blocker / should-fix / nit), each with `file:line` and a
  concrete suggestion.
