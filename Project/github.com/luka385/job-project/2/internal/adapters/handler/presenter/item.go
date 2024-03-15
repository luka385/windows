package presenter

import (
	"time"

	"github.com/luka385/job-project/2/internal/domain"
)

type jsonItem struct {
	Code        string    `json:"code"`
	Tittle      string    `json:"tittle"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

func Item(i *domain.Item) *jsonItem {
	var itemResponse jsonItem

	itemResponse.Code = i.Code
	itemResponse.Tittle = i.Tittle
	itemResponse.Description = i.Description
	itemResponse.Price = i.Price
	itemResponse.Stock = i.Stock
	itemResponse.Status = i.Status
	itemResponse.CreatedAt = i.CreatedAt
	itemResponse.UpdateAt = i.UpdatedAt

	return &itemResponse
}

func Items(items domain.MapRepo) map[domain.ID]*jsonItem {
	mJson := make(map[domain.ID]*jsonItem)
	for id, i := range items {
		mJson[id] = Item(i)
	}

	return mJson
}
