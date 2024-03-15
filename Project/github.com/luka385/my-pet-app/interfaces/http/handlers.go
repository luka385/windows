package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/my-pet-app/application"
	"github.com/luka385/my-pet-app/domain"
)

type PetHandler struct {
	usecase *application.PetUsecase
}

func NewPetHandler(usecase *application.PetUsecase) *PetHandler {
	return &PetHandler{
		usecase: usecase,
	}
}

func (h *PetHandler) CreatePet(c *gin.Context) {
	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.CreatePet(&pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pet)
}

func (h *PetHandler) GetPet(c *gin.Context) {
	petID := c.Param("id")

	pet, err := h.usecase.GetPetByID(petID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) UpdatePet(c *gin.Context) {
	petID := c.Param("id")

	var updatedPet domain.Pet
	if err := c.ShouldBindJSON(&updatedPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.UpdatePet(petID, &updatedPet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPet)
}

func (h *PetHandler) DetelePet(c *gin.Context) {
	petID := c.Param("id")

	err := h.usecase.DeletePet(petID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ped Deleted successfully"})
}
