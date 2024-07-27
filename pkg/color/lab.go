package color

import (
	"math"

	"github.com/sandricoprovo/fran/pkg/util"
)

// This is a comment

type LabColor struct{}

var Lab LabColor

func (LabColor) ToLch(l, a, b float64) (float64, float64, float64) {
	c := math.Sqrt(a*a + b*b)
	h := util.NormalizeHue((math.Atan2(b, a) * 180) / math.Pi)

	return (l * 100), c, h
}

func (LabColor) ToLchString(l, a, b float64) string {
	l, c, h := Lab.ToLch(l, a, b)
	return Lch.ToString(l, c, h)
}
