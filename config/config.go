package config

import "github.com/spf13/viper"

type Config struct {
	API   API
	MySQL MySQL
	Redis Redis
}

type API struct {
	Host string
	Port string
}

type MySQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Redis struct {
	Addr     string
	Password string
	DB       int
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
