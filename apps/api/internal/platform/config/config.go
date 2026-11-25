package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`
	PORT           string `env:"PORT" envDefault:"8080"`
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	DATABASE_URL string `env:"DATABASE_URL,required"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		LogLevel: getEnv("LOG_LEVEL", "info"),
		PORT:     getEnv("PORT", "8080"),
		DatabaseConfig: DatabaseConfig{
			DATABASE_URL: getEnv("DATABASE_URL", ""),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
