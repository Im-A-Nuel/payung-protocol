package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// syncIndexer catches up the indexer right after an admin-triggered oracle
// action, so the payout the action just caused shows up in the dashboard
// immediately instead of waiting for the next scheduled sync.
func (s *Server) syncIndexer(r *http.Request) {
	ctx := r.Context()
	latest, err := s.OracleChain.LatestBlock(ctx)
	if err != nil {
		return
	}
	_ = s.Indexer.Sync(ctx, latest)
}

type simulateRainRequest struct {
	ZoneID int16  `json:"zoneId"`
	Date   string `json:"date"`
	Mm     int    `json:"mm"`
}

type simulateRainResponse struct {
	SubmitTx string `json:"submitTx"`
	SettleTx string `json:"settleTx"`
	Payouts  int64  `json:"payouts"`
}

func (s *Server) handleSimulateRain(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req simulateRainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "INTERNAL", "Body permintaan tidak valid.")
		return
	}

	zone, err := s.Queries.GetZone(ctx, req.ZoneID)
	if err != nil || !zone.Active {
		writeError(w, http.StatusBadRequest, "ZONE_INACTIVE", "Zona tidak aktif atau tidak ditemukan.")
		return
	}

	dayIndex, err := chain.DayIndexFromDateWIB(req.Date)
	if err != nil {
		writeError(w, http.StatusBadRequest, "INTERNAL", "Format tanggal harus YYYY-MM-DD.")
		return
	}

	result := s.Oracle.ProcessZoneDay(ctx, req.ZoneID, dayIndex, req.Mm, "simulated")
	if result.Err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", result.Err.Error())
		return
	}

	s.syncIndexer(r)

	payouts, err := s.Queries.CountPayoutsForZoneDay(ctx, store.CountPayoutsForZoneDayParams{
		ZoneID: req.ZoneID, DayIndex: int32(dayIndex),
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal menghitung payout.")
		return
	}

	writeJSON(w, http.StatusOK, simulateRainResponse{
		SubmitTx: result.SubmitTx,
		SettleTx: result.SettleTx,
		Payouts:  payouts,
	})
}

type oracleRunZoneResult struct {
	ZoneID   int16  `json:"zoneId"`
	Status   string `json:"status"`
	SubmitTx string `json:"submitTx,omitempty"`
	SettleTx string `json:"settleTx,omitempty"`
	Error    string `json:"error,omitempty"`
}

func (s *Server) handleOracleRun(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dayIndex := chain.DayIndex(time.Now().Unix()) - 1 // default: yesterday
	if v := r.URL.Query().Get("date"); v != "" {
		dari, err := chain.DayIndexFromDateWIB(v)
		if err != nil {
			writeError(w, http.StatusBadRequest, "INTERNAL", "Format tanggal harus YYYY-MM-DD.")
			return
		}
		dayIndex = dari
	}

	results, err := s.Oracle.RunDaily(ctx, dayIndex)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", err.Error())
		return
	}

	s.syncIndexer(r)

	out := make([]oracleRunZoneResult, len(results))
	for i, res := range results {
		out[i] = oracleRunZoneResult{ZoneID: res.ZoneID, Status: res.Status, SubmitTx: res.SubmitTx, SettleTx: res.SettleTx}
		if res.Err != nil {
			out[i].Error = res.Err.Error()
		}
	}

	writeJSON(w, http.StatusOK, out)
}
