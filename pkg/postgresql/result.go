package postgresql

import (
	"beca"
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ResultService struct {
	Ctx context.Context
	DB  *pgxpool.Pool
}

func (s *ResultService) Result(id int) (*beca.Result, error) {
	var r beca.Result

	err := pgxscan.Get(s.Ctx, s.DB, &r, "SELECT * FROM result WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &r, nil
}

func (s *ResultService) Results() ([]*beca.Result, error) {
	var rs []*beca.Result

	err := pgxscan.Select(s.Ctx, s.DB, &rs, "SELECT * FROM result ORDER BY id DESC")

	return rs, err
}

func (s *ResultService) CreateResult(dto beca.CreateResultDTO) (*beca.Result, error) {
	var id int

	err := s.DB.QueryRow(s.Ctx, "INSERT INTO result(status, repository_name, repository_url, queued_at) VALUES ($1, $2, $3, NOW()) RETURNING id", dto.Status, dto.RepositoryName, dto.RepositoryURL).Scan(&id)
	if err != nil {
		return nil, err
	}

	r, err := s.Result(id)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ResultService) UpdateResult(r *beca.Result) (err error) {
	tx, err := s.DB.Begin(s.Ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(s.Ctx)
		} else {
			tx.Commit(s.Ctx)
		}
	}()

	_, err = tx.Exec(s.Ctx, "UPDATE result SET status=$2, findings=$3, scanning_at=$4, finished_at=$5 WHERE id=$1", r.ID, r.Status, r.Findings, r.ScanningAt, r.FinishedAt)
	return err
}
