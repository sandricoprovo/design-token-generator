package util

import "math"

func NormalizeHue(h float64) float64 {
	hue := math.Mod(h, 360)

	if hue < 0 {
		return float64(hue + 360)
	} else {
		return float64(hue)
	}
}
