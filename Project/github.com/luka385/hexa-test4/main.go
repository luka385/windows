package main

import (
	"log"

	"github.com/luka385/hexa-test4/application"
	"github.com/luka385/hexa-test4/infrastructure/http"
	"github.com/luka385/hexa-test4/infrastructure/mysql"
)

func main() {
	db, err := mysql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := mysql.NewMySQLRepo(db)
	service := application.NewService(repo)
	r := http.SetupServer(service)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
