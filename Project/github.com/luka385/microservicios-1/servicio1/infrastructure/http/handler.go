package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/microservicios-1/servicio1/application"
	"github.com/luka385/microservicios-1/servicio1/domain"
)

type Handler struct {
	handler application.ServicePort
}

func NewHandler(handler application.ServicePort) Handler {
	return Handler{handler: handler}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.handler.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
