package client

import (
	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"github.com/go-resty/resty/v2"
)

func NewRestaurantClient(config *config.AppConfig) *resty.Client {
	client := resty.New()

	client.SetBaseURL(config.RestaurantServiceUrl)
	client.SetTimeout(30 * time.Second)

	return client
}
