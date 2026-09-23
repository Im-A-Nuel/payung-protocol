"use client";

import Link from "next/link";
import { useState } from "react";
import { useQuery } from "@tanstack/react-query";

import { getPolicy, getZones } from "@/lib/api";
import { useAkun } from "@/lib/akun";
import { formatRupiah, toWei } from "@/lib/format";
import { pesanGalat } from "@/lib/galat";
import { KartuZona } from "@/components/kartu-zona";
import { PanelBeli } from "@/components/panel-beli";
import { Galat, Kartu, Kosong, Memuat, Tombol } from "@/components/ui";

export default function Beranda() {
  const { siap, sudahMasuk, alamat, masuk, keluar, bisaMasuk } = useAkun();
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
      <header className="mb-5 flex items-center justify-between gap-3">
        <div className="flex items-center gap-2.5">
          <Logo />
          <span className="text-[22px] font-extrabold tracking-tight">Payung</span>
        </div>
        {sudahMasuk ? (
          <button
            onClick={keluar}
            className="min-h-[44px] rounded-xl px-3 text-[14px] font-bold text-abu active:text-tinta"
          >
            Keluar
          </button>
        ) : (
          <span className="text-[12px] font-semibold text-abu">opBNB Testnet</span>
        )}
      </header>

      {!sudahMasuk ? (
        <section className="mb-7 rounded-[28px] bg-tinta p-5 text-white sm:p-6">
          <p className="text-[13px] font-semibold text-[#aec8ea]">Perlindungan untuk pengemudi ojol</p>
          <h1 className="mt-3 max-w-[360px] text-[clamp(1.65rem,7vw,2.2rem)] leading-[1.12] font-extrabold tracking-tight">
            Hujan deras, order sepi.
          </h1>
          <p className="mt-3 max-w-[390px] text-[15px] leading-relaxed text-[#d7e3f1]">
            Saat polismu aktif dan hujan harian di zona pilihanmu mencapai batas, bayaran masuk ke dompet tanpa perlu mengajukan klaim.
          </p>
          {aturanSeragam ? (
            <div className="mt-5 flex flex-wrap items-end justify-between gap-x-5 gap-y-3 border-t border-white/20 pt-4">
              <div>
                <p className="text-[12px] font-medium text-[#aec8ea]">Bayaran per hari hujan</p>
                <p className="angka mt-0.5 text-[27px] leading-tight font-extrabold">{formatRupiah(contoh.payoutPerDay)}</p>
              </div>
              <p className="pb-0.5 text-[13px] font-semibold text-[#d7e3f1]">
                Mulai {contoh.thresholdMm} mm per hari
              </p>
            </div>
          ) : null}
        </section>
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

      <div id="zona" className="mb-3 scroll-mt-5">
        <div className="flex items-end justify-between gap-3">
          <h2 className="text-[19px] leading-tight font-extrabold">Pilih zona narikmu</h2>
          {daftar.length > 0 ? <span className="shrink-0 text-[13px] font-semibold text-abu">{daftar.length} zona</span> : null}
        </div>
        {aturanSeragam ? (
          <p className="mt-1 text-[13px] leading-relaxed text-abu">
            Maksimal {contoh.maxDaysPerWeek} hari hujan dibayar tiap minggu. Pilih tempat kamu biasa narik untuk melihat preminya.
          </p>
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
      ) : daftar.length === 0 ? (
        <Kosong judul="Zona belum tersedia" pesan="Pilihan zona belum muncul. Coba muat ulang sebentar lagi." anak={<Tombol varian="kedua" onClick={() => zona.refetch()}>Muat ulang zona</Tombol>} />
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
        <PanelBeli zona={terpilih} terkunci={polis.data?.zoneId === terpilih.id} tampilkanNarasi={narasiSeragam} />
      ) : null}

      {!sudahMasuk && terpilih ? (
        <Kartu className="mt-5 border-[#bdd5f5] bg-langit-muda p-5">
          <p className="text-[13px] font-semibold text-langit">Zona pilihanmu</p>
          <div className="mt-1 flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
            <p className="text-[21px] font-extrabold">{terpilih.name}</p>
            <p className="angka text-[17px] font-extrabold">
              {toWei(terpilih.premiumPerWeek) > 0n
                ? formatRupiah(terpilih.premiumPerWeek)
                : "Premi belum tersedia"}{" "}
              {toWei(terpilih.premiumPerWeek) > 0n ? (
                <span className="text-[13px] font-medium text-abu">/ minggu</span>
              ) : null}
            </p>
          </div>
          {narasiSeragam && terpilih.premiumNarrative ? (
            <p className="mt-2 text-[13px] leading-relaxed text-abu">{terpilih.premiumNarrative}</p>
          ) : null}
          <p className="mt-2 text-[14px] leading-relaxed text-abu">Masuk saat siap membeli. Google atau email bisa dipakai, lalu dompet testnet dibuat untuk menerima polis dan bayaran.</p>
          {bisaMasuk ? (
            <Tombol className="mt-4" onClick={masuk} disabled={!siap || toWei(terpilih.premiumPerWeek) === 0n}>
              {toWei(terpilih.premiumPerWeek) > 0n ? "Masuk untuk beli polis" : "Premi zona belum tersedia"}
            </Tombol>
          ) : (
            <p className="mt-4 rounded-xl bg-kartu p-3 text-[13px] font-semibold text-abu">
              Mode pratinjau: login dan pembelian belum aktif di lingkungan ini.
            </p>
          )}
        </Kartu>
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
