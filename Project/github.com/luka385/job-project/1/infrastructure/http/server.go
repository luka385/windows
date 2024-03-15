package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/job-project/1/application"
)

func SetupServer(ser application.ServicePort) *gin.Engine {

	r := gin.Default()

	h := NewHandler(ser)

	gr := r.Group("/v2")
	{
		gr.POST("/person", h.Create)
		gr.DELETE("/person/:id", h.Delete)
	}
	return r
}
