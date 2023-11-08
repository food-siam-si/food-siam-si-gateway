package menu

import "github.com/go-resty/resty/v2"

type Service struct {
	client *resty.Client
}

type IService interface {
	CreateMenu() error
	UpdateMenu() error
	DeleteMenu(restaurantId uint32, menuId uint32) error
	GetMenus(restaurantId uint32) error
	GetRecommendMenu(restaurantId uint32) error
	UpdateRecommendMenu(restaurantId uint32, menuId uint32, newStatus bool) error
}

func NewService(client *resty.Client) IService {
	return &Service{
		client: client,
	}
}

func (s *Service) CreateMenu() error {
	return nil
}

func (s *Service) UpdateMenu() error {
	return nil
}

func (s *Service) DeleteMenu(restaurantId uint32, menuId uint32) error {
	return nil
}

func (s *Service) GetMenus(restaurantId uint32) error {
	return nil
}

func (s *Service) GetRecommendMenu(restaurantId uint32) error {
	return nil
}

func (s *Service) UpdateRecommendMenu(restaurantId uint32, menuId uint32, newStatus bool) error {
	return nil
}
