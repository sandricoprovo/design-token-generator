package utils

import (
	"errors"
	"fmt"
)


func ConvertTypeScaleToString(scale []float64) (string, error) {
	if scale == nil {
		error:= errors.New("that type scale is empty")
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