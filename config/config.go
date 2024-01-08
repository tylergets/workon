package config

import (
	"log"

	"github.com/spf13/viper"
)

type Workflow struct {
	Name     string   `json:"name"`
	Run      []string `json:"run"`
	Cwd      *string  `json:"cwd"`
	Windowed bool     `json:"windowed"`
	// Data is a map of key/value pairs that can be used to pass data between steps.
	Data map[string]string `json:"data"`
}

type Config struct {
	Workflows map[string][]Workflow `json:"workflows"`
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return configuration
}
