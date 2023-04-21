package utils

import (
	"sandricoprovo/design-token-builder/structs"

	"github.com/spf13/viper"
)

type Config struct {
	TypeScale structs.TypeScaleConfig `mapstructure:"typeScale"`
}

func LoadConfig() (Config, error) {
	var config Config
	vp := viper.New()


	// TODO: update config to work with a save path. this'll mean users don't have to constantly have to drag the file into a folder of their choice

	// Finds the config file
	vp.SetConfigName("design.config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("./config/")
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

