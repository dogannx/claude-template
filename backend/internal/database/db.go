// Package database opens and configures the bun DB handle (PostgreSQL).
package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/yourname/yourproject/backend/internal/config"
)

// Connect returns a ready bun.DB. Caller is responsible for Close().
func Connect(cfg config.Config) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DatabaseURL)))
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
