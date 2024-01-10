package services

import (
	"fmt"
	"github.com/popeskul/car-outliers-detection/internal/domain"
	"github.com/popeskul/car-outliers-detection/internal/utils"
	"sync"
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

	maxAgeMonths := 100 * 12 // Maximum acceptable age in months (e.g., 100 years)

	var outliers []*domain.Machine
	var mutex sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	for _, machine := range machines {
		wg.Add(1)
		go func(machine *domain.Machine) {
			defer wg.Done()

			isOutlier, err := checkMachineOutlier(machine, maxAgeMonths)
			if err != nil {
				select {
				case errChan <- err: // Send error to channel, non-blocking
				default: // If channel already has an error, do not block
				}
				return
			}

			if isOutlier {
				mutex.Lock()
				outliers = append(outliers, machine)
				mutex.Unlock()
			}
		}(machine)
	}

	wg.Wait()
	close(errChan)

	// Check if there were any errors
	if err, ok := <-errChan; ok {
		return nil, err
	}

	return outliers, nil
}

func checkMachineOutlier(machine *domain.Machine, maxAgeMonths int) (bool, error) {
	if !machine.Validate() {
		return false, fmt.Errorf("invalid machine: %v", machine)
	}

	ageInMonths, valid := utils.ConvertToMonths(machine.Age)
	if !valid {
		return false, fmt.Errorf("invalid age format: %s", machine.Age)
	}

	return ageInMonths > maxAgeMonths, nil
}
