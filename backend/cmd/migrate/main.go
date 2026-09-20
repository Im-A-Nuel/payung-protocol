// Command migrate applies the SQL files in backend/migrations, in filename
// order, tracking which ones already ran in a schema_migrations table so it
// is safe to run repeatedly (e.g. on every deploy).
package main

import (
	"context"
	"fmt"
	"log"
	"path"
	"sort"

	"github.com/jackc/pgx/v5"

	"github.com/im-a-nuel/payung-protocol/backend/internal/config"
	"github.com/im-a-nuel/payung-protocol/backend/migrations"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("migrate: %v", err)
	}
}

func run() error {
	cfg, err := config.LoadMigrateConfig()
	if err != nil {
		return fmt.Errorf("migrate.run: %w", err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("migrate.run: connect: %w", err)
	}
	defer conn.Close(ctx)

	if _, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			filename    TEXT PRIMARY KEY,
			applied_at  TIMESTAMPTZ NOT NULL DEFAULT now()
		)
	`); err != nil {
		return fmt.Errorf("migrate.run: create schema_migrations: %w", err)
	}

	entries, err := migrations.Files.ReadDir(".")
	if err != nil {
		return fmt.Errorf("migrate.run: read embedded migrations: %w", err)
	}

	var filenames []string
	for _, e := range entries {
		if !e.IsDir() && path.Ext(e.Name()) == ".sql" {
			filenames = append(filenames, e.Name())
		}
	}
	sort.Strings(filenames)

	for _, filename := range filenames {
		var already bool
		err := conn.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE filename = $1)`, filename).Scan(&already)
		if err != nil {
			return fmt.Errorf("migrate.run: check %s: %w", filename, err)
		}
		if already {
			log.Printf("migrate: %s already applied, skipping", filename)
			continue
		}

		sqlBytes, err := migrations.Files.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("migrate.run: read %s: %w", filename, err)
		}

		tx, err := conn.Begin(ctx)
		if err != nil {
			return fmt.Errorf("migrate.run: begin tx for %s: %w", filename, err)
		}

		if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("migrate.run: apply %s: %w", filename, err)
		}
		if _, err := tx.Exec(ctx, `INSERT INTO schema_migrations (filename) VALUES ($1)`, filename); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("migrate.run: record %s: %w", filename, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("migrate.run: commit %s: %w", filename, err)
		}

		log.Printf("migrate: applied %s", filename)
	}

	return nil
}
