package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test5/application"
)

func SetupServer(ser application.ServicePort) *gin.Engine {
	r := gin.Default()

	h := NewHandler(ser)

	r.Group("/v1")
	{
		r.POST("/person", h.Create)
		r.DELETE("/person/:id", h.Delete)
	}
	return r
}
