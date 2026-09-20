-- name: ListActiveZones :many
SELECT id, name, lat, lon, threshold_mm, payout_per_day, max_days_per_week, active
FROM zones
WHERE active = TRUE
ORDER BY id;

-- name: GetZone :one
SELECT id, name, lat, lon, threshold_mm, payout_per_day, max_days_per_week, active
FROM zones
WHERE id = $1;
