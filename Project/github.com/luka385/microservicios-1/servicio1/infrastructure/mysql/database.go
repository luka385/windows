package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:3858@tcp(localhost:3306)/agenda")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
