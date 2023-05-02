/*
Copyright © 2023 Sandrico Provo sprovo@outlook.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/sandricoprovo/denoken/builders"
	"github.com/sandricoprovo/denoken/structs"
	"github.com/sandricoprovo/denoken/utils"
	"log"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a css file with design tokens.",
	Long:  "Generates a css file within design tokens within your current directory if a config is present.",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
