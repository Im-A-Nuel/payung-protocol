/**
 * The only place IDRP wei becomes rupiah. Everything else — API responses,
 * props, state — carries wei strings untouched, so no rounding can sneak
 * into a number the contract also knows about.
 */

const SATU_IDRP = 10n ** 18n;

function toBigInt(wei: string | bigint | null | undefined): bigint {
  if (typeof wei === "bigint") return wei;
  if (!wei) return 0n;
  try {
    return BigInt(wei);
  } catch {
    return 0n;
  }
}

/** "25000000000000000000000" -> "Rp 25.000" */
export function formatRupiah(wei: string | bigint | null | undefined): string {
  const rupiah = toBigInt(wei) / SATU_IDRP;
  return `Rp ${new Intl.NumberFormat("id-ID").format(rupiah)}`;
}

/** Same as formatRupiah without the prefix, for when the label already says rupiah. */
export function formatAngka(wei: string | bigint | null | undefined): string {
  return new Intl.NumberFormat("id-ID").format(toBigInt(wei) / SATU_IDRP);
}

/**
 * Whole rupiah as a plain number, for the one case that needs arithmetic:
 * animating a total upward. Payout totals are thousands of rupiah, nowhere
 * near the safe-integer ceiling, and nothing derived from this goes on-chain.
 */
export function rupiahBulat(wei: string | bigint | null | undefined): number {
  return Number(toBigInt(wei) / SATU_IDRP);
}

/** Formats a value that is already whole rupiah, not wei. */
export function formatRupiahAngka(rupiah: number): string {
  return `Rp ${new Intl.NumberFormat("id-ID").format(Math.round(rupiah))}`;
}

export function isLessThan(a: string | bigint, b: string | bigint): boolean {
  return toBigInt(a) < toBigInt(b);
}

export function multiplyWei(wei: string | bigint, factor: number): string {
  return (toBigInt(wei) * BigInt(factor)).toString();
}

export function toWei(wei: string | bigint | null | undefined): bigint {
  return toBigInt(wei);
}

const BULAN_PANJANG = [
  "Januari",
  "Februari",
  "Maret",
  "April",
  "Mei",
  "Juni",
  "Juli",
  "Agustus",
  "September",
  "Oktober",
  "November",
  "Desember",
];

const BULAN_PENDEK = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"];

const HARI_PENDEK = ["Min", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"];

function pecahTanggal(iso: string): { hari: number; bulan: number; tahun: number } | null {
  const cocok = /^(\d{4})-(\d{2})-(\d{2})$/.exec(iso ?? "");
  if (!cocok) return null;
  return { tahun: Number(cocok[1]), bulan: Number(cocok[2]), hari: Number(cocok[3]) };
}

/** "2026-09-16" -> "16 Sep" */
export function formatTanggalPendek(iso: string): string {
  const t = pecahTanggal(iso);
  if (!t) return iso ?? "";
  return `${t.hari} ${BULAN_PENDEK[t.bulan - 1]}`;
}

/** "2026-09-16" -> "16 September 2026" */
export function formatTanggalPanjang(iso: string): string {
  const t = pecahTanggal(iso);
  if (!t) return iso ?? "";
  return `${t.hari} ${BULAN_PANJANG[t.bulan - 1]} ${t.tahun}`;
}

/**
 * Whole days between two WIB calendar dates. This counts days between two
 * dates the API already resolved; it is not day_index math and must never
 * be used to derive one (see CLAUDE.md).
 */
export function selisihHari(dari: string, sampai: string): number {
  const a = pecahTanggal(dari);
  const b = pecahTanggal(sampai);
  if (!a || !b) return 0;
  const ms = Date.UTC(b.tahun, b.bulan - 1, b.hari) - Date.UTC(a.tahun, a.bulan - 1, a.hari);
  return Math.round(ms / 86_400_000);
}

/**
 * "2026-09-16" -> "Rab". Uses Date.UTC on a plain calendar date, so the
 * browser's own timezone can never shift the weekday.
 */
export function formatHari(iso: string): string {
  const t = pecahTanggal(iso);
  if (!t) return "";
  const hari = new Date(Date.UTC(t.tahun, t.bulan - 1, t.hari)).getUTCDay();
  return HARI_PENDEK[hari];
}
