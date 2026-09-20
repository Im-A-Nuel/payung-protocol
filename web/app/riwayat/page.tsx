"use client";

import Link from "next/link";
import type { CSSProperties } from "react";
import { useQuery } from "@tanstack/react-query";

import { getPayouts, type Payout } from "@/lib/api";
import { useAkun } from "@/lib/akun";
import { urlTransaksi } from "@/lib/contracts";
import { formatRupiah, formatRupiahAngka, formatTanggalPanjang, rupiahBulat, toWei } from "@/lib/format";
import { useHitungNaik } from "@/lib/hitung-naik";
import { pesanGalat } from "@/lib/galat";
import { Galat, JudulHalaman, Kartu, Kosong, Memuat } from "@/components/ui";

export default function HalamanRiwayat() {
  const { sudahMasuk, alamat } = useAkun();

  const payout = useQuery({
    queryKey: ["payout", alamat],
    queryFn: () => getPayouts(alamat as string),
    enabled: Boolean(alamat),
  });

  if (!sudahMasuk) {
    return (
      <>
        <JudulHalaman judul="Riwayat bayaran" />
        <Kosong
          judul="Masuk dulu"
          pesan="Setiap bayaran yang pernah masuk ke dompetmu tercatat di sini."
          anak={
            <Link
              href="/"
              className="inline-flex min-h-[44px] items-center rounded-xl px-4 text-[15px] font-bold text-langit"
            >
              Ke halaman depan
            </Link>
          }
        />
      </>
    );
  }

  if (payout.isLoading) {
    return (
      <>
        <JudulHalaman judul="Riwayat bayaran" />
        <div className="flex flex-col gap-3">
          <Memuat tinggi="h-28" />
          <Memuat tinggi="h-28" />
        </div>
      </>
    );
  }

  if (payout.isError) {
    return (
      <>
        <JudulHalaman judul="Riwayat bayaran" />
        <Galat pesan={pesanGalat(payout.error)} onCoba={() => payout.refetch()} />
      </>
    );
  }

  const daftar = payout.data ?? [];

  if (daftar.length === 0) {
    return (
      <>
        <JudulHalaman judul="Riwayat bayaran" />
        <Kosong
          judul="Belum ada yang masuk"
          pesan="Begitu hujan di zonamu lewat ambang, bayarannya langsung dikirim dan tercatat di sini."
        />
      </>
    );
  }

  const total = daftar.reduce((jumlah, p) => jumlah + toWei(p.amount), 0n);

  return (
    <>
      <JudulHalaman judul="Riwayat bayaran" />

      <TotalMasuk totalWei={total} jumlahHari={daftar.length} />

      <ul className="flex flex-col gap-3">
        {daftar.map((p, i) => (
          <li
            key={`${p.txHash}-${p.date}`}
            className="masuk"
            style={{ "--tunda": `${120 + i * 70}ms` } as CSSProperties}
          >
            <BarisPayout payout={p} terbaru={i === 0} />
          </li>
        ))}
      </ul>
    </>
  );
}

function TotalMasuk({ totalWei, jumlahHari }: { totalWei: bigint; jumlahHari: number }) {
  const berjalan = useHitungNaik(rupiahBulat(totalWei));

  return (
    <Kartu className="masuk mb-4 border-uang/30 bg-uang-muda">
      <p className="text-[13px] font-semibold text-uang">Total yang sudah masuk</p>
      <p className="angka mt-1 text-[30px] leading-tight font-extrabold text-uang">
        {formatRupiahAngka(berjalan)}
      </p>
      <p className="mt-1 text-[13px] text-abu">dari {jumlahHari} hari hujan yang kena ambang</p>
    </Kartu>
  );
}

function BarisPayout({ payout, terbaru }: { payout: Payout; terbaru: boolean }) {
  return (
    <Kartu className={terbaru ? "border-uang/40" : ""}>
      {terbaru ? (
        <span className="mb-2 inline-block rounded-full bg-uang px-2 py-0.5 text-[11px] font-bold text-white">
          Terbaru
        </span>
      ) : null}

      <div className="flex items-start justify-between gap-3">
        <div className="min-w-0">
          <p className="text-[15px] font-bold">{formatTanggalPanjang(payout.date)}</p>
          <p className="mt-0.5 text-[13px] text-abu">
            {payout.zoneName}, hujan {payout.mm} mm
          </p>
        </div>
        <p className="angka shrink-0 text-[18px] font-extrabold text-uang">+{formatRupiah(payout.amount)}</p>
      </div>

      <p className="mt-3 border-t border-garis pt-3 text-[14px] leading-relaxed text-abu">
        {payout.explanation ?? "Penjelasan menyusul."}
      </p>

      <a
        href={urlTransaksi(payout.txHash)}
        target="_blank"
        rel="noopener noreferrer"
        className="mt-3 inline-flex min-h-[44px] items-center gap-1.5 text-[14px] font-bold text-langit"
      >
        Lihat buktinya
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M14 4h6v6M20 4l-8.5 8.5M18 14v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h5"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      </a>
    </Kartu>
  );
}
