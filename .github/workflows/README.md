# CI / CD

`ci.yml` enforces the rules that the docs only describe, so they can't silently rot:

| Job          | Enforces                                                        |
|--------------|----------------------------------------------------------------|
| os-health    | `scripts/check-system.sh` — surfaces memory/notes decay        |
| commit-lint  | Conventional Commits (`docs/conventions.md`)                   |
| backend      | gofmt, `go vet`, build, tests                                  |
| frontend     | lint, build, tests                                             |

Both the backend and frontend scaffolds ship with passing tests, so CI is green on a
fresh copy of the template.

## Recommended branch protection (Settings → Branches → main)

- Require PRs before merging.
- Require status checks: `backend`, `frontend`, `commit-lint`.
- (Optional) Require a review approval — for teams.

## Local equivalent

The git hooks in `.githooks/` run the same checks before commit. Install with
`make hooks-install`. CI is the backstop; hooks give fast local feedback.
