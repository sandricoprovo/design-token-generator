package builders

import (
	"errors"
	"sandricoprovo/design-token-builder/structs"
	"sandricoprovo/design-token-builder/utils"
	"sort"
)

func generateSmallFontSizes(scale float64, base int, steps int) ([]float64, error) {
	var smallFontSizes []float64

	if steps == 0 {
		return smallFontSizes, nil
	}

	previousSize := float64(base) // converts int to float64

	for i := 0; i < steps; i++ {
		size := previousSize / scale
		smallFontSizes = append(smallFontSizes, size)

		// Update the previous font size
		previousSize = size
	};

	return smallFontSizes, nil
}

func generateLargeFontSizes[T int | float64](scale float64, base T, steps int) ([]float64, error) {
	var largeFontSizes []float64

	if steps == 0 {
		return largeFontSizes, nil
	}

	previousSize := float64(base)

	for i := 0; i < steps; i++ {
		size := previousSize * scale
		largeFontSizes = append(largeFontSizes, size)

		previousSize = size
	}

	return largeFontSizes, nil
}

func GenerateTypeScale(scale float64, steps structs.Steps, baseFontSize int) ([]float64, error) {
	if steps.Large == 0 && steps.Small == 0 {
		return nil, errors.New("the small and large steps can't both be zero")
	}

	initialScale := []float64{float64(baseFontSize)}

	belowBaseSizes, err := generateSmallFontSizes(scale, baseFontSize, steps.Small)

	if err != nil {
		return nil, err
	}

	largerBaseSizes, err := generateLargeFontSizes(scale, baseFontSize, steps.Large)

	if err != nil {
		return nil, err
	}

	// Concats the three slices and sorts then in ascending order
	fontScale := append(initialScale, belowBaseSizes...) // appends smaller fonts
	fontScale = append(fontScale, largerBaseSizes...) // appends larger fonts
	sort.Float64s(fontScale)

	// Rounds each float to the nearest .05 and fixes to a number of decimal points
	for i, size := range fontScale {
		roundedSize := utils.RoundToPrecision(size, 0.01)
		fontScale[i] = utils.ToFixed(roundedSize, 2)
	}

	return fontScale, nil;
}