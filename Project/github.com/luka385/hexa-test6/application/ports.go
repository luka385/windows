package application

import "github.com/luka385/hexa-test6/domain"

type RepoPort interface {
	Create(*domain.User) error
	Delete(string) error
}

type ServicePort interface {
	CreateUser(*domain.User) error
	DeleteUser(string) error
}
