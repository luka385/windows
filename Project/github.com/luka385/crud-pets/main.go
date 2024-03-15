package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/crud-pets/application"
	"github.com/luka385/crud-pets/infrastructure/mongodb"
	"github.com/luka385/crud-pets/interfaces/http"
)

func main() {
	db, err := mongodb.NewMongoDB("mongodb://localhost:27017", "petapp", "pets")
	if err != nil {
		panic(err)
	}

	petUsecase := application.NewPetUsecase(db)
	ph := http.NewPetHandler(petUsecase)

	r := gin.Default()

	r.POST("/pets", ph.CreatePet)
	r.GET("/pets/:id", ph.GetPet)
	r.PUT("/pets/:id", ph.UpdatePet)
	r.DELETE("/pets/:id", ph.DeletePet)

	r.Run(":8080")
}
