package http

import (
	"github.com/gin-gonic/gin"
	"github.com/luka385/excel2/application/usecases"
	"github.com/luka385/excel2/infra/adapters/excel"
)

func SetupServer(uc usecases.UseCases, e excel.ExcelAdapter) *gin.Engine {
	r := gin.Default()

	h := NewHandler(uc, e)

	v1 := r.Group("v1")
	{
		v1.POST("/upload", h.UploadExcel)
	}

	return r
}
