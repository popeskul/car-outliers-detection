package utils

import (
	"testing"
)

func TestConvertToMonths(t *testing.T) {
	testCases := []struct {
		age      string
		expected int
		valid    bool
	}{
		{"10 years", 120, true},
		{"1 year", 12, true},
		{"6 months", 6, true},
		{"3 days", 0, true}, // 3 days approximated to 0 months
		{"invalid", 0, false},
		{"2", 0, false},
		{"year", 0, false},
	}

	for _, tc := range testCases {
		result, valid := ConvertToMonths(tc.age)
		if result != tc.expected || valid != tc.valid {
			t.Errorf("ConvertToMonths(%s): expected %d, %v, got %d, %v", tc.age, tc.expected, tc.valid, result, valid)
		}
	}
}
