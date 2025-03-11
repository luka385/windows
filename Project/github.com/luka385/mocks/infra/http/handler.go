package http

import (
	"net/http"
	"primer-api/application/usecase"
	"primer-api/domain"

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
	id := c.Param("id")

	user, err := h.Usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var newUser *domain.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
