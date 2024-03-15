package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/microservicios-1/servicio2/application"
)

type Handler struct {
	UserSer application.UserSerPort
}

func NewHandler(UserSer application.UserSerPort) Handler {
	return Handler{UserSer: UserSer}
}

func (h *Handler) GetAllTable(c *gin.Context) {
	data, err := h.UserSer.GetAllTable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, data)
}
