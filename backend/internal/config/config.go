// Package config loads configuration from environment variables (12-factor).
package config

import "os"

type Config struct {
	AppEnv      string
	AppPort     string
	DatabaseURL string
}

func Load() Config {
	return Config{
		AppEnv:      env("APP_ENV", "development"),
		AppPort:     env("APP_PORT", "8080"),
		DatabaseURL: env("DATABASE_URL", "postgres://app:change-me@localhost:5432/app_dev?sslmode=disable"),
	}
}

func env(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
