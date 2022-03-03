package postgresql

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var repositoryService *RepositoryService
var resultService *ResultService

func TestMain(m *testing.M) {
	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("PGX_TEST_DATABASE"))
	if err != nil {
		os.Exit(1)
	}

	repositoryService = &RepositoryService{
		Ctx: context.Background(),
		DB:  dbPool,
	}
	resultService = &ResultService{
		Ctx: context.Background(),
		DB:  dbPool,
	}

	code := m.Run()

	dbPool.Close()

	os.Exit(code)
}
