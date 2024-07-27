package util

import (
	"fmt"
	"testing"
)

type normalizeHueTest struct {
	input    float64
	expected float64
}

func TestNormalizeHue(t *testing.T) {
	tests := []normalizeHueTest{
		{0, 0},
		{360, 0},
		{720, 0},
		{-360, 0},
		{180, 180},
		{-180, 180},
		{450, 90},
		{-450, 270},
		{1080, 0},
		{-1080, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Hue_%f", test.input), func(t *testing.T) {
			result := NormalizeHue(test.input)
			if result != test.expected {
				t.Errorf("For input %f, expected %f but got %f", test.input, test.expected, result)
			}
		})
	}
}
