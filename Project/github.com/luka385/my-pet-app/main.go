package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/my-pet-app/application"
	"github.com/luka385/my-pet-app/infrastructure/mongodb"
	"github.com/luka385/my-pet-app/interfaces/http"
)

func main() {
	db, err := mongodb.NewMongoDB("mongodb://localhost:27017", "petapp", "pets")
	if err != nil {
		panic(err)
	}

	petUsecase := application.NewPetUsecase(db)
	petHandler := http.NewPetHandler(petUsecase)

	r := gin.Default()

	r.POST("/pets", petHandler.CreatePet)
	r.GET("/pets/:id", petHandler.GetPet)
	r.PUT("/pets/:id", petHandler.UpdatePet)
	r.DELETE("/pets/:id", petHandler.DetelePet)

	r.Run(":8080")

}
