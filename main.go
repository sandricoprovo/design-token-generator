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

	configPaths := structs.ConfigPaths{
		File:  "denoken.config",
		Paths: []string{".", "./config/"},
	}

	// Loads config file settings
	config, configErr := utils.LoadConfig(configPaths)
	if configErr != nil {
		fmt.Println(configErr)
	}
	// Builds the type scale struct to be used for generating this block of css
	typeScale, scaleGeneratorErr := builders.BuildTypeScale(config.TypeScale)
	if scaleGeneratorErr != nil {
		log.Fatal(scaleGeneratorErr)
	}

	cssBlocks := structs.CSSBlocks{
		TypeScale: typeScale,
	}

	createFileErr := utils.CreateGlobalCssFile(cssBlocks, config.Path)
	if createFileErr != nil {
		log.Fatal("🛑 There was an issue generating the css file. Please try again.")
	}

	fmt.Println("⭐ Successfully created your css file at /" + config.Path)
}
