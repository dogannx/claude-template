# 0001. Record architecture decisions

- **Status:** Accepted
- **Date:** 2026-06-11
- **Deciders:** Project owner

## Context

We need a durable, low-friction way to capture significant technical decisions so the
reasoning behind the system survives over time and across team members and AI sessions.

## Decision

We will use **Architecture Decision Records (ADRs)**, stored as numbered Markdown files
in `docs/architecture/decisions/`. Each significant decision gets one ADR. ADRs are
append-only: a changed decision is recorded as a new ADR that supersedes the old one.

## Alternatives considered

- **No formal record** — decisions live in chat/memory and are lost. Rejected.
- **Wiki / external tool** — drifts from the code and needs separate access. Rejected
  in favor of keeping decisions in-repo, versioned with the code.

## Consequences

- Positive: durable context; easier onboarding; less re-debate.
- Trade-off: small discipline cost to write one. Mitigated by the `/new-adr` command.
