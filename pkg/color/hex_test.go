package color

import (
	"fmt"
	"testing"
)

type MockRgb struct{}

var RgbMock MockRgb

func (MockRgb) ToLinear(value int) float64 {
	return float64(value) / 255.0
}

func (MockRgb) ToLabFromLinear(rl, gl, bl float64) (float64, float64, float64) {
	return rl * 100, gl * 100, bl * 100
}

type MockLab struct{}

var LabMock MockLab

func (MockLab) ToLch(l, a, b float64) (float64, float64, float64) {
	return l, a, b
}

func (MockLab) ToString(l, c, h float64) string {
	return fmt.Sprintf("lch(%.2f, %.3f, %.2f)", l, c, h)
}

type hexToRgbTest struct {
	hex       string
	expectedR int
	expectedG int
	expectedB int
}

func TestHexToRgb(t *testing.T) {
	tests := []hexToRgbTest{
		{"#ff5733", 255, 87, 51},
		{"ff5733", 255, 87, 51},
		{"#000000", 0, 0, 0},
		{"#ffffff", 255, 255, 255},
	}

	for _, test := range tests {
		t.Run(test.hex, func(t *testing.T) {
			r, g, b, _ := Hex.ToRgb(test.hex)
			if r != test.expectedR || g != test.expectedG || b != test.expectedB {
				t.Errorf("Expected (r, g, b): (%d, %d, %d), got: (%d, %d, %d)", test.expectedR, test.expectedG, test.expectedB, r, g, b)
			}
		})
	}
}

type hexToRgbStringTest struct {
	hex      string
	expected string
}

func TestHexToRgbString(t *testing.T) {
	tests := []hexToRgbStringTest{
		{"#ff5733", "rgb(255,87,51)"},
		{"ff5733", "rgb(255,87,51)"},
		{"#000000", "rgb(0,0,0)"},
		{"#ffffff", "rgb(255,255,255)"},
	}

	for _, test := range tests {
		t.Run(test.hex, func(t *testing.T) {
			result := Hex.ToRgbString(test.hex)
			if result != test.expected {
				t.Errorf("Expected: %s, got: %s", test.expected, result)
			}
		})
	}
}

type hexToLchTest struct {
	hex       string
	expectedL float64
	expectedC float64
	expectedH float64
	err       error
}

func TestHexToLch(t *testing.T) {
	tests := []hexToLchTest{
		{"#ff5733", 68.04, 0.210, 33.69, nil},
		{"#000000", 0.0, 0.0, 0.0, nil},
		{"#ffffff", 100.0, 0.000, 89.88, nil},
		{"", 0.0, 0.0, 0.0, LogColorErr("cant convert invalid hex to oklch")},
	}

	for _, test := range tests {
		t.Run(test.hex, func(t *testing.T) {
			l, c, h, err := Hex.ToLch(test.hex)

			if (err != nil) && err.Error() != test.err.Error() {
				t.Errorf("Expected error: %v, got: %v", test.err, err)
			}
			if l != test.expectedL || c != test.expectedC || h != test.expectedH {
				t.Errorf("Expected (L, C, H): (%.2f, %.3f, %.2f), got: (%.2f, %.3f, %.2f)", test.expectedL, test.expectedC, test.expectedH, l, c, h)
			}
		})
	}
}

type hexToOklchStringTest struct {
	hex      string
	expected string
}

func TestHexToLchString(t *testing.T) {
	tests := []hexToOklchStringTest{
		{"#ff5733", "oklch(68.04% 0.210 33.69)"},
		{"#000000", "oklch(0.00% 0.000 0.00)"},
		{"#ffffff", "oklch(100.00% 0.000 89.88)"},
	}

	for _, test := range tests {
		t.Run(test.hex, func(t *testing.T) {
			result := Hex.ToOklchString(test.hex)
			if result != test.expected {
				t.Errorf("Expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
