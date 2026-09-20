import type { Zona } from "@/lib/api";
import { formatRupiah } from "@/lib/format";

export function KartuZona({
  zona,
  terpilih,
  terkunci,
  tampilkanAturan,
  tampilkanNarasi,
  onPilih,
}: {
  zona: Zona;
  terpilih: boolean;
  terkunci: boolean;
  /** False when every zone shares the same payout rule and it is stated once above the list. */
  tampilkanAturan: boolean;
  /** False when every zone carries the same narrative and it is stated once above the list. */
  tampilkanNarasi: boolean;
  onPilih: (id: number) => void;
}) {
  return (
    <button
      type="button"
      onClick={() => onPilih(zona.id)}
      aria-pressed={terpilih}
      className={`tekan w-full rounded-2xl border p-4 text-left transition-colors ${
        terpilih ? "border-langit bg-langit-muda" : "border-garis bg-kartu"
      }`}
    >
      <div className="flex items-start justify-between gap-3">
        <div className="min-w-0">
          <p className="flex items-center gap-1.5 text-[17px] font-extrabold">
            {terpilih ? <Centang /> : null}
            {zona.name}
          </p>
          {tampilkanAturan ? (
            <p className="mt-0.5 text-[13px] text-abu">
              Bayar {formatRupiah(zona.payoutPerDay)} tiap hari hujan lewat {zona.thresholdMm} mm
            </p>
          ) : null}
        </div>
        <div className="shrink-0 text-right">
          <p className="angka text-[17px] leading-tight font-extrabold">{formatRupiah(zona.premiumPerWeek)}</p>
          <p className="text-[12px] text-abu">per minggu</p>
        </div>
      </div>

      {tampilkanNarasi && zona.premiumNarrative ? (
        <p className="mt-3 border-t border-garis pt-3 text-[13px] leading-relaxed text-abu">
          {zona.premiumNarrative}
        </p>
      ) : null}

      {terkunci ? (
        <p className="mt-3 rounded-xl bg-kertas px-3 py-2 text-[13px] font-semibold text-abu">
          Polismu di zona ini masih jalan
        </p>
      ) : null}
    </button>
  );
}

/** Selection needs a shape, not just a tint, for anyone who reads colour poorly. */
function Centang() {
  return (
    <span className="lingkar flex h-5 w-5 shrink-0 items-center justify-center rounded-full bg-langit">
      <svg width="12" height="12" viewBox="0 0 24 24" fill="none" aria-hidden="true">
        <path
          d="M5 12.5 10 17.5 19 7.5"
          stroke="white"
          strokeWidth="3.4"
          strokeLinecap="round"
          strokeLinejoin="round"
        />
      </svg>
    </span>
  );
}
