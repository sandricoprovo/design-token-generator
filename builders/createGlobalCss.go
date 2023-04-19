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
		:root {
			/* === TYPOGRAPHY === */
			/* Scale */
			--scale: {{.FontScale.Scale}};
			{{.FontScale.CSS}}

			/* Clamps */
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