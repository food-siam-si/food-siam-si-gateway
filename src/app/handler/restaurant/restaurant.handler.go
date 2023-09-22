package restaurant

import (
	"strconv"

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

	user := ctx.Locals("user").(dto.UserToken)

	_err := h.service.CreateRestaurant(&body, &user)

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusCreated)
	return nil
}

func (h *Handler) GetCurrentRestaurant(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

	res, _err := h.service.GetCurrentRestaurant(&user)

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(res)
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

	// Convert id to uint32
	idUint, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid id",
		})
		return nil
	}

	res, _err := h.service.ViewRestaurantById(uint32(idUint))

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(res)
	return nil
}

func (h *Handler) UpdateRestaurantInfo(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

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

	_err := h.service.UpdateRestaurantInfo(&user, &body)

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)

	return nil
}

func (h *Handler) ViewRestaurantType(ctx *fiber.Ctx) error {
	res, err := h.service.ViewRestaurantType()

	if err != nil {
		ctx.Status(err.Code)
		ctx.JSON(dto.DTOError{
			Message: err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(res)

	return nil
}

func (h *Handler) RandomRestaurant(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotImplemented)
	ctx.JSON(dto.DTOError{
		Message: "Not Implemented",
	})
	return nil
}
