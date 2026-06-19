---
description: "Notlarını al ve oturumu güvenle kapat — her şeyi otomatik kaydeder, projeden gönül rahatlığıyla çıkabilirsin."
---

# /kaydet — Save everything & close the session safely

Tetikleyiciler: **"notlarını al"**, **"kaydet"**, **"oturumu kapat"**,
**"projeden çıkıyorum"**, **"işim bitti"**, **"bugünlük bu kadar"**.

Amaç: kullanıcı hiçbir şeyi elle yazmadan, çıktığında hiçbir bağlam kaybolmasın.
Aşağıdakileri **sırayla ve otomatik** yap, sonra tek bir "çıkabilirsin" özeti ver.

1. **Ne değişti, tespit et:** `git status` + `git diff --stat` ve bu oturumdaki
   konuşma/işleri tara. Neyin yapıldığını, neyin yarım kaldığını çıkar.

2. **Oturum notu yaz:** `notes/sessions/YYYY-MM-DD-<slug>.md` —
   `.claude/templates/session-template.md` şablonuyla. Hedef, yapılanlar, kararlar,
   değişen dosyalar, açık işler, "sonraki sefer için notlar".

3. **TODO + küçük kararlar (checklist — hepsini yap):**
   - [ ] `todo/TODO.md`'de biten görevleri `[x]` ile işaretle.
   - [ ] Her `[x]` satırını `todo/done.md` altına tarihli olarak **gerçekten taşı**
         (TODO'da bırakma — taşımayı atlamak en sık unutulan adım).
   - [ ] Bu oturumda çıkan yeni görevleri TODO'ya ekle.
   - [ ] ADR'ye değmeyen ama hatırlanmaya değer küçük teknik karar varsa
         → `notes/decisions/README.md`'ye tarihli tek satır ekle.
         (Eşik: [`notes/decisions/README.md`](../../notes/decisions/README.md))

4. **İlerleme kaydet:** `notes/progress.md`'ye tarihli tek satır ekle.

5. **Memory güncelle (gerekirse):** auth/gerçek değiştiyse `project-facts.md`;
   yeni bug → `known-issues.md`; bilinçli kestirme → `tech-debt.md`; yeni sert kural →
   `constraints.md`. Doğruladığın dosyaların `last-verified` alanını **bugüne** çek.

6. **Mimari karar var mıydı?** Varsa `/karar` ile ADR oluşturmayı öner (zorla yazma,
   sor).

7. **Sağlık kontrolü:** `bash scripts/check-system.sh` çalıştır, uyarı varsa göster.

8. **"Çıkabilirsin" özeti:** ne kaydedildi, ne commit'lenmeyi bekliyor, neyin elle
   ilgilenmesi gerekiyor (varsa) — net ve kısa. Kullanıcı bunu okuyup huzurla çıkabilmeli.
