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
	ViewRestaurantById(id uint32) (*dto.Restaurant, *dto.DTOErrorWithCode)
	GetCurrentRestaurant(user *dto.UserToken) (*dto.Restaurant, *dto.DTOErrorWithCode)
	UpdateRestaurantInfo(user *dto.UserToken, body *dto.UpdateRestaurantRequest) *dto.DTOErrorWithCode
	ViewRestaurantType() (*proto.GetRestaurantTypeResponse, *dto.DTOErrorWithCode)
	RandomRestaurant(params *dto.RandomRestaurantRequest) (*dto.Restaurant, *dto.DTOErrorWithCode)
}

func NewService(resClient proto.RestaurantServiceClient, resTypeClient proto.RestaurantTypeServiceClient) IService {
	return &Service{
		resClient:     resClient,
		resTypeClient: resTypeClient,
	}
}

func (s *Service) GetCurrentRestaurant(user *dto.UserToken) (*dto.Restaurant, *dto.DTOErrorWithCode) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.resClient.GetCurrent(ctx, &proto.GetCurrentRestaurantRequest{
		User: &proto.User{
			Id:   user.Id,
			Name: user.Name,
			Type: proto.UserType(proto.UserType_value[string(user.Type)]),
		},
	})

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

	restaurantType := make([]dto.RestaurantType, 0)

	for _, v := range res.RestaurantType {
		restaurantType = append(restaurantType, dto.RestaurantType{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return &dto.Restaurant{
		Id:             res.Id,
		Name:           res.Name,
		Description:    res.Description,
		PhoneNumber:    res.PhoneNumber,
		LocationLat:    res.LocationLat,
		LocationLong:   res.LocationLong,
		Rating:         res.AverageScore,
		AveragePrice:   dto.AveragePrice(proto.AveragePrice_name[int32(res.AveragePrice)]),
		ImageUrl:       res.ImageUrl,
		RestaurantType: restaurantType,
	}, nil
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
		AveragePrice:      proto.AveragePrice(proto.AveragePrice_value[string(body.AveragePrice)]),
		ImageUrl:          body.ImageUrl,
		RestaurantTypeIds: body.RestaurantTypeIds,
		User: &proto.User{
			Id:   user.Id,
			Name: user.Name,
			Type: proto.UserType(proto.UserType_value[string(user.Type)]),
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

func (s *Service) ViewRestaurantById(id uint32) (*dto.Restaurant, *dto.DTOErrorWithCode) {
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

	restaurantType := make([]dto.RestaurantType, 0)

	for _, v := range res.RestaurantType {
		restaurantType = append(restaurantType, dto.RestaurantType{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return &dto.Restaurant{
		Id:             res.Id,
		Name:           res.Name,
		Description:    res.Description,
		PhoneNumber:    res.PhoneNumber,
		LocationLat:    res.LocationLat,
		LocationLong:   res.LocationLong,
		Rating:         res.AverageScore,
		AveragePrice:   dto.AveragePrice(proto.AveragePrice_name[int32(res.AveragePrice)]),
		ImageUrl:       res.ImageUrl,
		RestaurantType: restaurantType,
	}, nil
}

func (s *Service) UpdateRestaurantInfo(user *dto.UserToken, body *dto.UpdateRestaurantRequest) *dto.DTOErrorWithCode {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	protoBody := proto.UpdateCurrentRestaurantRequest{
		Name:              body.Name,
		Description:       body.Description,
		LocationLat:       body.LocationLat,
		LocationLong:      body.LocationLong,
		PhoneNumber:       body.PhoneNumber,
		AveragePrice:      proto.AveragePrice(proto.AveragePrice_value[string(body.AveragePrice)]),
		ImageUrl:          body.ImageUrl,
		IsInService:       body.IsInService,
		RestaurantTypeIds: body.RestaurantTypeIds,
		User: &proto.User{
			Id:   user.Id,
			Name: user.Name,
			Type: proto.UserType(proto.UserType_value[string(user.Type)]),
		},
	}

	_, err := s.resClient.UpdateCurrent(ctx, &protoBody)

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

func (s *Service) RandomRestaurant(params *dto.RandomRestaurantRequest) (*dto.Restaurant, *dto.DTOErrorWithCode) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	protoBody := proto.RandomRestaurantRequest{
		RestaurantTypeIds: params.RestaurantTypeIds,
		CurrentLat:        params.CurrentLat,
		CurrentLong:       params.CurrentLong,
		MaxDistanceKm:     params.MaxDistanceKm,
	}

	res, errRes := s.resClient.Random(ctx, &protoBody)

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

	restaurantType := make([]dto.RestaurantType, 0)

	for _, v := range res.RestaurantType {
		restaurantType = append(restaurantType, dto.RestaurantType{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return &dto.Restaurant{
		Id:             res.Id,
		Name:           res.Name,
		Description:    res.Description,
		PhoneNumber:    res.PhoneNumber,
		LocationLat:    res.LocationLat,
		LocationLong:   res.LocationLong,
		Rating:         res.AverageScore,
		AveragePrice:   dto.AveragePrice(proto.AveragePrice_name[int32(res.AveragePrice)]),
		ImageUrl:       res.ImageUrl,
		RestaurantType: restaurantType,
	}, nil
}
