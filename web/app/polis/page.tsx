"use client";

import { Suspense } from "react";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { useQuery } from "@tanstack/react-query";

import { getPolicy, getZoneRain } from "@/lib/api";
import { useAkun } from "@/lib/akun";
import { formatTanggalPanjang } from "@/lib/format";
import { pesanGalat } from "@/lib/galat";
import { GrafikHujan } from "@/components/grafik-hujan";
import { Galat, JudulHalaman, Kartu, Kosong, Memuat } from "@/components/ui";

export default function HalamanPolis() {
  return (
    <Suspense fallback={<Memuat tinggi="h-40" />}>
      <IsiPolis />
    </Suspense>
  );
}

function IsiPolis() {
  const { sudahMasuk, alamat } = useAkun();
  const baruBeli = useSearchParams().get("baru") === "1";

  const polis = useQuery({
    queryKey: ["polis", alamat],
    queryFn: () => getPolicy(alamat as string),
    enabled: Boolean(alamat),
    // Right after a purchase the indexer may not have caught the event yet.
    refetchInterval: (q) => (baruBeli && !q.state.data ? 3000 : false),
  });

  const hujan = useQuery({
    queryKey: ["hujan", polis.data?.zoneId],
    queryFn: () => getZoneRain(polis.data!.zoneId, 7),
    enabled: Boolean(polis.data),
  });

  if (!sudahMasuk) {
    return (
      <>
        <JudulHalaman judul="Polis kamu" />
        <Kosong
          judul="Masuk dulu"
          pesan="Polismu muncul di sini setelah kamu masuk pakai akun Google."
          anak={<TautanBeranda label="Ke halaman depan" />}
        />
      </>
    );
  }

  if (polis.isLoading) {
    return (
      <>
        <JudulHalaman judul="Polis kamu" />
        <Memuat tinggi="h-40" />
      </>
    );
  }

  if (polis.isError) {
    return (
      <>
        <JudulHalaman judul="Polis kamu" />
        <Galat pesan={pesanGalat(polis.error)} onCoba={() => polis.refetch()} />
      </>
    );
  }

  if (!polis.data) {
    return (
      <>
        <JudulHalaman judul="Polis kamu" />
        {baruBeli ? (
          <Kosong
            judul="Polismu lagi disiapkan"
            pesan="Pembayaranmu sudah masuk. Catatannya muncul di sini sebentar lagi."
          />
        ) : (
          <Kosong
            judul="Belum ada polis aktif"
            pesan="Pilih zona tempat kamu narik, lalu ambil perlindungan mingguan."
            anak={<TautanBeranda label="Pilih zona" />}
          />
        )}
      </>
    );
  }

  const p = polis.data;
  const kuotaTerpakai = p.payoutsThisWeek;
  const kuotaTotal = p.maxDaysPerWeek;

  return (
    <>
      <JudulHalaman judul="Polis kamu" />

      <Kartu>
        <p className="text-[13px] font-semibold tracking-wide text-abu uppercase">{p.zoneName}</p>
        <p className="mt-1 text-[32px] leading-tight font-extrabold">
          Sisa {p.daysLeft} hari
        </p>
        <p className="mt-1 text-[14px] text-abu">
          Berlaku {formatTanggalPanjang(p.startDate)} sampai {formatTanggalPanjang(p.endDate)}
        </p>
      </Kartu>

      <Kartu className="mt-3">
        <div className="flex items-baseline justify-between gap-3">
          <p className="text-[15px] font-bold">Jatah bayar minggu ini</p>
          <p className="angka text-[15px] font-extrabold">
            {kuotaTerpakai} dari {kuotaTotal}
          </p>
        </div>
        <div className="mt-3 flex gap-1.5" aria-hidden="true">
          {Array.from({ length: kuotaTotal }, (_, i) => (
            <span
              key={i}
              className={`h-2.5 flex-1 rounded-full ${i < kuotaTerpakai ? "bg-uang" : "bg-garis"}`}
            />
          ))}
        </div>
        <p className="mt-3 text-[13px] leading-relaxed text-abu">
          Paling banyak {kuotaTotal} hari hujan dibayar tiap minggu. Hitungannya ikut minggu berjalan, bukan
          tanggal belimu.
        </p>
      </Kartu>

      <Kartu className="mt-3">
        <p className="text-[15px] font-bold">Hujan 7 hari terakhir</p>
        <p className="mb-3 text-[13px] text-abu">di {p.zoneName}</p>

        {hujan.isLoading ? (
          <Memuat tinggi="h-36" />
        ) : hujan.isError ? (
          <Galat pesan={pesanGalat(hujan.error)} onCoba={() => hujan.refetch()} />
        ) : (
          <GrafikHujan
            hari={hujan.data?.days ?? []}
            ambangMm={hujan.data?.thresholdMm ?? 0}
            namaZona={p.zoneName}
          />
        )}
      </Kartu>
    </>
  );
}

function TautanBeranda({ label }: { label: string }) {
  return (
    <Link
      href="/"
      className="inline-flex min-h-[44px] items-center rounded-xl px-4 text-[15px] font-bold text-langit"
    >
      {label}
    </Link>
  );
}
