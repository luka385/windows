package ports

import "Project/github.com/luka385/mocks2/domain"

type RepositoryPort interface {
	GetUsers() ([]*domain.User, error)
	GetUserById(string) (*domain.User, error)
	CreateUser(*domain.User) error
	UpdateUser(string, *domain.User) error
	DeleteUser(string) error
}
