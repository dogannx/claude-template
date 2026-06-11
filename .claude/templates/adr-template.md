# NNNN. <Decision title>

- **Status:** Proposed | Accepted | Superseded by [NNNN](NNNN-...md) | Deprecated
- **Date:** YYYY-MM-DD
- **Deciders:** <names>

## Context

What problem are we solving? What forces are at play (technical, product, team)?
Keep it factual; describe the situation that makes a decision necessary.

## Decision

The change we are making, stated clearly in active voice.
"We will ..."

## Alternatives considered

- **Option A** — pros / cons, why rejected.
- **Option B** — pros / cons, why rejected.

## Consequences

- Positive: ...
- Negative / trade-offs: ...
- Follow-ups: migrations, docs, refactors this implies.

## Memory impact (REQUIRED — keeps decisions alive)

How this ADR changes day-to-day behavior. An archived decision nobody acts on is dead.

- **`memory/project-facts.md`:** what fact to add/update (e.g. new auth method) — or "none".
- **`memory/constraints.md`:** any new hard rule this creates — or "none".
- **Code/process:** what must change so the decision is actually in force.

When this ADR is **superseded**, set Status to `Superseded by NNNN`, and update/revert the
memory entries above so `memory/` always reflects only decisions currently in force.
