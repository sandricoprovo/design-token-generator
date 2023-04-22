package main

import (
	"fmt"
	"log"
	"sandricoprovo/denoken/builders"
	"sandricoprovo/denoken/structs"
	"sandricoprovo/denoken/utils"
)

func main() {
	// TODO:
	// - Convert type scale to use rems optionally

	// Loads config file settings
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	// Builds the type scale struct to be used for generating this block of css
	typeScale, scaleGeneratorErr := builders.BuildTypeScale(config.TypeScale)
	if scaleGeneratorErr != nil {
		log.Fatal(scaleGeneratorErr)
	}

	cssBlocks := structs.CSSBlocks{
		TypeScale: typeScale,
	}

	builders.CreateGlobalCSS(cssBlocks)
}