-- name: GetLastBlock :one
SELECT value FROM indexer_state WHERE key = 'last_block';

-- name: SetLastBlock :exec
INSERT INTO indexer_state (key, value)
VALUES ('last_block', $1)
ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value;
