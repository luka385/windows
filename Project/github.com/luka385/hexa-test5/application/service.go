package application

import "github.com/luka385/hexa-test5/domain"

type Service struct {
	repo RepoPort
}

func NewService(repo RepoPort) ServicePort {
	return &Service{repo: repo}
}

func (s *Service) Create(p *domain.Person) error {
	return s.repo.CreatePerson(p)
}

func (s *Service) Delete(id string) error {
	return s.repo.DeletePerson(id)
}
