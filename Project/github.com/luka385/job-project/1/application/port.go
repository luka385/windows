package application

import "github.com/luka385/job-project/1/domain"

type ServicePort interface {
	Create(*domain.User) error
	Delete(string) error
}

type RepoPort interface {
	CreateUser(*domain.User) error
	DeleteUser(string) error
}
