-- name: GetPremiumForZoneMonth :one
SELECT zone_id, month, expected_rain_days, premium_per_week, narrative_id, computed_at, on_chain_tx
FROM premiums
WHERE zone_id = $1 AND month = $2;
