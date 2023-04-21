package builders

import (
	"fmt"
	"os"
	"sandricoprovo/denoken/structs"
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
	Denoken - CSS Design Token Generator
	Built by Sandrico Provo https://www.sandricoprovo.dev/
*/

		:root {
			/* === TYPOGRAPHY === */
			/* Config */
			--scale: {{.TypeScale.Multiplier}};
			--shrink: {{.TypeScale.Shrink}};

			/* Scale */
			{{.TypeScale.Scale}}

			/* NOTE: The vw unit used within the clamps below is a placeholder, so please tweak as needed for your use case. */
			/* Clamps */
			{{.TypeScale.Clamps}}
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