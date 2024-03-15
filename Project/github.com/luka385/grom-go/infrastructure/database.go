package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/luka385/grom-go/domain"
)

func SetupDataBase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "mytaskapp.db")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.Task{})

	return db, nil
}
