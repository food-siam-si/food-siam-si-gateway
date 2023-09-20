package restaurant

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
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
	body := dto.CreateRestaurantRequest{}

	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid request body" + err.Error(),
		})
		return nil
	}

	if err := h.v.Validate(body); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *err,
		})

		return nil
	}

	ctx.Status(fiber.StatusNotImplemented)
	ctx.JSON(dto.DTOError{
		Message: "Not Implemented",
	})
	return nil
}

func (h *Handler) ViewRestaurantById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Missing id",
		})
		return nil
	}

	ctx.Status(fiber.StatusNotImplemented)
	ctx.JSON(dto.DTOError{
		Message: "Not Implemented",
	})
	return nil
}

func (h *Handler) UpdateRestaurantInfo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Missing id",
		})
		return nil
	}

	body := dto.UpdateRestaurantRequest{}

	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid request body",
		})
		return nil
	}

	if err := h.v.Validate(body); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *err,
		})

		return nil
	}

	ctx.Status(fiber.StatusNotImplemented)
	ctx.JSON(dto.DTOError{
		Message: "Not Implemented",
	})
	return nil
}

func (h *Handler) ViewRestaurantType(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotImplemented)
	ctx.JSON(dto.DTOError{
		Message: "Not Implemented",
	})
	return nil
}

func (h *Handler) RandomRestaurant(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotImplemented)
	ctx.JSON(dto.DTOError{
		Message: "Not Implemented",
	})
	return nil
}
