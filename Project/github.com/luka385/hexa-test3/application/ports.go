package application

import "github.com/luka385/hexa-test3/domain"

type RepoPort interface {
	Create(*domain.Person) error
	Delete(string) error
}

type ServicePort interface {
	CreatePerson(*domain.Person) error
	DeletePerson(string) error
}
