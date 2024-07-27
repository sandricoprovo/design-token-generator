package color

import (
	"fmt"
	"strconv"

	"github.com/sandricoprovo/fran/pkg/util"
)

type HexColor struct{}

var Hex HexColor
var HexParserMap = ParserMap{
	"rgb":   Hex.ToRgbString,
	"oklch": Hex.ToOklchString,
}

func (HexColor) ToRgb(h string) (int, int, int, error) {
	if len(h) <= 0 {
		return 0.0, 0.0, 0.0, LogColorErr("cant convert invalid hex to oklch")
	}

	if h[0] == '#' {
		h = h[1:]
	}
	r, _ := strconv.ParseInt(h[0:2], 16, 64)
	g, _ := strconv.ParseInt(h[2:4], 16, 64)
	b, _ := strconv.ParseInt(h[4:6], 16, 64)

	return int(r), int(g), int(b), nil
}

func (HexColor) ToRgbString(h string) string {
	r, g, b, err := Hex.ToRgb(h)
	if err != nil {
		util.PanicCheck(err)
	}

	return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
}

func (HexColor) ToLch(h string) (float64, float64, float64, error) {
	if len(h) <= 0 {
		return 0.0, 0.0, 0.0, LogColorErr("cant convert invalid hex to oklch")
	}

	r, g, b, err := Hex.ToRgb(h)
	if err != nil {
		util.PanicCheck(err)
	}

	rl := Rgb.ToLinear(r)
	gl := Rgb.ToLinear(g)
	bl := Rgb.ToLinear(b)
	labL, labA, labB := Rgb.ToLabFromLinear(rl, gl, bl)
	lchL, lchC, lchH := Lab.ToLch(labL, labA, labB)

	return util.Round(lchL, 2), util.Round(lchC, 3), util.Round(lchH, 2), nil
}

func (HexColor) ToOklchString(h string) string {
	okL, okC, okH, err := Hex.ToLch(h)
	if err != nil {
		util.PanicCheck(err)
	}

	return Lch.ToString(okL, okC, okH)
}
