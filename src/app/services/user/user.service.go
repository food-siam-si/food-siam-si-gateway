package user

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
}

type IService interface {
	GetCurrentUser(token string) *dto.DTOErrorWithCode
	CreateUser(req *dto.CreateUserRequest) *dto.DTOErrorWithCode
	Signin(req *dto.LoginRequest) (string, *dto.DTOErrorWithCode)
}

func NewService(client *resty.Client) IService {
	return &Service{client}
}

func (s *Service) GetCurrentUser(token string) *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) CreateUser(req *dto.CreateUserRequest) *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) Signin(req *dto.LoginRequest) (string, *dto.DTOErrorWithCode) {
	return "hello world", nil
}
