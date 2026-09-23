"use client";

import { createContext, useContext, useMemo, type ReactNode } from "react";
import { usePrivy } from "@privy-io/react-auth";
import { useAccount } from "wagmi";

/**
 * One shape for "who is signed in", so screens never touch Privy directly.
 * Two implementations back it: the real Privy one, and a read-only preview
 * used when no Privy app ID is configured yet (buying is disabled there).
 */
export type Akun = {
  /** False while the auth library is still booting. */
  siap: boolean;
  sudahMasuk: boolean;
  alamat: `0x${string}` | null;
  masuk: () => void;
  keluar: () => void;
  /** False in read-only preview mode, where the sign-in action is unavailable. */
  bisaMasuk: boolean;
  /** False in preview mode: reads work, sending anything on-chain does not. */
  bisaKirim: boolean;
};

const AkunContext = createContext<Akun | null>(null);

export function useAkun(): Akun {
  const akun = useContext(AkunContext);
  if (!akun) throw new Error("useAkun dipakai di luar AkunProvider");
  return akun;
}

export function AkunPrivy({ children }: { children: ReactNode }) {
  const { ready, authenticated, login, logout, user } = usePrivy();
  const { address } = useAccount();

  const nilai = useMemo<Akun>(() => {
    // wagmi's address is the one a write would actually come from; Privy's
    // user record fills the gap while the embedded wallet connects.
    const alamat = (address ?? user?.wallet?.address ?? null) as `0x${string}` | null;
    return {
      siap: ready,
      sudahMasuk: authenticated,
      alamat: authenticated ? alamat : null,
      masuk: login,
      keluar: logout,
      bisaMasuk: true,
      bisaKirim: authenticated && Boolean(address),
    };
  }, [ready, authenticated, login, logout, user?.wallet?.address, address]);

  return <AkunContext.Provider value={nilai}>{children}</AkunContext.Provider>;
}

export function AkunPratinjau({ alamat, children }: { alamat: string; children: ReactNode }) {
  const nilai = useMemo<Akun>(
    () => ({
      siap: true,
      sudahMasuk: Boolean(alamat),
      alamat: (alamat || null) as `0x${string}` | null,
      masuk: () => {},
      keluar: () => {},
      bisaMasuk: false,
      bisaKirim: false,
    }),
    [alamat],
  );

  return <AkunContext.Provider value={nilai}>{children}</AkunContext.Provider>;
}
