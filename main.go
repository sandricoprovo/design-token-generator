package main

import (
	"fmt"
	"log"
	"sandricoprovo/design-token-builder/builders"
	"sandricoprovo/design-token-builder/structs"
)

func main() {
	scale := 1.414
	base := 16
	steps := structs.Steps{
		Small: 2,
		Large: 5,
	}

	typeScale, err := builders.GenerateTypeScale(scale, steps, base)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(typeScale)
}