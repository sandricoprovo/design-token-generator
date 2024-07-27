package color

import (
	"fmt"
	"testing"
)

type oklchToString struct {
	l      float64
	c      float64
	h      float64
	output string
	err    string
}

func TestOklchToString(t *testing.T) {
	tests := []oklchToString{
		{100.00, 0.00, 89.88, "oklch(100.00% 0.00 89.88)", ""},
		{58.46, 0.162, 147.20, "oklch(58.46% 0.162 147.2)", ""},
	}

	for _, test := range tests {
		title := fmt.Sprintf("(%b, %b, %b)", test.l, test.c, test.h)

		t.Run(title, func(t *testing.T) {
			oklchStr := OklchColor{}.ToString(test.l, test.c, test.h)

			if oklchStr == "N/A" {
				t.Errorf("Expected error: %v, got: %v", test.err, "N/A")
			}
		})
	}
}
