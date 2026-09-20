# Requirements: Payung

Terakhir diperbarui: 17 September 2026

## Problem Statement

Driver ojol kehilangan sebagian besar pendapatan harian saat hujan deras: order turun, perjalanan lebih lambat, risiko kecelakaan naik. Dalam sebulan musim hujan, seorang driver bisa kehilangan 6 sampai 10 hari kerja efektif. Tidak ada produk asuransi yang melayani segmen ini karena proses klaim manual (foto, formulir, verifikasi) lebih mahal daripada premi yang sanggup dibayar driver.

Payung menghapus klaim sepenuhnya. Polis bersifat parametrik: pemicunya adalah angka curah hujan di zona driver, bukan laporan kerugian. Saat ambang terlewati, kontrak membayar. Driver tidak perlu membuktikan apa pun.

## Goals & Non-Goals

### Goals (In Scope)
- Driver bisa membeli polis mingguan untuk satu zona dalam kurang dari 60 detik dari pertama buka aplikasi, tanpa memegang BNB untuk gas.
- Kontrak membayar otomatis ke semua polis aktif di zona pada hari yang curah hujannya >= ambang, paling lambat 6 jam setelah data hari itu tersedia.
- Setiap payout bisa diverifikasi di BscScan dan dijelaskan ke driver dalam bahasa Indonesia sederhana.
- Premi per zona per bulan dihitung dari data hujan historis 3 tahun, bukan angka tebakan, dan alasannya bisa dibaca driver.
- Seluruh alur (beli polis, hujan, payout, penjelasan) bisa didemokan end-to-end di opBNB testnet dalam video 4 menit.

### Non-Goals (Out of Scope untuk MVP)
- Oracle terdesentralisasi atau multi-sumber. MVP memakai satu signer backend.
- Data hujan per kecamatan. MVP memakai level kota/kabupaten.
- Pricing dinamis real-time atau per individu.
- Deteksi fraud atau anti-sybil.
- Penyedia likuiditas pihak ketiga (LP token, yield). Pool dimodali oleh satu penjamin.
- Integrasi dengan aplikasi Gojek/Grab atau verifikasi status driver.
- Mainnet, KYC, atau kepatuhan OJK.
- Aplikasi native Android/iOS. MVP adalah PWA.

## Functional Requirements

### FR-01: Zona dan Parameter Polis
- Deskripsi: Sistem harus menyimpan daftar zona dengan parameter polis masing-masing.
- Input: id zona, nama, koordinat (lat, lon), premi per minggu, ambang curah hujan (mm/hari), payout per hari hujan, maksimum hari payout per minggu.
- Output: zona bisa dibaca dari kontrak dan API.
- Nilai awal MVP: 5 zona (Yogyakarta, Sleman, Bantul, Surakarta, Semarang). Ambang 20 mm/hari. Payout Rp 25.000/hari. Maksimum 3 hari/minggu. Premi awal Rp 5.000/minggu, ditimpa oleh tabel premi AI (FR-06).
- Priority: High

### FR-02: Beli Polis (Gasless)
- Deskripsi: Driver harus bisa membeli polis untuk satu zona untuk 1 sampai 4 minggu.
- Input: alamat wallet driver, id zona, jumlah minggu.
- Output: polis tersimpan on-chain dengan `start` (awal hari berikutnya, 00:00 WIB) dan `end`; premi berpindah ke pool; event `PolicyBought`.
- Aturan: satu wallet maksimal satu polis aktif per zona. Transaksi disponsori gas-nya (MegaFuel atau relayer). Driver harus mendapatkan mock IDRP dari faucet in-app sebelum beli (testnet saja).
- Priority: High

### FR-03: Submit Curah Hujan (Oracle)
- Deskripsi: Backend oracle harus mengirim curah hujan harian per zona ke kontrak.
- Input: id zona, indeks hari (hari sejak epoch, zona waktu WIB), curah hujan dalam mm (integer, dibulatkan).
- Output: nilai tersimpan on-chain, event `RainfallReported`. Jika mm >= ambang zona, hari ditandai `isRainDay`.
- Aturan: hanya alamat dengan role `ORACLE` yang boleh memanggil. Satu hari per zona hanya bisa disubmit sekali. Oracle berjalan setiap hari pukul 06:00 WIB untuk data hari sebelumnya.
- Priority: High

### FR-04: Settlement Otomatis
- Deskripsi: Sistem harus membayar semua polis aktif di zona untuk hari hujan.
- Input: id zona, indeks hari.
- Output: transfer IDRP ke setiap pemegang polis yang memenuhi syarat, event `PayoutSent` per polis.
- Syarat bayar: hari itu `isRainDay`, polis aktif pada hari itu (`start <= day < end`), jumlah payout polis di minggu kalender itu < maksimum, dan hari itu belum dibayar untuk polis tersebut.
- Aturan: siapa pun boleh memanggil `settle` (MVP: backend memanggil langsung setelah submit). Jumlah polis aktif per zona dibatasi 50 di kontrak agar loop tidak kehabisan gas. Jika saldo pool tidak cukup untuk satu payout, payout itu dilewati dengan event `PayoutSkipped`, bukan revert seluruh batch.
- Priority: High

### FR-05: Dashboard Driver
- Deskripsi: Driver harus bisa melihat status polis, hujan 7 hari terakhir di zonanya, dan riwayat payout.
- Input: alamat wallet (dari sesi login).
- Output: polis aktif (zona, sisa hari), grafik/tabel curah hujan 7 hari dengan penanda ambang, daftar payout dengan tanggal, mm hujan, nominal, dan link BscScan.
- Aturan: data dibaca dari API backend (yang mengindeks event kontrak), bukan langsung dari RPC, agar halaman terbuka < 2 detik.
- Priority: High

### FR-06: Tabel Premi Berbasis Data (AI Pricing)
- Deskripsi: Sistem harus menghasilkan premi per zona per bulan dari data hujan historis.
- Input: 3 tahun data curah hujan harian per zona dari Open-Meteo Archive.
- Output: untuk setiap zona dan bulan: ekspektasi hari hujan per minggu, premi yang direkomendasikan, dan narasi 1 sampai 2 kalimat dalam bahasa Indonesia. Premi ditulis ke kontrak via `setPremium`.
- Metode: perhitungan aktuaria deterministik di Go (ekspektasi payout per minggu dikali loading factor 1,3, dibulatkan ke Rp 500 terdekat). LLM hanya menulis narasi dan memeriksa kewajaran, tidak menentukan angka.
- Priority: High

### FR-07: Penjelasan Payout (AI Explainer)
- Deskripsi: Setelah setiap payout, sistem harus menghasilkan pesan penjelasan untuk driver.
- Input: event `PayoutSent`, data hujan hari itu, status polis, hitungan payout minggu ini.
- Output: pesan maksimal 3 kalimat, bahasa Indonesia, menyebut tanggal, zona, mm hujan, nominal, dan sisa kuota minggu ini. Tersimpan di DB dan tampil di dashboard.
- Contoh: "Kemarin hujan 31 mm di Bantul, di atas batas 20 mm. Polis kamu aktif, jadi Rp 25.000 sudah masuk ke dompetmu. Ini payout ke-2 dari maksimal 3 minggu ini."
- Priority: High

### FR-08: Simulasi Hujan (Demo dan Testing)
- Deskripsi: Admin harus bisa menyuntikkan nilai curah hujan untuk zona dan hari tertentu tanpa menunggu cuaca.
- Input: id zona, tanggal, mm, admin key.
- Output: alur FR-03 dan FR-04 berjalan dengan nilai tersebut, ditandai `source = simulated` di DB.
- Aturan: hanya aktif jika `ENV != production`.
- Priority: High (kritis untuk video demo)

### FR-09: Pendanaan Pool
- Deskripsi: Penjamin harus bisa menyetor IDRP ke pool sebagai modal awal.
- Input: jumlah IDRP.
- Output: saldo pool bertambah, event `PoolFunded`.
- Priority: Medium

### FR-10: Faucet Testnet
- Deskripsi: Driver harus bisa mendapatkan IDRP testnet dari dalam aplikasi.
- Input: alamat wallet.
- Output: 50.000 IDRP dikirim, maksimal sekali per 24 jam per alamat.
- Priority: Medium

## Non-Functional Requirements

- Performance: dashboard terbuka < 2 detik pada koneksi 4G; beli polis selesai (tx confirmed) < 15 detik di opBNB.
- Security: private key oracle dan relayer hanya di environment backend, tidak pernah di frontend. Kontrak memakai `AccessControl` untuk role `ORACLE` dan `ADMIN`. `ReentrancyGuard` pada `settle` dan `buyPolicy`. Semua transfer memakai `SafeERC20`.
- Reliability: oracle job idempotent; jika gagal, bisa dijalankan ulang untuk hari yang sama tanpa duplikasi payout.
- Observability: setiap submit dan settle dicatat di tabel `oracle_runs` dengan tx hash dan status.
- Scalability (MVP): 50 polis aktif per zona, 5 zona. Skala lebih besar ada di roadmap (pull-based).
- Mobile-first: semua layar dirancang untuk lebar 360 px, bisa di-install sebagai PWA.

## Constraints

- Satu developer, 13 hari kalender (17 sampai 30 September 2026), deadline submission 30 September 23:59 WIB.
- Wajib ada kontrak yang terdeploy dan terverifikasi di BSC atau opBNB (syarat lomba).
- Data BMKG observasi tidak tersedia lewat API publik tanpa registrasi, sehingga MVP memakai Open-Meteo.
- Tidak ada IDR stablecoin resmi di opBNB testnet, sehingga memakai mock ERC-20.

## Assumptions

- Curah hujan harian level kota cukup berkorelasi dengan penurunan order ojol untuk keperluan MVP.
- Driver mau membayar premi Rp 5.000 sampai 10.000 per minggu jika payout terlihat nyata dan cepat.
- Juri menerima oracle terpusat selama dinyatakan jujur dan ada roadmap.
- Open-Meteo tetap tersedia dan gratis selama periode lomba.
