package utils

import (
	"github.com/sandricoprovo/denoken/structs"
	"os"
	"testing"
)

var typeScale = structs.TypeScale{
	Base:       16,
	Multiplier: 1.141,
	Shrink:     0.6,
	Scale:      "",
	Clamps:     "",
}
var cssBlocks = structs.CSSBlocks{
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

	t.Run("should return error if path has characters but no file extension", func(t *testing.T) {
		var path = "global"
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

	t.Run("should trim whitespace from path and create file", func(t *testing.T) {
		var path = "     styles/global.css   "
		createFileErr := CreateGlobalCssFile(cssBlocks, path)

		if createFileErr != nil {
			t.Fail()
		}

		os.Remove("./styles/global.css")
		os.Remove("./styles/")
	})

	t.Run("should return error when path is empty", func(t *testing.T) {
		var path = ""
		createFileErr := CreateGlobalCssFile(cssBlocks, path)

		if createFileErr == nil {
			t.Fail()
		}
	})
}
