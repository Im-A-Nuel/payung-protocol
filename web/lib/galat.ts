import { ApiError } from "./api";

/**
 * Turns anything thrown by wagmi/viem or the API into one sentence a driver
 * can act on. Raw revert names, hex data and English library text never
 * reach the screen.
 */

const PETA: Array<[RegExp, string]> = [
  [/PolicyAlreadyActive/i, "Kamu sudah punya polis aktif di zona ini. Tunggu sampai masa berlakunya habis."],
  [/ZoneInactive/i, "Zona ini sedang tidak menerima polis baru."],
  [/ZoneFull/i, "Kuota zona ini sudah penuh. Coba zona lain dulu."],
  [/InvalidWeeks/i, "Lama perlindungan cuma bisa 1 sampai 4 minggu."],
  [/ERC20InsufficientAllowance|InsufficientAllowance/i, "Persetujuan saldo belum selesai. Coba tekan sekali lagi."],
  [
    /ERC20InsufficientBalance|exceeds balance|InsufficientBalance/i,
    "Saldo IDRP kamu kurang untuk premi ini. Ambil saldo testnet dulu.",
  ],
  [/user rejected|user denied|rejected the request/i, "Kamu membatalkan permintaannya."],
  [
    /insufficient funds|gas required|out of gas/i,
    "Saldo tBNB di dompet ini belum cukup untuk biaya jaringan opBNB Testnet. Isi tBNB lalu coba lagi.",
  ],
  [/timeout|timed out/i, "Jaringan lagi lambat. Coba lagi sebentar lagi."],
];

const UMUM = "Ada yang gagal. Coba lagi sebentar lagi.";

export function pesanGalat(err: unknown): string {
  if (err instanceof ApiError) return err.message;

  const teks =
    err instanceof Error ? `${err.message} ${(err as { details?: string }).details ?? ""}` : String(err ?? "");

  for (const [pola, pesan] of PETA) {
    if (pola.test(teks)) return pesan;
  }
  return UMUM;
}
