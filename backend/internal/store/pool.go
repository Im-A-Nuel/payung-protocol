package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectPool opens a pgx connection pool and verifies it with a ping, so
// callers fail fast at startup rather than on the first query.
func ConnectPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("store.ConnectPool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("store.ConnectPool: ping: %w", err)
	}
	return pool, nil
}
