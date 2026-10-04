package main

import (
	"better_auth/internal/config"
	"better_auth/internal/service/repository/postgres"

	e "github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	e := e.New()

	db, err := postgres.InitPostgresConnection(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	e.Start(":8080")
}
