package config

import (
	"os"

	"github.com/joho/godotenv"
)

type FelxibleConfig struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresHost     string
	PostgresSSLMode  string
}

type AppConfig struct {
	AppPort         string
	AppEnv          string
	FleixibleConfig FelxibleConfig
}

func LoadConfig() *AppConfig {
	_ = godotenv.Load()
	return &AppConfig{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
		FleixibleConfig: FelxibleConfig{
			PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
			PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
			PostgresDBName:   getEnv("POSTGRES_DBNAME", "mydb"),
			PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
			PostgresSSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
