package main

import (
	"github.com/luka385/hexa-test/application"
	"github.com/luka385/hexa-test/infrastructure/http"
	"github.com/luka385/hexa-test/infrastructure/mysqldb"
)

func main() {
	db, err := mysqldb.OpenDB()
	if err != nil {
		panic("ssss")
	}

	repo := mysqldb.NewMySQLPersonRepository(db)
	serv := application.NewService(repo)
	handler := http.NewHandler(serv)
	
	//
}