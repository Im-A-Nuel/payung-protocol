// Package oracle implements the "report rainfall, settle if it crossed the
// threshold" flow shared by the daily cron job (cmd/oracle) and the demo
// admin endpoint (POST /admin/simulate-rain): both are the same operation,
// just with mm coming from Open-Meteo or a manual value.
package oracle

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/rain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// Statuses stored in oracle_runs.status. "settled" and "no_rain" are
// terminal: ProcessZoneDay treats them as already done and returns early.
const (
	StatusPending   = "pending"
	StatusSubmitted = "submitted"
	StatusSettled   = "settled"
	StatusNoRain    = "no_rain"
	StatusFailed    = "failed"
)

// ChainOracle is the subset of *chain.Client the oracle service needs,
// declared as an interface so it can be exercised with a fake in tests.
type ChainOracle interface {
	RainReported(ctx context.Context, zoneID uint16, dayIndex uint32) (bool, error)
	SubmitRainfall(ctx context.Context, zoneID uint16, dayIndex uint32, mm uint16) (string, error)
	Settle(ctx context.Context, zoneID uint16, dayIndex uint32) (string, error)
}

type Service struct {
	Queries *store.Queries
	Chain   ChainOracle
	Rain    *rain.Client
}

func New(queries *store.Queries, chainClient ChainOracle, rainClient *rain.Client) *Service {
	return &Service{Queries: queries, Chain: chainClient, Rain: rainClient}
}

// Result summarizes what happened for one (zone, day).
type Result struct {
	ZoneID   int16
	DayIndex int32
	Status   string
	SubmitTx string
	SettleTx string
	Err      error
}

// RunDaily processes dayIndex (normally "yesterday") for every active zone,
// fetching mm from Open-Meteo. It never returns early on a single zone's
// error; each zone's outcome is reported in the returned slice so one bad
// zone doesn't block the rest.
func (s *Service) RunDaily(ctx context.Context, dayIndex uint32) ([]Result, error) {
	zones, err := s.Queries.ListActiveZones(ctx)
	if err != nil {
		return nil, fmt.Errorf("oracle.RunDaily: list zones: %w", err)
	}

	dateWIB := chain.DateStringWIB(dayIndex)
	results := make([]Result, 0, len(zones))

	for _, zone := range zones {
		points, err := s.Rain.ForecastDaily(ctx, zone.Lat, zone.Lon)
		if err != nil {
			results = append(results, Result{ZoneID: zone.ID, DayIndex: int32(dayIndex), Status: StatusFailed, Err: fmt.Errorf("fetch rain: %w", err)})
			continue
		}
		point, ok := rain.PointForDate(points, dateWIB)
		if !ok {
			results = append(results, Result{ZoneID: zone.ID, DayIndex: int32(dayIndex), Status: StatusFailed, Err: fmt.Errorf("no rain data for %s yet", dateWIB)})
			continue
		}

		res := s.ProcessZoneDay(ctx, zone.ID, dayIndex, point.Mm, "open-meteo")
		results = append(results, res)
	}

	return results, nil
}

// ProcessZoneDay upserts the rain observation, submits it on-chain (unless
// already reported), and settles if it crossed the zone's threshold. It is
// idempotent: calling it again for a (zone, day) that already reached
// "settled" or "no_rain" is a no-op that returns the prior result.
func (s *Service) ProcessZoneDay(ctx context.Context, zoneID int16, dayIndex uint32, mm int, source string) Result {
	result := Result{ZoneID: zoneID, DayIndex: int32(dayIndex)}

	existing, err := s.Queries.GetOracleRun(ctx, store.GetOracleRunParams{ZoneID: zoneID, DayIndex: int32(dayIndex)})
	if err == nil && (existing.Status == StatusSettled || existing.Status == StatusNoRain) {
		result.Status = existing.Status
		result.SubmitTx = textOrEmpty(existing.SubmitTx)
		result.SettleTx = textOrEmpty(existing.SettleTx)
		return result
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		result.Status = StatusFailed
		result.Err = fmt.Errorf("oracle.ProcessZoneDay: check existing run: %w", err)
		return result
	}

	if err := s.Queries.UpsertRainObservation(ctx, store.UpsertRainObservationParams{
		ZoneID: zoneID, DayIndex: int32(dayIndex), Mm: int32(mm), Source: source,
	}); err != nil {
		return s.fail(ctx, zoneID, dayIndex, fmt.Errorf("upsert rain observation: %w", err))
	}

	zone, err := s.Queries.GetZone(ctx, zoneID)
	if err != nil {
		return s.fail(ctx, zoneID, dayIndex, fmt.Errorf("load zone: %w", err))
	}

	var submitTx string
	already, err := s.Chain.RainReported(ctx, uint16(zoneID), dayIndex)
	if err != nil {
		return s.fail(ctx, zoneID, dayIndex, fmt.Errorf("check rainReported: %w", err))
	}
	if !already {
		submitTx, err = s.Chain.SubmitRainfall(ctx, uint16(zoneID), dayIndex, uint16(mm))
		if err != nil {
			return s.fail(ctx, zoneID, dayIndex, fmt.Errorf("submitRainfall: %w", err))
		}
	}

	isRainDay := int32(mm) >= zone.ThresholdMm
	if !isRainDay {
		run, err := s.Queries.UpsertOracleRun(ctx, store.UpsertOracleRunParams{
			ZoneID: zoneID, DayIndex: int32(dayIndex),
			SubmitTx: toText(submitTx), Status: StatusNoRain,
		})
		if err != nil {
			result.Status = StatusFailed
			result.Err = fmt.Errorf("oracle.ProcessZoneDay: record no_rain: %w", err)
			return result
		}
		result.Status = run.Status
		result.SubmitTx = textOrEmpty(run.SubmitTx)
		return result
	}

	settleTx, err := s.Chain.Settle(ctx, uint16(zoneID), dayIndex)
	if err != nil {
		return s.fail(ctx, zoneID, dayIndex, fmt.Errorf("settle: %w", err))
	}

	run, err := s.Queries.UpsertOracleRun(ctx, store.UpsertOracleRunParams{
		ZoneID: zoneID, DayIndex: int32(dayIndex),
		SubmitTx: toText(submitTx), SettleTx: toText(settleTx), Status: StatusSettled,
	})
	if err != nil {
		result.Status = StatusFailed
		result.Err = fmt.Errorf("oracle.ProcessZoneDay: record settled: %w", err)
		return result
	}
	result.Status = run.Status
	result.SubmitTx = textOrEmpty(run.SubmitTx)
	result.SettleTx = textOrEmpty(run.SettleTx)
	return result
}

func (s *Service) fail(ctx context.Context, zoneID int16, dayIndex uint32, cause error) Result {
	_, _ = s.Queries.UpsertOracleRun(ctx, store.UpsertOracleRunParams{
		ZoneID: zoneID, DayIndex: int32(dayIndex),
		Status: StatusFailed, Error: toText(cause.Error()),
	})
	return Result{ZoneID: zoneID, DayIndex: int32(dayIndex), Status: StatusFailed, Err: cause}
}

func toText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{}
	}
	return pgtype.Text{String: s, Valid: true}
}

func textOrEmpty(t pgtype.Text) string {
	if !t.Valid {
		return ""
	}
	return t.String
}
