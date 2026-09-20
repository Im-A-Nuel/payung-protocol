-- name: GetLastClaim :one
SELECT address, last_claim_at FROM faucet_claims WHERE address = $1;

-- name: UpsertClaim :exec
INSERT INTO faucet_claims (address, last_claim_at)
VALUES ($1, now())
ON CONFLICT (address) DO UPDATE SET last_claim_at = EXCLUDED.last_claim_at;
