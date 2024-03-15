package application

type UserRepoPort interface {
	GetAllDataTable() ([]interface{}, error)
}

type UserSerPort interface {
	GetAllTable() ([]interface{}, error)
}
