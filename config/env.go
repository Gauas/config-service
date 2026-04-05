package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func load() AppConfig {
	_ = godotenv.Load()

	return AppConfig{
		Port:      getEnvOrDefault("PORT", "8080"),
		DBUrl:     os.Getenv("DB_URL"),
		SecretKey: os.Getenv("CONFIG_SECRET_KEY"),
	}
}

func validate(appConfig AppConfig) {
	if appConfig.DBUrl == "" {
		log.Fatal("DB_URL is required")
	}
	if appConfig.SecretKey == "" {
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
