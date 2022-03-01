package postgresql

import (
	"beca"
	"context"
	"errors"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryService struct {
	Ctx context.Context
	DB  *pgxpool.Pool
}

func (s *RepositoryService) Repository(id int) (*beca.Repository, error) {
	var r beca.Repository

	err := pgxscan.Get(s.Ctx, s.DB, &r, "SELECT id, name, url FROM repository WHERE id=$1", id)
	fmt.Println(err)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &r, nil
}

func (s *RepositoryService) Repositories() ([]*beca.Repository, error) {
	var repositories []*beca.Repository

	err := pgxscan.Select(s.Ctx, s.DB, &repositories, "SELECT * FROM repository")

	return repositories, err
}

func (s *RepositoryService) CreateRepository(r beca.CreateRepositoryDTO) (*beca.Repository, error) {
	var id int

	err := s.DB.QueryRow(s.Ctx, `INSERT INTO repository (name, url) VALUES ($1, $2) RETURNING id`, r.Name, r.Url).Scan(&id)
	if err != nil {
		return nil, err
	}

	repository, err := s.Repository(id)
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (s *RepositoryService) DeleteRepository(id int) error {
	_, err := s.DB.Exec(s.Ctx, "DELETE FROM repository WHERE id=$1", id)

	return err
}
