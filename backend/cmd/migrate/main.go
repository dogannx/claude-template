// Command migrate applies and rolls back the SQL migrations in ./migrations.
//
// Usage: go run ./cmd/migrate [up|down|status|unlock]   (default: up)
//
//	up      apply all pending migrations as one group
//	down    roll back the last applied group (which may hold several migrations)
//	status  list applied and pending migrations
//	unlock  clear a stale lock left by a crashed run
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"

	"github.com/dogannx/claude-template/backend/internal/config"
	"github.com/dogannx/claude-template/backend/internal/database"
	"github.com/dogannx/claude-template/backend/migrations"
)

func main() {
	cmd := "up"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	db, err := database.Connect(config.Load())
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	if err := run(context.Background(), newMigrator(db, migrations.Migrations), cmd); err != nil {
		log.Fatal(err)
	}
}

// newMigrator records a migration as applied only after its SQL succeeded, so a failed
// migration is retried on the next run instead of being silently skipped.
func newMigrator(db *bun.DB, ms *migrate.Migrations, opts ...migrate.MigratorOption) *migrate.Migrator {
	return migrate.NewMigrator(db, ms, append([]migrate.MigratorOption{migrate.WithMarkAppliedOnSuccess(true)}, opts...)...)
}

func run(ctx context.Context, m *migrate.Migrator, cmd string) (err error) {
	switch cmd {
	case "up", "down", "status", "unlock":
	default:
		return fmt.Errorf("unknown command %q (use up|down|status|unlock)", cmd)
	}

	// Init is idempotent: it creates bun's bookkeeping tables if they don't exist.
	if err := m.Init(ctx); err != nil {
		return fmt.Errorf("init: %w", err)
	}
	if cmd == "unlock" {
		if err := m.Unlock(ctx); err != nil {
			return fmt.Errorf("unlock: %w", err)
		}
		log.Println("lock cleared")
		return nil
	}
	if err := m.Lock(ctx); err != nil {
		return fmt.Errorf("lock (if a previous run crashed, run `migrate unlock`): %w", err)
	}
	defer func() {
		if uerr := m.Unlock(ctx); uerr != nil {
			err = errors.Join(err, fmt.Errorf("unlock: %w", uerr))
		}
	}()

	switch cmd {
	case "up":
		group, err := m.Migrate(ctx)
		if err != nil {
			return fmt.Errorf("migrate: %w", err)
		}
		if group.IsZero() {
			log.Println("nothing to migrate")
		} else {
			log.Printf("migrated to %s", group)
		}
	case "down":
		group, err := m.Rollback(ctx)
		if err != nil {
			return fmt.Errorf("rollback: %w", err)
		}
		if group.IsZero() {
			log.Println("nothing to roll back")
		} else {
			log.Printf("rolled back %s", group)
		}
	case "status":
		ms, err := m.MigrationsWithStatus(ctx)
		if err != nil {
			return fmt.Errorf("status: %w", err)
		}
		log.Printf("applied: %s", ms.Applied())
		log.Printf("pending: %s", ms.Unapplied())
	}
	return nil
}
