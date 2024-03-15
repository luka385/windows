package main

import (
	"log"

	"github.com/luka385/microservicios-1/servicio2/application"
	"github.com/luka385/microservicios-1/servicio2/infrastructure/http"
	"github.com/luka385/microservicios-1/servicio2/infrastructure/mysql"
)

func main() {
	db, err := mysql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := mysql.NewMySQL(db)
	ser := application.NewService(repo)
	r := http.SetupServer(ser)

	err = r.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
