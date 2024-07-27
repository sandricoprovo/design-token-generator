package color

import (
	"fmt"

	"github.com/sandricoprovo/fran/pkg/util"
)

type OklchColor struct{}

var Lch OklchColor

func (OklchColor) ToString(l, c, h float64) string {
	if l < 0 || c < 0 || h < 0 {
		util.PanicCheck(LogColorErr("invalid l,c, or h value passed"))
	}

	preciseL := fmt.Sprintf("%.2f", l)
	preciseC := fmt.Sprintf("%.3f", c)
	preciseH := fmt.Sprintf("%.2f", h)

	return fmt.Sprintf("oklch(%v%% %v %v)", preciseL, preciseC, preciseH)
}
