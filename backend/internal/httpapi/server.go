// Package httpapi implements the Go API described in docs/SCHEMA.md: public
// read endpoints backed by Postgres (never RPC, so the dashboard stays
// fast), the faucet write, and non-production admin/demo endpoints.
package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/im-a-nuel/payung-protocol/backend/internal/chain"
	"github.com/im-a-nuel/payung-protocol/backend/internal/indexer"
	"github.com/im-a-nuel/payung-protocol/backend/internal/oracle"
	"github.com/im-a-nuel/payung-protocol/backend/internal/store"
)

// Server holds every dependency the HTTP handlers need. FaucetChain is a
// separate signer from Oracle's chain client because the faucet address
// only needs IDRP's FAUCET_ROLE, not ORACLE_ROLE.
type Server struct {
	Queries      *store.Queries
	FaucetChain  *chain.Client
	OracleChain  *chain.Client // also used for indexer.Sync's LatestBlock lookups
	Oracle       *oracle.Service
	Indexer      *indexer.Service
	AdminKey     string
	IsProduction bool

	FaucetCooldownHours int
}

func NewRouter(s *Server) http.Handler {
	r := chi.NewRouter()
	r.Use(corsMiddleware)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/zones", s.handleListZones)
		r.Get("/zones/{id}/rain", s.handleZoneRain)
		r.Get("/drivers/{address}/policy", s.handleDriverPolicy)
		r.Get("/drivers/{address}/payouts", s.handleDriverPayouts)
		r.Post("/faucet", s.handleFaucet)

		if !s.IsProduction {
			r.Route("/admin", func(r chi.Router) {
				r.Use(s.requireAdminKey)
				r.Post("/simulate-rain", s.handleSimulateRain)
				r.Post("/oracle/run", s.handleOracleRun)
			})
		}
	})

	return r
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Admin-Key")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) requireAdminKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Admin-Key") != s.AdminKey {
			writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Admin key tidak valid.")
			return
		}
		next.ServeHTTP(w, r)
	})
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorEnvelope struct {
	Error apiError `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errorEnvelope{Error: apiError{Code: code, Message: message}})
}
