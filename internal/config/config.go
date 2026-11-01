package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgreConfig struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresHost     string
	PostgresSSLMode  string
}

type RedisConfig struct {
	RedisHost string
	RedisPort string
	RedisPassword string
	RedisDB int
}

type AppConfig struct {
	AppPort         string
	AppEnv          string
	PostgreConfig PostgreConfig
	RedisConfig RedisConfig
}

func getEnvAsInt(key string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultVal
}

func LoadConfig() *AppConfig {
	_ = godotenv.Load()
	return &AppConfig{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
		PostgreConfig: PostgreConfig{
			PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
			PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
			PostgresDBName:   getEnv("POSTGRES_DBNAME", "mydb"),
			PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
			PostgresSSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		RedisConfig: RedisConfig{
			RedisHost:     getEnv("REDIS_HOST", "postgres"),
			RedisPort: getEnv("REDIS_PORT", "password"),
			RedisPassword:   getEnv("REDIS_PASSWORD", "mydb"),
			RedisDB:     getEnvAsInt("REDIS_DB", 0),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
