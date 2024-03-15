package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test6/application"
	"github.com/luka385/hexa-test6/domain"
)

type Handler struct {
	Servi application.ServicePort
}

func NewHandler(ser application.ServicePort) Handler {
	return Handler{Servi: ser}
}

func (h *Handler) Create(c *gin.Context) {
	var u domain.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.Servi.CreateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.Servi.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}
