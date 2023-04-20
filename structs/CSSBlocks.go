package structs

type TypeScale struct {
	Base int
	Multiplier float64
	Shrink float64
	Scale string
	Clamps string
}

type CSSBlocks struct {
	FontScale TypeScale
}