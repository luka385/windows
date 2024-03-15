package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luka385/job-project/2/internal/adapters/handler/presenter"
	"github.com/luka385/job-project/2/internal/domain"
	ctypes "github.com/luka385/job-project/2/internal/platform/custom-types"
	"github.com/luka385/job-project/2/internal/usecase"
)

type ItemHandler struct {
	usecase usecase.ItemUsecasePort
}

func NewHandler(u usecase.ItemUsecasePort) *ItemHandler {
	return &ItemHandler{
		usecase: u,
	}
}

func (h *ItemHandler) SaveItem(c *gin.Context) {
	var dto itemDTO
	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item := dtoToItem(&dto)
	saveItem, err := h.usecase.SaveItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, presenter.Item(saveItem))
}

func (h *ItemHandler) GetAllItems(c *gin.Context) {
	items, err := h.usecase.GetAllItems()
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, presenter.Items(items))
}

func (h *ItemHandler) GetItemsByID(c *gin.Context) {
	id, err := stringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	item, err := h.usecase.GetItemByID(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, presenter.Item(item))
}

func stringToID(s string) (domain.ID, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return domain.ID(id), nil
}

func handleError(c *gin.Context, err error) {
	if err.Error() == ctypes.ErrItemNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
