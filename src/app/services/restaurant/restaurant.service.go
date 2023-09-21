package restaurant

import (
	"context"
	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/food-siam-si/food-siam-si-gateway/src/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Service struct {
	resClient     proto.RestaurantServiceClient
	resTypeClient proto.RestaurantTypeServiceClient
}

type IService interface {
	CreateRestaurant(body *dto.CreateRestaurantRequest, user *dto.UserToken) *dto.DTOErrorWithCode
	ViewRestaurantById(id uint32) (*proto.Restaurant, *dto.DTOErrorWithCode)
	UpdateRestaurantInfo(id uint32, user *dto.UserToken, body *dto.UpdateRestaurantRequest) *dto.DTOErrorWithCode
	ViewRestaurantType() (*proto.GetRestaurantTypeResponse, *dto.DTOErrorWithCode)
	RandomRestaurant() *dto.DTOErrorWithCode
}

func NewService(resClient proto.RestaurantServiceClient, resTypeClient proto.RestaurantTypeServiceClient) IService {
	return &Service{
		resClient:     resClient,
		resTypeClient: resTypeClient,
	}
}

func (s *Service) CreateRestaurant(body *dto.CreateRestaurantRequest, user *dto.UserToken) *dto.DTOErrorWithCode {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, errRes := s.resClient.Create(ctx, &proto.CreateRestaurantRequest{
		Name:              body.Name,
		Description:       body.Description,
		LocationLat:       body.LocationLat,
		LocationLong:      body.LocationLong,
		PhoneNumber:       body.PhoneNumber,
		AveragePrice:      body.AveragePrice,
		ImageUrl:          body.ImageUrl,
		RestaurantTypeIds: body.RestaurantTypeIds,
		User: &proto.User{
			Id:   user.Id,
			Name: user.Name,
			Type: user.Type,
		},
	})

	if errRes != nil {
		st, ok := status.FromError(errRes)

		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return &dto.DTOErrorWithCode{
					Code:    fiber.StatusBadRequest,
					Message: st.Message(),
				}
			case codes.NotFound:
				return &dto.DTOErrorWithCode{
					Code:    fiber.StatusNotFound,
					Message: st.Message(),
				}
			default:
				return &dto.DTOErrorWithCode{
					Code:    fiber.StatusInternalServerError,
					Message: st.Message(),
				}
			}
		}
	}

	return nil
}

func (s *Service) ViewRestaurantById(id uint32) (*proto.Restaurant, *dto.DTOErrorWithCode) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	protoBody := wrapperspb.UInt32Value{Value: id}

	res, errRes := s.resClient.FindById(ctx, &protoBody)

	if errRes != nil {
		st, ok := status.FromError(errRes)

		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusBadRequest,
					Message: st.Message(),
				}
			case codes.NotFound:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusNotFound,
					Message: st.Message(),
				}
			default:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusInternalServerError,
					Message: st.Message(),
				}
			}
		}
	}

	return res, nil
}

func (s *Service) UpdateRestaurantInfo(id uint32, user *dto.UserToken, body *dto.UpdateRestaurantRequest) *dto.DTOErrorWithCode {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	protoBody := proto.UpdateRestaurantRequest{
		Id:                id,
		Name:              body.Name,
		Description:       body.Description,
		LocationLat:       body.LocationLat,
		LocationLong:      body.LocationLong,
		PhoneNumber:       body.PhoneNumber,
		AveragePrice:      body.AveragePrice,
		ImageUrl:          body.ImageUrl,
		IsInService:       body.IsInService,
		RestaurantTypeIds: body.RestaurantTypeIds,
		User: &proto.User{
			Id:   user.Id,
			Name: user.Name,
			Type: user.Type,
		},
	}

	_, err := s.resClient.Update(ctx, &protoBody)

	if err != nil {
		st, ok := status.FromError(err)

		if ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return &dto.DTOErrorWithCode{
					Code:    fiber.StatusBadRequest,
					Message: st.Message(),
				}
			case codes.NotFound:
				return &dto.DTOErrorWithCode{
					Code:    fiber.StatusNotFound,
					Message: st.Message(),
				}
			default:
				return &dto.DTOErrorWithCode{
					Code:    fiber.StatusInternalServerError,
					Message: st.Message(),
				}
			}
		}
	}

	return nil
}

func (s *Service) ViewRestaurantType() (*proto.GetRestaurantTypeResponse, *dto.DTOErrorWithCode) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.resTypeClient.GetAll(ctx, &emptypb.Empty{})

	if errRes != nil {
		st, ok := status.FromError(errRes)

		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusNotFound,
					Message: st.Message(),
				}
			default:
				return nil, &dto.DTOErrorWithCode{
					Code:    fiber.StatusInternalServerError,
					Message: st.Message(),
				}
			}
		}
	}

	return res, nil
}

func (s *Service) RandomRestaurant() *dto.DTOErrorWithCode {
	return nil
}
