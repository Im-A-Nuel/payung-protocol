"use client";

import { useState } from "react";

import type { Zona } from "@/lib/api";
import { formatRupiah, multiplyWei, toWei } from "@/lib/format";

const PILIHAN_MINGGU = [1, 2, 3, 4];

export function SimulasiPolis({
  daftar,
  zona,
  onPilihZona,
}: {
  daftar: Zona[];
  zona: Zona;
  onPilihZona: (id: number) => void;
}) {
  const [hujanMm, setHujanMm] = useState(zona.thresholdMm);
  const [minggu, setMinggu] = useState(1);

  const batas = zona.thresholdMm;
  const maksimum = Math.max(60, batas * 2);
  const posisiBatas = (batas / maksimum) * 100;
  const posisiHujan = (hujanMm / maksimum) * 100;
  const memenuhiSyarat = hujanMm >= batas;
  const premiTersedia = toWei(zona.premiumPerWeek) > 0n;
  const contohHujan = [0, batas, Math.min(maksimum, batas + 15)];

  return (
    <section id="simulasi" className="payung-simulator scroll-mt-5 overflow-hidden">
      <div className="langit-payung px-5 pb-6 pt-6 text-white">
        <p className="text-[13px] font-semibold text-white">Simulasi perlindungan</p>
        <h2 className="mt-2 max-w-[300px] text-[clamp(1.8rem,8vw,2.3rem)] leading-[1.06] font-extrabold tracking-tight">
          Hari basah, tetap ada pegangan.
        </h2>
        <p className="mt-3 max-w-[310px] text-[14px] leading-relaxed text-white">
          Geser angka hujan untuk melihat aturan bayaran di zona narikmu.
        </p>
        <div className="mt-6 rounded-[22px] border border-white/30 bg-[#09468d]/35 p-4">
          <p className="text-[12px] font-semibold text-white">Bayaran saat ambang terlewati</p>
          <div className="mt-1 flex items-end justify-between gap-3">
            <p className="angka whitespace-nowrap text-[25px] leading-none font-extrabold min-[360px]:text-[27px]">{formatRupiah(zona.payoutPerDay)}</p>
            <p className="shrink-0 whitespace-nowrap pb-0.5 text-right text-[12px] font-semibold text-white">mulai {batas} mm</p>
          </div>
        </div>
      </div>

      <div className="bg-white p-5">
        <label htmlFor="zona-simulasi" className="text-[14px] font-bold text-tinta">
          Zona tempat narik
        </label>
        <select
          id="zona-simulasi"
          value={zona.id}
          onChange={(event) => onPilihZona(Number(event.target.value))}
          className="mt-2 min-h-[54px] w-full rounded-2xl border border-[#b8cbe4] bg-[#fbfdff] px-4 text-[16px] font-semibold text-tinta shadow-[inset_0_1px_0_rgba(255,255,255,0.9)]"
        >
          {daftar.map((pilihan) => (
            <option key={pilihan.id} value={pilihan.id}>
              {pilihan.name}
            </option>
          ))}
        </select>

        <div className="mt-6 flex items-end justify-between gap-3">
          <label htmlFor="hujan-simulasi" className="text-[14px] font-bold text-tinta">
            Hujan dalam sehari
          </label>
          <output htmlFor="hujan-simulasi" className="angka text-[31px] leading-none font-extrabold text-[#165bc4]">
            {hujanMm} <span className="text-[15px] font-bold">mm</span>
          </output>
        </div>

        <div className="relative mt-5 h-3 rounded-full bg-[#abc5df]" aria-hidden="true">
          <div
            className="h-full rounded-full bg-[#165bc4] transition-[width] duration-150"
            style={{ width: `${posisiHujan}%` }}
          />
          <span
            className="absolute top-[-5px] h-[22px] w-[3px] rounded-full bg-[#163552]"
            style={{ left: `calc(${posisiBatas}% - 1px)` }}
          />
        </div>
        <input
          id="hujan-simulasi"
          type="range"
          min="0"
          max={maksimum}
          step="1"
          value={hujanMm}
          onChange={(event) => setHujanMm(Number(event.target.value))}
          aria-valuetext={`${hujanMm} milimeter per hari`}
          className="rain-range mt-[-27px] block h-[44px] w-full"
        />
        <div className="flex items-center justify-between text-[12px] font-semibold text-abu">
          <span>0 mm</span>
          <span>Ambang {batas} mm</span>
          <span>{maksimum} mm</span>
        </div>

        <div
          className={`mt-5 rounded-2xl p-4 ${
            memenuhiSyarat ? "bg-uang-muda" : "bg-[#eff6ff]"
          }`}
        >
          <p className="text-[13px] font-semibold text-abu">Hasil simulasi satu hari</p>
          <p
            role="status"
            aria-live="polite"
            aria-atomic="true"
            className={`angka mt-1 text-[27px] leading-tight font-extrabold ${
              memenuhiSyarat ? "text-uang" : "text-tinta"
            }`}
          >
            {memenuhiSyarat ? formatRupiah(zona.payoutPerDay) : "Belum ada bayaran"}
          </p>
          <p className="mt-1 text-[13px] leading-relaxed text-abu">
            {memenuhiSyarat
              ? `${hujanMm} mm mencapai ambang ${batas} mm. Berlaku jika polis sudah aktif dan jatah ${zona.maxDaysPerWeek} hari minggu ini belum habis.`
              : `${hujanMm} mm belum mencapai ambang ${batas} mm di ${zona.name}.`}
          </p>
        </div>

        <div className="mt-4 flex flex-wrap gap-2" aria-label="Contoh curah hujan">
          {contohHujan.map((contoh) => (
            <button
              key={contoh}
              type="button"
              onClick={() => setHujanMm(contoh)}
              aria-pressed={hujanMm === contoh}
              className={`min-h-[44px] rounded-xl border px-3 text-[13px] font-bold transition-colors ${
                hujanMm === contoh
                  ? "border-[#1a62c8] bg-[#e6f0ff] text-[#1457bb]"
                  : "border-[#b8cbe4] bg-white text-tinta active:bg-[#eff6ff]"
              }`}
            >
              {contoh} mm
            </button>
          ))}
        </div>

        <div className="mt-6 border-t border-[#dbe7f3] pt-5">
          <p className="text-[14px] font-bold">Lama polis</p>
          <div className="mt-2 grid grid-cols-2 gap-2 min-[360px]:grid-cols-4" role="group" aria-label="Lama polis simulasi">
            {PILIHAN_MINGGU.map((pilihan) => (
              <button
                key={pilihan}
                type="button"
                onClick={() => setMinggu(pilihan)}
                aria-pressed={minggu === pilihan}
                className={`min-h-[48px] rounded-xl border text-[12px] font-bold transition-colors sm:text-[14px] ${
                  minggu === pilihan
                  ? "border-[#1a62c8] bg-[#1a62c8] text-white"
                  : "border-[#b8cbe4] bg-white text-tinta active:bg-[#eff6ff]"
                }`}
              >
                {pilihan} minggu
              </button>
            ))}
          </div>
        </div>

        <div className="mt-5 flex items-end justify-between gap-3 rounded-[20px] bg-[#f1f6fd] px-4 py-4">
          <p className="text-[13px] font-medium text-abu">Premi untuk {minggu} minggu</p>
          <p className="angka text-right text-[21px] font-extrabold">
            {premiTersedia ? formatRupiah(multiplyWei(zona.premiumPerWeek, minggu)) : "Belum tersedia"}
          </p>
        </div>
      </div>

      <p className="mx-5 mb-5 border-t border-[#e4edf7] pt-4 text-[12px] leading-relaxed text-abu">
        Ini contoh perhitungan, bukan data hujan hari ini. Simulasi tidak membuat polis atau mengirim transaksi.
      </p>
    </section>
  );
}
