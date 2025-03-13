package main

import (
	"log"
	"primer-api/application/usecase"
	dbs "primer-api/infra/dbs/mongodb"
	"primer-api/infra/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("userdb")

	repo := dbs.NewUserRepository(db)
	usecase := usecase.NewUseCase(repo)
	r := http.SetupsServer(usecase)

	r.Run(":8080")

}
