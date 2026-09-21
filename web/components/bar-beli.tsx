"use client";

import { useEffect, useState } from "react";

import type { Zona } from "@/lib/api";
import { formatRupiah } from "@/lib/format";

/**
 * The purchase panel sits below five zone cards, which is about two screens
 * down on a 360px phone. This bar keeps the way forward in reach, and only
 * appears while the real panel is off screen so there are never two competing
 * buttons on the same view.
 */
export function BarBeli({ zona, terkunci }: { zona: Zona; terkunci: boolean }) {
  const [tampil, setTampil] = useState(false);

  useEffect(() => {
    const panel = document.getElementById("panel-beli");
    if (!panel) return;

    // Discount the bottom strip of the viewport: the tab bar and this bar sit
    // there, so a panel merely peeking above them is not yet usable and the
    // shortcut should stay.
    const pengamat = new IntersectionObserver(([masuk]) => setTampil(!masuk.isIntersecting), {
      rootMargin: "0px 0px -180px 0px",
    });
    pengamat.observe(panel);
    return () => pengamat.disconnect();
  }, []);

  if (!tampil || terkunci) return null;

  return (
    <div
      className="naik fixed inset-x-0 z-10 border-t border-garis bg-kartu px-4 py-3 shadow-[0_-6px_20px_rgba(12,24,38,0.07)]"
      style={{ bottom: "calc(60px + env(safe-area-inset-bottom, 0px))" }}
    >
      <div className="mx-auto flex w-full max-w-[520px] items-center gap-3">
        <div className="min-w-0 flex-1">
          <p className="truncate text-[15px] font-extrabold">{zona.name}</p>
          <p className="angka text-[13px] text-abu">{formatRupiah(zona.premiumPerWeek)} per minggu</p>
        </div>
        <button
          type="button"
          onClick={() =>
            document.getElementById("panel-beli")?.scrollIntoView({ behavior: "smooth", block: "center" })
          }
          className="tekan min-h-[48px] shrink-0 rounded-xl bg-langit px-5 text-[15px] font-bold text-white active:bg-langit-tua"
        >
          Lanjut
        </button>
      </div>
    </div>
  );
}
