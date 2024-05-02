package service

type Service struct {
	TaxLevel TaxLevelService
}

func New() *Service {
	return &Service{
		TaxLevel: &taxLevelService{},
	}
}
