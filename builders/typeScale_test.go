package builders

import (
	"reflect"
	"sandricoprovo/design-token-builder/structs"
	"testing"
)

var scale = 1.414
var base = 16
var shrink = 0.6

func TestTypeScaleBuilder(t *testing.T) {
	t.Run("should return a valid 1.414 type scale with 7 steps", func (t *testing.T) {
		var steps = structs.Steps{
			Small: 2,
			Large: 5,
		}
		var correctScale = []float64{8, 11.32, 16, 22.62, 31.99, 45.23, 63.96, 90.44}

		typeScale, err := BuildTypeScale(scale, steps, base, shrink)

		if err != nil {
			t.Fatalf("An error occurred while generating the type scale.")
		}

		if !reflect.DeepEqual(correctScale, typeScale) {
			t.Fatalf("The generated type scale is incorrect")
		}

	})

	t.Run("should return an empty scale if steps are 0", func (t *testing.T) {
		var steps = structs.Steps{
			Small: 0,
			Large: 0,
		}

		_, err := BuildTypeScale(scale, steps, base, shrink)

		if err == nil {
			t.Fatalf("The generated type scale is not empty")
		}

	})
}
