// Package indexer polls PayungPool events into Postgres so the API can
// serve driver dashboards from the database instead of the chain (reads
// must stay fast on a public testnet RPC). It never talks to RPC for the
// API's request path — only this background sync does.
package indexer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jackc/pgx/v5"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// blockChunkSize caps how many blocks are scanned per FilterLogs call, so a
// long-unsynced deployment doesn't attempt one huge range query.
const blockChunkSize = 5000

type Service struct {
	Queries    *store.Queries
	Pool       *chain.PayungPool
	StartBlock uint64 // used the first time, when indexer_state has no last_block yet (e.g. the pool's deploy block)
}

func New(queries *store.Queries, pool *chain.PayungPool, startBlock uint64) *Service {
	return &Service{Queries: queries, Pool: pool, StartBlock: startBlock}
}

// Sync scans from the last processed block up to latestBlock (inclusive),
// in chunks, persisting progress after each chunk so a crash mid-sync
// resumes rather than rescanning from the start.
func (s *Service) Sync(ctx context.Context, latestBlock uint64) error {
	from, err := s.lastBlock(ctx)
	if err != nil {
		return fmt.Errorf("indexer.Sync: %w", err)
	}

	for from <= latestBlock {
		to := from + blockChunkSize
		if to > latestBlock {
			to = latestBlock
		}

		if err := s.processRange(ctx, from, to); err != nil {
			return fmt.Errorf("indexer.Sync: process range [%d,%d]: %w", from, to, err)
		}
		if err := s.Queries.SetLastBlock(ctx, int64(to+1)); err != nil {
			return fmt.Errorf("indexer.Sync: save progress: %w", err)
		}
		from = to + 1
	}

	return nil
}

func (s *Service) lastBlock(ctx context.Context) (uint64, error) {
	v, err := s.Queries.GetLastBlock(ctx)
	if err == nil {
		return uint64(v), nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return s.StartBlock, nil
	}
	return 0, fmt.Errorf("lastBlock: %w", err)
}

func (s *Service) processRange(ctx context.Context, from, to uint64) error {
	opts := &bind.FilterOpts{Start: from, End: &to, Context: ctx}

	if err := s.indexPolicyBought(ctx, opts); err != nil {
		return err
	}
	if err := s.indexPayoutSent(ctx, opts); err != nil {
		return err
	}
	logPayoutSkipped(s.Pool, opts)
	logPolicyExpired(s.Pool, opts)

	return nil
}

func (s *Service) indexPolicyBought(ctx context.Context, opts *bind.FilterOpts) error {
	it, err := s.Pool.FilterPolicyBought(opts, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("filter PolicyBought: %w", err)
	}
	defer it.Close()

	for it.Next() {
		ev := it.Event
		err := s.Queries.UpsertPolicy(ctx, store.UpsertPolicyParams{
			PolicyID:    ev.PolicyId.Int64(),
			Holder:      strings.ToLower(ev.Holder.Hex()),
			ZoneID:      int16(ev.ZoneId),
			StartDay:    int32(ev.StartDay),
			EndDay:      int32(ev.EndDay),
			PremiumPaid: store.WeiFromBigInt(ev.Premium),
			TxHash:      strings.ToLower(ev.Raw.TxHash.Hex()),
			BlockNumber: int64(ev.Raw.BlockNumber),
		})
		if err != nil {
			return fmt.Errorf("upsert policy %d: %w", ev.PolicyId, err)
		}
	}
	return it.Error()
}

func (s *Service) indexPayoutSent(ctx context.Context, opts *bind.FilterOpts) error {
	it, err := s.Pool.FilterPayoutSent(opts, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("filter PayoutSent: %w", err)
	}
	defer it.Close()

	for it.Next() {
		ev := it.Event

		var mm int32
		obs, err := s.Queries.GetRainObservation(ctx, store.GetRainObservationParams{
			ZoneID: int16(ev.ZoneId), DayIndex: int32(ev.DayIndex),
		})
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return fmt.Errorf("load rain observation for payout policy=%d day=%d: %w", ev.PolicyId, ev.DayIndex, err)
			}
			log.Printf("indexer: no rain_observations row for zone=%d day=%d, recording payout with mm=0", ev.ZoneId, ev.DayIndex)
		} else {
			mm = obs.Mm
		}

		err = s.Queries.UpsertPayout(ctx, store.UpsertPayoutParams{
			PolicyID: ev.PolicyId.Int64(),
			Holder:   strings.ToLower(ev.Holder.Hex()),
			ZoneID:   int16(ev.ZoneId),
			DayIndex: int32(ev.DayIndex),
			Mm:       mm,
			Amount:   store.WeiFromBigInt(ev.Amount),
			TxHash:   strings.ToLower(ev.Raw.TxHash.Hex()),
		})
		if err != nil {
			return fmt.Errorf("upsert payout policy=%d day=%d: %w", ev.PolicyId, ev.DayIndex, err)
		}
	}
	return it.Error()
}

// logPayoutSkipped and logPolicyExpired have no dedicated table (SCHEMA.md
// only asks the indexer to write policies/payouts); they're surfaced as log
// lines for operator visibility. "Active" status is derived from end_day at
// query time, so PolicyExpired needs no write either.

func logPayoutSkipped(pool *chain.PayungPool, opts *bind.FilterOpts) {
	it, err := pool.FilterPayoutSkipped(opts, nil, nil)
	if err != nil {
		log.Printf("indexer: filter PayoutSkipped: %v", err)
		return
	}
	defer it.Close()
	for it.Next() {
		ev := it.Event
		log.Printf("indexer: PayoutSkipped policyId=%d day=%d reason=%q", ev.PolicyId, ev.DayIndex, ev.Reason)
	}
	if err := it.Error(); err != nil {
		log.Printf("indexer: iterate PayoutSkipped: %v", err)
	}
}

func logPolicyExpired(pool *chain.PayungPool, opts *bind.FilterOpts) {
	it, err := pool.FilterPolicyExpired(opts, nil)
	if err != nil {
		log.Printf("indexer: filter PolicyExpired: %v", err)
		return
	}
	defer it.Close()
	for it.Next() {
		log.Printf("indexer: PolicyExpired policyId=%d", it.Event.PolicyId)
	}
	if err := it.Error(); err != nil {
		log.Printf("indexer: iterate PolicyExpired: %v", err)
	}
}
