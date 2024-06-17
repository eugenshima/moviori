package main

import (
	"context"
	"fmt"

	"github.com/eugenshima/moviori/internal/handlers"
	repo "github.com/eugenshima/moviori/internal/repository"
	"github.com/eugenshima/moviori/internal/service"
	movieProto "github.com/eugenshima/moviori_movies/proto"
	"google.golang.org/grpc"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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

	MovieConn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer func() {
		err = MovieConn.Close()
		if err != nil {
			fmt.Println("error closing price service connection")
		}
	}()

	movieCLient := movieProto.NewMovioriMoviesClient(MovieConn)

	rps := repo.NewAuthRepository(dbpool)
	rpsm := repo.NewMovieRepository(movieCLient)
	srv := service.NewAuthService(rps, rpsm)
	hnd := handlers.NewAuthHandler(srv)

	auth := e.Group("/auth")
	{
		auth.POST("/login", hnd.Login)
		auth.POST("/signup", hnd.Signup)
	}

	movies := e.Group("/movies")
	{
		movies.POST("/getbyid", hnd.GetMovieByName)
	}

	// profile := e.Group("/profile")
	// {
	// 	profile.GET("/get-profile")
	// 	profile.PATCH("/update-profile")
	// }

	e.Logger.Fatal(e.Start(":8080"))
}
