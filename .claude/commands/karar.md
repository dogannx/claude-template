---
description: "Bir kararı kaydet — büyükse ADR, küçükse hızlı karar günlüğü; memory'ye de yansıt."
argument-hint: <karar başlığı/özeti>
---

# /karar — Record a decision

Tetikleyiciler: **"karar: …"**, **"şunu karar olarak yaz"**, **"buna karar verdik"**,
**"kararı kaydet"**.

Karar: **$ARGUMENTS**

1. **Boyutu belirle:**
   - **Mimari / kalıcı etki** → `/new-adr` akışını çalıştır: ADR oluştur, "Memory impact"
     bölümünü doldur ve gerçekten `memory/`'ye yansıt; index'i güncelle.
   - **Küçük karar** → `notes/decisions/README.md`'ye tarihli tek satır ekle.
2. Karar bir gerçeği/kısıtı değiştiriyorsa ilgili `memory/` dosyasını güncelle.
3. Ne kaydettiğini ve nereye yazdığını göster.
