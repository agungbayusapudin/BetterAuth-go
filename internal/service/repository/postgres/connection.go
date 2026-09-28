package postgres

import (
	"database/sql"
	"fmt"
)

func initPostgresConnection(cfg *config.AppConfig) (*sql.DB, error) {
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

	db.SetMaxOpenConns(cfg.FleixibleConfig.PostgresMaxOpenConns)
	db.SetMaxIdleConns(cfg.FleixibleConfig.PostgresMaxIdleConns)
	db.SetConnMaxLifetime(cfg.FleixibleConfig.PostgresConnMaxLifetime)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	} else {
		fmt.Printf("postgress database connection pool initialized successfully!")
	}

	return db, nil

}
