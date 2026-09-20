// Package config loads configuration from environment variables only.
// Every Load* function fails fast (returns an error) if a required
// variable is missing, rather than falling back to a default that would
// hide a misconfigured deployment.
package config

import (
	"fmt"
	"os"
	"strconv"
)

func requireEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("config.requireEnv: missing required env var %s", key)
	}
	return v, nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Common holds settings shared by every binary.
type Common struct {
	DatabaseURL string
	Env         string // "development", "production", etc. Defaults to "development".
}

func loadCommon() (Common, error) {
	dbURL, err := requireEnv("DATABASE_URL")
	if err != nil {
		return Common{}, err
	}
	return Common{
		DatabaseURL: dbURL,
		Env:         envOr("ENV", "development"),
	}, nil
}

// IsProduction reports whether admin/demo-only endpoints must stay disabled.
func (c Common) IsProduction() bool {
	return c.Env == "production"
}

// ChainConfig holds the on-chain connection details shared by the oracle
// and the API's faucet/relay paths.
type ChainConfig struct {
	RPCURL      string
	ChainID     int64
	PoolAddress string
	IDRPAddress string
	DeployBlock uint64 // block PayungPool was deployed at; indexer's starting point when indexer_state is empty
}

func loadChainConfig() (ChainConfig, error) {
	rpcURL, err := requireEnv("OPBNB_RPC")
	if err != nil {
		return ChainConfig{}, err
	}
	poolAddr, err := requireEnv("POOL_ADDRESS")
	if err != nil {
		return ChainConfig{}, err
	}
	idrpAddr, err := requireEnv("IDRP_ADDRESS")
	if err != nil {
		return ChainConfig{}, err
	}
	chainIDStr := envOr("OPBNB_CHAIN_ID", "5611")
	chainID, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		return ChainConfig{}, fmt.Errorf("config.loadChainConfig: invalid OPBNB_CHAIN_ID %q: %w", chainIDStr, err)
	}
	deployBlockStr := envOr("DEPLOY_BLOCK", "0")
	deployBlock, err := strconv.ParseUint(deployBlockStr, 10, 64)
	if err != nil {
		return ChainConfig{}, fmt.Errorf("config.loadChainConfig: invalid DEPLOY_BLOCK %q: %w", deployBlockStr, err)
	}
	return ChainConfig{
		RPCURL:      rpcURL,
		ChainID:     chainID,
		PoolAddress: poolAddr,
		IDRPAddress: idrpAddr,
		DeployBlock: deployBlock,
	}, nil
}

// MigrateConfig is the config needed by cmd/migrate.
type MigrateConfig struct {
	Common
}

func LoadMigrateConfig() (MigrateConfig, error) {
	common, err := loadCommon()
	if err != nil {
		return MigrateConfig{}, err
	}
	return MigrateConfig{Common: common}, nil
}

// OracleConfig is the config needed by cmd/oracle.
type OracleConfig struct {
	Common
	ChainConfig
	OraclePrivateKey string
}

func LoadOracleConfig() (OracleConfig, error) {
	common, err := loadCommon()
	if err != nil {
		return OracleConfig{}, err
	}
	chain, err := loadChainConfig()
	if err != nil {
		return OracleConfig{}, err
	}
	pk, err := requireEnv("ORACLE_PRIVATE_KEY")
	if err != nil {
		return OracleConfig{}, err
	}
	return OracleConfig{Common: common, ChainConfig: chain, OraclePrivateKey: pk}, nil
}

// APIConfig is the config needed by cmd/api.
type APIConfig struct {
	Common
	ChainConfig
	ListenAddr          string
	AdminKey            string
	FaucetPrivateKey    string
	OraclePrivateKey    string
	FaucetCooldownHours int
}

func LoadAPIConfig() (APIConfig, error) {
	common, err := loadCommon()
	if err != nil {
		return APIConfig{}, err
	}
	chain, err := loadChainConfig()
	if err != nil {
		return APIConfig{}, err
	}
	adminKey, err := requireEnv("ADMIN_KEY")
	if err != nil {
		return APIConfig{}, err
	}
	faucetPK, err := requireEnv("FAUCET_PRIVATE_KEY")
	if err != nil {
		return APIConfig{}, err
	}
	oraclePK, err := requireEnv("ORACLE_PRIVATE_KEY")
	if err != nil {
		return APIConfig{}, err
	}
	return APIConfig{
		Common:              common,
		ChainConfig:         chain,
		ListenAddr:          ":" + envOr("PORT", "8080"),
		AdminKey:            adminKey,
		FaucetPrivateKey:    faucetPK,
		OraclePrivateKey:    oraclePK,
		FaucetCooldownHours: 24,
	}, nil
}
