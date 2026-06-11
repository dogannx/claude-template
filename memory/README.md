# memory/ — Project Truth Layer

The **always-true, always-loaded** facts an AI assistant (or human) must know before
touching this project. This is NOT documentation and NOT history — it is the small,
curated set of current realities that change rarely but matter constantly.

## How it differs from `docs/` and `notes/`

| Layer      | Question it answers        | Size       | Read when        | Changes      |
|------------|----------------------------|------------|------------------|--------------|
| `memory/`  | "What is true RIGHT NOW?"  | tiny       | **every time**   | rarely       |
| `docs/`    | "How does this work?"      | large      | on demand        | occasionally |
| `notes/`   | "What happened?"           | grows      | for history      | append-only  |

If `memory/` grows large, you're putting reference material in it — move that to `docs/`.
Memory must stay scannable in under a minute.

## Files

- `project-facts.md` — invariant truths: auth method, environments, integrations, the
  decisions currently in effect.
- `known-issues.md` — bugs, gotchas, and footguns to avoid right now.
- `tech-debt.md` — intentional shortcuts and what it'll take to repay them.
- `constraints.md` — hard limits: compliance, performance budgets, "do not touch" zones.

## Freshness (`last-verified`)

Each file carries frontmatter:

```yaml
---
last-verified: 2026-06-11   # set to today whenever you confirm the file is still accurate
review-cadence: 30d         # how often it should be re-checked
---
```

`scripts/check-system.sh` (run by the SessionStart hook and CI) warns when a file hasn't
been verified within its cadence. **Review cadence by file:**

| File              | Cadence | Why                                            |
|-------------------|---------|------------------------------------------------|
| `known-issues.md` | 14 days | bugs come and go fast                          |
| `project-facts.md`| 30 days | core facts; check monthly                      |
| `tech-debt.md`    | 60 days | debt moves slowly                              |
| `constraints.md`  | 90 days | hard limits rarely change                      |

## Maintenance rule

Update `memory/` the moment a fact changes (new auth flow, resolved bug, new constraint),
and bump `last-verified` when you confirm a file is still accurate. A stale memory file is
worse than none — the AI will confidently act on wrong facts. The `/session-end` flow and
the SessionStart health check both prompt you.
