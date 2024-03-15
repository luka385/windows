package repository

import "github.com/luka385/job-project/3/internal/domain"

type Repository struct {
	repo domain.ItemRepositoryPort
}

func NewRepository(repo domain.ItemRepositoryPort) Repository {
	return Repository{repo: repo}
}
