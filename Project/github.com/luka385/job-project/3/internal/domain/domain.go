package domain

import "time"

type ID uint

type ItemMap map[*ID]item

type item struct {
	Code        string
	Tittle      string
	Description string
	Price       float64
	Stock       int
	Status      string
	CreateAt    time.Time
	UpdateAt    time.Time
}

type ItemRepositoryPort interface {
	SaveItem(*item) (*item, error)
	GetAllItem() (ItemMap, error)
	GetItemByCode(string) (*item, error)
	GetItemByID(ID) (*item, error)
}
