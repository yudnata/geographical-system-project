package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	Mode        string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		_ = godotenv.Load("../.env")
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", "secret"),
		Mode:        getEnv("MODE", "development"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
