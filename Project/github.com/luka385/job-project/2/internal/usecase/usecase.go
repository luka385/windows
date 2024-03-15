package usecase

import (
	"fmt"

	"github.com/luka385/job-project/2/internal/domain"
	ctypes "github.com/luka385/job-project/2/internal/platform/custom-types"
)

const (
	noStock = iota
	Instock
)

const (
	activeStatus   = "ACTIVE"
	inactiveStatus = "INACTIVE"
)

type ItemUsecasePort interface {
	SaveItem(*domain.Item) (*domain.Item, error)
	GetAllItems() (domain.MapRepo, error)
	GetItemByID(domain.ID) (*domain.Item, error)
}

type ItemUsecase struct {
	repository domain.ItemRepositoyPort
}

func NewItemUsecase(repo domain.ItemRepositoyPort) ItemUsecasePort {
	return &ItemUsecase{
		repository: repo,
	}
}

func (u *ItemUsecase) SaveItem(item *domain.Item) (*domain.Item, error) {
	if ExistingItem, err := u.repository.GetItemByCode(item.Code); err == nil && ExistingItem != nil {
		return nil, fmt.Errorf("code must be unique %s", item.Code)
	}
	item.Status = inactiveStatus
	if item.Stock > noStock {
		item.Status = activeStatus
	}

	savedItem, err := u.repository.SaveItem(item)
	if err != nil {
		return nil, fmt.Errorf("error saving domain.Item: %w", err)
	}

	return savedItem, nil
}

func (u *ItemUsecase) GetAllItems() (domain.MapRepo, error) {
	items, err := u.repository.GetAllItem()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	if len(items) == 0 {
		return nil, ctypes.NewCustomType(1, ctypes.ErrItemNotFound)
	}
	return items, nil
}

func (u *ItemUsecase) GetItemByID(id domain.ID) (*domain.Item, error) {
	item, err := u.repository.GetItemByID(id)
	if err != nil {
		return nil, ctypes.NewCustomType(1, ctypes.ErrItemNotFound)
	}
	return item, nil
}
