package excel

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/luka385/excel2/infra/adapters/http/dto"
	"github.com/xuri/excelize/v2"
)

type ExcelAdapter struct{}

func NewExcelAdapter() *ExcelAdapter {
	return &ExcelAdapter{}
}

func (a *ExcelAdapter) ParseExcel(r io.Reader) ([]dto.ExcelPerson, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, fmt.Errorf("error opening excel: %w", err)
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, fmt.Errorf("error reading rows: %w", err)
	}

	var (
		persons []dto.ExcelPerson
		errs    []string
	)
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 4 {
			errs = append(errs, fmt.Sprintf("incomplete %d row", i+1))
			continue
		}

		age, err := strconv.Atoi(row[2])
		if err != nil {
			errs = append(errs, fmt.Sprintf("edad inválida en fila %d: %v", i+1, err))
			continue
		}

		person := dto.ExcelPerson{
			FirstName: row[0],
			LastName:  row[1],
			Age:       age,
			Phone:     row[3],
		}

		persons = append(persons, person)

		if len(errs) > 0 {
			return persons, fmt.Errorf("error in the file:\n%s", strings.Join(errs, "\n"))
		}
	}

	return persons, nil
}
