package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetAPIKey() (string, error) {
	apiKey := viper.GetString(apiKeyConfigKey)
	if apiKey == "" {
		return "", fmt.Errorf("API key not found")
	}
	return apiKey, nil
}
