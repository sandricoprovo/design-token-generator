package util

import (
	"fmt"
	"testing"
)

type roundTest struct {
	value     float64
	precision int
	expected  float64
}

func TestRound(t *testing.T) {
	tests := []roundTest{
		{1.2345, 0, 1},
		{1.2345, 1, 1.2},
		{1.2345, 2, 1.23},
		{1.2345, 3, 1.235},
		{1.2345, 4, 1.2345},
		{-1.2345, 0, -1},
		{-1.2345, 1, -1.2},
		{-1.2345, 2, -1.23},
		{-1.2345, 3, -1.235},
		{-1.2345, 4, -1.2345},
		{0, 0, 0},
		{0.5, 0, 1},
		{-0.5, 0, -1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Round_%f_precision_%d", test.value, test.precision), func(t *testing.T) {
			result := Round(test.value, test.precision)
			if result != test.expected {
				t.Errorf("For value %f with precision %d, expected %f but got %f", test.value, test.precision, test.expected, result)
			}
		})
	}
}
