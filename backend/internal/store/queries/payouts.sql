-- name: UpsertPayout :exec
INSERT INTO payouts (policy_id, holder, zone_id, day_index, mm, amount, tx_hash)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (policy_id, day_index) DO NOTHING;

-- name: CountPayoutsForZoneDay :one
SELECT COUNT(*)::bigint
FROM payouts
WHERE zone_id = $1 AND day_index = $2;

-- name: ListPayoutsForHolder :many
SELECT p.id, p.policy_id, p.holder, p.zone_id, p.day_index, p.mm, p.amount, p.tx_hash, p.explanation, p.explained_at,
       z.name AS zone_name
FROM payouts p
JOIN zones z ON z.id = p.zone_id
WHERE p.holder = $1
ORDER BY p.day_index DESC;
