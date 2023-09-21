package user

import (
	"fmt"

	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	client *resty.Client
}

type IService interface {
	GetCurrentUser(token string) *dto.DTOErrorWithCode
	CreateUser(req *dto.CreateUserRequest) *dto.DTOErrorWithCode
	Signin(req *dto.LoginRequest) (*dto.LoginResponse, *dto.DTOErrorWithCode)
}

func NewService(client *resty.Client) IService {
	return &Service{client}
}

func (s *Service) GetCurrentUser(token string) *dto.DTOErrorWithCode {
	res, err := s.client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %v", token)).EnableTrace().Get("/users/verify")

	if err != nil {
		return &dto.DTOErrorWithCode{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if res.StatusCode() >= 400 {
		return &dto.DTOErrorWithCode{
			Code:    res.StatusCode(),
			Message: res.String(),
		}
	}

	return nil
}

func (s *Service) CreateUser(req *dto.CreateUserRequest) *dto.DTOErrorWithCode {
	res, err := s.client.R().SetBody(req).EnableTrace().Post("/users/register")

	if err != nil {
		return &dto.DTOErrorWithCode{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if res.StatusCode() >= 400 {
		return &dto.DTOErrorWithCode{
			Code:    res.StatusCode(),
			Message: res.String(),
		}
	}

	return nil
}

func (s *Service) Signin(req *dto.LoginRequest) (*dto.LoginResponse, *dto.DTOErrorWithCode) {
	body := dto.LoginResponse{}

	res, err := s.client.R().SetBody(req).SetResult(&body).EnableTrace().Post("/users/login")

	if err != nil {
		return nil, &dto.DTOErrorWithCode{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if res.StatusCode() != 200 {
		return nil, &dto.DTOErrorWithCode{
			Code:    res.StatusCode(),
			Message: res.String(),
		}
	}

	return &body, nil
}
