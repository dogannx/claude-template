# Architecture Decision Records (ADRs)

This directory holds **ADRs** — short documents capturing significant architectural
decisions, their context, and consequences. They are immutable history: instead of
editing a past decision, supersede it with a new one.

## Why

- New teammates (and AI assistants) understand *why* the system is the way it is.
- We avoid re-litigating settled decisions.
- Context survives team changes.

## How to add one

Run `/new-adr <title>` or copy `.claude/templates/adr-template.md` to the next number.

## Status values

- **Proposed** — under discussion, not yet binding.
- **Accepted** — in force. Its "Memory impact" must be reflected in `memory/`.
- **Superseded by NNNN** — replaced; no longer in force. Remove its effects from `memory/`.
- **Deprecated** — no longer relevant, not replaced.

> Rule: `memory/` reflects only **Accepted** decisions. When status changes, update `memory/`.

## Index

| #    | Title                                  | Status   | Memory touched          |
|------|----------------------------------------|----------|-------------------------|
| 0001 | Record architecture decisions          | Accepted | none (process only)     |
