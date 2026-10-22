---
description: "Notlarını al ve oturumu güvenle kapat — oturum notu, TODO, progress, memory. Tetikleyiciler: \"notlarını al\", \"kaydet\", \"oturumu kapat\", \"projeden çıkıyorum\", \"işim bitti\", \"bugünlük bu kadar\"."
---

# /kaydet — Her şeyi kaydet, oturumu kapat

Amaç: kullanıcı hiçbir dosyayı elle düzenlemeden çıkabilsin, bağlam kaybolmasın.

1. Bu oturumda ne yapıldığını, ne kararlaştırıldığını ve neyin yarım kaldığını çıkar
   (konuşma + `git status` / `git diff --stat`).
2. `notes/sessions/YYYY-MM-DD-<slug>.md` yaz (`.claude/templates/session-template.md`).
3. Geri kalanını CLAUDE.md §3 "Where to write what" tablosuna göre yerleştir: biten
   görevler `done.md`'ye, yeni görevler TODO'ya, küçük kararlar decision log'a, yeni
   gerçek/bug/borç/kısıt ilgili `memory/` dosyasına; `notes/progress.md`'ye tek satır.
   Mimari bir karar çıktıysa ADR'yi öner, kendiliğinden yazma.
4. `bash scripts/check-system.sh` çalıştır.
5. Kısa bir "çıkabilirsin" özeti ver: ne nereye kaydedildi, ne commit bekliyor, elle
   ilgilenilmesi gereken bir şey varsa o.
