-- name: UpsertPolicy :exec
INSERT INTO policies (policy_id, holder, zone_id, start_day, end_day, premium_paid, tx_hash, block_number)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (policy_id) DO NOTHING;

-- name: GetActivePolicyForHolder :one
SELECT po.policy_id, po.holder, po.zone_id, po.start_day, po.end_day, z.name AS zone_name, z.max_days_per_week
FROM policies po
JOIN zones z ON z.id = po.zone_id
WHERE po.holder = $1 AND po.end_day > $2
ORDER BY po.start_day DESC
LIMIT 1;

-- name: CountPayoutsForPolicyInWeek :one
SELECT COUNT(*)::bigint
FROM payouts
WHERE policy_id = $1 AND (day_index / 7) = $2;
