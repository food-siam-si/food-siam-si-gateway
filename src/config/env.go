package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port                 string
	RestaurantServiceUrl string
	UserServiceUrl       string
	ReviewServiceUrl     string
	MenuServiceUrl       string
}

func LoadEnv() *AppConfig {
	config := &AppConfig{}

	// Load .env file
	godotenv.Load(".env")

	config.Port = os.Getenv("PORT")
	config.RestaurantServiceUrl = os.Getenv("RESTAURANT_SERVICE_URL")
	config.UserServiceUrl = os.Getenv("USER_SERVICE_URL")
	config.ReviewServiceUrl = os.Getenv("REVIEW_SERVICE_URL")
	config.MenuServiceUrl = os.Getenv("MENU_SERVICE_URL")

	return config
}
