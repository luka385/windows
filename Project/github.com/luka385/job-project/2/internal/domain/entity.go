package domain

import "time"

type ID uint

type MapRepo map[ID]*Item

type Item struct {
	Code        string
	Tittle      string
	Description string
	Price       float64
	Stock       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemRepositoyPort interface {
	SaveItem(*Item) (*Item, error)
	GetAllItem() (MapRepo, error)
	GetItemByCode(string) (*Item, error)
	GetItemByID(ID) (*Item, error)
}
