package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"

	"github.com/dogannx/claude-template/backend/migrations"
)

// isolatedDB returns a connection whose search_path is a fresh schema that is dropped
// after the test, so these tests never touch existing tables — even on a dev database.
// They run when TEST_DATABASE_URL is set (CI sets it), e.g.:
//
//	TEST_DATABASE_URL=postgres://app:change-me@localhost:5432/app_dev?sslmode=disable go test ./cmd/migrate
func isolatedDB(t *testing.T) *bun.DB {
	t.Helper()
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}
	ctx := context.Background()
	schema := fmt.Sprintf("migtest_%d", time.Now().UnixNano())

	admin := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))), pgdialect.New())
	if _, err := admin.ExecContext(ctx, "CREATE SCHEMA "+schema); err != nil {
		t.Fatalf("create schema: %v", err)
	}
	t.Cleanup(func() {
		admin.ExecContext(ctx, "DROP SCHEMA "+schema+" CASCADE") //nolint:errcheck
		admin.Close()
	})

	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
		pgdriver.WithConnParams(map[string]interface{}{"search_path": schema}),
	)), pgdialect.New())
	t.Cleanup(func() { db.Close() })
	return db
}

func tableExists(t *testing.T, db *bun.DB, name string) bool {
	t.Helper()
	var reg sql.NullString
	if err := db.QueryRowContext(context.Background(), "SELECT to_regclass(?)::text", name).Scan(&reg); err != nil {
		t.Fatalf("query: %v", err)
	}
	return reg.Valid
}

// The migrator itself, against a fixture registry so it doesn't depend on which
// migrations the project currently has.
func TestMigratorRoundTrip(t *testing.T) {
	db := isolatedDB(t)
	fixture := migrate.NewMigrations()
	if err := fixture.Discover(fstest.MapFS{
		"20000101000000_fixture.tx.up.sql":   {Data: []byte("CREATE TABLE fixture_items (id int)")},
		"20000101000000_fixture.tx.down.sql": {Data: []byte("DROP TABLE fixture_items")},
	}); err != nil {
		t.Fatalf("discover: %v", err)
	}
	m := newMigrator(db, fixture)

	// up twice proves idempotence; down must actually undo the schema change.
	for _, step := range []struct {
		cmd  string
		want bool
	}{{"up", true}, {"up", true}, {"down", false}, {"up", true}} {
		if err := run(context.Background(), m, step.cmd); err != nil {
			t.Fatalf("%s: %v", step.cmd, err)
		}
		if got := tableExists(t, db, "fixture_items"); got != step.want {
			t.Fatalf("after %s: fixture_items exists = %v, want %v", step.cmd, got, step.want)
		}
	}
}

// A failed migration must not be recorded as applied, so the next run retries it.
func TestFailedMigrationIsRetried(t *testing.T) {
	db := isolatedDB(t)
	broken := migrate.NewMigrations()
	if err := broken.Discover(fstest.MapFS{
		"20000101000000_broken.tx.up.sql":   {Data: []byte("CREATE TABLE ok_part (id int); SELECT no_such_column FROM ok_part")},
		"20000101000000_broken.tx.down.sql": {Data: []byte("DROP TABLE IF EXISTS ok_part")},
	}); err != nil {
		t.Fatalf("discover: %v", err)
	}
	m := newMigrator(db, broken)
	ctx := context.Background()

	if err := run(ctx, m, "up"); err == nil {
		t.Fatal("want error from broken migration")
	}
	ms, err := m.MigrationsWithStatus(ctx)
	if err != nil {
		t.Fatalf("status: %v", err)
	}
	if n := len(ms.Unapplied()); n != 1 {
		t.Fatalf("pending migrations = %d, want 1 (failed one must stay pending)", n)
	}
	if tableExists(t, db, "ok_part") {
		t.Fatal("transactional migration left partial schema behind")
	}
}

// The project's real migrations apply cleanly and fully roll back.
func TestProjectMigrationsApplyAndRollBack(t *testing.T) {
	db := isolatedDB(t)
	m := newMigrator(db, migrations.Migrations)
	ctx := context.Background()

	if err := run(ctx, m, "up"); err != nil {
		t.Fatalf("up: %v", err)
	}
	if err := run(ctx, m, "down"); err != nil {
		t.Fatalf("down: %v", err)
	}
	ms, err := m.MigrationsWithStatus(ctx)
	if err != nil {
		t.Fatalf("status: %v", err)
	}
	if applied := ms.Applied(); len(applied) != 0 {
		t.Fatalf("still applied after down: %s", applied)
	}
}

func TestRunRejectsUnknownCommand(t *testing.T) {
	if err := run(context.Background(), nil, "sideways"); err == nil {
		t.Fatal("want error for unknown command")
	}
}
