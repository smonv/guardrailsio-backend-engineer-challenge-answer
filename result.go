package beca

import "time"

type Result struct {
	ID             int        `json:"id"`
	Status         string     `json:"status"`
	RepositoryName string     `json:"repositoryName" db:"repository_name"`
	RepositoryURL  string     `json:"repositoryUrl" db:"repository_url"`
	Findings       []*Finding `json:"findings"`
	QueuedAt       *time.Time `json:"queuedAt" db:"queued_at"`
	ScanningAt     *time.Time `json:"scanningAt" db:"scanning_at"`
	FinishedAt     *time.Time `json:"finishedAt" db:"finished_at"`
}

type CreateResultDTO struct {
	Status         string `faker:"oneof: Queued, In Process, Success, Failure"`
	RepositoryName string `faker:"name"`
	RepositoryURL  string `faker:"url"`
}

type ResultService interface {
	Result(id int) (*Result, error)
	Results() ([]*Result, error)
	CreateResult(r CreateResultDTO) (*Result, error)
	UpdateResult(r *Result) error
}
