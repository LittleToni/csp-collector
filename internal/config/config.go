package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration values for the service.
type Config struct {
	Port string
}

// Populates the Config struct.
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Create config and assign environment variables to struct fields.
	config := &Config{
		Port: os.Getenv("PORT"),
	}

	// Set the configuration fallbacks.
	if config.Port == "" {
		config.Port = "8080"
	}

	return config, nil
}
