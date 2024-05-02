package service

type Service struct {
	CalculatorService *Calculator
}

func NewService() *Service {
	return &Service{
		CalculatorService: NewCalculator(),
	}
}
