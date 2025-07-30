package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"db"`
}

var AppConfig *Config

func LoadConfig() {
	data, err := os.ReadFile("resources/application.yml")
	if err != nil {
		log.Fatalf("Error reading YAML config file: %v", err)
	}

	AppConfig = &Config{}
	if err := yaml.Unmarshal(data, AppConfig); err != nil {
		log.Fatalf("Error parsing YAML config: %v", err)
	}

	fmt.Println(" Config loaded")
}

func GetConfig() *Config {
	return AppConfig
}
