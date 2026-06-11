# Reusable Skills

Store project-specific or cross-project **skills** here. A skill is a folder with a
`SKILL.md` that teaches Claude how to perform a repeatable, specialized task in this
codebase (e.g. "add a new Fiber endpoint end-to-end", "create a bun migration",
"scaffold a Vue feature module").

## Structure

```
.claude/skills/
└── <skill-name>/
    └── SKILL.md      # name, description (when to use), and the procedure
```

## SKILL.md frontmatter

```markdown
---
name: add-endpoint
description: Scaffold a new Fiber endpoint with handler, route, bun model, and test.
---

Step-by-step instructions Claude should follow...
```

## Included skills (ready to use)

| Skill              | When to use                                              |
|--------------------|---------------------------------------------------------|
| `add-endpoint`     | New Fiber endpoint end-to-end (model→repo→service→handler) |
| `create-migration` | Safe bun/PostgreSQL schema migration (up + down)        |
| `review-pr`        | Review the current diff/PR for bugs, security, tests    |
| `generate-tests`   | Write meaningful tests (happy + error + edge paths)     |
| `debug-production` | Systematic incident analysis: symptom → root cause → fix |
| `manage-release`   | Verify → version → changelog → tag → build → record     |

Keep skills small and composable. Promote a skill to a shared/global location
(`~/.claude/skills/`) once it proves useful across multiple projects.
