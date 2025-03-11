package http

import (
	"primer-api/application/usecase"

	"github.com/gin-gonic/gin"
)

func SetupsServer(usecase *usecase.Usecase) *gin.Engine {
	r := gin.Default()

	h := NewHandler(usecase)

	v1 := r.Group("v1")
	{
		v1.GET("/users/:id", h.GetUserById)
		v1.POST("users", h.CreateUser)
	}

	return r
}
