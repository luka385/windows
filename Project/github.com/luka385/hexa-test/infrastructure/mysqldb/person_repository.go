package mysqldb

import (
	"database/sql"

	"github.com/luka385/hexa-test/application"
	"github.com/luka385/hexa-test/domain"
)

type MySQLPersonRepository struct {
	db *sql.DB
}

func NewMySQLPersonRepository(db *sql.DB) application.RepoPort {
	return &MySQLPersonRepository{
		db: db,
	}
}

func (r *MySQLPersonRepository) CreatePerson(person *domain.Person) error {
	query := "INSERT INTO person(id_person, name_person , email_person) VALUES(?, ?, ?);"

	_, err := r.db.Exec(query, person.ID, person.Name, person.Email)

	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLPersonRepository) DeletePerson(id string) error {
	query := "DELETE FROM person WHERE id_person = ?"

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
