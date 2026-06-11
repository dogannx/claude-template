# Commands & Trigger Phrases

The personal command set for working with the AI on this project. Two ways to invoke
each: type the **slash command**, or just say one of the **natural phrases** — the AI
recognizes both (see the trigger table in root `CLAUDE.md`). Goal: record everything with
**zero manual file editing**.

## The commands

### `/kaydet` — "Notlarını al" (the peace-of-mind command)
**Say:** "notlarını al" · "kaydet" · "oturumu kapat" · "projeden çıkıyorum" · "işim bitti" · "bugünlük bu kadar"
**Does:** detects what changed → writes a session note → updates `todo/` → appends to
`notes/progress.md` → updates `memory/` (+ bumps `last-verified`) → suggests an ADR if
needed → runs the health check → gives a clear "you can leave" summary.
**Use it when:** you're done for the session and want to walk away with nothing lost.

### `/basla` — Start the session
**Say:** "başlayalım" · "neredeyiz" · "durum ne" · "kaldığımız yer" · "bugün ne yapıyoruz"
**Does:** reads `memory/`, `progress.md`, `todo/`, last session note, git status →
briefs where you are and what to focus on today. Waits for your OK before coding.

### `/not` — Smart quick-capture
**Say:** "bunu not et: …" · "şunu hatırla: …" · "yapılacak: …" · "sorun var: …" · "bug: …" · "borç: …" · "fikir: …"
**Does:** routes the note to the right file automatically — task→`todo/TODO.md`,
bug→`memory/known-issues.md`, debt→`memory/tech-debt.md`, fact→`memory/project-facts.md`,
constraint→`memory/constraints.md`, decision→`/karar`, else→today's session note. Tells
you where it filed it.

### `/karar` — Record a decision
**Say:** "karar: …" · "şunu karar olarak yaz" · "buna karar verdik" · "kararı kaydet"
**Does:** big/architectural → creates an ADR and reflects it into `memory/`; small →
one line in `notes/decisions/`.

### `/ozet` — Summarize (no writing)
**Say:** "özetle" · "ne yaptık" · "nereye geldik"
**Does:** quick summary of this session's work, open items, next steps. Does **not** save —
use `/kaydet` to persist.

## Also available (from earlier)

- `/new-feature <ad>` — plan a feature before coding.
- `/new-adr <başlık>` — create an ADR directly (also reachable via `/karar`).
- `/session-start`, `/session-end` — English equivalents of `/basla` and `/kaydet`.

## Typical day

```
"başlayalım"            → /basla   (brief)
… geliştirme …
"bunu not et: refresh token süresi 7 gün"   → /not   (→ memory/project-facts.md)
"bug: login sonrası boş sayfa"              → /not   (→ memory/known-issues.md)
"karar: auth için JWT kullanıyoruz"         → /karar (→ ADR + memory)
…
"notlarını al, çıkıyorum"                   → /kaydet (her şeyi kaydet → çık)
```

## Adding your own

Drop a new `.claude/commands/<isim>.md` (frontmatter `description`, body = instructions),
then add a row to the trigger table in root `CLAUDE.md` so the natural phrase works too.
Keep this file in sync — it's the catalog you'll actually read.
