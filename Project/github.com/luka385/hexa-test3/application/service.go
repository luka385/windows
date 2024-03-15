package application

import "github.com/luka385/hexa-test3/domain"

type Service struct {
	repo RepoPort
}

func NewService(repo RepoPort) ServicePort {
	return &Service{repo: repo}
}

func (s *Service) CreatePerson(p *domain.Person) error {
	return s.repo.Create(p)
}

func (s *Service) DeletePerson(id string) error {
	return s.repo.Delete(id)
}
