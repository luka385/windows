package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luka385/job-project/2/internal/adapters/handler"
)

const port = ":8080"

func NewHTTPServer(h *handler.ItemHandler) error {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/items", h.SaveItem)
	v1.GET("/items", h.GetAllItems)
	v1.GET("/items/:id", h.GetItemsByID)

	log.Println("Server started at http://localhost" + port)

	err := router.Run(port)
	if err != nil {
		return err
	}
	return nil
}
