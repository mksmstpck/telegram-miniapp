package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PsqlDns       string `mapstructure:"PSQL_DNS"`
	TelegramToken string `mapstructure:"TELEGRAM_TOKEN"`
	GinUrl        string `mapstructure:"GIN_URL"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	// Initialize the configuration structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Fatalf("Unable to decode into struct: %v", err)
		return nil, err
	}

	return &config, nil
}
