---
name: manage-release
description: Run the release process — verify, version, changelog, tag, build, and record — consistently and safely.
---

# Manage a release

Use to cut a release. Follow semantic versioning (`docs/conventions.md`).

## Pre-flight

- `main` is green: `make test`, `make lint` pass.
- All intended changes merged; `memory/known-issues.md` has no blockers.
- Migrations included and tested (up + down).

## Steps

1. **Decide version:** major/minor/patch from the changes since the last tag
   (`git log --oneline <last-tag>..HEAD`).
2. **Changelog:** summarize changes (group by feat/fix/etc.). Add to `CHANGELOG.md` if
   present, else the release notes.
3. **Bump version** wherever it's declared.
4. **Tag:** `git tag -a vX.Y.Z -m "..."` and push the tag.
5. **Build artifacts:** `scripts/release.sh` (build images, etc.).
6. **Deploy** per environment process; run migrations as part of deploy.
7. **Record:** add an entry to `notes/progress.md` and move shipped items to
   `todo/done.md`.

## Safety

- Never tag with failing CI. Confirm before any production deploy step.
- If a release introduces a known limitation, log it in `memory/tech-debt.md`.
