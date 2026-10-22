// Command api is the HTTP server entrypoint. Keep it thin: load config, wire
// dependencies, start the server. Logic lives in internal/.
package main

import (
	"log"

	"github.com/dogannx/claude-template/backend/internal/config"
	"github.com/dogannx/claude-template/backend/internal/database"
	apihttp "github.com/dogannx/claude-template/backend/internal/http"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	app := apihttp.NewRouter(db, cfg.CORSOrigins)

	log.Printf("listening on :%s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("server: %v", err)
	}
}
