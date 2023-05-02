package builders

import (
	"errors"
	"fmt"
	"sort"

	"github.com/sandricoprovo/denoken/structs"
	"github.com/sandricoprovo/denoken/utils"
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
	}

	return smallFontSizes, nil
}

func generateLargeFontSizes(scale float64, base float64, steps int) ([]float64, error) {
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

func convertTypeScaleToString(scale []float64) (string, error) {
	if scale == nil {
		error := errors.New("that type scale is empty")
		return "", error
	}

	baseKey := 100
	scaleString := ""

	for _, size := range scale {
		response := fmt.Sprintf("--font-%d: %gpx;\n", baseKey, size)
		scaleString += response

		// Increases the key value
		baseKey += 100
	}

	return scaleString, nil
}

func generateCssClamps(scale []float64, base int) (string, error) {
	if scale == nil {
		return "", errors.New("the font scale cannot be empty")
	}

	clampsBlockString := ""

	for i, size := range scale {
		previousSizeKey := i
		sizeKey := i + 1

		if size < float64(base) {
			// Avoids clamp for font sizes smaller than base
			cssString := fmt.Sprintf("--fs-%d00: var(--font-%d00);\n", sizeKey, sizeKey)
			clampsBlockString += cssString
		} else if size == float64(base) {
			// Adds comment to tag base font size with camps
			cssString := fmt.Sprintf("--fs-%d00: var(--font-%d00); /* Base */\n", sizeKey, sizeKey)
			clampsBlockString += cssString
		} else {
			// Adds CSS clamp to font sizes larger than base
			clampString := fmt.Sprintf("--fs-%d00: clamp(calc(var(--font-%d00) * var(--scale) * var(--shrink)), 12vw, var(--font-%d00));\n", sizeKey, previousSizeKey, sizeKey)
			clampsBlockString += clampString
		}
	}

	return clampsBlockString, nil
}

func BuildTypeScale(typeScaleConfig structs.TypeScaleConfig) (structs.TypeScale, error) {
	baseSize := typeScaleConfig.Base
	scaleMultiplier := typeScaleConfig.Multiplier
	fontShrink := typeScaleConfig.Shrink
	scaleSteps := typeScaleConfig.Steps

	if scaleSteps.Small == 0 && scaleSteps.Large == 0 {
		return structs.TypeScale{}, errors.New("the small and large steps can't both be zero")
	}

	initialScale := []float64{float64(baseSize)}

	belowBaseSizes, belowSizeErr := generateSmallFontSizes(scaleMultiplier, baseSize, scaleSteps.Small)
	if belowSizeErr != nil {
		return structs.TypeScale{}, belowSizeErr
	}

	largerBaseSizes, largeSizeError := generateLargeFontSizes(scaleMultiplier, float64(baseSize), scaleSteps.Large)
	if largeSizeError != nil {
		return structs.TypeScale{}, largeSizeError
	}

	// Concats the three slices and sorts then in ascending order
	fontScale := append(initialScale, belowBaseSizes...) // appends smaller fonts
	fontScale = append(fontScale, largerBaseSizes...)    // appends larger fonts
	sort.Float64s(fontScale)

	// Rounds each float to the nearest .05 and fixes to a number of decimal points
	for i, size := range fontScale {
		roundedSize := utils.RoundToPrecision(size, 0.01)
		fontScale[i] = utils.ToFixed(roundedSize, 2)
	}

	// Creates the type scale string to be added to the css file
	scaleCssBlock, scaleToStringErr := convertTypeScaleToString(fontScale)
	if scaleToStringErr != nil {
		return structs.TypeScale{}, scaleToStringErr
	}

	// Creates the css clamps based on the type scale
	clampsCssBlock, clampsErr := generateCssClamps(fontScale, baseSize)
	if clampsErr != nil {
		return structs.TypeScale{}, clampsErr
	}

	// Forms the returned typeScale struct
	typeScale := structs.TypeScale{
		Base:       baseSize,
		Multiplier: scaleMultiplier,
		Shrink:     fontShrink,
		Scale:      scaleCssBlock,
		Clamps:     clampsCssBlock,
	}

	return typeScale, nil
}
