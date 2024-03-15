package mysql

import (
	"database/sql"

	"github.com/luka385/hexa-test6/application"
	"github.com/luka385/hexa-test6/domain"
)

type RepoMySQL struct {
	db *sql.DB
}

func NewRepoMySQL(db *sql.DB) application.RepoPort {
	return &RepoMySQL{db: db}
}

func (r *RepoMySQL) Create(u *domain.User) error {
	Query := "INSERT INTO person(id_person, name_person, email_person) VALUES(?, ? , ?)"

	_, err := r.db.Exec(Query, u.ID, u.Name, u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepoMySQL) Delete(id string) error {
	Query := "DELETE FROM person WHERE id_person = ?"

	_, err := r.db.Exec(Query, id)
	if err != nil {
		return err
	}

	return nil
}
