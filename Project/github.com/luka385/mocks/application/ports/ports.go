package ports

import "primer-api/domain"

type UseCasePort interface {
	GetUserByID(string) (*domain.User, error)
	CreateUser(*domain.User) error
}

type RepositoryPort interface {
	GetById(string) (*domain.User, error)
	Create(*domain.User) error
}
