package main

import (
	"log"

	"Project/github.com/luka385/mocks2/application/usecase"
	dbs "Project/github.com/luka385/mocks2/infra/adapters/dbs/mongodb"
	"Project/github.com/luka385/mocks2/infra/adapters/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("userdb")

	repo := dbs.NewRepository(db)
	usecase := usecase.NewUserUseCase(repo)
	r := http.SetupsServer(usecase)

	r.Run(":8080")
}
