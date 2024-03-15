package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test4/application"
	"github.com/luka385/hexa-test4/domain"
)

type Handler struct {
	ser application.ServicePort
}

func NewHandler(ser application.ServicePort) Handler {
	return Handler{ser: ser}
}

func (h *Handler) CreatePerson(c *gin.Context) {
	var p domain.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.ser.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (h *Handler) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	if err := h.ser.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
