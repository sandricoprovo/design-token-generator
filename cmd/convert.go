/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sandricoprovo/fran/pkg/color"
	"github.com/sandricoprovo/fran/pkg/file"
	"github.com/spf13/cobra"
)

var format string
var path string
var exts = []string{".css", ".scss", ".svelte"} // supported extensions

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert colors to different color formats.",
	Long: `Convert colors to different color formats. Currently supports converting:
- oklch: from hex, rgb
- rgb: from hex
	`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: check if unsupported format
		// panic("Sorry, that color format is not supported. Conversions currently support: rgb().")

		parser := func() func(string) string {
			return func(l string) string {
				return color.LineParser(format, l)
			}
		}

		if path != "" {
			file.ParseFile(path, parser())
		} else {
			file.ParseFileRecursive(exts, parser())
		}
	},
}

func init() {
	convertCmd.Flags().StringVarP(&format, "format", "f", "hex", "Format to convert too.")
	convertCmd.Flags().StringVarP(&path, "path", "p", "", "Path to css file.")

	rootCmd.AddCommand(convertCmd)
}
