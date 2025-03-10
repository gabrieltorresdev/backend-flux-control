package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() (*viper.Viper, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return viper.GetViper(), nil
}
