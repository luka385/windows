package application

// tengo que hacer que la primera letra este en mayuscula
import (
	"github.com/luka385/hexa-test4/domain"
)

type Service struct {
	repo RepoPort
	veri Verificar
}

func NewService(r RepoPort) ServicePort {
	return &Service{repo: r}
}

func (s *Service) Create(p *domain.Person) error {
	pv := s.veri.VerificarPerson(p)
	return s.repo.CreatePerson(pv)
}

func (s *Service) Delete(id string) error {
	return s.repo.DeletePerson(id)
}
