package usecase

import (
	"Project/github.com/luka385/mocks2/application/ports"
	"Project/github.com/luka385/mocks2/domain"
)

type UserUseCase struct {
	repo ports.RepositoryPort
}

func NewUserUseCase(repo ports.RepositoryPort) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) GetUsers() ([]*domain.User, error) {
	return u.repo.GetUsers()
}

func (u *UserUseCase) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUseCase) GetUserById(id string) (*domain.User, error) {
	return u.repo.GetUserById(id)
}

func (u *UserUseCase) UpdateUser(id string, newUser *domain.User) error {
	return u.repo.UpdateUser(id, newUser)
}

func (u *UserUseCase) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}
