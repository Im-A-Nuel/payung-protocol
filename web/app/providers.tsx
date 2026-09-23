"use client";

import { type ReactNode, useState } from "react";
import { PrivyProvider } from "@privy-io/react-auth";
import { WagmiProvider, createConfig } from "@privy-io/wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { WagmiProvider as WagmiProviderPolos, http } from "wagmi";

import { AkunPratinjau, AkunPrivy } from "@/lib/akun";
import { rantai } from "@/lib/contracts";

// Privy validates app IDs synchronously and only accepts its 25-character
// dashboard ID. Keep an invalid deployment in preview mode so an accidental
// pasted secret or `NAME=value` string cannot fail Next's static build.
const PRIVY_APP_ID = (process.env.NEXT_PUBLIC_PRIVY_APP_ID ?? "").trim();
const ALAMAT_PRATINJAU = process.env.NEXT_PUBLIC_PREVIEW_ADDRESS ?? "";

/** True until a valid Privy dashboard App ID is present. */
export const MODE_PRATINJAU = PRIVY_APP_ID.length !== 25;

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
          loginMessage: "Pakai Google atau email. Payung akan menyiapkan dompet testnet untuk polismu.",
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
