package config

import (
	"log"

	"github.com/spf13/viper"
)

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

// Функция для загрузки конфигурации из файла config.yaml
func LoadConfig() *Config {
	// Устанавливаем имя и тип конфигурационного файла
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Чтение конфигурационного файла
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка при чтении конфигурационного файла: %v", err)
	}

	// Структура для хранения конфигурации
	var cfg Config
	// Декодируем конфигурацию в структуру
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Невозможно декодировать конфигурацию в структуру: %v", err)
	}

	return &cfg
}
