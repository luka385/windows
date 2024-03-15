package application

import "github.com/luka385/hexa-test4/domain"

type RepoPort interface {
	CreatePerson(*domain.Person) error
	DeletePerson(string) error
}

type ServicePort interface {
	Create(*domain.Person) error
	Delete(string) error
}

type Verificar interface {
	VerificarPerson(*domain.Person) *domain.Person
}
