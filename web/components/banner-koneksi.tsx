"use client";

import { useEffect, useState } from "react";

/**
 * Drivers lose signal under bridges and inside basements, and the dashboard
 * then shows an error per card with no shared explanation. One honest line
 * at the top beats three identical failures below it.
 */
export function BannerKoneksi() {
  const [putus, setPutus] = useState(false);

  useEffect(() => {
    const perbarui = () => setPutus(!navigator.onLine);
    perbarui();

    window.addEventListener("online", perbarui);
    window.addEventListener("offline", perbarui);
    return () => {
      window.removeEventListener("online", perbarui);
      window.removeEventListener("offline", perbarui);
    };
  }, []);

  if (!putus) return null;

  return (
    <div
      role="status"
      className="naik sticky top-0 z-30 -mx-4 mb-4 bg-tinta px-4 py-2.5 text-center text-[13px] font-semibold text-white"
    >
      Sinyal putus. Angka di layar ini mungkin belum yang terbaru.
    </div>
  );
}
