package main

import (
	"better_auth/internal/adapter/repository/postgres"
	"better_auth/internal/adapter/repository/redis"
	"better_auth/internal/config"

	e "github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	e := e.New()

	db, err := postgres.InitPostgresConnection(cfg)
	if err != nil {
		panic(err)
	}

	err = redis.InitialRedisConnection(cfg)
	if err != nil {
		panic(err)
	}

	// this code will move to the container.go to centralize the dependency injection and to make the main.go cleaner and more maintainable
	// authRepo := postgres.NewAuthRepository(db)

	defer db.Close()
	e.Start(":8080")
}
