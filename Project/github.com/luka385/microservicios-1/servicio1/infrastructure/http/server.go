package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/microservicios-1/servicio1/application"
)

func SetupServer(service application.ServicePort) *gin.Engine {

	r := gin.Default()

	h := NewHandler(service)

	r.POST("/user", h.CreateUser)

	return r
}
