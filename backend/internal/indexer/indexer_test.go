package indexer

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// testQueries returns Queries backed by the local Postgres, and clears
// indexer_state.last_block first since it's a single shared row (not
// scoped per test the way oracle_runs rows are by day_index).
func testQueries(t *testing.T) *store.Queries {
	t.Helper()
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("TEST_DATABASE_URL not set, skipping integration test")
	}
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(pool.Close)

	if _, err := pool.Exec(context.Background(), `DELETE FROM indexer_state WHERE key = 'last_block'`); err != nil {
		t.Fatalf("clear indexer_state: %v", err)
	}

	return store.New(pool)
}

func TestLastBlock_FallsBackToStartBlockWhenUnset(t *testing.T) {
	q := testQueries(t)
	svc := New(q, nil, 12345)

	got, err := svc.lastBlock(context.Background())
	if err != nil {
		t.Fatalf("lastBlock: %v", err)
	}
	if got != 12345 {
		t.Errorf("lastBlock = %d, want fallback 12345", got)
	}
}

func TestLastBlock_ReadsPersistedValue(t *testing.T) {
	q := testQueries(t)
	ctx := context.Background()

	if err := q.SetLastBlock(ctx, 999); err != nil {
		t.Fatalf("SetLastBlock: %v", err)
	}

	svc := New(q, nil, 1)
	got, err := svc.lastBlock(ctx)
	if err != nil {
		t.Fatalf("lastBlock: %v", err)
	}
	if got != 999 {
		t.Errorf("lastBlock = %d, want 999", got)
	}
}
