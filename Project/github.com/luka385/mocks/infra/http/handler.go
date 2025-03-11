package http

import (
	"primer-api/application/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) Handler {
	return Handler{Usecase: usecase}
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.
}
