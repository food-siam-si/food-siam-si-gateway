package hello

import (
	"context"
	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/food-siam-si/food-siam-si-gateway/src/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	HelloWorld(text string) (*proto.HelloWorldResponse, *dto.DTOErrorWithCode)
}

type Service struct {
	client proto.HelloServiceClient
}

func NewService(client proto.HelloServiceClient) IService {
	return &Service{
		client,
	}
}

func (s *Service) HelloWorld(text string) (*proto.HelloWorldResponse, *dto.DTOErrorWithCode) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.client.HelloWorld(ctx, &proto.HelloWorldRequest{
		Text: text,
	})

	if errRes != nil {
		st, ok := status.FromError(errRes)

		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusNotFound,
					Message: "Not found",
				}
			default:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusInternalServerError,
					Message: "Internal server error",
				}
			}
		}
	}

	return res, nil
}
