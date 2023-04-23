package utils

import (
	"fmt"
	"sandricoprovo/denoken/structs"
	"testing"
)

var typeScale = structs.TypeScale {
	Base: 16,
	Multiplier: 1.141,
	Shrink: 0.6,
	Scale: "",
	Clamps: "",
}
var cssBlocks = structs.CSSBlocks {
	TypeScale: typeScale,
}

func TestCreateGlobalCssFile(t *testing.T) {
	t.Run("should return error if path is empty", func(t *testing.T) {
		var path = ""
		createFileErr := CreateGlobalCssFile(cssBlocks, path)

		if createFileErr == nil {
			fmt.Println("The returned error should be populated when passed an empty path.")
			t.Fail()
		}
	})
}