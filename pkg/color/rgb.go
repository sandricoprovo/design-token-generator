package color

import (
	"math"
	"strconv"

	"github.com/sandricoprovo/fran/pkg/util"
)

type RgbColor struct{}

var MAX_RANGE = 255

var Rgb RgbColor
var RgbParserMap = ParserMap{
	"oklch": Rgb.ToOklchString,
}

func (RgbColor) Parse(s string) (int, int, int, error) {
	// Ex result: [rgb(255, 255, 255) 255 255 255]
	matches := RGB_REGEX.FindStringSubmatch(s)

	if len(matches) != 4 {
		return 0, 0, 0, LogColorErr("invalid RGB format")
	}

	r, err := strconv.Atoi(matches[1])
	if err != nil || r > MAX_RANGE {
		return 0, 0, 0, LogColorErr("r is out of range or missing")
	}

	g, err := strconv.Atoi(matches[2])
	if err != nil || g > MAX_RANGE {
		return 0, 0, 0, LogColorErr("g is out of range or missing")
	}

	b, err := strconv.Atoi(matches[3])
	if err != nil || b > MAX_RANGE {
		return 0, 0, 0, LogColorErr("b is out of range or missing")
	}

	return r, g, b, nil
}

func (RgbColor) ToLinear(c int) float64 {
	cf := float64(c) / 255.0
	if cf <= 0.04045 {
		return cf / 12.92
	}
	return math.Pow((cf+0.055)/1.055, 2.4)
}

func (RgbColor) ToLabFromLinear(rgbR, rgbG, rgbB float64) (float64, float64, float64) {
	L := math.Cbrt(
		0.41222147079999993*rgbR + 0.5363325363*rgbG + 0.0514459929*rgbB,
	)
	M := math.Cbrt(
		0.2119034981999999*rgbR + 0.6806995450999999*rgbG + 0.1073969566*rgbB,
	)
	S := math.Cbrt(
		0.08830246189999998*rgbR + 0.2817188376*rgbG + 0.6299787005000002*rgbB,
	)

	l := 0.2104542553*L + 0.793617785*M - 0.0040720468*S
	a := 1.9779984951*L - 2.428592205*M + 0.4505937099*S
	b := 0.0259040371*L + 0.7827717662*M - 0.808675766*S

	return l, a, b
}

func (RgbColor) ToLch(s string) (float64, float64, float64, error) {
	r, g, b, err := Rgb.Parse(s)
	if err != nil {
		return 0.0, 0.0, 0.0, err
	}

	rl := Rgb.ToLinear(r)
	gl := Rgb.ToLinear(g)
	bl := Rgb.ToLinear(b)
	labL, labA, labB := Rgb.ToLabFromLinear(rl, gl, bl)
	lchL, lchC, lchH := Lab.ToLch(labL, labA, labB)

	return util.Round(lchL, 2), util.Round(lchC, 3), util.Round(lchH, 2), nil
}

func (RgbColor) ToOklchString(s string) string {
	l, c, h, err := Rgb.ToLch(s)
	if err != nil {
		util.PanicCheck(err)
	}

	return Lch.ToString(l, c, h)
}
