package mysql

import (
	"database/sql"

	"github.com/luka385/microservicios-1/servicio1/application"
	"github.com/luka385/microservicios-1/servicio1/domain"
)

type RepoMySQL struct {
	db *sql.DB
}

func NewRepoMySQL(db *sql.DB) application.RepoPort {
	return &RepoMySQL{db: db}
}

func (r *RepoMySQL) Create(user *domain.User) error {
	Query := "INSERT INTO usuarios(nombre, edad) VALUES(? ,?)"

	_, err := r.db.Exec(Query, user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}
