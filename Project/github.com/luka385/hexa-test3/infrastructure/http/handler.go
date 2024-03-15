package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test3/application"
	"github.com/luka385/hexa-test3/domain"
)

type Handler struct {
	ser application.ServicePort
}

func NewHandler(s application.ServicePort) Handler {
	return Handler{ser: s}
}

func (h *Handler) Create(c *gin.Context) {
	var p domain.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.ser.CreatePerson(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.ser.DeletePerson(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}
