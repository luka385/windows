package main

import (
	"log"

	"github.com/luka385/microservicios-1/servicio1/application"
	"github.com/luka385/microservicios-1/servicio1/infrastructure/http"
	"github.com/luka385/microservicios-1/servicio1/infrastructure/mysql"
)

func main() {
	db, err := mysql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := mysql.NewRepoMySQL(db)
	service := application.NewService(repo)
	r := http.SetupServer(service)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
