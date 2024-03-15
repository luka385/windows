package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test5/application"
	"github.com/luka385/hexa-test5/domain"
)

type Handler struct {
	s application.ServicePort
}

func NewHandler(s application.ServicePort) Handler {
	return Handler{s: s}
}

func (h *Handler) Create(c *gin.Context) {
	var p domain.Person

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.s.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.s.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}
