package color

import (
	"testing"
)

type TestCase struct {
	toFormat string
	input    string
	expected string
}

var testCases = []TestCase{
	{"lch", "", ""},
	{"rgb", "#FFFFFF", "rgb(255,255,255)"},
	{"lch", "#FFFFFF", "oklch(100.00% 0.000 89.88)"},
	{"lch", "rgb(255, 255, 255)", "oklch(100.00% 0.000 89.88)"},
}

func TestLineParser(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			result := LineParser(test.toFormat, test.input)
			if result != test.expected {
				t.Errorf("Expected %q but got %q", test.expected, result)
			}
		})
	}
}
