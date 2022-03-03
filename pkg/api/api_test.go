package api

import (
	"beca"
	"beca/pkg/postgresql"
	"context"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
)

var repositoryService beca.RepositoryService
var resultService beca.ResultService

var server *Server

func TestMain(m *testing.M) {
	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("PGX_TEST_DATABASE"))
	if err != nil {
		os.Exit(1)
	}
	defer dbPool.Close()

	repositoryService = &postgresql.RepositoryService{
		Ctx: context.Background(),
		DB:  dbPool,
	}
	resultService = &postgresql.ResultService{
		Ctx: context.Background(),
		DB:  dbPool,
	}

	server = &Server{
		RepositoryService: repositoryService,
		ResultService:     resultService,
		Validate:          validator.New(),
	}

	code := m.Run()

	os.Exit(code)
}
