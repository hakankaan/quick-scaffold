package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	apiKeyConfigKey = "api_key"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	configPath := filepath.Join(os.Getenv("HOME"), ".app")
	viper.AddConfigPath(configPath)

	// Create config file if it doesn't exist
	configFile := filepath.Join(configPath, "config.yaml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func SaveConfig() error {
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func EnsureConfigFileExists() error {
	configPath := filepath.Join(os.Getenv("HOME"), ".app", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return err
		}
		_, err = os.Create(configPath)
		if err != nil {
			return err
		}
	}
	return nil
}
