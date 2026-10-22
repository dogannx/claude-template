// Command migrate applies/rolls back database migrations.
// Usage: go run ./cmd/migrate up   |   go run ./cmd/migrate down
//
// This is a minimal placeholder. Wire it to bun's migrate package (or your tool of
// choice) and the migrations in ./migrations. See the create-migration skill.
package main

import (
	"log"
	"os"

	"github.com/dogannx/claude-template/backend/internal/config"
	"github.com/dogannx/claude-template/backend/internal/database"
)

func main() {
	cmd := "up"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	cfg := config.Load()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	switch cmd {
	case "up":
		log.Println("TODO: apply migrations from ./migrations")
	case "down":
		log.Println("TODO: roll back last migration")
	default:
		log.Fatalf("unknown command %q (use up|down)", cmd)
	}
}
