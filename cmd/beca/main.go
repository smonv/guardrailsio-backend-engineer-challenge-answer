package main

import (
	"beca/pkg/api"
	"beca/pkg/postgresql"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "beca/docs"
)

// @title  BECA API
func main() {
	ctx := context.Background()

	dbPool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	repositoryService := &postgresql.RepositoryService{
		Ctx: ctx, DB: dbPool,
	}

	s := api.Server{RepositoryService: repositoryService}

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/repositories", s.RepositoryIndex)
	e.GET("/repositories/:rid", s.RepositoryShow)
	e.POST("/repositories", s.RepositoryCreate)
	e.DELETE("/repositories/:rid", s.RepositoryDelete)

	e.GET("swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start("localhost:3000"))
}
