package main

import "github.com/spf13/viper"

type Config struct {
	API API
}

type API struct {
	Host string
	Port string
}

func Load() (*Config, error) {
	viper.SetConfigFile("./conf/config-local.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
