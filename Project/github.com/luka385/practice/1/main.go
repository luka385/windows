package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mouse struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

var Mouses []Mouse

func main() {

	r := gin.Default()

	r.GET("/mouses", GetMouse)
	r.POST("/mouses", PostMouse)

	r.Run(":8080")
}

func GetMouse(c *gin.Context) {
	c.JSON(http.StatusOK, Mouses)
}

func PostMouse(c *gin.Context) {
	var NewMouse Mouse

	err := c.BindJSON(&NewMouse)
	if err != nil {
		log.Fatal(err)
	}

	Mouses = append(Mouses, NewMouse)

	c.IndentedJSON(http.StatusCreated, Mouses)
}
