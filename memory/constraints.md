---
last-verified: TODO
review-cadence: 90d
---

# Hard Constraints

> Non-negotiable limits. The AI must respect these without being reminded.
> Violating one is a real-world failure (legal, money, data loss, outage).

## Do-not-touch zones

- TODO (e.g. "never modify `migrations/` files that already shipped — add new ones",
  "do not change the public API contract without versioning").

## Security & compliance

- TODO (e.g. "no PII in logs", "GDPR: support data deletion", "secrets only via env").

## Performance budgets

- TODO (e.g. "p95 API latency < 200ms", "no N+1 queries in hot paths").

## Operational

- TODO (e.g. "all writes that span tables use a transaction",
  "destructive scripts require explicit confirmation").

## Boundaries for AI assistants

- Never commit secrets or `.env`.
- Never run destructive commands (drop db, force push, `rm -rf`) without explicit ask.
- Schema changes are migrations, never manual DB edits.
