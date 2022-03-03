package postgresql

import (
	"beca"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateResultAndGet(t *testing.T) {
	cr := beca.CreateResultDTO{}
	err := faker.FakeData(&cr)
	require.NoError(t, err)

	result, err := resultService.CreateResult(cr)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, result.RepositoryName, cr.RepositoryName)
	assert.Equal(t, result.RepositoryURL, cr.RepositoryURL)

	rResult, err := resultService.Result(result.ID)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, rResult.ID, result.ID)
}

func TestCreateResultAndList(t *testing.T) {
	cr := beca.CreateResultDTO{}
	err := faker.FakeData(&cr)
	require.NoError(t, err)

	result, err := resultService.CreateResult(cr)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, result.RepositoryName, cr.RepositoryName)
	assert.Equal(t, result.RepositoryURL, cr.RepositoryURL)

	rResults, err := resultService.Results()
	require.NoError(t, err)
	require.NotNil(t, rResults)
	require.NotEmpty(t, rResults)

	rResult := rResults[0]

	assert.Equal(t, rResult.ID, result.ID)
}
