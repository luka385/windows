package application

import "github.com/luka385/hexa-test6/domain"

type Service struct {
	Repo RepoPort
}

func NewService(repo RepoPort) ServicePort {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(u *domain.User) error {
	return s.Repo.Create(u)
}

func (s *Service) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}
