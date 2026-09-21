# PROGRESS

Catatan keadaan repo untuk siapa pun yang baru masuk, termasuk Claude Code lokal dan Codex.
Baca ini dulu sebelum menyentuh kode. Aturan teknis yang mengikat ada di `CLAUDE.md`, rencana
per fase ada di `docs/ROADMAP.md`.

Terakhir diperbarui: 21 September 2026
Deadline submission: 30 September 2026, 23:59 WIB

## Status per fase

| Fase | Isi | Keadaan |
|---|---|---|
| 1. Kontrak | `PayungPool.sol`, `IDRP.sol`, 10 test wajib, `Deploy.s.sol` | Kode selesai, `forge test` hijau (17 test). Belum di-deploy ke opBNB testnet |
| 2. Oracle & API | Go API, oracle job, indexer, migrasi Postgres | Kode selesai, `go test ./...` hijau (21 test). Belum di-deploy ke Railway |
| 3. Frontend | 3 layar Next.js, PWA, semua teks bahasa Indonesia | Kode selesai, `pnpm build` hijau. Belum di-deploy ke Vercel. Gasless belum jalan |
| 4. AI | `internal/pricing`, `internal/ai`, narasi premi, penjelasan payout | Belum dikerjakan sama sekali |
| 5. Polish | Empty state, loading, error, seed, README final | Sebagian: empty state, loading, error, dan animasi UI sudah ada. Seed 3 wallet demo dan README final belum |
| 6. Video & submit | Rekaman 4 menit, isi form | Belum. Terhalang deploy dan gasless |

Yang dianggap "selesai" di tabel itu artinya: teruji lokal dan sudah dijalankan end to end di
chain lokal (anvil) plus browser sungguhan, bukan cuma lolos kompilasi.

## Cabang dan riwayat

`main` sudah berisi Phase 1 sampai 3 lewat PR #2 yang sudah di-merge.

Satu commit belum masuk `main`: `5ae837c` (polesan UI/UX, animasi, layar "Berhasil", sticky
CTA, penanda "Kemarin" di grafik). Ada di branch `claude/mulai-dari-yang-bisa-2n70mr` dan
sedang menunggu merge lewat PR terpisah. Kalau PR itu sudah di-merge, abaikan paragraf ini.

## Peta kode

```
contracts/          Foundry. PayungPool.sol (inti), IDRP.sol (mock token IDR)
  test/             17 test, termasuk 10 kasus wajib dari docs/SCHEMA.md
  script/Deploy.s.sol   deploy + 5 zona + grant ORACLE_ROLE + isi pool

backend/            Go 1.22, chi, pgx + sqlc, binding go-ethereum
  cmd/api           HTTP API (:8080)
  cmd/oracle        job harian, ada flag --once
  cmd/migrate       jalankan migrasi, aman diulang
  internal/chain    day_index, signer, kirim tx, binding hasil abigen
  internal/oracle    alur submit rainfall lalu settle, idempoten
  internal/indexer  event on-chain masuk Postgres
  internal/rain     klien Open-Meteo
  internal/httpapi  handler HTTP
  internal/store    hasil sqlc + file query SQL
  migrations/       0001 skema, 0002 seed 5 zona, 0003 seed premi awal

web/                Next.js 15 App Router, Tailwind 4, Privy, wagmi
  app/page.tsx      Layar 1: pilih zona, beli polis
  app/polis/        Layar 2: polis aktif, kuota, grafik hujan 7 hari
  app/riwayat/      Layar 3: riwayat payout
  lib/api.ts        semua panggilan ke backend
  lib/format.ts     satu-satunya tempat wei jadi rupiah
  lib/contracts.ts  ABI, alamat, konfigurasi chain

docs/               REQUIREMENTS, ARCHITECTURE, SCHEMA, ROADMAP
```

## Endpoint yang sudah ada

Publik: `GET /v1/zones`, `GET /v1/zones/:id/rain?days=7`, `GET /v1/drivers/:address/policy`,
`GET /v1/drivers/:address/payouts`, `GET /v1/drivers/:address/wallet`, `POST /v1/faucet`.

Admin (mati kalau `ENV=production`, butuh header `X-Admin-Key`):
`POST /v1/admin/simulate-rain`, `POST /v1/admin/oracle/run`.

`GET /drivers/:address/wallet` tidak ada di `docs/SCHEMA.md`. Itu tambahan yang dibutuhkan
Phase 3: frontend harus tahu saldo IDRP (untuk memutuskan menampilkan tombol faucet) dan
allowance (untuk tahu apakah perlu approve), sementara aturan proyek melarang frontend baca RPC
sendiri.

## Menjalankan semuanya di lokal

Butuh: Foundry, Go 1.22+, Node 22 + pnpm, PostgreSQL 16.

```bash
# 1. Postgres
createdb payung   # atau pakai docker
export DATABASE_URL="postgres://payung:payung@localhost:5432/payung?sslmode=disable"

# 2. Chain lokal
anvil --port 8545 --chain-id 31337
# anvil mencetak 10 akun beserta private key-nya saat start; pakai itu, jangan taruh di repo

# 3. Deploy kontrak (akun #0 sebagai deployer sekaligus oracle)
cd contracts
PRIVATE_KEY=<kunci akun #0> ORACLE_ADDRESS=<alamat akun #0> \
  forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --broadcast

# Alamat hasil deploy di anvil selalu sama:
#   IDRP        0x5FbDB2315678afecb367f032d93F642f64180aa3
#   PayungPool  0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512

# 4. Beri FAUCET_ROLE ke signer faucet (pakai akun #1)
cast send 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
  "grantRole(bytes32,address)" $(cast keccak "FAUCET_ROLE") <alamat akun #1> \
  --rpc-url http://127.0.0.1:8545 --private-key <kunci akun #0>

# 5. Backend
cd ../backend
go run ./cmd/migrate
# lalu isi env (lihat backend/.env.example) dan jalankan:
go run ./cmd/api

# 6. Frontend
cd ../web
cp .env.example .env.local   # arahkan ke API lokal dan alamat kontrak di atas
pnpm install && pnpm dev
```

### Bikin data demo (polis aktif + payout)

Polis mulai berlaku besok, jadi harus memajukan waktu chain sebelum ada payout.

```bash
# beli polis di zona 3 (Bantul)
cast send <IDRP> "approve(address,uint256)" <POOL> 5000000000000000000000 --private-key <kunci driver> --rpc-url http://127.0.0.1:8545
cast send <POOL> "buyPolicy(uint16,uint8)" 3 1 --private-key <kunci driver> --rpc-url http://127.0.0.1:8545

# majukan 2 hari supaya hari mulai polis sudah lewat
cast rpc evm_increaseTime 172800 --rpc-url http://127.0.0.1:8545
cast rpc evm_mine --rpc-url http://127.0.0.1:8545

# suntik hujan untuk tanggal yang ada di dalam masa polis
curl -X POST http://127.0.0.1:8080/v1/admin/simulate-rain \
  -H "Content-Type: application/json" -H "X-Admin-Key: <ADMIN_KEY>" \
  -d '{"zoneId":3,"date":"2026-09-22","mm":31}'
```

## Yang belum jalan, dan kenapa

1. **Belum ada yang di-deploy.** Kontrak belum ada di opBNB testnet, backend belum di Railway,
   web belum di Vercel. Semua butuh kredensial yang hanya dipegang pemilik repo.

2. **Gasless belum tersambung.** `buyPolicy` dikirim langsung dari embedded wallet Privy yang
   tidak punya tBNB, jadi pembelian sungguhan akan gagal. Dua jalan menurut `docs/SCHEMA.md`:
   MegaFuel (perlu sponsor policy di dashboard mereka) atau relayer fallback (perlu fungsi baru
   bergaya `buyPolicyFor` di kontrak, artinya membuka lagi Phase 1). Ini penghalang utama
   checkpoint hari ke-7.

3. **Phase 4 belum ada.** `GET /zones` menyajikan premi seed Rp 5.000 dengan narasi placeholder,
   dan `payouts.explanation` selalu null sehingga UI menampilkan "Penjelasan menyusul".

4. **Bundle web 950 kB.** Mayoritas dari Privy yang ikut memaketkan konektor dompet yang aplikasi
   ini tidak pernah tawarkan. Meleset dari target "dashboard terbuka < 2 detik di 4G" di
   `docs/REQUIREMENTS.md`. Memperbaikinya berarti merombak lapisan auth, jadi ditahan dulu.

5. **Tanpa `NEXT_PUBLIC_PRIVY_APP_ID`**, web jalan dalam mode pratinjau: baca saja, tombol beli
   mati, ada keterangan di layar. Alamat yang dibaca diambil dari `NEXT_PUBLIC_PREVIEW_ADDRESS`.

## Jebakan yang perlu diketahui sebelum mengubah apa pun

**`day_index` adalah fondasi semua perhitungan hari.** Rumusnya `floor((unix + 7*3600) / 86400)`
dan harus identik di Solidity (`today()`) dan Go (`internal/chain/day.go`). Jangan hitung hari
dengan cara lain di mana pun, termasuk di TypeScript.

**`DateStringWIB` dan `DayIndexFromDateWIB` wajib jadi kebalikan satu sama lain.** Pernah tidak:
versi lama mengurangi offset WIB dari sisi yang salah sehingga meleset satu hari. Akibatnya
oracle akan mengambil tanggal yang salah dari respons Open-Meteo dan menyelesaikan hari yang
sebenarnya tidak ditanggung siapa pun. Sudah diperbaiki dan dikunci oleh test round-trip di
`backend/internal/chain/day_test.go`. Kalau menyentuh dua fungsi itu, jalankan test tersebut.

**Test integrasi berbagi database yang sama dengan demo lokal.** Data fixture memakai rentang
jauh di masa depan (`day_index >= 900000`, `policy_id >= 910000`) dan setiap test membersihkan
barisnya sendiri. Kalau menambah test, ikuti pola itu. Baris sisa bertanggal tahun 4461 akan
muncul paling atas di query "7 hari terakhir" dan menutupi data asli.

**Waktu chain dan jam server bisa berbeda saat testing.** Kalau memajukan waktu anvil dengan
`evm_increaseTime`, API tetap memakai jam sistem, jadi `daysLeft` akan terlihat aneh. Itu
artefak setup uji, bukan bug. Di produksi keduanya sejalan.

**Indexer perlu berjalan supaya pembelian muncul di dashboard.** Sinkronisasi terjadi saat API
start, setiap 15 detik, dan sesudah aksi admin. Kalau `policies` atau `payouts` kosong padahal
transaksi on-chain sudah sukses, kemungkinan indexer belum sempat jalan.

**Migrasi 0003 menyeed premi Rp 5.000 per minggu.** Tanpa itu `GET /zones` melaporkan Rp 0
sementara kontrak tetap menagih Rp 5.000, jadi driver diberi harga yang tidak dihormati kontrak.
Phase 4 akan menimpa baris-baris ini dengan angka hasil perhitungan.

**Uang selalu wei IDRP (18 desimal)** di kontrak, di Postgres, dan di semua respons API.
Konversi ke rupiah hanya boleh terjadi di `web/lib/format.ts`.

**Alamat disimpan huruf kecil di Postgres.** Indexer melakukan `strings.ToLower` sebelum menulis.

## Perintah yang sering dipakai

```bash
cd contracts && forge test                 # 17 test
cd backend   && go test ./...              # 21 test, 3 file butuh TEST_DATABASE_URL
cd backend   && sqlc generate              # sesudah mengubah internal/store/queries/*.sql
cd backend   && ./scripts/gen-abi.sh       # sesudah mengubah contracts/src/*.sol
cd web       && pnpm build                 # sekaligus typecheck
```

Test Go yang butuh Postgres akan otomatis di-skip kalau `TEST_DATABASE_URL` tidak diisi, jadi
`go test ./...` tetap hijau tanpa database tapi cakupannya berkurang. CI mengisinya lewat service
container.

## Langkah berikutnya, berurutan

1. Putuskan jalur gasless (MegaFuel atau relayer). Ini menentukan apakah Phase 1 perlu dibuka lagi.
2. Deploy kontrak ke opBNB testnet, verifikasi di BscScan, catat alamatnya di `README.md` dan
   semua `.env.example`.
3. Deploy backend ke Railway beserta Postgres, pasang cron 06:00 WIB untuk `cmd/oracle`.
4. Deploy web ke Vercel dengan Root Directory `web`, isi `NEXT_PUBLIC_PRIVY_APP_ID` dan alamat
   kontrak hasil langkah 2.
5. Kerjakan Phase 4 (pricing deterministik lalu narasi AI dengan fallback).
6. Sisa Phase 5: seed 3 wallet demo, README final dengan diagram dan bagian batasan jujur.
7. Rekam video, isi form submission.

## Merawat file ini

Perbarui saat keadaan berubah, terutama tabel status, daftar penghalang, dan bagian jebakan.
Bagian jebakan yang paling berharga: setiap kali ada bug yang makan waktu lama untuk ketemu,
tulis di situ supaya tidak terulang.
