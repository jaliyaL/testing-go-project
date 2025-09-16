package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Port int
	}
	Databases map[string]struct {
		Driver string
		DSN    string
	}
	Cache struct {
		Driver   string
		Address  string
		Password string
		DB       int
	}
}

func LoadConfig() *AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config.yaml found, using only ENV: %v", err)
	}

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config unmarshal error: %v", err)
	}

	return &cfg
}
