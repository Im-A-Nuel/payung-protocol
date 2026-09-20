/**
 * Every backend call lives here. The app reads from the Go API and never
 * from an RPC node, so a slow public testnet endpoint can't make the
 * dashboard hang (see docs/ARCHITECTURE.md).
 */

const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/v1";

export type Zona = {
  id: number;
  name: string;
  thresholdMm: number;
  payoutPerDay: string;
  maxDaysPerWeek: number;
  premiumPerWeek: string;
  premiumNarrative: string;
};

export type HariHujan = {
  dayIndex: number;
  date: string;
  mm: number;
  isRainDay: boolean;
  source: string;
};

export type HujanZona = {
  thresholdMm: number;
  days: HariHujan[];
};

export type Polis = {
  policyId: number;
  zoneId: number;
  zoneName: string;
  startDate: string;
  endDate: string;
  daysLeft: number;
  payoutsThisWeek: number;
  maxDaysPerWeek: number;
};

export type Payout = {
  date: string;
  zoneName: string;
  mm: number;
  amount: string;
  txHash: string;
  explanation: string | null;
};

export type Dompet = {
  balance: string;
  allowance: string;
};

/** Carries the backend's own Indonesian message so screens can show it as-is. */
export class ApiError extends Error {
  readonly code: string;

  constructor(code: string, message: string) {
    super(message);
    this.name = "ApiError";
    this.code = code;
  }
}

const PESAN_GAGAL = "Koneksi bermasalah. Coba lagi sebentar lagi.";

async function minta<T>(path: string, init?: RequestInit): Promise<T> {
  let res: Response;
  try {
    res = await fetch(`${BASE_URL}${path}`, {
      ...init,
      headers: { "Content-Type": "application/json", ...init?.headers },
      cache: "no-store",
    });
  } catch {
    throw new ApiError("NETWORK", PESAN_GAGAL);
  }

  const teks = await res.text();
  const data = teks ? safeParse(teks) : null;

  if (!res.ok) {
    const err = (data as { error?: { code?: string; message?: string } } | null)?.error;
    throw new ApiError(err?.code ?? "INTERNAL", err?.message ?? PESAN_GAGAL);
  }

  return data as T;
}

function safeParse(teks: string): unknown {
  try {
    return JSON.parse(teks);
  } catch {
    return null;
  }
}

export function getZones(): Promise<Zona[]> {
  return minta<Zona[]>("/zones");
}

export function getZoneRain(zoneId: number, days = 7): Promise<HujanZona> {
  return minta<HujanZona>(`/zones/${zoneId}/rain?days=${days}`);
}

/** Returns null when the driver has no active policy. */
export function getPolicy(address: string): Promise<Polis | null> {
  return minta<Polis | null>(`/drivers/${address}/policy`);
}

export async function getPayouts(address: string): Promise<Payout[]> {
  return (await minta<Payout[] | null>(`/drivers/${address}/payouts`)) ?? [];
}

export function getWallet(address: string): Promise<Dompet> {
  return minta<Dompet>(`/drivers/${address}/wallet`);
}

export function claimFaucet(address: string): Promise<{ txHash: string }> {
  return minta<{ txHash: string }>("/faucet", {
    method: "POST",
    body: JSON.stringify({ address }),
  });
}
