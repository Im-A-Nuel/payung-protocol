package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// testPool connects to the local Postgres used by every backend integration
// test. Fixture rows use IDs far outside the real ranges, and every test
// removes its own afterwards: the same database backs local demo runs, and a
// leftover row dated in the year 4461 sorts ahead of real data in any
// "last N days" query.
func testPool(t *testing.T) *pgxpool.Pool {
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
	return pool
}

// bersihkan removes fixture rows once the test is done.
func bersihkan(t *testing.T, pool *pgxpool.Pool, perintah ...string) {
	t.Helper()
	t.Cleanup(func() {
		for _, sql := range perintah {
			if _, err := pool.Exec(context.Background(), sql); err != nil {
				t.Logf("cleanup failed for %q: %v", sql, err)
			}
		}
	})
}

// Bantul (zone 3) is seeded by migrations/0002_seed_zones.sql with
// threshold_mm=20.
const testZoneID = int16(3)

func TestHandleListZones_ReturnsSeededZones(t *testing.T) {
	pool := testPool(t)
	srv := &Server{Queries: store.New(pool)}

	req := httptest.NewRequest(http.MethodGet, "/v1/zones", nil)
	w := httptest.NewRecorder()
	srv.handleListZones(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var zones []zoneResponse
	if err := json.Unmarshal(w.Body.Bytes(), &zones); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(zones) < 5 {
		t.Fatalf("len(zones) = %d, want at least 5 seeded zones", len(zones))
	}
	for _, z := range zones {
		if z.PremiumNarrative == "" {
			t.Errorf("zone %d has empty narrative, want fallback text", z.ID)
		}
	}
}

func TestHandleZoneRain_ReturnsChronologicalOrder(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	srv := &Server{Queries: store.New(pool)}
	bersihkan(t, pool, "DELETE FROM rain_observations WHERE day_index >= 910000")

	const baseDay = 910000
	for i, mm := range []int32{5, 25, 0} {
		day := int32(baseDay + i)
		_, err := pool.Exec(ctx,
			`INSERT INTO rain_observations (zone_id, day_index, mm, source, fetched_at) VALUES ($1,$2,$3,'simulated', now())
			 ON CONFLICT (zone_id, day_index) DO UPDATE SET mm = EXCLUDED.mm`,
			testZoneID, day, mm)
		if err != nil {
			t.Fatalf("seed rain: %v", err)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/zones/3/rain?days=3", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "3")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	srv.handleZoneRain(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var resp zoneRainResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(resp.Days) != 3 {
		t.Fatalf("len(days) = %d, want 3", len(resp.Days))
	}
	if resp.Days[0].DayIndex != baseDay || resp.Days[2].DayIndex != baseDay+2 {
		t.Errorf("days not in chronological order: %+v", resp.Days)
	}
	if !resp.Days[1].IsRainDay {
		t.Errorf("day with mm=25 (threshold 20) should be a rain day: %+v", resp.Days[1])
	}
	if resp.Days[0].IsRainDay {
		t.Errorf("day with mm=5 should not be a rain day: %+v", resp.Days[0])
	}
}

func TestHandleDriverPolicy_NoPolicy_ReturnsNull(t *testing.T) {
	pool := testPool(t)
	srv := &Server{Queries: store.New(pool)}

	req := httptest.NewRequest(http.MethodGet, "/v1/drivers/0x00000000000000000000000000000000000000ff/policy", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("address", "0x00000000000000000000000000000000000000ff")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	srv.handleDriverPolicy(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if strings.TrimSpace(w.Body.String()) != "null" {
		t.Errorf("body = %q, want null", w.Body.String())
	}
}

func TestHandleDriverPolicy_ActivePolicy_ReturnsDetails(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	srv := &Server{Queries: store.New(pool)}

	const holder = "0x00000000000000000000000000000000000000ab"
	const policyID = 910001
	bersihkan(t, pool, "DELETE FROM payouts WHERE policy_id = 910001", "DELETE FROM policies WHERE policy_id = 910001")
	today := chain.DayIndex(time.Now().Unix())

	_, err := pool.Exec(ctx, `
		INSERT INTO policies (policy_id, holder, zone_id, start_day, end_day, premium_paid, tx_hash, block_number)
		VALUES ($1, $2, $3, $4, $5, 1000000000000000000000, '0xabc', 1)
		ON CONFLICT (policy_id) DO UPDATE SET
			holder = EXCLUDED.holder, zone_id = EXCLUDED.zone_id,
			start_day = EXCLUDED.start_day, end_day = EXCLUDED.end_day
	`, policyID, holder, testZoneID, today-1, today+6)
	if err != nil {
		t.Fatalf("seed policy: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/drivers/"+holder+"/policy", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("address", holder)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	srv.handleDriverPolicy(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var resp policyResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v, body=%s", err, w.Body.String())
	}
	if resp.PolicyID != policyID {
		t.Errorf("policyId = %d, want %d", resp.PolicyID, policyID)
	}
	if resp.ZoneName != "Bantul" {
		t.Errorf("zoneName = %q, want Bantul", resp.ZoneName)
	}
	if resp.DaysLeft != 6 {
		t.Errorf("daysLeft = %d, want 6", resp.DaysLeft)
	}
}

func TestHandleDriverPayouts_ReturnsHistory(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	srv := &Server{Queries: store.New(pool)}

	const holder = "0x00000000000000000000000000000000000000cd"
	const policyID = 910002
	bersihkan(t, pool, "DELETE FROM payouts WHERE policy_id = 910002", "DELETE FROM policies WHERE policy_id = 910002")
	today := chain.DayIndex(time.Now().Unix())

	_, err := pool.Exec(ctx, `
		INSERT INTO policies (policy_id, holder, zone_id, start_day, end_day, premium_paid, tx_hash, block_number)
		VALUES ($1, $2, $3, $4, $5, 1000000000000000000000, '0xabc', 1)
		ON CONFLICT (policy_id) DO UPDATE SET
			holder = EXCLUDED.holder, zone_id = EXCLUDED.zone_id,
			start_day = EXCLUDED.start_day, end_day = EXCLUDED.end_day
	`, policyID, holder, testZoneID, today-10, today-3)
	if err != nil {
		t.Fatalf("seed policy: %v", err)
	}
	_, err = pool.Exec(ctx, `
		INSERT INTO payouts (policy_id, holder, zone_id, day_index, mm, amount, tx_hash)
		VALUES ($1, $2, $3, $4, 31, 25000000000000000000000, '0xdef')
		ON CONFLICT (policy_id, day_index) DO UPDATE SET holder = EXCLUDED.holder, amount = EXCLUDED.amount
	`, policyID, holder, testZoneID, today-5)
	if err != nil {
		t.Fatalf("seed payout: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/drivers/"+holder+"/payouts", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("address", holder)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	srv.handleDriverPayouts(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var resp []payoutResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(resp) != 1 {
		t.Fatalf("len(payouts) = %d, want 1", len(resp))
	}
	if resp[0].Amount != "25000000000000000000000" {
		t.Errorf("amount = %q, want 25000000000000000000000", resp[0].Amount)
	}
	if resp[0].ZoneName != "Bantul" {
		t.Errorf("zoneName = %q, want Bantul", resp[0].ZoneName)
	}
	if resp[0].Explanation != nil {
		t.Errorf("explanation = %v, want nil before Phase 4", resp[0].Explanation)
	}
}

func TestHandleFaucet_InvalidAddress_Returns400(t *testing.T) {
	pool := testPool(t)
	srv := &Server{Queries: store.New(pool)}

	req := httptest.NewRequest(http.MethodPost, "/v1/faucet", strings.NewReader(`{"address":"not-an-address"}`))
	w := httptest.NewRecorder()
	srv.handleFaucet(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400, body = %s", w.Code, w.Body.String())
	}
}

func TestHandleFaucet_Cooldown_Returns429(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	srv := &Server{Queries: store.New(pool), FaucetCooldownHours: 24}

	const addr = "0x00000000000000000000000000000000000000ee"
	bersihkan(t, pool, "DELETE FROM faucet_claims WHERE address = '0x00000000000000000000000000000000000000ee'")
	_, err := pool.Exec(ctx, `
		INSERT INTO faucet_claims (address, last_claim_at) VALUES ($1, now())
		ON CONFLICT (address) DO UPDATE SET last_claim_at = now()
	`, addr)
	if err != nil {
		t.Fatalf("seed claim: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/faucet", strings.NewReader(`{"address":"`+addr+`"}`))
	w := httptest.NewRecorder()
	srv.handleFaucet(w, req)

	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("status = %d, want 429, body = %s", w.Code, w.Body.String())
	}
}
