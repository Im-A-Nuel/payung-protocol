"use client";

import { useEffect } from "react";

/** Registers the service worker in production only, so dev reloads stay predictable. */
export function DaftarSW() {
  useEffect(() => {
    if (process.env.NODE_ENV !== "production") return;
    if (!("serviceWorker" in navigator)) return;

    const daftar = () => navigator.serviceWorker.register("/sw.js").catch(() => {});
    if (document.readyState === "complete") daftar();
    else window.addEventListener("load", daftar, { once: true });
  }, []);

  return null;
}
