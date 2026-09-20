import { defineChain } from "viem";

/**
 * Chain and contract wiring. Defaults target opBNB testnet; the env
 * overrides exist so the app can be pointed at a local anvil node during
 * development without touching code.
 */

const CHAIN_ID = Number(process.env.NEXT_PUBLIC_CHAIN_ID ?? 5611);
const RPC_URL = process.env.NEXT_PUBLIC_RPC_URL ?? "https://opbnb-testnet-rpc.bnbchain.org";
const EXPLORER_URL = process.env.NEXT_PUBLIC_EXPLORER_URL ?? "https://opbnb-testnet.bscscan.com";

export const rantai = defineChain({
  id: CHAIN_ID,
  name: "opBNB Testnet",
  nativeCurrency: { name: "tBNB", symbol: "tBNB", decimals: 18 },
  rpcUrls: { default: { http: [RPC_URL] } },
  blockExplorers: { default: { name: "opBNBScan", url: EXPLORER_URL } },
  testnet: true,
});

export const POOL_ADDRESS = (process.env.NEXT_PUBLIC_POOL_ADDRESS ?? "") as `0x${string}`;
export const IDRP_ADDRESS = (process.env.NEXT_PUBLIC_IDRP_ADDRESS ?? "") as `0x${string}`;

/** Link a driver can open to see the payout for themselves. */
export function urlTransaksi(txHash: string): string {
  return `${EXPLORER_URL}/tx/${txHash}`;
}

/** Only the two calls the app actually sends. Reads go through the Go API. */
export const poolAbi = [
  {
    type: "function",
    name: "buyPolicy",
    stateMutability: "nonpayable",
    inputs: [
      { name: "zoneId", type: "uint16" },
      { name: "weeksCount", type: "uint8" },
    ],
    outputs: [{ name: "policyId", type: "uint256" }],
  },
] as const;

export const idrpAbi = [
  {
    type: "function",
    name: "approve",
    stateMutability: "nonpayable",
    inputs: [
      { name: "spender", type: "address" },
      { name: "value", type: "uint256" },
    ],
    outputs: [{ name: "", type: "bool" }],
  },
] as const;
