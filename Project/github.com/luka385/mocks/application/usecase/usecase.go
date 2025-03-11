package usecase

import (
	"primer-api/application/ports"
	"primer-api/domain"
)

type Usecase struct {
	repo *ports.RepositoryPort
}

func NewUseCase(repo *ports.RepositoryPort) Usecase {
	return Usecase{repo: repo}
}

func (u *Usecase) GetUserByID(id string) (*domain.User, error) {
	return u.GetUserByID(id)
}

func (u *Usecase) CreateUser(user *domain.User) error {
	return u.CreateUser(user)
}
