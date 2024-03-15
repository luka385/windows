package application

import "github.com/luka385/hexa-test2/domain"

type PersonRepository interface {
	CreatePerson(*domain.Person) error
	DeletePerson(string) error
}

type ServiceRepo interface {
	Create(*domain.Person) error
	Delete(string) error
}
