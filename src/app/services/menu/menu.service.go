package menu

import (
	"fmt"
	"log"
	"net/url"

	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
}

type IService interface {
	CreateMenu(restaurantId uint32, body *dto.CreateMenuRequestBody, userId uint32) *dto.DTOErrorWithCode
	UpdateMenu(menuId uint32, restaurantId uint32, body *dto.UpdateMenuRequestBody, userId uint32) *dto.DTOErrorWithCode
	DeleteMenu(restaurantId uint32, menuId uint32, userId uint32) *dto.DTOErrorWithCode
	GetMenus(restaurantId uint32) (*dto.GetMenusResponseService, *dto.DTOErrorWithCode)
	RandomMenu(restaurantId uint32, query *dto.RandomMenuRequest) (*dto.GetMenuResponseService, *dto.DTOErrorWithCode)
	GetRecommendMenu(restaurantId uint32) (*dto.GetRecommendMenuResponseService, *dto.DTOErrorWithCode)
	UpdateRecommendMenu(restaurantId uint32, menuId uint32, userId uint32, newStatus bool) *dto.DTOErrorWithCode
	ViewMenuType() (*dto.GetMenuTypeResponseService, *dto.DTOErrorWithCode)
	ViewMenuTypeByRestaurantId(restaurantId uint32) (*dto.GetMenuTypeResponseService, *dto.DTOErrorWithCode)
}

func NewService(client *resty.Client) IService {
	return &Service{
		client: client,
	}
}

func (s *Service) CreateMenu(restaurantId uint32, body *dto.CreateMenuRequestBody, userId uint32) *dto.DTOErrorWithCode {
	serviceBody := dto.CreateMenuRequestBodyService{
		UserId:      userId,
		Price:       body.Price,
		Title:       body.Title,
		Description: body.Description,
		IsRecom:     body.IsRecom,
		ImageUrl:    body.ImageUrl,
		Addons:      body.Addons,
		Types:       body.Types,
	}

	res, err := s.client.R().SetBody(&serviceBody).Post(fmt.Sprintf("/menus/%v", restaurantId))

	if err != nil {
		return &dto.DTOErrorWithCode{
			Message: "Failed to create menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return nil
}

func (s *Service) UpdateMenu(menuId uint32, restaurantId uint32, body *dto.UpdateMenuRequestBody, userId uint32) *dto.DTOErrorWithCode {
	serviceBody := dto.UpdateMenuRequestBodyService{
		UserId:      userId,
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		IsRecom:     body.IsRecom,
		ImageUrl:    body.ImageUrl,
		Addons:      body.Addons,
		MenuId:      menuId,
		Types:       body.Types,
	}

	res, err := s.client.R().SetBody(&serviceBody).Put(fmt.Sprintf("/menus/%v", restaurantId))

	if err != nil {
		return &dto.DTOErrorWithCode{
			Message: "Failed to update menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return nil
}

func (s *Service) DeleteMenu(restaurantId uint32, menuId uint32, userId uint32) *dto.DTOErrorWithCode {
	body := map[string]interface{}{
		"menu_id": menuId,
		"user_id": userId,
	}

	res, err := s.client.R().
		SetBody(body).
		Delete(fmt.Sprintf("/menus/%v", restaurantId))

	if err != nil {
		return &dto.DTOErrorWithCode{
			Message: "Failed to delete menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return nil
}

func (s *Service) RandomMenu(restaurantId uint32, query *dto.RandomMenuRequest) (*dto.GetMenuResponseService, *dto.DTOErrorWithCode) {
	var body dto.GetMenuResponseService

	types := []string{}

	for _, t := range query.Types {
		types = append(types, fmt.Sprint(t))
	}

	res, err := s.client.R().SetQueryParamsFromValues(url.Values{"types": types}).SetResult(&body).Get(fmt.Sprintf("/menus/%v/random", restaurantId))

	if err != nil {
		return nil, &dto.DTOErrorWithCode{
			Message: "Failed to get random menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return nil, &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return &body, nil
}

func (s *Service) GetMenus(restaurantId uint32) (*dto.GetMenusResponseService, *dto.DTOErrorWithCode) {
	body := dto.GetMenusResponseService{}

	res, err := s.client.R().SetResult(&body).Get(fmt.Sprintf("/menus/%v", restaurantId))

	if err != nil {
		log.Println(err)
		return nil, &dto.DTOErrorWithCode{
			Message: "Failed to get menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return nil, &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return &body, nil
}

func (s *Service) GetRecommendMenu(restaurantId uint32) (*dto.GetRecommendMenuResponseService, *dto.DTOErrorWithCode) {
	body := dto.GetRecommendMenuResponseService{}

	res, err := s.client.R().SetResult(&body).Get(fmt.Sprintf("/menus/%v/recommend", restaurantId))

	if err != nil {
		return nil, &dto.DTOErrorWithCode{
			Message: "Failed to get random menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return nil, &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return &body, nil
}

func (s *Service) UpdateRecommendMenu(restaurantId uint32, menuId uint32, userId uint32, newStatus bool) *dto.DTOErrorWithCode {
	body := dto.UpdateRecommendMenuRequestBodyService{
		IsRecom: newStatus,
		UserId:  userId,
		MenuId:  menuId,
	}

	res, err := s.client.R().SetBody(&body).Put(fmt.Sprintf("/menus/%v/recommend", restaurantId))

	if err != nil {
		return &dto.DTOErrorWithCode{
			Message: "Failed to update recommend menu",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return nil
}

func (s *Service) ViewMenuType() (*dto.GetMenuTypeResponseService, *dto.DTOErrorWithCode) {
	body := dto.GetMenuTypeResponseService{}

	res, err := s.client.R().SetResult(&body).Get("/menus/types")

	if err != nil {
		return nil, &dto.DTOErrorWithCode{
			Message: "Failed to get menu type",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return nil, &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return &body, nil
}

func (s *Service) ViewMenuTypeByRestaurantId(restaurantId uint32) (*dto.GetMenuTypeResponseService, *dto.DTOErrorWithCode) {
	body := dto.GetMenuTypeResponseService{}

	res, err := s.client.R().SetResult(&body).Get(fmt.Sprintf("/menus/%v/types", restaurantId))

	if err != nil {
		return nil, &dto.DTOErrorWithCode{
			Message: "Failed to get menu type",
			Code:    500,
		}
	}

	if res.StatusCode() >= 400 {
		return nil, &dto.DTOErrorWithCode{
			Message: res.String(),
			Code:    res.StatusCode(),
		}
	}

	return &body, nil
}
