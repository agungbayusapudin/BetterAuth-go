package main

import (
	"better_auth/internal/config"
	"better_auth/internal/service/repository/postgres"
)

func main() {
	cfg := config.LoadConfig()

	db, err := postgres.InitPostgresConnection(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
