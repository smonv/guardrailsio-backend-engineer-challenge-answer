package beca

type Repository struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CreateRepositoryDTO struct {
	Name string `json:"name" faker:"name"`
	Url  string `json:"url" faker:"url"`
}

type UpdateRepositoryDTO struct {
	Name *string `json:"name" faker:"name"`
	Url  *string `json:"url" faker:"url"`
}

type RepositoryService interface {
	Repository(id int) (*Repository, error)
	Repositories() ([]*Repository, error)
	CreateRepository(r CreateRepositoryDTO) (*Repository, error)
	UpdateRepository(r *Repository) error
	DeleteRepository(id int) error
}
