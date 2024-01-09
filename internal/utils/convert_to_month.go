package utils

import (
	"strconv"
	"strings"
)

func ConvertToMonths(age string) (int, bool) {
	parts := strings.Split(age, " ")
	if len(parts) != 2 {
		return 0, false
	}

	value, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, false
	}

	switch parts[1] {
	case "year", "years":
		return value * 12, true
	case "month", "months":
		return value, true
	case "day", "days":
		return value / 30, true // approximation
	default:
		return 0, false
	}
}
