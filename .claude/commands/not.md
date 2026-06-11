---
description: "Hızlı not yakala — içeriğe göre doğru dosyaya otomatik yerleştirir (todo, bug, borç, gerçek, fikir)."
argument-hint: <not içeriği>
---

# /not — Smart quick-capture

Tetikleyiciler: **"bunu not et: …"**, **"şunu hatırla: …"**, **"kaydet şunu: …"**,
**"yapılacak: …"**, **"sorun var: …"**, **"bug: …"**, **"borç: …"**, **"fikir: …"**.

İçerik: **$ARGUMENTS**

Notu **doğru yere** yönlendir (intent'i içerikten anla). Tek satır ekle, gerekirse
tarihle, sonra **nereye yazdığını** söyle:

| İçerik türü                      | Hedef dosya                  |
|----------------------------------|------------------------------|
| Yapılacak iş / görev             | `todo/TODO.md`               |
| Hata / bug / kırık davranış      | `memory/known-issues.md`     |
| Bilinçli teknik borç / kestirme  | `memory/tech-debt.md`        |
| Kalıcı gerçek (auth, ortam, vb.) | `memory/project-facts.md`    |
| Sert kısıt / kural               | `memory/constraints.md`      |
| Mimari karar                     | `/karar` öner (ADR)          |
| Geçici / sınıflandırılamayan     | bugünün `notes/sessions/` notu |

Belirsizse hangi dosyaya gideceğini **bir cümleyle sor**, sonra yaz. Memory dosyasına
yazdıysan `last-verified`'ı bugüne çek.
