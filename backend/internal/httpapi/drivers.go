package httpapi

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// parseAddress validates and lowercases a wallet address, matching the
// "addresses stored lowercase hex in Postgres" convention.
func parseAddress(raw string) (string, bool) {
	if !common.IsHexAddress(raw) {
		return "", false
	}
	return strings.ToLower(raw), true
}

type policyResponse struct {
	PolicyID        int64  `json:"policyId"`
	ZoneID          int16  `json:"zoneId"`
	ZoneName        string `json:"zoneName"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	DaysLeft        int32  `json:"daysLeft"`
	PayoutsThisWeek int64  `json:"payoutsThisWeek"`
	MaxDaysPerWeek  int16  `json:"maxDaysPerWeek"`
}

func (s *Server) handleDriverPolicy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	address, ok := parseAddress(chi.URLParam(r, "address"))
	if !ok {
		writeError(w, http.StatusBadRequest, "INVALID_ADDRESS", "Alamat wallet tidak valid.")
		return
	}

	today := chain.DayIndex(time.Now().Unix())

	row, err := s.Queries.GetActivePolicyForHolder(ctx, store.GetActivePolicyForHolderParams{
		Holder: address, EndDay: int32(today),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeJSON(w, http.StatusOK, nil)
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat polis.")
		return
	}

	weekIndex := chain.WeekIndex(today)
	payoutsThisWeek, err := s.Queries.CountPayoutsForPolicyInWeek(ctx, store.CountPayoutsForPolicyInWeekParams{
		PolicyID: row.PolicyID, DayIndex: int32(weekIndex),
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat kuota payout.")
		return
	}

	writeJSON(w, http.StatusOK, policyResponse{
		PolicyID:        row.PolicyID,
		ZoneID:          row.ZoneID,
		ZoneName:        row.ZoneName,
		StartDate:       chain.DateStringWIB(uint32(row.StartDay)),
		EndDate:         chain.DateStringWIB(uint32(row.EndDay)),
		DaysLeft:        row.EndDay - int32(today),
		PayoutsThisWeek: payoutsThisWeek,
		MaxDaysPerWeek:  row.MaxDaysPerWeek,
	})
}

type payoutResponse struct {
	Date        string  `json:"date"`
	ZoneName    string  `json:"zoneName"`
	Mm          int32   `json:"mm"`
	Amount      string  `json:"amount"`
	TxHash      string  `json:"txHash"`
	Explanation *string `json:"explanation"`
}

func (s *Server) handleDriverPayouts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	address, ok := parseAddress(chi.URLParam(r, "address"))
	if !ok {
		writeError(w, http.StatusBadRequest, "INVALID_ADDRESS", "Alamat wallet tidak valid.")
		return
	}

	rows, err := s.Queries.ListPayoutsForHolder(ctx, address)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat riwayat payout.")
		return
	}

	out := make([]payoutResponse, 0, len(rows))
	for _, row := range rows {
		var explanation *string
		if row.Explanation.Valid {
			explanation = &row.Explanation.String
		}
		out = append(out, payoutResponse{
			Date:        chain.DateStringWIB(uint32(row.DayIndex)),
			ZoneName:    row.ZoneName,
			Mm:          row.Mm,
			Amount:      store.WeiString(row.Amount),
			TxHash:      row.TxHash,
			Explanation: explanation,
		})
	}

	writeJSON(w, http.StatusOK, out)
}
