package ports

import "primer-api/domain"

type UseCasePort interface {
	GetUserByID(string) (*domain.User, error)
	CreateUser(*domain.User) error
}

type RepositoryPort interface {
	GetUserByID(string) (*domain.User, error)
	CreateUser(*domain.User) error
}
