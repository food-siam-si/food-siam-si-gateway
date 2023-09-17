package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port                 string
	RestaurantServiceUrl string
	HelloServiceUrl      string
}

func LoadEnv() *AppConfig {
	config := &AppConfig{}

	// Load .env file
	godotenv.Load(".env")

	config.Port = os.Getenv("PORT")
	config.RestaurantServiceUrl = os.Getenv("RESTAURNAT_SERVICE_URL")
	config.HelloServiceUrl = os.Getenv("HELLO_SERVICE_URL")

	return config
}
