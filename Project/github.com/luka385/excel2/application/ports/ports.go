package ports

import (
	"context"

	"github.com/luka385/excel2/domain"
	"github.com/luka385/excel2/infra/adapters/http/dto"
)

type Usecases interface {
	Process(context.Context, []dto.ExcelPerson) ([]string, error)
}

type PersonRepository interface {
	SavePerson(context.Context, []domain.Person) ([]string, error)
}
