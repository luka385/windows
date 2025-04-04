package application

import "github.com/luka385/hexa-test5/domain"

type RepoPort interface {
	CreatePerson(*domain.Person) error
	DeletePerson(string) error
}

type ServicePort interface {
	Create(*domain.Person) error
	Delete(string) error
}
