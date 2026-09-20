-- Initial schema for Payung, per docs/SCHEMA.md.

CREATE TABLE IF NOT EXISTS zones (
    id                SMALLINT PRIMARY KEY,
    name              TEXT NOT NULL,
    lat               DOUBLE PRECISION NOT NULL,
    lon               DOUBLE PRECISION NOT NULL,
    threshold_mm      INT NOT NULL,
    payout_per_day    NUMERIC(38, 0) NOT NULL,
    max_days_per_week SMALLINT NOT NULL,
    active            BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS premiums (
    zone_id             SMALLINT NOT NULL REFERENCES zones (id),
    month               SMALLINT NOT NULL CHECK (month BETWEEN 1 AND 12),
    expected_rain_days  NUMERIC(6, 3) NOT NULL,
    premium_per_week    NUMERIC(38, 0) NOT NULL,
    narrative_id        TEXT NOT NULL,
    computed_at         TIMESTAMPTZ NOT NULL,
    on_chain_tx         TEXT,
    PRIMARY KEY (zone_id, month)
);

CREATE TABLE IF NOT EXISTS rain_observations (
    zone_id    SMALLINT NOT NULL REFERENCES zones (id),
    day_index  INT NOT NULL,
    mm         INT NOT NULL,
    source     TEXT NOT NULL,
    fetched_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (zone_id, day_index)
);

CREATE TABLE IF NOT EXISTS oracle_runs (
    id         BIGSERIAL PRIMARY KEY,
    zone_id    SMALLINT NOT NULL,
    day_index  INT NOT NULL,
    submit_tx  TEXT,
    settle_tx  TEXT,
    status     TEXT NOT NULL,
    error      TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    UNIQUE (zone_id, day_index)
);

CREATE TABLE IF NOT EXISTS policies (
    policy_id     BIGINT PRIMARY KEY,
    holder        TEXT NOT NULL,
    zone_id       SMALLINT NOT NULL REFERENCES zones (id),
    start_day     INT NOT NULL,
    end_day       INT NOT NULL,
    premium_paid  NUMERIC(38, 0) NOT NULL,
    tx_hash       TEXT NOT NULL,
    block_number  BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_policies_holder ON policies (holder);

CREATE TABLE IF NOT EXISTS payouts (
    id            BIGSERIAL PRIMARY KEY,
    policy_id     BIGINT NOT NULL REFERENCES policies (policy_id),
    holder        TEXT NOT NULL,
    zone_id       SMALLINT NOT NULL,
    day_index     INT NOT NULL,
    mm            INT NOT NULL,
    amount        NUMERIC(38, 0) NOT NULL,
    tx_hash       TEXT NOT NULL,
    explanation   TEXT,
    explained_at  TIMESTAMPTZ,
    UNIQUE (policy_id, day_index)
);

CREATE INDEX IF NOT EXISTS idx_payouts_holder ON payouts (holder);

CREATE TABLE IF NOT EXISTS faucet_claims (
    address        TEXT PRIMARY KEY,
    last_claim_at  TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS indexer_state (
    key    TEXT PRIMARY KEY,
    value  BIGINT NOT NULL
);
