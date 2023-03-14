package core

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	Listener string `mapstructure:"LISTENER"`
}

func NewConfig() (config Config, err error) {
	viper.ReadInConfig()
	viper.AutomaticEnv()

	viper.AddConfigPath("resource/")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err

	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
