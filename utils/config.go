package utils

import (
	"sandricoprovo/denoken/structs"

	"github.com/spf13/viper"
)

type Config struct {
	Path      string                  `mapstructure:"path"`
	TypeScale structs.TypeScaleConfig `mapstructure:"typeScale"`
}

func LoadConfig(configInfo structs.ConfigPaths) (Config, error) {
	var config Config
	vp := viper.New()

	// Finds the config file
	vp.SetConfigName(configInfo.File)
	vp.SetConfigType("json")
	// Loops over and adds the default config paths
	for _, path := range configInfo.Paths {
		vp.AddConfigPath(path)
	}

	configErr := vp.ReadInConfig()
	if configErr != nil {
		return Config{}, configErr
	}

	err := vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
