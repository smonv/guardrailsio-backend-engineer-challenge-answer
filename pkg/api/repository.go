package api

import (
	"beca"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Tags     repositories
// @Summary  list repository
// @Produce  json
// @Success  200  {object}  beca.Repository
// @Failure  404
// @Failure  500
// @Router   /repositories [get]
func (s Server) RepositoryIndex(c echo.Context) error {
	repositories, err := s.RepositoryService.Repositories()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, repositories)
}

// @Tags     repositories
// @Summary  get a repository
// @Produce  json
// @Param    rid  path     int  true  "Repository ID"
// @Success  200  {array}  beca.Repository
// @Failure  500
// @Router   /repositories/{rid} [get]
func (s Server) RepositoryShow(c echo.Context) error {
	idStr := c.Param("rid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	repository, err := s.RepositoryService.Repository(id)
	if repository == nil {
		return echo.ErrNotFound
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, repository)
}

// @Tags     repositories
// @Summary  create a repository
// @Accept   json
// @Produce  json
// @Param    repository  body      beca.CreateRepositoryDTO  true  "CreateRepositoryDTO"
// @Success  200         {object}  beca.Repository
// @Failure  400
// @Failure  500
// @Router   /repositories [post]
func (s Server) RepositoryCreate(c echo.Context) error {
	r := new(beca.Repository)

	if err := c.Bind(r); err != nil {
		return err
	}

	repository := beca.CreateRepositoryDTO{
		Name: r.Name, Url: r.Url,
	}

	newRepository, err := s.RepositoryService.CreateRepository(repository)
	if newRepository == nil {
		return echo.ErrInternalServerError
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newRepository)
}

// @Tags     repositories
// @Summary  delete a repository
// @Produce  json
// @Success  200
// @Failure  500
// @Router   /repositories [delete]
func (s Server) RepositoryDelete(c echo.Context) error {
	idStr := c.Param("rid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = s.RepositoryService.DeleteRepository(id)

	return err
}
