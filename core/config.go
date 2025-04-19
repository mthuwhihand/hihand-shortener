package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	REDIS_HOST     string `mapstructure:"REDIS_HOST"`
	REDIS_PORT     string `mapstructure:"REDIS_PORT"`
	REDIS_USER     string `mapstructure:"REDIS_USER"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_URL      string `mapstructure:"REDIS_URL"`
	BASE62_CHARS   string `mapstructure:"BASE62_CHARS"`
	BASE_SHORT_URL string `mapstructure:"BASE_SHORT_URL"`
	APP_PORT       string `mapstructure:"APP_PORT"`
	BASE_URL       string `mapstructure:"BASE_URL"`
}

var AppConfig Config

// InitConfig loads configuration from .env file
func InitConfig(path string) {
	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(path)
	v.AutomaticEnv() // Load env variables

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	err = v.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// Printing loaded configuration for debugging
	fmt.Println("REDIS_HOST:", AppConfig.REDIS_HOST)
	fmt.Println("REDIS_PORT:", AppConfig.REDIS_PORT)
	fmt.Println("REDIS_USER:", AppConfig.REDIS_USER)
	fmt.Println("REDIS_PASSWORD:", AppConfig.REDIS_PASSWORD)
	fmt.Println("REDIS_URL:", AppConfig.REDIS_URL)
	fmt.Println("BASE62_CHARS:", AppConfig.BASE62_CHARS)
	fmt.Println("BASE_SHORT_URL:", AppConfig.BASE_SHORT_URL)
	fmt.Println("APP_PORT:", AppConfig.APP_PORT)
	fmt.Println("BASE_URL:", AppConfig.BASE_URL)
}
