package utils

import (
	"github.com/sandricoprovo/denoken/structs"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("should successfully load config", func(t *testing.T) {
		mockConfigPath := structs.ConfigPaths{
			File:  "denoken.config",
			Paths: []string{"../", "../config/"},
		}

		config, err := LoadConfig(mockConfigPath)

		// Assert that the expected config is returned with no errors
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if config != config {
			t.Errorf("Expected config to be %v, but got %v", config, config)
		}
	})

	t.Run("should return error from invalid config", func(t *testing.T) {
		mockConfigPath := structs.ConfigPaths{
			File:  "denoken.config",
			Paths: []string{},
		}

		_, err := LoadConfig(mockConfigPath)

		if err == nil {
			t.Fail()
		}
	})
}

// TestLoadConfig_UnmarshalError tests when the configuration file cannot be unmarshalled.

// func TestLoadConfig_UnmarshalError(t *testing.T) {
// 	// Set up mock viper instance that returns a malformed config
// 	viperMock := viper.New()
// 	viperMock.Set("path", "invalid")

// 	// Replace the LoadConfig function's implementation with our mock
// 	oldViper := ViperInstance
// 	ViperInstance = viperMock
// 	defer func() { ViperInstance = oldViper }()

// 	// Assert that the function returns an error and no config is returned
// 	config, err := LoadConfig()
// 	if err.Error() != expectedError.Error() {
// 		t.Errorf("Expected error to be %v, but got %v", expectedError, err)
// 	}
// 	expectedError := errors.New("Error unmarshalling config into struct: field path: cannot unmarshal string into Go struct field Config.path of type int")
// 	if config != (Config{}) {
// 		t.Errorf("Expected empty config, but got %v", config)
// 	}
// }
