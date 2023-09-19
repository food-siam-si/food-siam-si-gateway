package hello

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/hello"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service hello.IService
	v       validator.IValidator
}

type IHandler interface {
	HelloWorld(ctx *fiber.Ctx) error
}

func NewHandler(service hello.IService, v validator.IValidator) IHandler {
	return &Handler{
		service: service,
		v:       v,
	}
}

func (h *Handler) HelloWorld(ctx *fiber.Ctx) error {
	text := ctx.Query("text")

	if err := h.v.Validate(&dto.HelloWorldQuery{
		Text: text,
	}); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *err,
		})

		return nil
	}

	res, err := h.service.HelloWorld(text)

	if err != nil {
		ctx.Status(err.Code)
		ctx.JSON(dto.DTOError{
			Message: err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(dto.HelloWorldResponse{
		Message: res.Message,
	})
	return nil
}
