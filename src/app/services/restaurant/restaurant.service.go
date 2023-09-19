package restaurant

import "github.com/food-siam-si/food-siam-si-gateway/src/dto"

type Service struct {
}

type IService interface {
	CreateRestaurant() *dto.DTOErrorWithCode
	ViewRestaurantById() *dto.DTOErrorWithCode
	UpdateRestaurantInfo() *dto.DTOErrorWithCode
	ViewRestaurantType() *dto.DTOErrorWithCode
	RandomRestaurant() *dto.DTOErrorWithCode
}

func NewService() IService {
	return &Service{}
}

func (s *Service) CreateRestaurant() *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) ViewRestaurantById() *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) UpdateRestaurantInfo() *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) ViewRestaurantType() *dto.DTOErrorWithCode {
	return nil
}

func (s *Service) RandomRestaurant() *dto.DTOErrorWithCode {
	return nil
}
