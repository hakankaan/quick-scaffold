package config

import (
	"github.com/spf13/viper"
)

func SetAPIKey(apiKey string) error {
	viper.Set(apiKeyConfigKey, apiKey)
	return SaveConfig()
}
