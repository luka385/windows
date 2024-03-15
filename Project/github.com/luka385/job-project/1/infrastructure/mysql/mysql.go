package mysql

import (
	"database/sql"

	"github.com/luka385/job-project/1/application"
	"github.com/luka385/job-project/1/domain"
)

type RepoMysql struct {
	db *sql.DB
}

func NewRepoMysql(db *sql.DB) application.RepoPort {
	return &RepoMysql{db: db}
}

func (rdb *RepoMysql) CreateUser(p *domain.User) error {
	Query := "INSERT INTO person(id_person, name_person, email_person) VALUES(?, ? ,?)"

	_, err := rdb.db.Exec(Query, p.ID, p.Name, p.Email)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RepoMysql) DeleteUser(id string) error {
	QUERY := "DELETE FROM person WHERE id_person = ?"
	_, err := rdb.db.Exec(QUERY, id)
	if err != nil {
		return err
	}
	return nil
}
