package utils

import "testing"

func TestToFixed(t *testing.T) {
	t.Run("float should be precision of 2", func (t *testing.T) {
		num := 1.23456789
		precision := 2
		desiredNum := 1.23

		fixedNumber := ToFixed(num, precision)

		if (fixedNumber != desiredNum) {
			t.Fatalf("That test failed because %f is not equal to %f", fixedNumber, desiredNum)
		}
	})

	t.Run("float is correctly fixed if precision is more than decimal points", func (t *testing.T) {
		num := 2.0
		precision := 2
		desiredNum := 2.0

		fixedNumber := ToFixed(num, precision)

		if (fixedNumber != desiredNum) {
			t.Fatalf("That test failed because %f is not equal to %f", fixedNumber, desiredNum)
		}
	})

	t.Run("float is returned at the same if precision is 0", func (t *testing.T) {
		num := 1.0
		precision := 0
		desiredNum := 1.0

		fixedNumber := ToFixed(num, precision)

		if (fixedNumber != desiredNum) {
			t.Fatalf("That test failed because %f is not equal to %f", fixedNumber, desiredNum)
		}
	})
}