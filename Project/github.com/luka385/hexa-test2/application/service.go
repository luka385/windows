package application

import "github.com/luka385/hexa-test2/domain"

type Service struct {
	repo PersonRepository
}

func NewService(repo PersonRepository) ServiceRepo {
	return &Service{repo: repo}
}

func (s *Service) Create(person *domain.Person) error {
	return s.repo.CreatePerson(person)
}

func (s *Service) Delete(id string) error {
	return s.repo.DeletePerson(id)
}
