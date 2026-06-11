# references/ — Knowledge Base (use with discipline)

A place for **project-relevant** reference material: curated snippets, cheat-sheets,
and short tech summaries you reach for repeatedly *in this project*.

## ⚠️ Important distinction — read this first

There are two kinds of "personal knowledge", and mixing them is the #1 way this folder
rots:

1. **Project-specific references** → belong **here** (e.g. "our pagination snippet",
   "how our auth token is structured", "this project's Postgres tuning notes").
2. **Cross-project personal knowledge** (your years of generic Go/Vue snippets, generic
   tech summaries) → do **NOT** put these in every project repo. Keep them in ONE
   external place: a personal `knowledge` repo, a global `~/.claude/` directory, or a
   notes app. Otherwise each project carries a divergent, stale copy.

## Why keep `references/` at all

- **Pros:** fast recall of project-specific patterns; the AI can read them as context;
  versioned alongside the code they relate to.
- **Risks:** turns into a junk drawer; goes stale silently; duplicates `docs/`. If a
  reference describes how *this system* works, it probably belongs in `docs/` instead.

## Layout (optional)

```
references/
├── snippets/      # small, reusable code snippets specific to this project
└── cheatsheets/   # short command/pattern references
```

## Maintenance rule

This folder is **optional and on-demand**. Don't create it empty in every project.
Add it only when you have something concrete to put in it, and prune yearly.
