import type { HariHujan } from "@/lib/api";
import { formatHari, formatTanggalPendek } from "@/lib/format";

/**
 * Seven daily rainfall totals against the payout threshold. Bars, not a
 * line: each day is its own total, nothing flows between them.
 *
 * There is no hover layer because this is read on a phone, so every bar
 * carries a visible mm label. That also covers the light-blue fill's low
 * contrast against white, which needs visible labels as relief.
 */

const W = 320;
const H = 158;
const PAD_X = 10;
const TOP = 18; // room for the value label above the tallest bar
const BASELINE = 118;
const LABEL_HARI_Y = 134;
const LEBAR_BATANG = 26;

export function GrafikHujan({
  hari,
  ambangMm,
  namaZona,
}: {
  hari: HariHujan[];
  ambangMm: number;
  namaZona?: string;
}) {
  if (hari.length === 0) {
    return (
      <p className="py-6 text-center text-[15px] text-abu">
        Data hujan zona ini belum masuk. Biasanya terisi tiap pagi.
      </p>
    );
  }

  const tertinggi = Math.max(ambangMm, ...hari.map((h) => h.mm));
  const skalaMax = tertinggi * 1.15 || 1;
  const tinggiPlot = BASELINE - TOP;
  const slot = (W - PAD_X * 2) / hari.length;

  const yDari = (mm: number) => BASELINE - (mm / skalaMax) * tinggiPlot;
  const yAmbang = yDari(ambangMm);

  const ringkasan = [
    namaZona ? `Curah hujan tujuh hari terakhir di ${namaZona}.` : "Curah hujan tujuh hari terakhir.",
    `Ambang bayar ${ambangMm} mm.`,
    ...hari.map((h) => `${formatTanggalPendek(h.date)}: ${h.mm} mm${h.isRainDay ? ", dibayar" : ""}.`),
  ].join(" ");

  return (
    <figure className="m-0">
      <svg viewBox={`0 0 ${W} ${H}`} width="100%" role="img" aria-label={ringkasan} className="block">
        <line
          x1={PAD_X}
          x2={W - PAD_X}
          y1={BASELINE}
          y2={BASELINE}
          stroke="var(--color-garis)"
          strokeWidth="1"
        />

        <line
          x1={PAD_X}
          x2={W - PAD_X}
          y1={yAmbang}
          y2={yAmbang}
          stroke="var(--color-ambang)"
          strokeWidth="2"
          strokeDasharray="5 4"
          strokeLinecap="round"
        />

        {hari.map((h, i) => {
          const tengah = PAD_X + slot * i + slot / 2;
          const x = tengah - LEBAR_BATANG / 2;
          const y = yDari(h.mm);
          const tinggi = Math.max(BASELINE - y, h.mm > 0 ? 3 : 2);
          const atas = BASELINE - tinggi;

          return (
            <g key={h.dayIndex}>
              <path
                d={batangMembulat(x, atas, LEBAR_BATANG, tinggi, 4)}
                fill={h.isRainDay ? "var(--color-uang)" : "var(--color-hujan)"}
              />
              <text
                x={tengah}
                y={atas - 6}
                textAnchor="middle"
                fontSize="11"
                fontWeight={h.isRainDay ? 800 : 600}
                fill={h.isRainDay ? "var(--color-tinta)" : "var(--color-abu)"}
              >
                {h.mm}
              </text>
              <text
                x={tengah}
                y={LABEL_HARI_Y}
                textAnchor="middle"
                fontSize="11"
                fontWeight="600"
                fill="var(--color-abu)"
              >
                {formatHari(h.date)}
              </text>
            </g>
          );
        })}
      </svg>

      <figcaption className="mt-2 flex flex-wrap items-center gap-x-4 gap-y-1 text-[12px] text-abu">
        <span className="flex items-center gap-1.5">
          <span className="h-2.5 w-2.5 rounded-full bg-uang" aria-hidden="true" />
          Dibayar
        </span>
        <span className="flex items-center gap-1.5">
          <span className="h-2.5 w-2.5 rounded-full bg-hujan" aria-hidden="true" />
          Belum sampai ambang
        </span>
        <span className="flex items-center gap-1.5">
          <svg width="16" height="4" aria-hidden="true">
            <line
              x1="0"
              y1="2"
              x2="16"
              y2="2"
              stroke="var(--color-ambang)"
              strokeWidth="2"
              strokeDasharray="5 4"
            />
          </svg>
          Ambang {ambangMm} mm
        </span>
      </figcaption>
    </figure>
  );
}

/** A bar with only its top corners rounded, so it stays anchored to the axis. */
function batangMembulat(x: number, y: number, w: number, h: number, r: number): string {
  const jari = Math.min(r, h, w / 2);
  return [
    `M ${x} ${y + h}`,
    `L ${x} ${y + jari}`,
    `Q ${x} ${y} ${x + jari} ${y}`,
    `L ${x + w - jari} ${y}`,
    `Q ${x + w} ${y} ${x + w} ${y + jari}`,
    `L ${x + w} ${y + h}`,
    "Z",
  ].join(" ");
}
