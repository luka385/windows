package main

import (
	"log"

	"github.com/luka385/job-project/1/application"
	"github.com/luka385/job-project/1/infrastructure/http"
	"github.com/luka385/job-project/1/infrastructure/mysql"
)

func main() {
	db, err := mysql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := mysql.NewRepoMysql(db)
	ser := application.NewServices(repo)
	r := http.SetupServer(ser)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
