package http

import (
	"Project/github.com/luka385/mocks2/application/usecase"
	"Project/github.com/luka385/mocks2/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.UserUseCase
}

func NewHandler(usecase *usecase.UserUseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.usecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.usecase.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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

	if err := h.usecase.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var UpdatedUser *domain.User

	err := c.ShouldBindJSON(&UpdatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.usecase.UpdateUser(id, UpdatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}
