package utils

import (
	"testing"
)

func TestRoundToPrecision(t *testing.T) {
	t.Run("should return rounded number to correct precision", func(t *testing.T) {
		float := 15.9999
		precision := 0.01
		desiredNum := 16

		roundedFloat := RoundToPrecision(float, precision)

		if (desiredNum != int(roundedFloat)) {
			t.Fatalf("This test failed because %f is not equal to %d", roundedFloat, desiredNum)
		}
	})
	t.Run("nan should return the float passed in", func(t *testing.T) {
		float := 15.9999
		precision := 0.00

		roundedFloat := RoundToPrecision(float, precision)

		if float != roundedFloat {
			t.Fatalf("This test failed because %f is not equal to %f", roundedFloat, float)
		}
	})
}