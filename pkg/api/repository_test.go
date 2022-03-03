package api

import (
	"beca"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryCreate(t *testing.T) {
	e := echo.New()

	t.Run("create with invalid url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"echo", "url":"echo"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.RepositoryCreate(c)
		assert.Equal(t, echo.ErrBadRequest, err)
	})

	t.Run("create with missing name", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"url":"https://github.com/labstack/echo"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.RepositoryCreate(c)
		assert.Equal(t, echo.ErrBadRequest, err)
	})

	t.Run("create ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "echo", "url":"https://github.com/labstack/echo"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := server.RepositoryCreate(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body)

		var repository *beca.Repository
		err = json.Unmarshal(rec.Body.Bytes(), &repository)
		assert.NoError(t, err)
		assert.NotNil(t, repository)
		assert.Equal(t, repository.Name, "echo")
	})
}

func TestRepositoryIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/repositories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := server.RepositoryIndex(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)

	var repositories []*beca.Repository
	err = json.Unmarshal(rec.Body.Bytes(), &repositories)
	assert.NoError(t, err)
	assert.NotEmpty(t, repositories)
}
