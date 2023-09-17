package user

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/user"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service user.IService
}

func NewHandler(service user.IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetCurrentUser(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) Signin(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) SignOut(ctx *fiber.Ctx) error {
	return nil
}
