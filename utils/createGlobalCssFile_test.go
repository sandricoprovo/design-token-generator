package utils

import (
	"os"
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
			t.Fail()
		}
	})

	t.Run("should return error if path is has many characters but does not contain a file extension", func(t *testing.T) {
		var path = "   "
		createFileErr := CreateGlobalCssFile(cssBlocks, path)

		if createFileErr == nil {
			t.Fail()
		}
	})

	t.Run("should create css file within folder", func(t *testing.T) {
		var path = "styles/global.css"
		createFileErr := CreateGlobalCssFile(cssBlocks, path)

		if createFileErr != nil {
			t.Fail()
		}

		os.Remove("./styles/global.css")
		os.Remove("./styles/")
	})

	t.Run("should create css file at root of project", func(t *testing.T) {
		var path = "global.css"
		createFileErr := CreateGlobalCssFile(cssBlocks, path)

		if createFileErr != nil {
			t.Fail()
		}

		os.Remove("./global.css")
	})
}