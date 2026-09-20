// Command oracle runs the daily "fetch rain, submit, settle" job for every
// active zone. By default it loops forever, waking at 06:00 WIB each day;
// pass --once to run a single pass immediately and exit (used for manual
// runs and Railway/GitHub Actions cron, which invoke the process fresh).
package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/config"
	"github.com/im-a-nuel/payung-protocol/backend/internal/indexer"
	"github.com/im-a-nuel/payung-protocol/backend/internal/oracle"
	"github.com/im-a-nuel/payung-protocol/backend/internal/rain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

func main() {
	once := flag.Bool("once", false, "run a single pass for yesterday's rainfall and exit")
	flag.Parse()

	cfg, err := config.LoadOracleConfig()
	if err != nil {
		log.Fatalf("oracle: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := store.ConnectPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("oracle: %v", err)
	}
	defer pool.Close()

	chainClient, err := chain.Dial(
		ctx, cfg.RPCURL, cfg.ChainID, cfg.OraclePrivateKey,
		common.HexToAddress(cfg.PoolAddress), common.HexToAddress(cfg.IDRPAddress),
	)
	if err != nil {
		log.Fatalf("oracle: %v", err)
	}

	svc := oracle.New(store.New(pool), chainClient, rain.NewClient())
	indexerSvc := indexer.New(store.New(pool), chainClient.Pool, cfg.DeployBlock)

	if *once {
		runOnce(ctx, svc, indexerSvc, chainClient)
		return
	}

	for {
		sleepUntilNext6AMWIB()
		if ctx.Err() != nil {
			return
		}
		runOnce(ctx, svc, indexerSvc, chainClient)
	}
}

func runOnce(ctx context.Context, svc *oracle.Service, indexerSvc *indexer.Service, chainClient *chain.Client) {
	yesterday := chain.DayIndex(time.Now().Unix()) - 1
	log.Printf("oracle: processing day_index=%d (%s)", yesterday, chain.DateStringWIB(yesterday))

	results, err := svc.RunDaily(ctx, yesterday)
	if err != nil {
		log.Printf("oracle: RunDaily failed: %v", err)
		return
	}
	for _, r := range results {
		if r.Err != nil {
			log.Printf("oracle: zone=%d day=%d status=%s error=%v", r.ZoneID, r.DayIndex, r.Status, r.Err)
			continue
		}
		log.Printf("oracle: zone=%d day=%d status=%s submitTx=%s settleTx=%s", r.ZoneID, r.DayIndex, r.Status, r.SubmitTx, r.SettleTx)
	}

	// Per ARCHITECTURE.md: "the indexer catches up on every oracle run".
	latest, err := chainClient.LatestBlock(ctx)
	if err != nil {
		log.Printf("oracle: indexer sync skipped, could not read latest block: %v", err)
		return
	}
	if err := indexerSvc.Sync(ctx, latest); err != nil {
		log.Printf("oracle: indexer sync failed: %v", err)
	}
}

// sleepUntilNext6AMWIB blocks until the next 06:00 WIB (23:00 UTC the
// previous day, since WIB is UTC+7).
func sleepUntilNext6AMWIB() {
	now := time.Now().UTC()
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, time.UTC)
	if !now.Before(next) {
		next = next.Add(24 * time.Hour)
	}
	time.Sleep(time.Until(next))
}
