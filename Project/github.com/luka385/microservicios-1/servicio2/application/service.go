package application

type Service struct {
	repo UserRepoPort
}

func NewService(repo UserRepoPort) UserSerPort {
	return &Service{repo: repo}
}

func (s *Service) GetAllTable() ([]interface{}, error) {
	data, err := s.repo.GetAllDataTable()
	if err != nil {
		return nil, err
	}

	return data, nil
}
