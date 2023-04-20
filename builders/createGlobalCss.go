package builders

import (
	"fmt"
	"os"
	"sandricoprovo/design-token-builder/structs"
	"text/template"
)

func CreateGlobalCSS(pieces structs.CSSBlocks) {
	file, fileErr := os.Create("global.css")
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}

	defer file.Close()

	template, templateErr := template.New("GlobalCSSFile").Parse(`
/*
	Generated Global CSS File
	Generator built in Go by Sandrico Provo https://www.sandricoprovo.dev/
*/

		:root {
			/* === TYPOGRAPHY === */
			/* Config */
			--scale: {{.FontScale.Multiplier}};
			--shrink: {{.FontScale.Shrink}};

			/* Scale */
			{{.FontScale.Scale}}

			/* NOTE: The vw unit used within the clamps below is a placeholder, so please tweak as needed for your use case. */
			/* Clamps */
			{{.FontScale.Clamps}}
		};
	`)
	if templateErr != nil {
		return
	}

	fileWriteErr := template.Execute(file, pieces)
	if fileWriteErr != nil {
		return
	}

	file.Close()
}