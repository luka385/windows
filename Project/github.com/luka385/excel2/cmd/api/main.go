package main

import (
	"log"

	"github.com/luka385/excel2/application/usecases"
	dbs "github.com/luka385/excel2/infra/adapters/dbs/mongodb"
	"github.com/luka385/excel2/infra/adapters/excel"
	http "github.com/luka385/excel2/infra/adapters/http/handler"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("excel_test")

	repo := dbs.NewPersonRepository(db)
	usecase := usecases.NewUseCases(repo)
	excelAdapter := excel.NewExcelAdapter()
	r := http.SetupServer(*usecase, *excelAdapter)

	r.Run(":8080")

}
