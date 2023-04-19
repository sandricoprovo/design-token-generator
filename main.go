package main

import (
	"log"
	"sandricoprovo/design-token-builder/builders"
	"sandricoprovo/design-token-builder/structs"
	"sandricoprovo/design-token-builder/utils"
)

func main() {
	scale := 1.414
	base := 16
	steps := structs.Steps{
		Small: 2,
		Large: 5,
	}

	typeScale, scaleGeneratorErr := builders.GenerateTypeScale(scale, steps, base)
	if scaleGeneratorErr != nil {
		log.Fatal(scaleGeneratorErr)
	}

	typeScaleCssString, scaleToStringErr := utils.ConvertTypeScaleToString(typeScale)
	if scaleToStringErr != nil {
		log.Fatal(scaleToStringErr)
	}

	cssBlocks := structs.CSSBlocks{
		FontScale: structs.TypeScale{
			Base: base,
			Scale: scale,
			CSS: typeScaleCssString,
		},
	}

	builders.CreateGlobalCSS(cssBlocks)
}