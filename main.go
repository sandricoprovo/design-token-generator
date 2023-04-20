package main

import (
	"log"
	"sandricoprovo/design-token-builder/builders"
	"sandricoprovo/design-token-builder/structs"
)

func main() {
	// TODO:
	// - Update these so they can ge picked up from a '.designtokenrc' file, allowing the config to be shared if needed. This isn't really needed, but it'll be fun!
	// - Convert type scale to use rems optionally
	scale := 1.414
	base := 16
	fontShrink := 0.8
	steps := structs.Steps{
		Small: 2,
		Large: 5,
	}

	typeScale, scaleGeneratorErr := builders.BuildTypeScale(scale, steps, base, fontShrink)
	if scaleGeneratorErr != nil {
		log.Fatal(scaleGeneratorErr)
	}

	cssBlocks := structs.CSSBlocks{
		FontScale: typeScale,
	}

	builders.CreateGlobalCSS(cssBlocks)
}