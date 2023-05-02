package builders

import (
	"github.com/sandricoprovo/denoken/structs"
	"strconv"
	"strings"
	"testing"
)

var DEFAULT_CONFIG = structs.TypeScaleConfig{
	Base:       16,
	Multiplier: 1.414,
	Shrink:     0.6,
	Steps: structs.Steps{
		Small: 2,
		Large: 5,
	},
}
var validScale = []float64{8, 11.32, 16, 22.62, 31.99, 45.23, 63.96, 90.44}

func TestTypeScaleBuilder(t *testing.T) {
	t.Run("should not return an empty string", func(t *testing.T) {
		_, err := BuildTypeScale(DEFAULT_CONFIG)
		if err != nil {
			t.Fatalf("An error occurred while generating the type scale.")
		}
	})

	t.Run("should contain type scale string with valid numbers", func(t *testing.T) {
		typeScale, _ := BuildTypeScale(DEFAULT_CONFIG)
		if !strings.Contains(typeScale.Scale, strconv.FormatFloat(validScale[0], 'f', -1, 64)) {
			t.Fatalf("That number wasn't within the scale")
		}
	})

	t.Run("should return a populated error if steps are zero", func(t *testing.T) {
		var config = structs.TypeScaleConfig{
			Base:       16,
			Multiplier: 1.414,
			Shrink:     0.6,
			Steps: structs.Steps{
				Small: 0,
				Large: 0,
			},
		}

		_, err := BuildTypeScale(config)

		if err == nil {
			t.Fatalf("Does not correctly report an error if steps are empty")
		}
	})

	t.Run("should return no font sizes smaller than base", func(t *testing.T) {
		var config = structs.TypeScaleConfig{
			Base:       16,
			Multiplier: 1.414,
			Shrink:     0.6,
			Steps: structs.Steps{
				Small: 0,
				Large: 5,
			},
		}

		typeScale, err := BuildTypeScale(config)

		if err != nil || strings.Contains(typeScale.Scale, "8") {
			t.Fail()
		}
	})

	t.Run("should return no font sizes larger than base", func(t *testing.T) {
		var config = structs.TypeScaleConfig{
			Base:       16,
			Multiplier: 1.414,
			Shrink:     0.6,
			Steps: structs.Steps{
				Small: 2,
				Large: 0,
			},
		}

		typeScale, err := BuildTypeScale(config)

		if err != nil || strings.Contains(typeScale.Scale, "63.96") {
			t.Fail()
		}
	})
}
