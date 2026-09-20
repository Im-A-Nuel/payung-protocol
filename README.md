# Payung

Parametric rain insurance for Indonesian ride-hailing drivers. Pay a small weekly premium, get paid automatically when it rains, no claim button.

Built for Indonesia Web3 Hackathon 2026 (#WhereBuildersBuild). Tracks: Finance & Commerce, Consumer Apps. Chain: opBNB testnet.

## Overview

Ride-hailing drivers (ojol) lose most of their daily income when heavy rain hits: fewer orders, slower trips, higher accident risk. No insurer serves this segment because manual claims cost more than the premium. Payung removes the claim entirely. A driver buys a policy for their zone, the pool watches daily rainfall for that zone, and when rainfall crosses a threshold the contract pushes the payout to the driver's wallet. The driver never files anything.

The product has three parts. A `PayungPool` smart contract holds premiums and executes payouts by rule. A Go oracle service pulls daily rainfall per zone from Open-Meteo, submits it on-chain, and triggers settlement. An AI layer prices premiums per zone from three years of rainfall history and explains every payout to the driver in plain Indonesian.

Target user: ojol drivers in Yogyakarta, Sleman, Bantul, Solo, and Semarang for the MVP. The driver never sees gas, seed phrases, or hex addresses; social login creates the wallet and gas is sponsored.

## Tech Stack

| Layer | Choice |
|---|---|
| Smart contract | Solidity 0.8.x, Foundry, OpenZeppelin |
| Chain | opBNB testnet (chain id 5611), verified on BscScan |
| Gas sponsorship | opBNB MegaFuel paymaster, fallback: backend relayer |
| Backend / oracle | Go 1.22, chi router, sqlc + PostgreSQL, go-ethereum |
| Rain data | Open-Meteo Archive & Forecast API (no key required) |
| AI | Claude API (claude-sonnet-4-6): premium narrative, payout explanation |
| Frontend | Next.js 15 (App Router), TypeScript, Tailwind, wagmi + viem, Privy (social login + embedded wallet) |
| Hosting | Vercel (web), Railway (Go service + Postgres) |

## Quick Start

```bash
# 1. Contracts
cd contracts
forge install
forge test
cp .env.example .env            # PRIVATE_KEY, OPBNB_RPC, BSCSCAN_API_KEY
forge script script/Deploy.s.sol --rpc-url $OPBNB_RPC --broadcast --verify

# 2. Backend (oracle + API)
cd ../backend
cp .env.example .env            # DATABASE_URL, ORACLE_PRIVATE_KEY, POOL_ADDRESS, ANTHROPIC_API_KEY, ADMIN_KEY
docker compose up -d postgres
go run ./cmd/migrate
go run ./cmd/api                # http://localhost:8080
go run ./cmd/oracle --once      # fetch yesterday's rain, submit, settle

# 3. Frontend
cd ../web
cp .env.example .env.local      # NEXT_PUBLIC_POOL_ADDRESS, NEXT_PUBLIC_API_URL, NEXT_PUBLIC_PRIVY_APP_ID
pnpm install
pnpm dev                        # http://localhost:3000
```

Demo shortcut: `POST /admin/simulate-rain` injects rainfall for a zone so the payout flow can be shown without waiting for weather.

## Project Structure

```
payung/
  contracts/        Foundry project: PayungPool.sol, IDRP.sol (mock IDR token), tests, deploy script
  backend/
    cmd/api         HTTP API for the web app
    cmd/oracle      daily rain fetch + on-chain submit + settle
    cmd/migrate     DB migrations
    internal/
      chain/        go-ethereum bindings and tx sending
      rain/         Open-Meteo client
      pricing/      actuarial calc from rainfall history
      ai/           Claude API wrappers (premium narrative, payout explainer)
      store/        sqlc queries
  web/              Next.js app (mobile-first PWA)
  docs/             this folder
  CLAUDE.md         agent context file
```

## Features

MVP (must work in the demo video):

- Buy a weekly policy for a zone with a sponsored (gasless) transaction
- Daily rainfall per zone submitted on-chain by the oracle
- Automatic push payout to every active policy in a zone when rainfall >= threshold, capped per week
- Policy dashboard: active policy, rainfall history, payout history with tx links
- AI premium table per zone per month with a plain-language reason
- AI payout explanation delivered to the driver after every payout

Parked (see ROADMAP.md): multi-oracle, pull-based settlement at scale, dynamic pricing, fraud detection, mainnet IDR stablecoin.

## Honest Limitations

- The oracle is a single backend signer. This is stated in the pitch; the roadmap is multi-source attestation or Chainlink Functions.
- Rainfall granularity is city/regency level, not sub-district.
- The IDR token on testnet is a mock (`IDRP`). Mainnet target is a licensed IDR stablecoin.

## License

MIT
