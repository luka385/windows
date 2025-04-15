package domain

import (
	"errors"

	"github.com/luka385/excel2/infra/adapters/http/dto"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Phone     string
}

func NewPerson(dto dto.ExcelPerson) (*Person, error) {
	if dto.Age < 0 || dto.Age > 120 {
		return nil, errors.New("invalid age")
	}
	if dto.FirstName == "" || dto.LastName == "" {
		return nil, errors.New("firstname & lastname required")
	}

	return &Person{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       dto.Age,
		Phone:     dto.Phone,
	}, nil
}
