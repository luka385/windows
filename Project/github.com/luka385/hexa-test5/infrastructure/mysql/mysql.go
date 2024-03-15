package mysql

import (
	"database/sql"

	"github.com/luka385/hexa-test5/application"
	"github.com/luka385/hexa-test5/domain"
)

type MySqlRepo struct {
	db *sql.DB
}

func NewMySqlRepo(db *sql.DB) application.RepoPort {
	return MySqlRepo{db: db}
}

func (r *MySqlRepo) CreatePerson(p *domain.Person) error {
	Query := "INSERT INTO person(id_person, name_person, email_person) VALUES(?, ?, ?)"

	if _, err := r.db.Exec(Query, p.ID, p.Name, p.Email); err != nil {
		return nil
	}
	return nil
}

func (r *MySqlRepo) DeletePerson(id string) error {
	Query := "SELECT FROM person WHERE id_person=?"
	if _, err := r.db.Exec(Query, id); err != nil {
		return nil
	}

	return nil
}
