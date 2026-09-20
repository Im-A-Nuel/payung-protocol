# MVP Roadmap: Payung

Terakhir diperbarui: 17 September 2026
Deadline submission: Rabu, 30 September 2026, 23:59 WIB
Tim: solo

## MVP Definition

MVP dianggap jadi kalau satu alur ini bisa direkam tanpa edit: driver login dengan Google, ambil IDRP dari faucet, beli polis Bantul tanpa bayar gas, admin menyuntik hujan 31 mm untuk kemarin, payout Rp 25.000 masuk ke dompet driver, dashboard menampilkan payout dengan link BscScan dan penjelasan AI dalam bahasa Indonesia, dan tabel premi menunjukkan alasan kenapa Bantul lebih mahal dari Solo di bulan ini.

Kalau salah satu dari itu belum jalan di hari ke-7, semua fitur lain berhenti sampai itu jalan.

## Timeline Overview

| Fase | Hari | Tanggal | Goal |
|---|---|---|---|
| 1. Kontrak | 1 sampai 2 | Kam 17, Jum 18 Sep | PayungPool terdeploy dan terverifikasi, 10 test lolos |
| 2. Oracle & API | 3 sampai 4 | Sab 19, Min 20 Sep | Hujan masuk on-chain, settle jalan, API baca dari DB |
| 3. Frontend | 5 sampai 7 | Sen 21 sampai Rab 23 Sep | Alur beli polis sampai payout end-to-end (checkpoint) |
| 4. AI | 8 sampai 9 | Kam 24, Jum 25 Sep | Tabel premi dan penjelasan payout hidup |
| 5. Polish | 10 | Sab 26 Sep | UI, README, diagram, seed data |
| 6. Video & submit | 11 sampai 12 | Min 27, Sen 28 Sep | Video final, form terisi, edit code disimpan |
| 7. Buffer | 13 | Sel 29, Rab 30 Sep | Perbaikan, edit submission |

Catatan: hari 11 sampai 13 di atas sengaja diberi 4 tanggal untuk 3 slot agar ada satu hari kosong nyata untuk hal tak terduga.

---

## Phase 1: Kontrak
**Durasi**: 2 hari
**Goal**: Uang bisa masuk pool, hujan bisa dilaporkan, payout bisa dieksekusi, semuanya teruji.

### Tasks
- [ ] Init monorepo: `contracts/`, `backend/`, `web/`, `docs/`, `CLAUDE.md`
- [ ] Foundry init, install OpenZeppelin
- [ ] `IDRP.sol` dengan `faucet` (FAUCET_ROLE) dan `mint` (admin)
- [ ] `PayungPool.sol` sesuai SCHEMA.md: roles, Zone, Policy, buyPolicy, submitRainfall, settle, pruneExpired, fundPool, view functions, custom errors
- [ ] 10 test case di SCHEMA.md lolos, plus fuzz sederhana untuk `today()` dan week boundary
- [ ] `Deploy.s.sol`: deploy IDRP, deploy Pool, buat 5 zona, grant ORACLE_ROLE ke alamat oracle, mint 5.000.000 IDRP ke deployer, fundPool 2.000.000 IDRP
- [ ] Deploy ke opBNB testnet, verify di BscScan, simpan alamat di `docs/README.md` dan `.env.example`
- [ ] GitHub Actions: `forge test` on push

**Definition of done fase 1**: alamat kontrak terverifikasi bisa dibuka di BscScan dan `forge test` hijau.

---

## Phase 2: Oracle & API
**Durasi**: 2 hari
**Goal**: Backend Go bisa membaca hujan, menulis ke chain, memicu settle, dan menyajikan data ke frontend.

### Tasks
- [x] Go module, chi, pgx, sqlc, migrasi tabel di SCHEMA.md
- [x] `abigen` bindings dari ABI Foundry
- [x] `internal/rain`: client Open-Meteo Forecast (`past_days=2`, `daily=precipitation_sum`, `timezone=Asia/Jakarta`) dan Archive (3 tahun)
- [x] `internal/chain`: signer oracle, `SubmitRainfall`, `Settle`, `SetPremium`, `Faucet`, retry dengan backoff, tulis ke `oracle_runs`
- [x] `cmd/oracle`: untuk setiap zona, ambil hujan kemarin, upsert `rain_observations`, submit, settle jika >= ambang, idempotent lewat `oracle_runs` unique
- [x] `internal/indexer`: poll event `PolicyBought`, `PayoutSent`, `PayoutSkipped`, `PolicyExpired` dari `last_block`, tulis ke `policies` dan `payouts`
- [x] `cmd/api`: `GET /zones`, `GET /zones/:id/rain`, `GET /drivers/:address/policy`, `GET /drivers/:address/payouts`, `POST /faucet`
- [x] `POST /admin/simulate-rain` dan `POST /admin/oracle/run` dengan `X-Admin-Key`, nonaktif saat `ENV=production`
- [ ] Deploy ke Railway dengan Postgres, cron harian 06:00 WIB untuk `cmd/oracle`
- [x] GitHub Actions: `go test ./...`

**Definition of done fase 2**: `curl POST /admin/simulate-rain` menghasilkan `PayoutSent` di BscScan dan `GET /drivers/:address/payouts` menampilkannya.

---

## Phase 3: Frontend (Checkpoint Hari 7)
**Durasi**: 3 hari
**Goal**: Driver bisa menjalankan seluruh alur dari HP tanpa tahu apa itu gas.

### Tasks
- [ ] Next.js 15, Tailwind, Privy (Google + email), wagmi + viem, chain config opBNB testnet
- [ ] Layar 1 `/`: pilih zona (kartu dengan premi bulan ini dan narasi), pilih minggu, tombol "Lindungi minggu ini"
- [ ] Faucet in-app: tombol "Ambil IDRP testnet" memanggil `POST /faucet`, tampil hanya jika saldo < premi
- [ ] Gasless: integrasi MegaFuel sponsor policy untuk `buyPolicy`. Batas waktu setengah hari. Jika lewat, pindah ke fallback relayer (`POST /relay/buy-policy` + EIP-712 signing)
- [ ] Layar 2 `/polis`: polis aktif, sisa hari, hujan 7 hari dengan garis ambang, kuota payout minggu ini
- [ ] Layar 3 `/riwayat`: daftar payout, nominal rupiah, link BscScan, penjelasan (placeholder sampai fase 4)
- [ ] PWA manifest, ikon, service worker minimal
- [ ] Deploy ke Vercel
- [ ] **Checkpoint hari 7 malam**: rekam alur end-to-end di HP asli sebagai bukti. Kalau gagal, hari 8 dipakai untuk ini.

**Definition of done fase 3**: video kasar 60 detik alur lengkap dari HP.

---

## Phase 4: AI
**Durasi**: 2 hari
**Goal**: Premi punya alasan, payout punya penjelasan.

### Tasks
- [ ] `internal/pricing`: dari 3 tahun data harian per zona, hitung per bulan: P(hari hujan >= ambang), ekspektasi hari hujan per minggu (dibatasi max 3), ekspektasi payout = itu x payoutPerDay, premi = ekspektasi x 1,3 dibulatkan ke Rp 500. Unit test dengan data sintetis.
- [ ] `internal/ai`: client Claude API, prompt `premium_narrative` (input angka, output JSON `{narrative}` 1 sampai 2 kalimat bahasa Indonesia), fallback string deterministik
- [ ] `POST /admin/pricing/recompute`: hitung semua zona x bulan, simpan `premiums`, panggil `setPremium` untuk bulan berjalan
- [ ] `internal/ai` prompt `payout_explainer`: input (zona, tanggal, mm, ambang, nominal, payout ke-n dari max), output 3 kalimat, fallback template
- [ ] Indexer memanggil explainer setelah tiap `PayoutSent`, simpan ke `payouts.explanation`
- [ ] Frontend menampilkan narasi premi di layar 1 dan penjelasan di layar 3

**Definition of done fase 4**: premi Bantul dan Solo berbeda dengan alasan yang terbaca, dan payout baru punya penjelasan dalam < 30 detik.

---

## Phase 5: Polish
**Durasi**: 1 hari
**Goal**: Tidak ada layar yang bikin malu di video.

### Tasks
- [ ] Empty state: belum punya polis, belum ada payout, belum ada hujan
- [ ] Loading state dan error toast berbahasa Indonesia
- [ ] Seed: 3 wallet demo dengan polis aktif di Bantul agar simulate-rain membayar beberapa orang, bukan satu
- [ ] README final dengan alamat kontrak, diagram mermaid dari ARCHITECTURE.md, dan bagian "Honest Limitations"
- [ ] Cek ulang: rumus `day_index` Go dan Solidity identik (tes dengan 3 timestamp di sekitar tengah malam WIB)

---

## Phase 6: Video & Submission
**Durasi**: 2 hari

### Struktur video (maks 4 menit)
1. 0:00 sampai 0:30: masalah. Driver ojol, hujan, order hilang, tidak ada asuransi yang mau.
2. 0:30 sampai 2:30: demo live dari HP. Login Google, ambil IDRP, beli polis Bantul, layar admin suntik hujan 31 mm, kembali ke HP: notifikasi payout, riwayat, link BscScan, penjelasan AI.
3. 2:30 sampai 3:15: arsitektur (diagram), kontrak terverifikasi, kenapa AI di pricing bukan di pembayaran.
4. 3:15 sampai 3:45: batasan jujur (oracle terpusat) dan roadmap.
5. 3:45 sampai 4:00: tim dan ajakan.

### Tasks
- [ ] Rekam 3 kali, pilih terbaik, upload YouTube (unlisted boleh)
- [ ] Isi form submission: nama tim, nama project, track (Finance & Commerce dan Consumer Apps), alamat kontrak, problem statement, solusi, project detail (markdown + mermaid), repo publik, video, anggota
- [ ] Simpan edit code di password manager
- [ ] Pastikan registrasi Luma sudah dilakukan dengan email yang sama

---

## Phase 7: Buffer
**Durasi**: 1 hari
- [ ] Perbaiki apa pun dari feedback teman/mentor
- [ ] Edit submission sebelum 23:59 WIB 30 September

---

## Parking Lot (Post-MVP)
- Settlement pull-based dengan Merkle root per (zona, hari)
- Multi-oracle dengan median 3 sumber, atau Chainlink Functions
- Zona per kecamatan
- Pricing dinamis per musim dengan retraining bulanan
- Deteksi sybil dan verifikasi status driver (mitra Gojek/Grab)
- Liquidity provider dan model cadangan
- Notifikasi push/WhatsApp saat payout
- Polis bulanan dengan diskon
- Mainnet opBNB dengan IDR stablecoin berlisensi

## Definition of Done (umum)
- `forge test` dan `go test` hijau
- Jalan di deployment publik (Vercel + Railway), bukan hanya lokal
- Tidak ada langkah manual yang dibutuhkan selain yang ada di Quick Start
- Setiap tx yang disebut di UI punya link BscScan yang benar
