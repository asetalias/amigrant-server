package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT" required:"true"`
}

func LoadConfig() *Config {
	var config Config

	port := viper.GetString("PORT")
	if port == "" {
		log.Println("PORT not set, using default (8080)")
		config.Port = "8080"
	}

	return &config
}
