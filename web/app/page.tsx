"use client";

import Link from "next/link";
import { useState } from "react";
import { useQuery } from "@tanstack/react-query";

import { getPolicy, getZones } from "@/lib/api";
import { useAkun } from "@/lib/akun";
import { formatRupiah } from "@/lib/format";
import { pesanGalat } from "@/lib/galat";
import { KartuZona } from "@/components/kartu-zona";
import { PanelBeli } from "@/components/panel-beli";
import { Galat, Kartu, Memuat, Tombol } from "@/components/ui";

export default function Beranda() {
  const { siap, sudahMasuk, alamat, masuk, keluar } = useAkun();
  const [zonaDipilih, setZonaDipilih] = useState<number | null>(null);

  const zona = useQuery({ queryKey: ["zona"], queryFn: getZones });
  const polis = useQuery({
    queryKey: ["polis", alamat],
    queryFn: () => getPolicy(alamat as string),
    enabled: Boolean(alamat),
  });

  const daftar = zona.data ?? [];
  const idAktif = zonaDipilih ?? daftar[0]?.id ?? null;
  const terpilih = daftar.find((z) => z.id === idAktif) ?? null;
  const contoh = daftar[0];

  // Today every zone pays the same way and carries the same placeholder
  // narrative, so repeating both on five cards is noise. Each card states
  // its own again as soon as pricing makes them differ.
  const aturanSeragam =
    daftar.length > 0 &&
    daftar.every(
      (z) =>
        z.thresholdMm === contoh.thresholdMm &&
        z.payoutPerDay === contoh.payoutPerDay &&
        z.maxDaysPerWeek === contoh.maxDaysPerWeek,
    );
  const narasiSeragam =
    daftar.length > 0 && daftar.every((z) => z.premiumNarrative === contoh.premiumNarrative);

  return (
    <>
      <header className="mb-6 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Logo />
          <span className="text-[22px] font-extrabold tracking-tight">Payung</span>
        </div>
        {sudahMasuk ? (
          <button
            onClick={keluar}
            className="min-h-[44px] px-2 text-[14px] font-bold text-abu active:text-tinta"
          >
            Keluar
          </button>
        ) : null}
      </header>

      {!sudahMasuk ? (
        <Kartu className="mb-6 bg-langit text-white">
          <h1 className="text-[24px] leading-tight font-extrabold">Hujan deras, order sepi.</h1>
          <p className="mt-2 text-[15px] leading-relaxed text-white/90">
            {contoh
              ? `Begitu hujan sehari lewat ${contoh.thresholdMm} mm di zonamu, ${formatRupiah(contoh.payoutPerDay)} masuk ke dompetmu. Tidak perlu lapor, tidak perlu kirim foto.`
              : "Begitu hujan di zonamu melewati ambang, uangnya masuk sendiri ke dompetmu. Tidak perlu lapor, tidak perlu kirim foto."}
          </p>
          <Tombol
            varian="kedua"
            className="mt-4 border-transparent"
            onClick={masuk}
            disabled={!siap}
          >
            Masuk pakai Google
          </Tombol>
          <p className="mt-2 text-center text-[13px] text-white/80">
            Cukup akun Google. Tidak ada formulir.
          </p>
        </Kartu>
      ) : null}

      {polis.data ? (
        <Link href="/polis" className="mb-4 block">
          <Kartu className="flex items-center justify-between gap-3 border-langit bg-langit-muda">
            <p className="text-[14px] leading-snug font-semibold">
              Polismu di {polis.data.zoneName} masih jalan {polis.data.daysLeft} hari lagi.
            </p>
            <span aria-hidden="true" className="text-[18px] font-bold text-langit">
              &rsaquo;
            </span>
          </Kartu>
        </Link>
      ) : null}

      <div className="mb-3">
        <h2 className="text-[17px] font-extrabold">Pilih zona tempat kamu narik</h2>
        {aturanSeragam ? (
          <p className="mt-1 text-[13px] leading-relaxed text-abu">
            Semua zona bayar {formatRupiah(contoh.payoutPerDay)} per hari hujan di atas {contoh.thresholdMm}{" "}
            mm, paling banyak {contoh.maxDaysPerWeek} hari tiap minggu.
          </p>
        ) : null}
        {narasiSeragam ? (
          <p className="mt-1 text-[13px] leading-relaxed text-abu">{contoh.premiumNarrative}</p>
        ) : null}
      </div>

      {zona.isLoading ? (
        <div className="flex flex-col gap-3">
          <Memuat tinggi="h-28" />
          <Memuat tinggi="h-28" />
          <Memuat tinggi="h-28" />
        </div>
      ) : zona.isError ? (
        <Galat pesan={pesanGalat(zona.error)} onCoba={() => zona.refetch()} />
      ) : (
        <div className="flex flex-col gap-3">
          {daftar.map((z) => (
            <KartuZona
              key={z.id}
              zona={z}
              terpilih={z.id === idAktif}
              terkunci={polis.data?.zoneId === z.id}
              tampilkanAturan={!aturanSeragam}
              tampilkanNarasi={!narasiSeragam}
              onPilih={setZonaDipilih}
            />
          ))}
        </div>
      )}

      {sudahMasuk && terpilih ? (
        <PanelBeli zona={terpilih} terkunci={polis.data?.zoneId === terpilih.id} />
      ) : null}
    </>
  );
}

function Logo() {
  return (
    <svg width="26" height="26" viewBox="0 0 24 24" fill="none" aria-hidden="true" className="text-langit">
      <path
        d="M12 3v1.4M3.2 12.8a8.8 8.8 0 0 1 17.6 0Z"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path d="M12 12.8v6a2.2 2.2 0 0 1-4.4 0" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
    </svg>
  );
}
