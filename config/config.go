package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the application configuration from a YAML file.
// It reads the configuration from a file named "config.yaml" and unmarshals it into a Config struct.
// @return *Config returns the loaded configuration.
type Config struct {
	ServerPort string `yaml:"serverPort"`
	DBHost     string `yaml:"dbHost"`
	DBPort     string `yaml:"dbPort"`
	DBUser     string `yaml:"dbUser"`
	DBPassword string `yaml:"dbPassword"`
	DBName     string `yaml:"dbName"`
	RedisHost  string `yaml:"redisHost"`
	RedisPort  string `yaml:"redisPort"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("File configure error: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode configuration into structure: %v", err)
	}

	return &cfg
}
