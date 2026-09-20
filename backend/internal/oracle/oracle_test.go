package oracle

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// fakeChain lets the oracle service be tested without a real RPC endpoint.
type fakeChain struct {
	reported  map[string]bool
	submitErr error
	settleErr error
	submitN   int
	settleN   int
}

func newFakeChain() *fakeChain {
	return &fakeChain{reported: make(map[string]bool)}
}

func chainKey(zoneID uint16, dayIndex uint32) string {
	return fmt.Sprintf("%d:%d", zoneID, dayIndex)
}

func (f *fakeChain) RainReported(_ context.Context, zoneID uint16, dayIndex uint32) (bool, error) {
	return f.reported[chainKey(zoneID, dayIndex)], nil
}

func (f *fakeChain) SubmitRainfall(_ context.Context, zoneID uint16, dayIndex uint32, _ uint16) (string, error) {
	if f.submitErr != nil {
		return "", f.submitErr
	}
	f.reported[chainKey(zoneID, dayIndex)] = true
	f.submitN++
	return fmt.Sprintf("0xsubmit%d", f.submitN), nil
}

func (f *fakeChain) Settle(_ context.Context, _ uint16, _ uint32) (string, error) {
	if f.settleErr != nil {
		return "", f.settleErr
	}
	f.settleN++
	return fmt.Sprintf("0xsettle%d", f.settleN), nil
}

func testQueries(t *testing.T) *store.Queries {
	t.Helper()
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("TEST_DATABASE_URL not set, skipping integration test")
	}
	pool, err := store.ConnectPool(context.Background(), dbURL)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(pool.Close)
	return store.New(pool)
}

// Bantul (zone 3) is seeded with threshold_mm=20 by migrations/0002_seed_zones.sql.
const testZoneID = int16(3)

func TestProcessZoneDay_BelowThreshold_MarksNoRainAndSkipsSettle(t *testing.T) {
	q := testQueries(t)
	fc := newFakeChain()
	svc := New(q, fc, nil)
	ctx := context.Background()

	const dayIndex = uint32(900001) // far future, won't collide with other tests

	res := svc.ProcessZoneDay(ctx, testZoneID, dayIndex, 15, "simulated")
	if res.Err != nil {
		t.Fatalf("ProcessZoneDay: %v", res.Err)
	}
	if res.Status != StatusNoRain {
		t.Errorf("status = %s, want %s", res.Status, StatusNoRain)
	}
	if res.SubmitTx == "" {
		t.Error("expected a submit tx hash")
	}
	if res.SettleTx != "" {
		t.Errorf("expected no settle tx, got %s", res.SettleTx)
	}
	if fc.settleN != 0 {
		t.Errorf("settle should not have been called, settleN=%d", fc.settleN)
	}

	// Idempotent: re-running must not call the chain again.
	res2 := svc.ProcessZoneDay(ctx, testZoneID, dayIndex, 15, "simulated")
	if res2.Status != StatusNoRain {
		t.Errorf("second run status = %s, want %s", res2.Status, StatusNoRain)
	}
	if fc.submitN != 1 {
		t.Errorf("submit should have been called exactly once total, got %d", fc.submitN)
	}
}

func TestProcessZoneDay_AtOrAboveThreshold_Settles(t *testing.T) {
	q := testQueries(t)
	fc := newFakeChain()
	svc := New(q, fc, nil)
	ctx := context.Background()

	const dayIndex = uint32(900002)

	res := svc.ProcessZoneDay(ctx, testZoneID, dayIndex, 25, "simulated")
	if res.Err != nil {
		t.Fatalf("ProcessZoneDay: %v", res.Err)
	}
	if res.Status != StatusSettled {
		t.Errorf("status = %s, want %s", res.Status, StatusSettled)
	}
	if res.SubmitTx == "" || res.SettleTx == "" {
		t.Errorf("expected both tx hashes set, got submit=%q settle=%q", res.SubmitTx, res.SettleTx)
	}

	// Idempotent: re-running a settled day must not call the chain again.
	_ = svc.ProcessZoneDay(ctx, testZoneID, dayIndex, 25, "simulated")
	if fc.submitN != 1 || fc.settleN != 1 {
		t.Errorf("chain should be called exactly once each, got submit=%d settle=%d", fc.submitN, fc.settleN)
	}
}

func TestProcessZoneDay_AlreadyReportedOnChain_SkipsSubmit(t *testing.T) {
	q := testQueries(t)
	fc := newFakeChain()
	svc := New(q, fc, nil)
	ctx := context.Background()

	const dayIndex = uint32(900003)
	fc.reported[chainKey(uint16(testZoneID), dayIndex)] = true // pretend another process already submitted

	res := svc.ProcessZoneDay(ctx, testZoneID, dayIndex, 25, "simulated")
	if res.Err != nil {
		t.Fatalf("ProcessZoneDay: %v", res.Err)
	}
	if res.SubmitTx != "" {
		t.Errorf("expected no submit tx since already reported, got %q", res.SubmitTx)
	}
	if fc.submitN != 0 {
		t.Errorf("submitRainfall should not have been called, submitN=%d", fc.submitN)
	}
	if res.Status != StatusSettled || res.SettleTx == "" {
		t.Errorf("expected settled with a settle tx, got status=%s settleTx=%q", res.Status, res.SettleTx)
	}
}
