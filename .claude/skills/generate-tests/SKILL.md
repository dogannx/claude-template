---
name: generate-tests
description: Write meaningful tests for given Go code (or Vue components) — happy paths, error paths, and edge cases — matching project test conventions.
---

# Generate tests

Use to add or improve tests. Goal is **meaningful coverage**, not coverage numbers.
Tests must actually assert behavior and be deterministic.

## Steps

1. Identify the unit under test and its dependencies; check existing test style in the
   package to match conventions.
2. Enumerate cases: happy path, each error/edge case, boundary values, and any invariant
   from `memory/project-facts.md`.
3. **Go:** table-driven tests where it fits; name `TestXxx`; use subtests (`t.Run`).
   Mock/stub at the repository boundary; use a test DB or fakes for integration.
4. **Vue:** component tests for rendering + interaction; mock the `src/api` layer.
5. Avoid testing implementation details; assert observable behavior.
6. Run them: `go test ./...` / `yarn test`. Ensure they fail when the code is wrong
   (sanity-check by breaking the code mentally).

## Definition of done

- Tests pass, are deterministic, and cover the error paths — not just the happy path.
