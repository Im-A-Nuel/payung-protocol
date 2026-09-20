// Package chain wraps go-ethereum access to PayungPool and IDRP: a signer,
// retrying transaction sends, and the day_index/week_index formulas shared
// with Solidity.
package chain

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	defaultRetryAttempts = 3
	defaultRetryBaseWait = 2 * time.Second
)

// Client bundles a JSON-RPC connection with bound PayungPool/IDRP contracts
// and a single signer used for every outgoing transaction.
type Client struct {
	eth  *ethclient.Client
	Pool *PayungPool
	IDRP *IDRP
	auth *bind.TransactOpts
}

// Dial connects to rpcURL and binds both contracts, signing every
// transaction with privateKeyHex.
func Dial(ctx context.Context, rpcURL string, chainID int64, privateKeyHex string, poolAddress, idrpAddress common.Address) (*Client, error) {
	ec, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("chain.Dial: connect %s: %w", rpcURL, err)
	}

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("chain.Dial: parse private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(chainID))
	if err != nil {
		return nil, fmt.Errorf("chain.Dial: build transactor: %w", err)
	}

	pool, err := NewPayungPool(poolAddress, ec)
	if err != nil {
		return nil, fmt.Errorf("chain.Dial: bind PayungPool: %w", err)
	}
	idrp, err := NewIDRP(idrpAddress, ec)
	if err != nil {
		return nil, fmt.Errorf("chain.Dial: bind IDRP: %w", err)
	}

	return &Client{eth: ec, Pool: pool, IDRP: idrp, auth: auth}, nil
}

// Address is the signer's own address (the oracle or faucet signer).
func (c *Client) Address() common.Address {
	return c.auth.From
}

func (c *Client) txOpts(ctx context.Context) *bind.TransactOpts {
	opts := *c.auth
	opts.Context = ctx
	return &opts
}

// sendWithRetry submits a transaction via send, waits for it to be mined,
// and retries the whole attempt (new nonce, new tx) up to defaultRetryAttempts
// times with exponential backoff. It does not retry on-chain reverts that
// are certain to fail again (e.g. RainAlreadyReported) — callers should
// check chain state first where that matters (see cmd/oracle).
func sendWithRetry(ctx context.Context, eth *ethclient.Client, send func() (*types.Transaction, error)) (*types.Transaction, error) {
	var lastErr error
	wait := defaultRetryBaseWait

	for attempt := 1; attempt <= defaultRetryAttempts; attempt++ {
		tx, err := send()
		if err == nil {
			if _, err := bind.WaitMined(ctx, eth, tx); err != nil {
				lastErr = fmt.Errorf("wait mined: %w", err)
			} else {
				return tx, nil
			}
		} else {
			lastErr = err
		}

		if attempt == defaultRetryAttempts {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(wait):
		}
		wait *= 2
	}

	return nil, fmt.Errorf("giving up after %d attempts: %w", defaultRetryAttempts, lastErr)
}

// RainReported reports whether submitRainfall has already been called for
// (zoneID, dayIndex), letting callers avoid a doomed resubmission.
func (c *Client) RainReported(ctx context.Context, zoneID uint16, dayIndex uint32) (bool, error) {
	reported, err := c.Pool.RainReported(&bind.CallOpts{Context: ctx}, zoneID, dayIndex)
	if err != nil {
		return false, fmt.Errorf("chain.RainReported: %w", err)
	}
	return reported, nil
}

// SubmitRainfall calls PayungPool.submitRainfall, retrying on transient
// failures, and returns the mined transaction hash.
func (c *Client) SubmitRainfall(ctx context.Context, zoneID uint16, dayIndex uint32, mm uint16) (string, error) {
	tx, err := sendWithRetry(ctx, c.eth, func() (*types.Transaction, error) {
		return c.Pool.SubmitRainfall(c.txOpts(ctx), zoneID, dayIndex, mm)
	})
	if err != nil {
		return "", fmt.Errorf("chain.SubmitRainfall: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// Settle calls PayungPool.settle, retrying on transient failures, and
// returns the mined transaction hash.
func (c *Client) Settle(ctx context.Context, zoneID uint16, dayIndex uint32) (string, error) {
	tx, err := sendWithRetry(ctx, c.eth, func() (*types.Transaction, error) {
		return c.Pool.Settle(c.txOpts(ctx), zoneID, dayIndex)
	})
	if err != nil {
		return "", fmt.Errorf("chain.Settle: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// SetPremium calls PayungPool.setPremium, retrying on transient failures,
// and returns the mined transaction hash.
func (c *Client) SetPremium(ctx context.Context, zoneID uint16, premiumPerWeek *big.Int) (string, error) {
	tx, err := sendWithRetry(ctx, c.eth, func() (*types.Transaction, error) {
		return c.Pool.SetPremium(c.txOpts(ctx), zoneID, premiumPerWeek)
	})
	if err != nil {
		return "", fmt.Errorf("chain.SetPremium: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// Faucet calls IDRP.faucet(to), retrying on transient failures, and returns
// the mined transaction hash.
func (c *Client) Faucet(ctx context.Context, to common.Address) (string, error) {
	tx, err := sendWithRetry(ctx, c.eth, func() (*types.Transaction, error) {
		return c.IDRP.Faucet(c.txOpts(ctx), to)
	})
	if err != nil {
		return "", fmt.Errorf("chain.Faucet: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// LatestBlock returns the current block number, used by the indexer to
// bound its scan range.
func (c *Client) LatestBlock(ctx context.Context) (uint64, error) {
	n, err := c.eth.BlockNumber(ctx)
	if err != nil {
		return 0, fmt.Errorf("chain.LatestBlock: %w", err)
	}
	return n, nil
}

// EthClient exposes the underlying ethclient for the indexer's log
// filtering, which needs lower-level access than the bound contracts give.
func (c *Client) EthClient() *ethclient.Client {
	return c.eth
}
