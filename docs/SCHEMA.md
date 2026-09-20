# Database, API & Smart Contract Design: Payung

Last updated: 17 September 2026

## Conventions

- `day_index`: integer, days since Unix epoch in WIB (`floor((unix + 7*3600) / 86400)`). Used identically in Go and Solidity so both sides agree on "which day".
- `week_index`: `day_index / 7` (integer division). Payout caps are per week_index, not per policy start.
- Money: IDRP with 18 decimals on-chain; stored in Postgres as `NUMERIC(38,0)` in wei; formatted as rupiah only in the UI.
- Rainfall: integer millimeters, rounded half-up.

## Database Schema (PostgreSQL)

### zones
| Field | Type | Constraint | Description |
|---|---|---|---|
| id | SMALLINT | PK | matches contract zoneId |
| name | TEXT | NOT NULL | "Yogyakarta" |
| lat | DOUBLE | NOT NULL | city center |
| lon | DOUBLE | NOT NULL | |
| threshold_mm | INT | NOT NULL | mirrors contract |
| payout_per_day | NUMERIC(38,0) | NOT NULL | wei |
| max_days_per_week | SMALLINT | NOT NULL | |
| active | BOOL | NOT NULL DEFAULT true | |

### premiums
| Field | Type | Constraint | Description |
|---|---|---|---|
| zone_id | SMALLINT | FK zones, part of PK | |
| month | SMALLINT | 1..12, part of PK | |
| expected_rain_days | NUMERIC(6,3) | NOT NULL | per week, from history |
| premium_per_week | NUMERIC(38,0) | NOT NULL | wei, after loading + rounding |
| narrative_id | TEXT | NOT NULL | Indonesian explanation from AI |
| computed_at | TIMESTAMPTZ | NOT NULL | |
| on_chain_tx | TEXT | NULL | tx hash of setPremium |

### rain_observations
| Field | Type | Constraint | Description |
|---|---|---|---|
| zone_id | SMALLINT | FK, part of PK | |
| day_index | INT | part of PK | |
| mm | INT | NOT NULL | |
| source | TEXT | NOT NULL | `open-meteo` or `simulated` |
| fetched_at | TIMESTAMPTZ | NOT NULL | |

### oracle_runs
| Field | Type | Constraint | Description |
|---|---|---|---|
| id | BIGSERIAL | PK | |
| zone_id | SMALLINT | NOT NULL | |
| day_index | INT | NOT NULL | |
| submit_tx | TEXT | NULL | |
| settle_tx | TEXT | NULL | |
| status | TEXT | NOT NULL | `pending`, `submitted`, `settled`, `no_rain`, `failed` |
| error | TEXT | NULL | |
| created_at | TIMESTAMPTZ | NOT NULL | |
| UNIQUE | | (zone_id, day_index) | idempotency |

### policies (indexed from chain)
| Field | Type | Constraint | Description |
|---|---|---|---|
| policy_id | BIGINT | PK | contract policyId |
| holder | TEXT | NOT NULL, indexed | lowercase hex |
| zone_id | SMALLINT | FK | |
| start_day | INT | NOT NULL | inclusive |
| end_day | INT | NOT NULL | exclusive |
| premium_paid | NUMERIC(38,0) | NOT NULL | |
| tx_hash | TEXT | NOT NULL | |
| block_number | BIGINT | NOT NULL | |

### payouts (indexed from chain)
| Field | Type | Constraint | Description |
|---|---|---|---|
| id | BIGSERIAL | PK | |
| policy_id | BIGINT | FK | |
| holder | TEXT | NOT NULL, indexed | |
| zone_id | SMALLINT | NOT NULL | |
| day_index | INT | NOT NULL | |
| mm | INT | NOT NULL | copied for display |
| amount | NUMERIC(38,0) | NOT NULL | |
| tx_hash | TEXT | NOT NULL | |
| explanation | TEXT | NULL | filled by AI explainer |
| explained_at | TIMESTAMPTZ | NULL | |
| UNIQUE | | (policy_id, day_index) | |

### faucet_claims
| Field | Type | Constraint | Description |
|---|---|---|---|
| address | TEXT | PK | |
| last_claim_at | TIMESTAMPTZ | NOT NULL | |

### indexer_state
| Field | Type | Constraint | Description |
|---|---|---|---|
| key | TEXT | PK | `last_block` |
| value | BIGINT | NOT NULL | |

### Relationships
zones 1..n premiums, zones 1..n rain_observations, zones 1..n policies, policies 1..n payouts. `oracle_runs` is an operational log keyed by (zone, day).

## API Design

### Base URL
`https://api.payung.app/v1` (local: `http://localhost:8080/v1`)

### Authentication
- Public read endpoints: none. Address is a path param; nothing sensitive is exposed.
- Admin endpoints: header `X-Admin-Key`. Disabled when `ENV=production`.
- Relay endpoint (fallback only): body carries an EIP-712 signature; no session.

### Endpoints

#### Zones

**GET /zones**
- Description: list active zones with current month premium and parameters
- Response:
```json
[{"id":1,"name":"Yogyakarta","thresholdMm":20,"payoutPerDay":"25000000000000000000000","maxDaysPerWeek":3,"premiumPerWeek":"6500000000000000000000","premiumNarrative":"September di Yogyakarta rata-rata 1,2 hari hujan lebat per minggu, jadi preminya Rp 6.500."}]
```

**GET /zones/:id/rain?days=7**
- Description: last N days of rainfall for a zone, with threshold
- Response: `{"thresholdMm":20,"days":[{"dayIndex":20713,"date":"2026-09-16","mm":31,"isRainDay":true,"source":"open-meteo"}]}`

#### Driver

**GET /drivers/:address/policy**
- Description: active policy for this address, or null
- Response: `{"policyId":12,"zoneId":1,"zoneName":"Yogyakarta","startDate":"2026-09-18","endDate":"2026-09-25","daysLeft":6,"payoutsThisWeek":1,"maxDaysPerWeek":3}`
- Error codes: 400 invalid address

**GET /drivers/:address/payouts**
- Description: payout history, newest first
- Response: `[{"date":"2026-09-16","zoneName":"Bantul","mm":31,"amount":"25000000000000000000000","txHash":"0x...","explanation":"Kemarin hujan 31 mm di Bantul..."}]`

**POST /faucet**
- Description: send 50,000 IDRP to address (testnet)
- Request: `{"address":"0x..."}`
- Response: `{"txHash":"0x..."}`
- Error codes: 429 cooldown active, 400 invalid address

#### Relay (fallback path, only if MegaFuel is not used)

**POST /relay/buy-policy**
- Request: `{"address":"0x...","zoneId":1,"weeks":1,"deadline":1758200000,"nonce":3,"signature":"0x..."}`
- Response: `{"txHash":"0x..."}`
- Error codes: 400 bad signature, 409 policy already active, 402 insufficient IDRP

#### Admin (non-production)

**POST /admin/simulate-rain**
- Request: `{"zoneId":2,"date":"2026-09-16","mm":31}`
- Description: writes rain_observations with `source=simulated`, then runs submit + settle for that (zone, day)
- Response: `{"submitTx":"0x...","settleTx":"0x...","payouts":3}`

**POST /admin/pricing/recompute**
- Description: refetch 3-year history, recompute premiums for all zones and months, generate narratives, push `setPremium` for the current month
- Response: summary per zone

**POST /admin/oracle/run?date=2026-09-16**
- Description: manual trigger of the daily job for a date

### Error Response Format
```json
{"error":{"code":"POLICY_ACTIVE","message":"Kamu sudah punya polis aktif di zona ini."}}
```
Codes: `INVALID_ADDRESS`, `POLICY_ACTIVE`, `INSUFFICIENT_BALANCE`, `FAUCET_COOLDOWN`, `BAD_SIGNATURE`, `ZONE_INACTIVE`, `INTERNAL`.

## Smart Contract Interface

### IDRP (mock IDR)
- ERC-20, name "Payung IDR (testnet)", symbol IDRP, 18 decimals
- `faucet(address to)`: mints 50,000 IDRP, only `FAUCET_ROLE` (backend)
- `mint(address,uint256)`: only `ADMIN`, for funding the pool

### PayungPool
- Address (opBNB testnet): `TBD after deploy, verified on opBNB BscScan`

Roles: `DEFAULT_ADMIN_ROLE`, `ORACLE_ROLE`.

Storage:
```solidity
struct Zone {
  string name;
  uint256 premiumPerWeek;   // IDRP wei
  uint16  thresholdMm;
  uint256 payoutPerDay;     // IDRP wei
  uint8   maxDaysPerWeek;
  bool    active;
}
struct Policy {
  address holder;
  uint16  zoneId;
  uint32  startDay;         // inclusive, day_index
  uint32  endDay;           // exclusive
  bool    closed;
}
mapping(uint16 => Zone) zones;
mapping(uint256 => Policy) policies;               // policyId => Policy
mapping(uint16 => uint256[]) activePolicyIds;      // zoneId => list, max 50
mapping(address => mapping(uint16 => uint256)) activePolicyOf; // holder => zone => policyId (0 = none)
mapping(uint16 => mapping(uint32 => uint16)) rainfallMm;       // zone => day => mm
mapping(uint16 => mapping(uint32 => bool)) rainReported;
mapping(uint256 => mapping(uint32 => bool)) paidForDay;        // policyId => day => paid
mapping(uint256 => mapping(uint32 => uint8)) payoutsInWeek;    // policyId => week => count
uint256 constant MAX_ACTIVE_PER_ZONE = 50;
```

Functions:
- `createZone(uint16 id, string name, uint256 premiumPerWeek, uint16 thresholdMm, uint256 payoutPerDay, uint8 maxDaysPerWeek)`: admin
- `setPremium(uint16 zoneId, uint256 premiumPerWeek)`: admin (called by pricing engine via admin key)
- `setZoneActive(uint16 zoneId, bool active)`: admin
- `fundPool(uint256 amount)`: anyone; `transferFrom` IDRP into contract; emits `PoolFunded`
- `buyPolicy(uint16 zoneId, uint8 weeks) returns (uint256 policyId)`: requires zone active, `1 <= weeks <= 4`, no active policy for (holder, zone), active count < 50. Premium = `premiumPerWeek * weeks`, pulled via `transferFrom`. `startDay = today() + 1`, `endDay = startDay + 7*weeks`. Emits `PolicyBought`.
- `submitRainfall(uint16 zoneId, uint32 dayIndex, uint16 mm)`: `ORACLE_ROLE`, requires `!rainReported[zone][day]` and `dayIndex < today()`. Emits `RainfallReported(zoneId, dayIndex, mm, isRainDay)`.
- `settle(uint16 zoneId, uint32 dayIndex)`: anyone, `nonReentrant`. Requires `rainReported` and `mm >= threshold`. Loops `activePolicyIds[zoneId]`; for each policy: skip if not active on day, skip if `paidForDay`, skip if `payoutsInWeek[pid][day/7] >= maxDaysPerWeek`. If pool balance < payoutPerDay emit `PayoutSkipped` and continue. Else mark paid, increment week count, `safeTransfer`, emit `PayoutSent`. Also prunes expired policies from the array (swap-and-pop) and emits `PolicyExpired`.
- `pruneExpired(uint16 zoneId)`: anyone; removes policies with `endDay <= today()` so slots free up without waiting for rain.
- `today() view returns (uint32)`: `uint32((block.timestamp + 7 hours) / 1 days)`
- `getActivePolicies(uint16 zoneId) view returns (uint256[])`
- `getPolicy(uint256 id) view returns (Policy)`
- `poolBalance() view returns (uint256)`

Events:
- `ZoneCreated(uint16 zoneId, string name)`
- `PremiumUpdated(uint16 zoneId, uint256 premiumPerWeek)`
- `PoolFunded(address from, uint256 amount)`
- `PolicyBought(uint256 policyId, address holder, uint16 zoneId, uint32 startDay, uint32 endDay, uint256 premium)`
- `RainfallReported(uint16 zoneId, uint32 dayIndex, uint16 mm, bool isRainDay)`
- `PayoutSent(uint256 policyId, address holder, uint16 zoneId, uint32 dayIndex, uint256 amount)`
- `PayoutSkipped(uint256 policyId, uint32 dayIndex, string reason)`
- `PolicyExpired(uint256 policyId)`

Errors (custom): `ZoneInactive()`, `PolicyAlreadyActive()`, `ZoneFull()`, `InvalidWeeks()`, `RainAlreadyReported()`, `DayNotFinished()`, `NotRainDay()`, `RainNotReported()`.

Foundry test cases (must pass before day 2 ends):
1. buy policy, premium moves to pool, event emitted
2. buy second policy same zone reverts `PolicyAlreadyActive`
3. rain 15 mm (below 20): settle reverts `NotRainDay`
4. rain 25 mm: 3 holders paid, pool balance decreases by 3 x payout
5. same day settle twice: second call pays nobody
6. 4 rain days in one week: 4th day pays nobody (cap 3)
7. policy bought today is not paid for today (starts tomorrow)
8. expired policy not paid and is pruned
9. pool underfunded: `PayoutSkipped` emitted, no revert
10. non-oracle calling `submitRainfall` reverts
