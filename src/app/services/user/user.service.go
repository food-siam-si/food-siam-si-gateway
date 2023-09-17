package user

type Service struct {
}

type IService interface {
	GetCurrentUser() error
	CreateUser() error
	Signin() error
	SignOut() error
}

func NewService() IService {
	return &Service{}
}

func (s *Service) GetCurrentUser() error {
	return nil
}

func (s *Service) CreateUser() error {
	return nil
}

func (s *Service) Signin() error {
	return nil
}

func (s *Service) SignOut() error {
	return nil
}
