# <Project Name>

> Short tagline describing the product.

This repository is bootstrapped from an **AI-assisted project template** designed to
work well with Claude (and other AI assistants) while staying clean for human teams.

## Quick start

```bash
cp .env.example .env        # fill in values
make up                     # start everything with docker-compose
```

- Backend runs on `http://localhost:8080`
- Frontend runs on `http://localhost:5173`

## Structure

```
.
├── CLAUDE.md        # AI context entry point (read this first)
├── backend/         # Go + Fiber API
├── frontend/        # Vue SPA
├── docs/            # Architecture, standards, ADRs
├── notes/           # Session notes, decisions, progress
├── todo/            # Task lists
└── .claude/         # Reusable commands, skills, templates
```

## Documentation

| Topic               | File                                            |
|---------------------|-------------------------------------------------|
| Onboarding          | `docs/onboarding.md`                             |
| Architecture        | `docs/architecture/overview.md`                 |
| Tech stack          | `docs/tech-stack.md`                             |
| Coding standards    | `docs/coding-standards.md`                       |
| Conventions         | `docs/conventions.md`                            |
| Decisions (ADRs)    | `docs/architecture/decisions/`                  |

## Working with AI assistants

This repo uses `CLAUDE.md` files as context anchors. When you (or an assistant) start
work, begin with the root `CLAUDE.md`, then the relevant sub-project `CLAUDE.md`.

## License

TODO
