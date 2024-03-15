package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/hexa-test/application"
	"github.com/luka385/hexa-test/domain"
)

type HandlerRepo struct {
	handler application.ServicePort
}

func NewHandler(h application.ServicePort) *HandlerRepo {
	return &HandlerRepo{handler: h}
}

func (hr *HandlerRepo) Create(c *gin.Context) {
	var person domain.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := hr.handler.Create(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (hr *HandlerRepo) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := hr.handler.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)

}
