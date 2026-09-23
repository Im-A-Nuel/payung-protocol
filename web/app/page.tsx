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
import { SimulasiPolis } from "@/components/simulasi-polis";
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
    <div className="home-payung">
      <header className="mb-5 flex items-center justify-between gap-3 px-1">
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
          <span className="rounded-full border border-white/80 bg-white/70 px-3 py-1.5 text-[12px] font-semibold text-[#48647f]">opBNB Testnet</span>
        )}
      </header>

      {!sudahMasuk ? (
        <section className="mb-5 px-1">
          <p className="text-[14px] leading-relaxed text-[#516f8b]">
            Perlindungan hujan untuk pengemudi ojol, tanpa klaim manual.
          </p>
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

      {zona.isLoading ? (
        <div id="simulasi" className="scroll-mt-5">
          <Memuat tinggi="h-96" />
        </div>
      ) : zona.isError ? (
        <div id="simulasi" className="scroll-mt-5">
          <Galat pesan={pesanGalat(zona.error)} onCoba={() => zona.refetch()} />
        </div>
      ) : daftar.length === 0 ? (
        <div id="simulasi" className="scroll-mt-5">
          <Kosong
            judul="Zona belum tersedia"
            pesan="Pilihan zona belum muncul. Coba muat ulang sebentar lagi."
            anak={<Tombol varian="kedua" onClick={() => zona.refetch()}>Muat ulang zona</Tombol>}
          />
        </div>
      ) : (
        <SimulasiPolis daftar={daftar} zona={terpilih!} onPilihZona={setZonaDipilih} />
      )}

      {sudahMasuk && terpilih ? (
        <PanelBeli zona={terpilih} terkunci={polis.data?.zoneId === terpilih.id} tampilkanNarasi={narasiSeragam} />
      ) : null}

      {!sudahMasuk && terpilih ? (
        <section className="payung-next mt-5 p-5">
          <p className="text-[13px] font-semibold text-[#356caa]">Langkah berikutnya</p>
          <h2 className="mt-1 text-[22px] font-extrabold">Lindungi {terpilih.name}.</h2>
          <p className="mt-2 text-[14px] leading-relaxed text-[#496784]">
            Premi {formatRupiah(terpilih.premiumPerWeek)} per minggu. Polis mulai besok setelah pembelian berhasil.
          </p>
          {bisaMasuk ? (
            <Tombol
              varian="kedua"
              className="mt-4 border-transparent shadow-[0_12px_25px_rgba(20,87,187,0.22)]"
              onClick={masuk}
              disabled={!siap || toWei(terpilih.premiumPerWeek) === 0n}
            >
              {toWei(terpilih.premiumPerWeek) > 0n ? "Masuk untuk beli polis" : "Premi zona belum tersedia"}
            </Tombol>
          ) : (
            <p className="mt-4 rounded-2xl bg-[#eaf3ff] p-3 text-[13px] leading-relaxed text-[#315778]">
              Simulasi bisa dicoba sekarang. Pembelian testnet belum aktif karena login belum dikonfigurasi.
            </p>
          )}
        </section>
      ) : null}

      {daftar.length > 0 ? (
        <details className="group mt-5 rounded-[24px] border border-white/90 bg-white/80 shadow-[0_10px_30px_rgba(78,120,169,0.08)]">
          <summary className="flex min-h-[60px] cursor-pointer list-none items-center justify-between gap-3 px-5 text-[15px] font-bold text-tinta [&::-webkit-details-marker]:hidden">
            Bandingkan {daftar.length} zona
            <span aria-hidden="true" className="text-[23px] leading-none text-langit transition-transform group-open:rotate-45">+</span>
          </summary>
          <div className="flex flex-col gap-2 px-3 pb-3">
            {daftar.map((pilihan) => (
              <KartuZona
                key={pilihan.id}
                zona={pilihan}
                terpilih={pilihan.id === idAktif}
                terkunci={polis.data?.zoneId === pilihan.id}
                tampilkanAturan={!aturanSeragam}
                tampilkanNarasi={!narasiSeragam}
                onPilih={(id) => {
                  setZonaDipilih(id);
                  document.getElementById("simulasi")?.scrollIntoView({ block: "start" });
                }}
              />
            ))}
          </div>
        </details>
      ) : null}
    </div>
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
