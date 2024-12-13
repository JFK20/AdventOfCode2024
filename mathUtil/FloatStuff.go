package mathUtil

import (
	"fmt"
	"math"
)

func IsNearInt(value float64) bool {
	epsilon := 1e-4
	nearestInt := math.Round(value)
	return math.Abs(value-nearestInt) <= epsilon
}

func GetNearestInt(value float64) (int, error) {
	if !IsNearInt(value) {
		return 0, fmt.Errorf("value is near %f", value)
	}
	return int(math.Round(value)), nil
}
