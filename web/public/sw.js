/*
 * Minimal service worker: make the shell survive a dead signal, and never
 * touch the API. Policy and payout data always comes from the network, so a
 * driver can't be shown a stale balance or a payout that already changed.
 */
const CACHE = "payung-v1";
const HALAMAN_OFFLINE = "/offline";
const PRECACHE = [HALAMAN_OFFLINE, "/manifest.webmanifest", "/icon.svg"];

self.addEventListener("install", (event) => {
  event.waitUntil(
    caches
      .open(CACHE)
      .then((cache) => cache.addAll(PRECACHE))
      .then(() => self.skipWaiting()),
  );
});

self.addEventListener("activate", (event) => {
  event.waitUntil(
    caches
      .keys()
      .then((kunci) => Promise.all(kunci.filter((k) => k !== CACHE).map((k) => caches.delete(k))))
      .then(() => self.clients.claim()),
  );
});

self.addEventListener("fetch", (event) => {
  const req = event.request;
  if (req.method !== "GET") return;

  const url = new URL(req.url);
  // The Go API lives on another origin; leave every cross-origin request alone.
  if (url.origin !== self.location.origin) return;

  if (req.mode === "navigate") {
    event.respondWith(fetch(req).catch(() => caches.match(HALAMAN_OFFLINE)));
    return;
  }

  // Build output is content-hashed, so serving it from cache is always safe.
  if (url.pathname.startsWith("/_next/static/") || PRECACHE.includes(url.pathname)) {
    event.respondWith(
      caches.match(req).then(
        (hit) =>
          hit ??
          fetch(req).then((res) => {
            const salinan = res.clone();
            caches.open(CACHE).then((cache) => cache.put(req, salinan));
            return res;
          }),
      ),
    );
  }
});
