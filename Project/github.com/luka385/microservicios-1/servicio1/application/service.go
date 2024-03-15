package application

import "github.com/luka385/microservicios-1/servicio1/domain"

type Service struct {
	repo RepoPort
}

func NewService(repo RepoPort) ServicePort {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(usr *domain.User) error {
	return s.repo.Create(usr)
}
