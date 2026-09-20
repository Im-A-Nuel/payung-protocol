-- name: GetOracleRun :one
SELECT id, zone_id, day_index, submit_tx, settle_tx, status, error, created_at
FROM oracle_runs
WHERE zone_id = $1 AND day_index = $2;

-- name: UpsertOracleRun :one
INSERT INTO oracle_runs (zone_id, day_index, submit_tx, settle_tx, status, error, created_at)
VALUES ($1, $2, $3, $4, $5, $6, now())
ON CONFLICT (zone_id, day_index) DO UPDATE SET
    submit_tx = COALESCE(EXCLUDED.submit_tx, oracle_runs.submit_tx),
    settle_tx = COALESCE(EXCLUDED.settle_tx, oracle_runs.settle_tx),
    status    = EXCLUDED.status,
    error     = EXCLUDED.error
RETURNING id, zone_id, day_index, submit_tx, settle_tx, status, error, created_at;
