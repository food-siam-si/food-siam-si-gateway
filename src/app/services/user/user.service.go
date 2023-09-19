package user

import "github.com/food-siam-si/food-siam-si-gateway/src/dto"

type Service struct {
}

type IService interface {
	GetCurrentUser(token string) error
	CreateUser(req *dto.CreateUserRequest) *dto.DTOErrorWithCode
	Signin(req *dto.LoginRequest) (string, *dto.DTOErrorWithCode)
}

func NewService() IService {
	return &Service{}
}

func (s *Service) GetCurrentUser(token string) error {
	return nil
}

func (s *Service) CreateUser(req *dto.CreateUserRequest) *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) Signin(req *dto.LoginRequest) (string, *dto.DTOErrorWithCode) {
	return "hello world", nil
}
