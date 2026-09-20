-- name: UpsertRainObservation :exec
INSERT INTO rain_observations (zone_id, day_index, mm, source, fetched_at)
VALUES ($1, $2, $3, $4, now())
ON CONFLICT (zone_id, day_index) DO UPDATE SET
    mm         = EXCLUDED.mm,
    source     = EXCLUDED.source,
    fetched_at = EXCLUDED.fetched_at;

-- name: GetRainObservation :one
SELECT zone_id, day_index, mm, source, fetched_at
FROM rain_observations
WHERE zone_id = $1 AND day_index = $2;

-- name: ListRecentRain :many
SELECT zone_id, day_index, mm, source, fetched_at
FROM rain_observations
WHERE zone_id = $1
ORDER BY day_index DESC
LIMIT $2;
