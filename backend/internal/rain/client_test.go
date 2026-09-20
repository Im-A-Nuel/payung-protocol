package rain

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForecastDaily_ParsesAndRounds(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("timezone") != "Asia/Jakarta" {
			t.Errorf("timezone = %q, want Asia/Jakarta", q.Get("timezone"))
		}
		if q.Get("past_days") != "2" {
			t.Errorf("past_days = %q, want 2", q.Get("past_days"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"daily": {
				"time": ["2026-09-14", "2026-09-15", "2026-09-16"],
				"precipitation_sum": [0, 30.5, 19.4]
			}
		}`))
	}))
	defer srv.Close()

	c := NewClient()
	c.ForecastURL = srv.URL

	points, err := c.ForecastDaily(context.Background(), -7.7956, 110.3695)
	if err != nil {
		t.Fatalf("ForecastDaily: %v", err)
	}
	if len(points) != 3 {
		t.Fatalf("len(points) = %d, want 3", len(points))
	}
	if points[1].Mm != 31 { // 30.5 rounds half-up to 31
		t.Errorf("points[1].Mm = %d, want 31", points[1].Mm)
	}
	if points[2].Mm != 19 { // 19.4 rounds down to 19
		t.Errorf("points[2].Mm = %d, want 19", points[2].Mm)
	}

	p, ok := PointForDate(points, "2026-09-15")
	if !ok || p.Mm != 31 {
		t.Errorf("PointForDate(2026-09-15) = %+v, ok=%v, want Mm=31, ok=true", p, ok)
	}

	_, ok = PointForDate(points, "2099-01-01")
	if ok {
		t.Error("PointForDate for missing date should return ok=false")
	}
}

func TestArchiveDaily_UsesDateRange(t *testing.T) {
	var gotStart, gotEnd string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotStart = r.URL.Query().Get("start_date")
		gotEnd = r.URL.Query().Get("end_date")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"daily": {"time": [], "precipitation_sum": []}}`))
	}))
	defer srv.Close()

	c := NewClient()
	c.ArchiveURL = srv.URL

	_, err := c.ArchiveDaily(context.Background(), -7.7956, 110.3695, "2023-09-17", "2026-09-17")
	if err != nil {
		t.Fatalf("ArchiveDaily: %v", err)
	}
	if gotStart != "2023-09-17" || gotEnd != "2026-09-17" {
		t.Errorf("start=%q end=%q, want 2023-09-17 / 2026-09-17", gotStart, gotEnd)
	}
}

func TestRoundHalfUp(t *testing.T) {
	cases := map[float64]int{
		0:    0,
		0.4:  0,
		0.5:  1,
		19.4: 19,
		19.5: 20,
		30.5: 31,
		-5:   0,
	}
	for in, want := range cases {
		if got := roundHalfUp(in); got != want {
			t.Errorf("roundHalfUp(%v) = %d, want %d", in, got, want)
		}
	}
}
