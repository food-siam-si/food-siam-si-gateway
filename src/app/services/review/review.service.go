package review

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
}

type IService interface {
	CreateReview(body *dto.CreateReviewRequest, user *dto.UserToken, restaurantId uint32) *dto.DTOErrorWithCode
	GetReview(restaurantId uint32) *dto.DTOErrorWithCode
}

func NewService(client *resty.Client) IService {
	return &Service{client}
}

func (s *Service) CreateReview(body *dto.CreateReviewRequest, user *dto.UserToken, restaurantId uint32) *dto.DTOErrorWithCode {
	bodySrv := dto.CreateReviewRequestService{
		Description:  body.Description,
		Rating:       body.Rating,
		UserId:       user.Id,
		RestaurantId: restaurantId,
	}

	res, err := s.client.R().SetBody(&bodySrv).EnableTrace().Post("/reviews")

	if err != nil {
		return &dto.DTOErrorWithCode{
			Code:    500,
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

func (s *Service) GetReview(restaurantId uint32) *dto.DTOErrorWithCode {
	return nil
}
