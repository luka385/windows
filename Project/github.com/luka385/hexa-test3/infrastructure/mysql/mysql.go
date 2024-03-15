package mysql

import (
	"database/sql"

	"github.com/luka385/hexa-test3/application"
	"github.com/luka385/hexa-test3/domain"
)

type MySQLPersonRepo struct {
	db *sql.DB
}

func NewMySQLPersonRepo(db *sql.DB) application.RepoPort {
	return &MySQLPersonRepo{
		db: db,
	}
}

func (r *MySQLPersonRepo) Create(p *domain.Person) error {
	Query := "INSERT INTO person(id_person, name_person, email_person) VALUES(?, ?, ?)"

	_, err := r.db.Exec(Query, p.ID, p.Name, p.Email)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLPersonRepo) Delete(id string) error {
	Query := "DELETE FROM person WHERE id_person = ?"

	_, err := r.db.Exec(Query, id)
	if err != nil {
		return err
	}

	return nil
}
