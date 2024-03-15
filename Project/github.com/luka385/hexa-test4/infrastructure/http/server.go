package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test4/application"
)

func SetupServer(service application.ServicePort) *gin.Engine {
	r := gin.Default()

	h := NewHandler(service)

	r.Group("v1")
	{
		r.POST("/person", h.CreatePerson)
		r.DELETE("/person/:id", h.DeletePerson)
	}

	return r
}
