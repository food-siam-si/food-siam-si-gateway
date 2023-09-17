package restaurant

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service restaurant.IService
}

func NewHandler(service restaurant.IService) *Handler {
	return &Handler{
		service: service,
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
