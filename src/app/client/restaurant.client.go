package client

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRestaurantClient(config *config.AppConfig) (*grpc.ClientConn, error) {
	return grpc.Dial(config.RestaurantServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
