package color

import (
	"testing"
)

type RgbTestCase struct {
	input     string
	expectedR int
	expectedG int
	expectedB int
	expectErr error
}

func TestRgbParser(t *testing.T) {
	tests := []RgbTestCase{
		{"rgb(255, 255, 255)", 255, 255, 255, nil},
		{"rgb(0, 0, 0)", 0, 0, 0, nil},
		{"rgb(128, 128, 128)", 128, 128, 128, nil},
		{"rgb(256, 0, 0)", 0, 0, 0, LogColorErr("r is out of range or missing")}, // Invalid RGB value
		{"rgb(255, 255)", 0, 0, 0, LogColorErr("invalid RGB format")},            // Invalid format
		{"not a color", 0, 0, 0, LogColorErr("invalid RGB format")},              // Invalid format
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			r, g, b, err := RgbColor{}.Parse(test.input)

			if (err != nil) && err.Error() != test.expectErr.Error() {
				t.Errorf("Expected error: %v, got: %v", test.expectErr, err)
			}
			if r != test.expectedR || g != test.expectedG || b != test.expectedB {
				t.Errorf("Expected (r, g, b): (%d, %d, %d), got: (%d, %d, %d)", test.expectedR, test.expectedG, test.expectedB, r, g, b)
			}
		})
	}
}

type toLchTest struct {
	input       string
	expectedL   float64
	expectedC   float64
	expectedH   float64
	expectedErr error
}

func TestRgbToLch(t *testing.T) {
	tests := []toLchTest{
		{"rgb(255, 255, 255)", 100.00, 0.000, 89.88, nil},
		{"rgb(28, 148, 60)", 58.46, 0.162, 147.20, nil},
		{"rgb(28, 148)", 0, 0, 0, LogColorErr("invalid RGB format")},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			l, c, h, err := RgbColor{}.ToLch(test.input)

			if (err != nil) && err.Error() != test.expectedErr.Error() {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}
			if l != test.expectedL || c != test.expectedC || h != test.expectedH {
				t.Errorf("Expected (l, c, h): (%b, %b, %b), got: (%b, %b, %b)", test.expectedL, test.expectedC, test.expectedH, l, c, h)
			}
		})
	}
}
