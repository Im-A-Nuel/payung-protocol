// Command api serves the HTTP API described in docs/SCHEMA.md: public
// zone/driver reads backed by Postgres, the faucet write, and (outside
// production) demo/admin endpoints.
package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/config"
	"github.com/im-a-nuel/payung-protocol/backend/internal/httpapi"
	"github.com/im-a-nuel/payung-protocol/backend/internal/indexer"
	"github.com/im-a-nuel/payung-protocol/backend/internal/oracle"
	"github.com/im-a-nuel/payung-protocol/backend/internal/rain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

const httpShutdownTimeout = 5 * time.Second

func main() {
	cfg, err := config.LoadAPIConfig()
	if err != nil {
		log.Fatalf("api: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := store.ConnectPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("api: %v", err)
	}
	defer pool.Close()
	queries := store.New(pool)

	poolAddr := common.HexToAddress(cfg.PoolAddress)
	idrpAddr := common.HexToAddress(cfg.IDRPAddress)

	oracleChain, err := chain.Dial(ctx, cfg.RPCURL, cfg.ChainID, cfg.OraclePrivateKey, poolAddr, idrpAddr)
	if err != nil {
		log.Fatalf("api: dial oracle chain client: %v", err)
	}
	faucetChain, err := chain.Dial(ctx, cfg.RPCURL, cfg.ChainID, cfg.FaucetPrivateKey, poolAddr, idrpAddr)
	if err != nil {
		log.Fatalf("api: dial faucet chain client: %v", err)
	}

	oracleSvc := oracle.New(queries, oracleChain, rain.NewClient())
	indexerSvc := indexer.New(queries, oracleChain.Pool, cfg.DeployBlock)

	server := &httpapi.Server{
		Queries:             queries,
		FaucetChain:         faucetChain,
		OracleChain:         oracleChain,
		Oracle:              oracleSvc,
		Indexer:             indexerSvc,
		AdminKey:            cfg.AdminKey,
		IsProduction:        cfg.IsProduction(),
		FaucetCooldownHours: cfg.FaucetCooldownHours,
	}

	// Catch up before serving so the very first requests already reflect
	// on-chain state, per ARCHITECTURE.md ("the indexer catches up on
	// every oracle run and on API cold start").
	if latest, err := oracleChain.LatestBlock(ctx); err != nil {
		log.Printf("api: initial indexer sync skipped, could not read latest block: %v", err)
	} else if err := indexerSvc.Sync(ctx, latest); err != nil {
		log.Printf("api: initial indexer sync failed: %v", err)
	}

	httpServer := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: httpapi.NewRouter(server),
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), httpShutdownTimeout)
		defer cancel()
		_ = httpServer.Shutdown(shutdownCtx)
	}()

	log.Printf("api: listening on %s (env=%s)", cfg.ListenAddr, cfg.Env)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("api: %v", err)
	}
}
