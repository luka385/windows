package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test3/application"
)

func SetupServer(service application.ServicePort) *gin.Engine {
	r := gin.Default()

	UH := NewHandler(service)

	v := r.Group("/v1")
	{
		v.POST("/person", UH.Create)
		v.DELETE("/person/:id", UH.Delete)
	}
	return r
}
