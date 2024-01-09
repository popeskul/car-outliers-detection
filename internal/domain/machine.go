package domain

//go:generate mockgen -destination=mocks/mock_machine.go -package=mocks github.com/popeskul/car-outliers-detection/internal/domain MachineInterface
type MachineInterface interface {
	Validate() bool
}

type Machine struct {
	ID  string `json:"id"`
	Age string `json:"age"`
}

func (m *Machine) Validate() bool {
	if m == nil {
		return false
	}

	if m.ID == "" {
		return false
	}

	if m.Age == "" {
		return false
	}

	return true
}
