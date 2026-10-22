// Package config loads configuration from environment variables (12-factor).
package config

import "os"

type Config struct {
	AppEnv      string
	AppPort     string
	DatabaseURL string
	// CORSOrigins is a comma-separated list of origins allowed to call the API.
	CORSOrigins string
}

func Load() Config {
	return Config{
		AppEnv:      env("APP_ENV", "development"),
		AppPort:     env("APP_PORT", "8080"),
		DatabaseURL: env("DATABASE_URL", "postgres://app:change-me@localhost:5432/app_dev?sslmode=disable"),
		CORSOrigins: env("CORS_ORIGINS", "http://localhost:5173"),
	}
}

func env(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
