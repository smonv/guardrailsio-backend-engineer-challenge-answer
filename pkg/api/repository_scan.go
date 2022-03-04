package api

import (
	"beca"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Tags     repositories
// @Summary  trigger a scan again a repository
// @Accept   json
// @Produce  json
// @Param    rid  path      int  true  "Repository ID"
// @Success  200  {object}  beca.Result
// @Failure  400
// @Failure  404
// @Failure  500
// @Router   /repositories/{rid}/scans [post]
func (s Server) RepositoryScanCreate(c echo.Context) error {
	idStr := c.Param("rid")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	repository, err := s.RepositoryService.Repository(id)
	if repository == nil {
		return echo.ErrNotFound
	}
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	dto := beca.CreateResultDTO{
		Status:         "Queued",
		RepositoryName: repository.Name,
		RepositoryURL:  repository.Url,
	}

	result, err := s.ResultService.CreateResult(dto)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	s.JobChan <- result.ID

	return c.JSON(http.StatusOK, result)
}
