package dto

type ExcelPerson struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
}

/*func (ep *ExcelPerson) ToDomain() domain.Person {
	return domain.Person{
		FirstName: ep.FirstName,
		LastName:  ep.LastName,
		Age:       ep.Age,
		Phone:     ep.Phone,
	}
}*/
