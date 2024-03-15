package application

import "github.com/luka385/microservicios-1/servicio1/domain"

type RepoPort interface {
	Create(*domain.User) error
}

type ServicePort interface {
	CreateUser(*domain.User) error
}
