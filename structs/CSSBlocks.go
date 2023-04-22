package structs

// Type Scale Structs
type TypeScale struct {
	Base int
	Multiplier float64
	Shrink float64
	Scale string
	Clamps string
}

type TypeScaleConfig struct {
	Base int `mapstructure:"base"`
    Multiplier float64 `mapstructure:"multiplier"`
    Shrink float64 `mapstructure:"shrink"`
	Steps Steps `mapstructure:"steps"`
}

type CSSBlocks struct {
	TypeScale TypeScale
}