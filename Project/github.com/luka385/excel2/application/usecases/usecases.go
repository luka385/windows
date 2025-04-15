package usecases

import (
	"context"
	"fmt"

	"github.com/luka385/excel2/application/ports"
	"github.com/luka385/excel2/domain"
	"github.com/luka385/excel2/infra/adapters/http/dto"
)

type UseCases struct {
	repo ports.PersonRepository
}

func NewUseCases(r ports.PersonRepository) *UseCases {
	return &UseCases{repo: r}
}

func (uc *UseCases) Procces(ctx context.Context, dtos []dto.ExcelPerson) ([]string, error) {
	var errs []string
	var validPersons []domain.Person

	for i, dto := range dtos {
		person, err := domain.NewPerson(dto)
		if err != nil {
			errs = append(errs, fmt.Sprintf("fila %d, %v", i+1, err))
			continue
		}
		validPersons = append(validPersons, *person)
	}

	if len(validPersons) > 0 {
		if _, err := uc.repo.SavePerson(ctx, validPersons); err != nil {
			return errs, err
		}
	}

	return errs, nil
}
