package utils

import "math"

func RoundToPrecision(float, precision float64) float64 {
	roundedFloat := math.Round(float / precision) * precision

	if math.IsNaN(roundedFloat) {
		return float
	}

	return roundedFloat
}