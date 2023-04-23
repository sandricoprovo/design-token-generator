package utils

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"sandricoprovo/denoken/structs"
	"strings"
)

func createFileAtPath(path string) (error) {
	var filePath string
	splitPath := strings.Split(path, "/")
	directories := splitPath[:len(splitPath) - 1]

	if path == "" {
		return errors.New("the path cannot be empty");
	}

	for _, directory := range directories {
		// Skips any empty values due to splitting the path
		if directory == "" {
			continue
		}

		// Creates each directory based off of the previous directories as parents
		filePath += directory + "/"

		// Attempts to create the folder at the deepest level of the path
		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			return errors.New("failed to create the set directory. Please try again");
		}
	}

	return nil
}

func CreateGlobalCssFile(pieces structs.CSSBlocks, path string) (error) {
	if path == "" {
		return errors.New("a valid path is needed to create the css file. please set a valid path")
	}

	// Conditionally tries to create the directory from the path argument
	if strings.Contains(path, "/") {
		err := createFileAtPath(path)
		if err != nil {
			return errors.New("there was an error creating directory structure. please try again")
		}
	}

	file, fileErr := os.Create(path)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return errors.New("there was an error creating that file. please try again")
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
		return errors.New("there was an error creating the css file. please try again")
	}

	// Writes the template to the new css file
	fileWriteErr := template.Execute(file, pieces)
	if fileWriteErr != nil {
		return errors.New("there was an error creating the css file. please try again")
	}

	file.Close()
	return nil;
}