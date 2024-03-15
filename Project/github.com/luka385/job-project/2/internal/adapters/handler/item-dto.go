package handler

import "github.com/luka385/job-project/2/internal/domain"

type itemDTO struct {
	Code        string  `json:"code"`
	Tittle      string  `json:"tittle"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
}

func dtoToItem(dto *itemDTO) *domain.Item {
	return &domain.Item{
		Code:        dto.Code,
		Tittle:      dto.Tittle,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Status:      dto.Status,
	}
}
