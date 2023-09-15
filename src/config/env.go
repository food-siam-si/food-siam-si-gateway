package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port                 string `mapstructure:"PORT"`
	RestaurantServiceUrl string `mapstructure:"RESTAURANT_SERVICE_URL"`
}

func LoadEnv() (*AppConfig, error) {
	config := &AppConfig{}

	viper.SetConfigFile("./.env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return config, nil
}
