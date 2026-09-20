# CLAUDE.md

## Project
Payung: parametric rain insurance for Indonesian ride-hailing drivers on opBNB. Buy a weekly policy, get paid automatically when daily rainfall in your zone crosses a threshold. No claim button.

Hackathon: Indonesia Web3 Hackathon 2026. Deadline 30 Sep 2026 23:59 WIB. Solo developer. Read `docs/ROADMAP.md` for the current phase.

## Stack
- Contracts: Solidity 0.8.26, Foundry, OpenZeppelin (AccessControl, ReentrancyGuard, SafeERC20)
- Chain: opBNB testnet, chain id 5611
- Backend: Go 1.22, chi, pgx + sqlc, go-ethereum abigen bindings, PostgreSQL 16
- AI: Claude API (claude-sonnet-4-6) via official Go SDK, JSON outputs only
- Frontend: Next.js 15 App Router, TypeScript, Tailwind, wagmi + viem, Privy
- Rain data: Open-Meteo (no API key)

## Commands
```bash
# Contracts
cd contracts && forge build && forge test -vvv
forge script script/Deploy.s.sol --rpc-url $OPBNB_RPC --broadcast --verify

# Backend
cd backend && go test ./...
go run ./cmd/migrate
go run ./cmd/api                 # :8080
go run ./cmd/oracle --once       # yesterday's rain for all zones
sqlc generate                    # after editing internal/store/queries/*.sql
./scripts/gen-abi.sh              # after editing contracts/src/*.sol; regenerates internal/chain/*.go

# Frontend
cd web && pnpm install && pnpm dev
pnpm build
```

## Project Structure
```
contracts/src/PayungPool.sol     core contract
contracts/src/IDRP.sol           mock IDR token
contracts/test/                  Foundry tests (10 required cases in docs/SCHEMA.md)
backend/cmd/{api,oracle,migrate}
backend/internal/chain           signer, tx sending, abigen bindings
backend/internal/rain            Open-Meteo client
backend/internal/pricing         pure actuarial functions
backend/internal/ai              Claude prompts + fallbacks
backend/internal/indexer         event -> Postgres
backend/internal/store           sqlc generated + queries
web/app/{page,polis,riwayat}     three screens
web/lib/api.ts                   all backend calls
web/lib/contracts.ts             ABI + addresses
docs/                            REQUIREMENTS, ARCHITECTURE, SCHEMA, ROADMAP
```

## Key Conventions
- `day_index` = `floor((unix + 7*3600) / 86400)`. Same formula in Go (`internal/chain/day.go`) and Solidity (`today()`). Never compute days any other way.
- `week_index` = `day_index / 7`. Payout caps are per week_index.
- Money on-chain and in Postgres is IDRP wei (18 decimals). Format to rupiah only in `web/lib/format.ts`.
- Rainfall is integer mm everywhere.
- Addresses stored lowercase hex in Postgres.
- Go: errors wrapped with `fmt.Errorf("pkg.Func: %w", err)`. No panics outside `main`.
- Go: config from env only (`internal/config`), fail fast at startup if missing.
- Solidity: custom errors, not revert strings. Events for every state change.
- Frontend: all UI copy in Indonesian. Never show hex addresses, gas, or "transaction" to the driver; say "sedang diproses" and "berhasil".
- Commits: `contracts:`, `backend:`, `web:`, `docs:` prefixes.

## Architecture Notes
- The contract decides who gets paid. Backend only supplies rainfall and calls `settle`. AI never touches money.
- Premium numbers come from `internal/pricing` (deterministic). Claude only writes narratives. If the Claude call fails, use the fallback string; never block the flow.
- Frontend reads from the Go API, never from RPC. The only frontend write is `buyPolicy` (plus faucet), gas-sponsored.
- Settlement is push-based on purpose (demo). `MAX_ACTIVE_PER_ZONE = 50` protects the loop. Pull-based is post-MVP; do not implement it now.
- Oracle is a single signer with `ORACLE_ROLE`. This is a known, documented limitation. Do not add oracle decentralization during the hackathon.
- `POST /admin/*` endpoints are for demo and must be disabled when `ENV=production`.
- Gasless: try opBNB MegaFuel first. If it takes more than half a day, switch to the relayer fallback in `docs/SCHEMA.md` and move on.

## Do NOT
- Do not add features not listed in `docs/ROADMAP.md` phases 1 to 5 without asking.
- Do not change `day_index` math, threshold, payout, or cap constants without updating both Solidity and Go and their tests.
- Do not put private keys, admin keys, or Anthropic keys anywhere except env vars.
- Do not call the RPC from React components.
- Do not let the LLM output decide any number that goes on-chain.
- Do not use revert strings in Solidity; use custom errors.
- Do not install new dependencies in `web/` without a reason in the commit message.
- Do not write UI copy in English.

## Current Focus
Phase 1 (contracts) and Phase 2 (oracle & API) are code-complete and tested locally (Foundry tests green; backend `go test ./...` green against Postgres; full vertical slice verified end-to-end on a local anvil chain: buy policy -> admin simulate-rain -> submit+settle -> indexer -> driver dashboard endpoints).

Still needed before Phase 3 (frontend):
- Deploy `PayungPool`/`IDRP` to opBNB testnet and verify on BscScan (needs real `PRIVATE_KEY`, `BSCSCAN_API_KEY`).
- Deploy backend (`cmd/api`, `cmd/oracle`) to Railway with Postgres, oracle cron at 06:00 WIB (needs Railway credentials, `ORACLE_PRIVATE_KEY`, `FAUCET_PRIVATE_KEY`).
- Point `POOL_ADDRESS`/`IDRP_ADDRESS`/`DEPLOY_BLOCK` at the real deployment once it exists.

Phase 4 (AI pricing/explanations) is intentionally not implemented yet: `GET /zones` returns a fallback premium/narrative until `internal/pricing` and `internal/ai` land.
