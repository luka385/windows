package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/job-project/1/application"
	"github.com/luka385/job-project/1/domain"
)

type Handler struct {
	Serv application.ServicePort
}

func NewHandler(ser application.ServicePort) Handler {
	return Handler{Serv: ser}
}

func (h *Handler) Create(c *gin.Context) {
	var newUser domain.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := h.Serv.Create(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.Serv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.Status(http.StatusNoContent)
}
