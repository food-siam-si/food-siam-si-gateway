package hello

import (
	"context"
	"errors"
	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	HelloWorld(text string) (*proto.HelloWorldResponse, error)
}

type Service struct {
	client proto.HelloServiceClient
}

func NewService(client proto.HelloServiceClient) IService {
	return &Service{
		client,
	}
}

func (s *Service) HelloWorld(text string) (*proto.HelloWorldResponse, error) {
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
				return nil, errors.New("not found")
			default:
				return nil, errors.New("internal server error")
			}
		}
	}

	return res, nil
}
