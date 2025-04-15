package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luka385/excel2/application/usecases"
	"github.com/luka385/excel2/infra/adapters/excel"
)

type Handler struct {
	usecases usecases.UseCases
	excel    excel.ExcelAdapter
}

func NewHandler(uc usecases.UseCases, e excel.ExcelAdapter) *Handler {
	return &Handler{
		usecases: uc,
		excel:    e,
	}
}

func (h *Handler) UploadExcel(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error opening file"})
		return
	}
	defer f.Close()

	// llama al adaptador Excel
	persons, err := h.excel.ParseExcel(f)
	if err != nil {
		log.Printf("Error parsing excel: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal Error"})
		return
	}

	//llama al caso de uso
	if len(persons) > 0 {
		if errList, err := h.usecases.Procces(c.Request.Context(), persons); err != nil {
			log.Printf("Error processing data: %v", errList)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "error processing data",
				"details": errList,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "file processed successfully"})
	}
}
