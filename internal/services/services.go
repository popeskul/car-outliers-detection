package services

//go:generate mockgen -destination=mocks/mock_service.go -package=mocks github.com/popeskul/car-outliers-detection/internal/services IService
type IService interface {
	OutlierService() IOutlierService
}

type Service struct {
	outlierService *OutlierService
}

func (s *Service) OutlierService() IOutlierService {
	return s.outlierService
}

func NewService() (*Service, error) {
	return &Service{
		outlierService: NewOutlierService(),
	}, nil
}
