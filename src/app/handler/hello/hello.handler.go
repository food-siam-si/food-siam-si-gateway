package hello

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/hello"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service hello.IService
}

type IHandler interface {
	HelloWorld(ctx *fiber.Ctx) error
}

func NewHandler(service hello.IService) IHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HelloWorld(ctx *fiber.Ctx) error {
	text := ctx.Query("text")

	res, err := h.service.HelloWorld(text)

	if err != nil {
		ctx.Status(err.Code)
		ctx.JSON(fiber.Map{
			"message": err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(fiber.Map{
		"text": res.Message,
	})
	return nil
}
