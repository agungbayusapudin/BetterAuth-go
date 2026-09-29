package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"better_auth/internal/config"

	_ "github.com/lib/pq"
)

func InitPostgresConnection(cfg *config.AppConfig) (*sql.DB, error) {
	var err error
	var db *sql.DB

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		cfg.FleixibleConfig.PostgresUser,
		cfg.FleixibleConfig.PostgresPassword,
		cfg.FleixibleConfig.PostgresDBName,
		cfg.FleixibleConfig.PostgresHost,
		cfg.FleixibleConfig.PostgresSSLMode,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(300 * time.Second)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	} else {
		fmt.Printf("postgress database connection pool initialized successfully!")
	}

	return db, nil

}
