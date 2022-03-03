package postgresql

import (
	"beca"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRepositoryAndGet(t *testing.T) {
	var cr = beca.CreateRepositoryDTO{}
	err := faker.FakeData(&cr)
	require.NoError(t, err)

	repository, err := repositoryService.CreateRepository(cr)
	require.NoError(t, err)
	require.NotNil(t, repository)

	assert.Equal(t, repository.Name, cr.Name)
	assert.Equal(t, repository.Url, cr.Url)

	rRepository, err := repositoryService.Repository(repository.ID)
	require.NoError(t, err)
	require.NotNil(t, rRepository)

	assert.Equal(t, rRepository.ID, repository.ID)
}

func TestCreateRepositoryAndList(t *testing.T) {
	var cr = beca.CreateRepositoryDTO{}
	err := faker.FakeData(&cr)
	require.NoError(t, err)

	repository, err := repositoryService.CreateRepository(cr)
	require.NoError(t, err)
	require.NotNil(t, repository)

	assert.Equal(t, repository.Name, cr.Name)
	assert.Equal(t, repository.Url, cr.Url)

	repositories, err := repositoryService.Repositories()
	require.NoError(t, err)
	require.NotNil(t, repository)
	require.NotEmpty(t, repositories)

	rRepository := repositories[0]

	assert.Equal(t, rRepository.ID, repository.ID)
}

func TestCreateRepositoryAndUpdate(t *testing.T) {
	var cr = beca.CreateRepositoryDTO{}
	err := faker.FakeData(&cr)
	require.NoError(t, err)

	repository, err := repositoryService.CreateRepository(cr)
	require.NoError(t, err)
	require.NotNil(t, repository)

	newName := faker.Name()
	newUrl := faker.URL()

	repository.Name = newName

	err = repositoryService.UpdateRepository(repository)
	require.NoError(t, err)

	rRepository, err := repositoryService.Repository(repository.ID)
	require.NoError(t, err)
	require.NotNil(t, rRepository)

	assert.Equal(t, rRepository.Name, newName)
	assert.Equal(t, rRepository.Url, repository.Url)
	assert.NotEqual(t, rRepository.Url, newUrl)
}

func TestCreateRepositoryAndDelete(t *testing.T) {
	var cr = beca.CreateRepositoryDTO{}
	err := faker.FakeData(&cr)
	require.NoError(t, err)

	repository, err := repositoryService.CreateRepository(cr)
	require.NoError(t, err)
	require.NotNil(t, repository)

	assert.Equal(t, repository.Name, cr.Name)
	assert.Equal(t, repository.Url, cr.Url)

	err = repositoryService.DeleteRepository(repository.ID)
	require.NoError(t, err)

	rRepository, err := repositoryService.Repository(repository.ID)
	require.NoError(t, err)
	require.Nil(t, rRepository)
}
