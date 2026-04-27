package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func load() Config {
	_ = godotenv.Load()

	return Config{
		Port:      getEnvOrDefault("PORT", "8080"),
		DBUrl:     os.Getenv("DB_URL"),
		SecretKey: os.Getenv("CONFIG_SECRET_KEY"),
	}
}

func validate(cfg Config) {
	if cfg.DBUrl == "" {
		log.Fatal("DB_URL is required")
	}
	if cfg.SecretKey == "" {
		log.Fatal("CONFIG_SECRET_KEY is required")
	}
}

func getEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
