package restaurant

type Service struct {
}

type IService interface {
	CreateRestaurant() error
	ViewRestaurantById() error
	UpdateRestaurantInfo() error
	ViewRestaurantType() error
	RandomRestaurant() error
}

func NewService() IService {
	return &Service{}
}

func (s *Service) CreateRestaurant() error {
	return nil
}

func (s *Service) ViewRestaurantById() error {
	return nil
}

func (s *Service) UpdateRestaurantInfo() error {
	return nil
}

func (s *Service) ViewRestaurantType() error {
	return nil
}

func (s *Service) RandomRestaurant() error {
	return nil
}
