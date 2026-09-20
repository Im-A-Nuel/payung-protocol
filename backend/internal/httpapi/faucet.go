package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5"
)

type faucetRequest struct {
	Address string `json:"address"`
}

type faucetResponse struct {
	TxHash string `json:"txHash"`
}

func (s *Server) handleFaucet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req faucetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_ADDRESS", "Alamat wallet tidak valid.")
		return
	}

	address, ok := parseAddress(req.Address)
	if !ok {
		writeError(w, http.StatusBadRequest, "INVALID_ADDRESS", "Alamat wallet tidak valid.")
		return
	}

	claim, err := s.Queries.GetLastClaim(ctx, address)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal memeriksa status faucet.")
		return
	}
	if err == nil {
		cooldownEnds := claim.LastClaimAt.Time.Add(time.Duration(s.FaucetCooldownHours) * time.Hour)
		if time.Now().Before(cooldownEnds) {
			writeError(w, http.StatusTooManyRequests, "FAUCET_COOLDOWN", "Kamu sudah mengambil IDRP dalam 24 jam terakhir.")
			return
		}
	}

	txHash, err := s.FaucetChain.Faucet(ctx, common.HexToAddress(address))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "Gagal mengirim IDRP dari faucet.")
		return
	}

	if err := s.Queries.UpsertClaim(ctx, address); err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL", "IDRP terkirim tetapi gagal mencatat waktu klaim.")
		return
	}

	writeJSON(w, http.StatusOK, faucetResponse{TxHash: txHash})
}
