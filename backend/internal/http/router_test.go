package http

import (
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// unreachableDB returns a bun.DB whose pings fail fast, so the real router can be
// exercised without a running PostgreSQL.
func unreachableDB(t *testing.T) *bun.DB {
	t.Helper()
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN("postgres://u:p@127.0.0.1:1/db?sslmode=disable&dial_timeout=200ms"),
	))
	db := bun.NewDB(sqldb, pgdialect.New())
	t.Cleanup(func() { db.Close() })
	return db
}

func TestRouter(t *testing.T) {
	app := NewRouter(unreachableDB(t), "http://localhost:5173")

	tests := []struct {
		name       string
		path       string
		wantStatus int
		check      func(t *testing.T, body map[string]any)
	}{
		{
			name:       "health reports degraded when DB is down",
			path:       "/api/v1/health",
			wantStatus: fiber.StatusOK,
			check: func(t *testing.T, body map[string]any) {
				if body["status"] != "degraded" {
					t.Errorf("status = %v, want degraded", body["status"])
				}
			},
		},
		{
			name:       "unknown route uses the shared error envelope",
			path:       "/api/v1/nope",
			wantStatus: fiber.StatusNotFound,
			check: func(t *testing.T, body map[string]any) {
				e, ok := body["error"].(map[string]any)
				if !ok {
					t.Fatalf("missing error envelope: %v", body)
				}
				if e["code"] != float64(fiber.StatusNotFound) {
					t.Errorf("error.code = %v, want 404", e["code"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, tt.path, nil), 5000)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Fatalf("status code = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
			var body map[string]any
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			tt.check(t, body)
		})
	}
}

func TestRouterCORS(t *testing.T) {
	app := NewRouter(unreachableDB(t), "http://localhost:5173")

	for origin, want := range map[string]string{
		"http://localhost:5173": "http://localhost:5173",
		"http://evil.example":   "",
	} {
		req := httptest.NewRequest(fiber.MethodGet, "/api/v1/health", nil)
		req.Header.Set("Origin", origin)
		resp, err := app.Test(req, 5000)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		resp.Body.Close()
		if got := resp.Header.Get("Access-Control-Allow-Origin"); got != want {
			t.Errorf("origin %s: Allow-Origin = %q, want %q", origin, got, want)
		}
	}
}
