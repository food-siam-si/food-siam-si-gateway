package client

import (
	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"github.com/go-resty/resty/v2"
)

func NewUserClient(config *config.AppConfig) *resty.Client {
	client := resty.New()

	client.SetBaseURL(config.UserServiceUrl)
	client.SetTimeout(30 * time.Second)
	client.EnableTrace()

	return client
}
