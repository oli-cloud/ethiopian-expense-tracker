package config

import "os"

type Config struct{ Port, DatabasePath, GinMode, CORSOrigins string }

func Load() Config {
	return Config{Port: getenv("PORT", "8080"), DatabasePath: getenv("DATABASE_PATH", "birrwise.db"), GinMode: getenv("GIN_MODE", "debug"), CORSOrigins: getenv("CORS_ORIGINS", "http://localhost:5173")}
}
func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
