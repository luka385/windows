package http

import (
	"Project/github.com/luka385/mocks2/application/usecase"

	"github.com/gin-gonic/gin"
)

func SetupsServer(usecase *usecase.UserUseCase) *gin.Engine {
	r := gin.Default()

	h := NewHandler(usecase)

	v1 := r.Group("v1")
	{
		v1.GET("/users/:id", h.GetUserById)
		v1.GET("/users", h.GetUsers)
		v1.POST("/users", h.CreateUser)
		v1.DELETE("/users/:id", h.DeleteUser)
		v1.PUT("/users/:id", h.UpdateUser)
	}

	return r
}
