package services

import (
	"github.com/golang/mock/gomock"
	"github.com/popeskul/car-outliers-detection/internal/domain"
	"testing"
)

func TestOutlierService_DetectOutliers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		machines      []*domain.Machine
		expectedCount int
		expectError   bool
	}{
		{
			name: "Valid Machines Within Age Limit",
			machines: []*domain.Machine{
				{ID: "1", Age: "60 months"}, // Equivalent to 5 years
				{ID: "2", Age: "96 months"}, // Equivalent to 8 years
			},
			expectedCount: 0,
			expectError:   false,
		},
		{
			name: "Machines With Outliers",
			machines: []*domain.Machine{
				{ID: "1", Age: "1440 months"}, // Equivalent to 120 years
				{ID: "2", Age: "120 months"},  // Equivalent to 10 years
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "Invalid Machines",
			machines: []*domain.Machine{
				{ID: "1", Age: "1440 months"}, // Invalid: 120 years
				{ID: "2", Age: "1 month"},     // Valid
				{ID: "3", Age: "10 months"},   // Valid
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "Invalid age format",
			machines: []*domain.Machine{
				{ID: "4", Age: "10"}, // Invalid
			},
			expectedCount: 0,
			expectError:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewOutlierService()

			outliers, err := service.DetectOutliers(tc.machines)
			if (err != nil) != tc.expectError {
				t.Errorf("Test '%s' failed: expected error status %v, got %v", tc.name, tc.expectError, err != nil)
			}
			if len(outliers) != tc.expectedCount {
				t.Errorf("Test '%s' failed: expected %d outliers, got %d", tc.name, tc.expectedCount, len(outliers))
			}
		})
	}
}
