package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/microservicios-1/servicio2/application"
)

func SetupServer(ser application.UserSerPort) *gin.Engine {

	r := gin.Default()

	h := NewHandler(ser)

	r.GET("/userdata", h.GetAllTable)

	return r
}
