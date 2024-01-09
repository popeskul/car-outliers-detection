package services

import (
	"fmt"
	"github.com/popeskul/car-outliers-detection/internal/domain"
	"github.com/popeskul/car-outliers-detection/internal/utils"
)

//go:generate mockgen -destination=mocks/mock_outlier_service.go -package=mocks github.com/popeskul/car-outliers-detection/internal/services IOutlierService
type IOutlierService interface {
	DetectOutliers(machines []*domain.Machine) ([]*domain.Machine, error)
}

type OutlierService struct {
	Machines []domain.Machine
}

func NewOutlierService() *OutlierService {
	return &OutlierService{}
}

func (s *OutlierService) DetectOutliers(machines []*domain.Machine) ([]*domain.Machine, error) {
	if s == nil {
		return nil, fmt.Errorf("outlier service is nil")
	}

	if machines == nil {
		return nil, fmt.Errorf("machines is nil")
	}

	if len(machines) == 0 {
		return nil, fmt.Errorf("machines is empty")
	}

	maxAgeMonths := 100 * 12 // Максимальный допустимый возраст в месяцах (например, 100 лет)

	var outliers []*domain.Machine
	for _, machine := range machines {
		if !machine.Validate() {
			return nil, fmt.Errorf("invalid machine: %v", machine)
		}

		ageInMonths, valid := utils.ConvertToMonths(machine.Age)
		if !valid {
			return nil, fmt.Errorf("invalid age format: %s", machine.Age)
		}

		if ageInMonths > maxAgeMonths {
			outliers = append(outliers, machine)
		}
	}

	return outliers, nil
}
