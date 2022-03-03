package api

import (
	"beca"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags     results
// @Summary  list result
// @Accept   json
// @Produce  json
// @Success  200  {array}  beca.Result
// @Failure  500
// @Router   /results [get]
func (s Server) ResultIndex(c echo.Context) error {
	results, err := s.ResultService.Results()
	if results == nil {
		return c.JSON(http.StatusOK, []*beca.Result{})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, results)
}
