# Claude Code hooks

Hook scripts that run on Claude Code lifecycle events. Configured in `.claude/settings.json`.

## Active hooks

- **SessionStart** → `session-start.sh` — runs `scripts/check-system.sh` and surfaces
  stale `memory/` and uncommitted facts at the start of every session.

## Optional / parked hooks

The Stop hook is **off by default** because it fires every turn, which is noisy. If you
want a gentle "did you `/kaydet`?" nudge at session end, paste this block inside the
`hooks` object in `.claude/settings.json`:

```json
"Stop": [
  {
    "hooks": [
      { "type": "command", "command": "bash .claude/hooks/session-end-reminder.sh" }
    ]
  }
]
```

The script (`session-end-reminder.sh`) already exists; it just isn't wired up.

## Adding a new hook

1. Drop the script under `.claude/hooks/<name>.sh` and `chmod +x` it.
2. Reference it under the correct event in `.claude/settings.json` — only real Claude
   Code event names are accepted (`SessionStart`, `Stop`, `PreToolUse`, etc.).
   Don't stash comments or "disabled" blocks under fake event keys — `/doctor` flags
   them. Use this README for any documentation that can't live in JSON.
