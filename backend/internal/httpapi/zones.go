package httpapi

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

type zoneResponse struct {
	ID               int16  `json:"id"`
	Name             string `json:"name"`
	ThresholdMm      int32  `json:"thresholdMm"`
	PayoutPerDay     string `json:"payoutPerDay"`
	MaxDaysPerWeek   int16  `json:"maxDaysPerWeek"`
	PremiumPerWeek   string `json:"premiumPerWeek"`
	PremiumNarrative string `json:"premiumNarrative"`
}

// fallbackNarrative is shown before the AI pricing engine (Phase 4) has
// computed a premium for the current month, so the frontend never shows a
// broken or empty explanation.
const fallbackNarrative = "Premi bulan ini belum dihitung."

func (s *Server) handleListZones(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	zones, err := s.Queries.ListActiveZones(ctx)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat daftar zona.")
		return
	}

	month := int16(time.Now().Month())
	out := make([]zoneResponse, 0, len(zones))
	for _, z := range zones {
		resp := zoneResponse{
			ID:               z.ID,
			Name:             z.Name,
			ThresholdMm:      z.ThresholdMm,
			PayoutPerDay:     store.WeiString(z.PayoutPerDay),
			MaxDaysPerWeek:   z.MaxDaysPerWeek,
			PremiumPerWeek:   "0",
			PremiumNarrative: fallbackNarrative,
		}

		premium, err := s.Queries.GetPremiumForZoneMonth(ctx, store.GetPremiumForZoneMonthParams{ZoneID: z.ID, Month: month})
		switch {
		case err == nil:
			resp.PremiumPerWeek = store.WeiString(premium.PremiumPerWeek)
			resp.PremiumNarrative = premium.NarrativeID
		case errors.Is(err, pgx.ErrNoRows):
			// Not priced yet; keep the fallback above.
		default:
			writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat premi zona.")
			return
		}

		out = append(out, resp)
	}

	writeJSON(w, http.StatusOK, out)
}

type rainDayResponse struct {
	DayIndex  int32  `json:"dayIndex"`
	Date      string `json:"date"`
	Mm        int32  `json:"mm"`
	IsRainDay bool   `json:"isRainDay"`
	Source    string `json:"source"`
}

type zoneRainResponse struct {
	ThresholdMm int32             `json:"thresholdMm"`
	Days        []rainDayResponse `json:"days"`
}

func (s *Server) handleZoneRain(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	zoneID, err := parseZoneID(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "ZONE_INACTIVE", "Zona tidak ditemukan.")
		return
	}

	zone, err := s.Queries.GetZone(ctx, zoneID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "ZONE_INACTIVE", "Zona tidak ditemukan.")
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat zona.")
		return
	}

	days := int32(7)
	if v := r.URL.Query().Get("days"); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 32); err == nil && parsed > 0 {
			days = int32(parsed)
		}
	}

	rows, err := s.Queries.ListRecentRain(ctx, store.ListRecentRainParams{ZoneID: zoneID, Limit: days})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memuat data hujan.")
		return
	}

	// rows are newest-first (see ListRecentRain); reverse to chronological
	// order, which is what a 7-day chart on the frontend wants.
	out := make([]rainDayResponse, len(rows))
	for i, row := range rows {
		out[len(rows)-1-i] = rainDayResponse{
			DayIndex:  row.DayIndex,
			Date:      chain.DateStringWIB(uint32(row.DayIndex)),
			Mm:        row.Mm,
			IsRainDay: row.Mm >= zone.ThresholdMm,
			Source:    row.Source,
		}
	}

	writeJSON(w, http.StatusOK, zoneRainResponse{ThresholdMm: zone.ThresholdMm, Days: out})
}

func parseZoneID(s string) (int16, error) {
	v, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}
