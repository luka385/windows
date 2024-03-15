package mysql

import "database/sql"

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:3858@tcp(localhost:3306)/superm")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
