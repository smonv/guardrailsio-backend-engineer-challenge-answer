package api

import (
	"beca"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s Server) RepositoryIndex(c echo.Context) error {
	repositories, err := s.RepositoryService.Repositories()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, repositories)
}

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

func (s Server) RepositoryDelete(c echo.Context) error {
	idStr := c.Param("rid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = s.RepositoryService.DeleteRepository(id)

	return err
}
