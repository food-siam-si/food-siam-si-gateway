package restaurant

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service restaurant.IService
	v       validator.IValidator
}

func NewHandler(service restaurant.IService, v validator.IValidator) *Handler {
	return &Handler{
		service: service,
		v:       v,
	}
}

func (h *Handler) CreateRestaurant(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) ViewRestaurantById(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateRestaurantInfo(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) ViewRestaurantType(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) RandomRestaurant(ctx *fiber.Ctx) error {
	return nil
}
