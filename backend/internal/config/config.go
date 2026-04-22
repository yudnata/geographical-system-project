package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	Mode             string
	DefaultAvatarURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		_ = godotenv.Load("../.env")
	}

	return &Config{
		Port:             getEnv("PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		JWTSecret:        getEnv("JWT_SECRET", "secret"),
		Mode:             getEnv("MODE", "development"),
		DefaultAvatarURL: getEnv("DEFAULT_AVATAR_URL", "https://res.cloudinary.com/di7pw938v/image/upload/v1706764503/user_3_pvyc8x.png"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
