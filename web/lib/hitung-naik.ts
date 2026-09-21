"use client";

import { useEffect, useRef, useState } from "react";

/**
 * Counts a number up to its target once, on mount. Used for the payout
 * total, where watching the figure climb is the point: it is the moment the
 * product's promise becomes visible.
 *
 * Returns the target immediately when the viewer asked for less motion, or
 * when there is nothing to count.
 */
export function useHitungNaik(target: number, lamaMs = 900): number {
  const kurangGerak =
    typeof window !== "undefined" && window.matchMedia?.("(prefers-reduced-motion: reduce)").matches;

  const [nilai, setNilai] = useState(() => (kurangGerak ? target : 0));
  const frameRef = useRef<number | null>(null);

  useEffect(() => {
    if (kurangGerak || target <= 0) {
      setNilai(target);
      return;
    }

    const mulai = performance.now();
    const langkah = (sekarang: number) => {
      const maju = Math.min((sekarang - mulai) / lamaMs, 1);
      // Ease out: fast at first, settling into the final figure.
      setNilai(Math.round(target * (1 - Math.pow(1 - maju, 3))));
      if (maju < 1) frameRef.current = requestAnimationFrame(langkah);
    };
    frameRef.current = requestAnimationFrame(langkah);

    return () => {
      if (frameRef.current !== null) cancelAnimationFrame(frameRef.current);
    };
  }, [target, lamaMs, kurangGerak]);

  return nilai;
}
