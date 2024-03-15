package mysql

import (
	"database/sql"

	"github.com/luka385/microservicios-1/servicio2/application"
)

type MySQLRepo struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) application.UserRepoPort {
	return &MySQLRepo{db: db}
}

func (r *MySQLRepo) GetAllDataTable() ([]interface{}, error) {
	rows, err := r.db.Query("SELECT * FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []interface{}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	for rows.Next() {

		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		rowValues := make([]interface{}, len(values))
		for i, v := range values {
			rowValues[i] = *v.(*interface{})
		}

		results = append(results, rowValues)
	}

	return results, nil

}
