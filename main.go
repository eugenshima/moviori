package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/eugenshima/moviori/internal/handlers"
	repo "github.com/eugenshima/moviori/internal/repository"
	"github.com/eugenshima/moviori/internal/service"
)

func PgxConnection(Url string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(Url)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	return dbpool, nil
}

func main() {
	e := echo.New()

	dbpool, err := PgxConnection("postgres://eugen:ur2qly1ini@localhost:5432/moviori")
	if err != nil {
		logrus.Printf("Failed to connect PotgreSQL: %v", err)
	}

	rps := repo.NewAuthRepository(dbpool)
	srv := service.NewAuthService(rps)
	hnd := handlers.NewAuthHandler(srv)

	auth := e.Group("/auth")
	auth.POST("/login", hnd.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
