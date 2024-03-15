package application

import "github.com/luka385/job-project/1/domain"

type Service struct {
	Repo RepoPort
}

func NewServices(repo RepoPort) ServicePort {
	return &Service{Repo: repo}
}

func (s *Service) Create(u *domain.User) error {
	return s.Create(u)
}

func (s *Service) Delete(st string) error {
	return s.Delete(st)
}
