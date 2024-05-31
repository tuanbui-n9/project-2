package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ENV             string
	PORT            string
	TRUSTED_PROXIES string
}

func LoadConfig() (*AppConfig, error) {
	// Load variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("no .env file found")
	}

	viper.AutomaticEnv()

	var cfg AppConfig
	cfg.PORT = viper.GetString("PORT")
	cfg.ENV = viper.GetString("PORT")
	cfg.TRUSTED_PROXIES = viper.GetString("TRUSTED_PROXIES")

	return &cfg, nil
}
