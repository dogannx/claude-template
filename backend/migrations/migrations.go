// Package migrations holds the SQL migrations, embedded into the migrate binary.
//
// File naming (bun): <YYYYMMDDHHMMSS>_<name>.tx.up.sql / .tx.down.sql. The ".tx" part
// runs the file inside a transaction. See the create-migration skill.
package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

// Migrations is the registry discovered from the embedded *.sql files.
var Migrations = migrate.NewMigrations()

// "*" rather than "*.sql" so the build still works with zero migrations; Discover
// ignores files that are not *.up.sql / *.down.sql.
//
//go:embed *
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
