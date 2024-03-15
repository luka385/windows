package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test6/application"
)

func SetupServer(s application.ServicePort) *gin.Engine {
	r := gin.Default()

	h := NewHandler(s)

	v := r.Group("/v1")
	{
		v.POST("/user", h.Create)
		v.DELETE("/user/:id", h.Delete)
	}

	return r
}
