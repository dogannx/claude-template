# Onboarding

Welcome. This gets a new teammate (or AI assistant) productive fast.

## 1. Read these, in order

1. Root `CLAUDE.md` — the map of the repo.
2. `docs/architecture/overview.md` — how the system fits together.
3. `docs/tech-stack.md` — what we use and why.
4. `docs/coding-standards.md` + `docs/conventions.md` — how we write & ship code.
5. Skim `notes/progress.md` and the latest `notes/sessions/` note — current state.

## 2. Local setup

```bash
cp .env.example .env     # fill in secrets
make up                  # start the stack
make migrate             # apply DB migrations
make test                # confirm everything is green
```

- Backend: http://localhost:8080
- Frontend: http://localhost:5173

## 3. How we work day-to-day (AI-assisted)

This project is built to work with Claude. The loop is:

> Tip: you can use natural Turkish phrases instead of slash commands — e.g. say
> "notlarını al" to save & close, or "başlayalım" to start. Full list:
> [`docs/commands.md`](commands.md).

### Session start
- Run `/basla` (or say "başlayalım"). Claude reads progress + todos and proposes a focus.
- Confirm the focus before coding.

### During development
- Use `/new-feature <name>` to plan before building.
- Keep diffs small; follow the standards.
- Significant architectural choice? Run `/new-adr <title>`.

### Session end
- Run `/kaydet` (or say "notlarını al" / "projeden çıkıyorum"). Claude writes a session
  note, updates `todo/`, `progress.md`, and `memory/`, then gives a "safe to leave" summary.
- Open a PR using the PR template.

## 4. Where things live

| You want to...                  | Go to                                   |
|---------------------------------|-----------------------------------------|
| Understand the architecture     | `docs/architecture/`                    |
| Know why a decision was made    | `docs/architecture/decisions/`          |
| See current tasks               | `todo/TODO.md`                          |
| See recent work                 | `notes/sessions/`, `notes/progress.md`  |
| Add reusable AI commands/skills | `.claude/commands/`, `.claude/skills/`  |

## 5. Definition of done

- Code follows standards, tests pass, migrations included.
- Docs/ADR updated if behavior or architecture changed.
- Session note + todo updated. PR opened with template filled.
