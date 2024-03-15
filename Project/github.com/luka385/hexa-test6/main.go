package main

import (
	"log"

	"github.com/luka385/hexa-test6/application"
	"github.com/luka385/hexa-test6/infrastructure/http"
	"github.com/luka385/hexa-test6/infrastructure/mysql"
)

func main() {
	db, err := mysql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := mysql.NewRepoMySQL(db)
	ser := application.NewService(repo)
	r := http.SetupServer(ser)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
