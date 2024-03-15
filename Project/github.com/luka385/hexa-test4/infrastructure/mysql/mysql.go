package mysql

import (
	"database/sql"

	"github.com/luka385/hexa-test4/application"
	"github.com/luka385/hexa-test4/domain"
)

type MySQLRepo struct {
	db *sql.DB
}

func NewMySQLRepo(db *sql.DB) application.RepoPort {
	return &MySQLRepo{db: db}
}

func (r *MySQLRepo) CreatePerson(p *domain.Person) error {
	Query := "INSERT INTO person(id_person, name_person, email_person) VALUES(?, ? , ?)"
	_, err := r.db.Exec(Query, p.ID, p.Name, p.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepo) DeletePerson(id string) error {
	Query := "DELETE FROM person WHERE id_person = ?"

	_, err := r.db.Query(Query, id)
	if err != nil {
		return nil
	}
	return nil
}
