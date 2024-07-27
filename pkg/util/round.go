package util

import "math"

func Round(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	rounded := math.Round(val*ratio) / ratio

	return rounded
}
