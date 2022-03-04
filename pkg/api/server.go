package api

import (
	"beca"

	"github.com/go-playground/validator/v10"
)

type Server struct {
	Validate          *validator.Validate
	RepositoryService beca.RepositoryService
	ResultService     beca.ResultService
	JobChan           chan int
}
