// Package http wires Fiber routes, middleware, and handlers.
// Handlers stay thin and delegate to services. See backend/CLAUDE.md.
package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

// NewRouter builds the Fiber app with all routes registered.
func NewRouter(db *bun.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler, // consistent error shape for all routes
	})

	api := app.Group("/api/v1")
	api.Get("/health", healthHandler(db))

	// Register more domain route groups here.
	return app
}

// healthHandler reports liveness and DB connectivity.
func healthHandler(db *bun.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		status := "ok"
		if err := db.PingContext(c.Context()); err != nil {
			status = "degraded"
		}
		return c.JSON(fiber.Map{"status": status})
	}
}

// errorHandler renders a consistent JSON error envelope.
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": fiber.Map{"code": code, "message": err.Error()},
	})
}
