import type { CSSProperties } from "react";

/**
 * Decorative rain behind the hero copy. It states the product's premise
 * before anyone reads a word, and it is the one place a driver looks while
 * deciding whether this is worth their Rp 5.000.
 *
 * Deliberately cheap: a dozen absolutely positioned lines translating on the
 * compositor. No canvas, no library, and it stops entirely under
 * prefers-reduced-motion (see .tetes in globals.css).
 */

// Fixed values rather than Math.random(), so the server and client markup
// match and React never has to patch it on hydration.
const TETES = [
  { kiri: 6, lama: 1.5, tunda: 0, tinggi: 14 },
  { kiri: 14, lama: 1.2, tunda: 420, tinggi: 10 },
  { kiri: 23, lama: 1.7, tunda: 180, tinggi: 16 },
  { kiri: 31, lama: 1.3, tunda: 900, tinggi: 11 },
  { kiri: 40, lama: 1.6, tunda: 320, tinggi: 15 },
  { kiri: 48, lama: 1.25, tunda: 700, tinggi: 9 },
  { kiri: 57, lama: 1.55, tunda: 120, tinggi: 14 },
  { kiri: 65, lama: 1.35, tunda: 980, tinggi: 12 },
  { kiri: 74, lama: 1.65, tunda: 540, tinggi: 16 },
  { kiri: 82, lama: 1.28, tunda: 260, tinggi: 10 },
  { kiri: 90, lama: 1.5, tunda: 820, tinggi: 13 },
  { kiri: 96, lama: 1.4, tunda: 60, tinggi: 11 },
];

export function HujanAnimasi() {
  return (
    <div aria-hidden="true" className="pointer-events-none absolute inset-0 overflow-hidden">
      {TETES.map((t) => (
        <span
          key={t.kiri}
          className="tetes absolute top-0 block w-px rounded-full bg-white"
          style={
            {
              left: `${t.kiri}%`,
              height: `${t.tinggi}px`,
              "--lama": `${t.lama}s`,
              "--tunda": `${t.tunda}ms`,
            } as CSSProperties
          }
        />
      ))}
    </div>
  );
}
