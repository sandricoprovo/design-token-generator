package utils

import "math"

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// Fixes float numbers to the number of decimals based on the precision
// EX: ToFixed(1.2345, 2) = 1.23
func ToFixed(num float64, precision int) float64 {
	fixedNum := math.Pow(10, float64(precision))
	return float64(round(num * fixedNum)) / fixedNum
}