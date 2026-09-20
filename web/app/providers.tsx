"use client";

import { type ReactNode, useState } from "react";
import { PrivyProvider } from "@privy-io/react-auth";
import { WagmiProvider, createConfig } from "@privy-io/wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { WagmiProvider as WagmiProviderPolos, http } from "wagmi";

import { AkunPratinjau, AkunPrivy } from "@/lib/akun";
import { rantai } from "@/lib/contracts";

const PRIVY_APP_ID = process.env.NEXT_PUBLIC_PRIVY_APP_ID ?? "";
const ALAMAT_PRATINJAU = process.env.NEXT_PUBLIC_PREVIEW_ADDRESS ?? "";

/** True when Privy isn't configured: the app renders read-only instead of crashing. */
export const MODE_PRATINJAU = PRIVY_APP_ID === "";

const wagmiConfig = createConfig({
  chains: [rantai],
  transports: { [rantai.id]: http() },
});

export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            // A driver checking their phone under a shelter should see fresh
            // numbers without hammering the API on every focus change.
            staleTime: 15_000,
            retry: 1,
            refetchOnWindowFocus: true,
          },
        },
      }),
  );

  if (MODE_PRATINJAU) {
    // Still mount plain wagmi so components that call its hooks render the
    // same way they will once Privy is configured. With no connector there
    // is no account, which matches this mode's read-only promise.
    return (
      <QueryClientProvider client={queryClient}>
        <WagmiProviderPolos config={wagmiConfig}>
          <AkunPratinjau alamat={ALAMAT_PRATINJAU}>{children}</AkunPratinjau>
        </WagmiProviderPolos>
      </QueryClientProvider>
    );
  }

  return (
    <PrivyProvider
      appId={PRIVY_APP_ID}
      config={{
        loginMethods: ["google", "email"],
        embeddedWallets: { ethereum: { createOnLogin: "users-without-wallets" } },
        defaultChain: rantai,
        supportedChains: [rantai],
        appearance: {
          theme: "light",
          accentColor: "#0b5fd0",
          landingHeader: "Masuk ke Payung",
          loginMessage: "Pakai akun Google kamu. Tidak perlu isi data apa pun.",
          showWalletLoginFirst: false,
          walletList: [],
        },
      }}
    >
      <QueryClientProvider client={queryClient}>
        <WagmiProvider config={wagmiConfig}>
          <AkunPrivy>{children}</AkunPrivy>
        </WagmiProvider>
      </QueryClientProvider>
    </PrivyProvider>
  );
}
