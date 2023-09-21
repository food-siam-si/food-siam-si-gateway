package user

import (
	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service user.IService
	v       validator.IValidator
}

func NewHandler(service user.IService, v validator.IValidator) *Handler {
	return &Handler{
		service: service,
		v:       v,
	}
}

func (h *Handler) GetCurrentUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

	ctx.Status(fiber.StatusOK)
	ctx.JSON(user)

	return nil
}

func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	body := dto.CreateUserRequest{}

	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid request body",
		})
		return nil
	}

	errVal := h.v.Validate(&body)

	if errVal != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *errVal,
		})

		return nil
	}

	_err := h.service.CreateUser(&body)

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

func (h *Handler) Signin(ctx *fiber.Ctx) error {
	body := dto.LoginRequest{}

	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid request body",
		})
		return nil
	}

	errVal := h.v.Validate(&body)

	if errVal != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *errVal,
		})

		return nil
	}

	token, _err := h.service.Signin(&body)

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})

		return nil
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "token"
	cookie.Value = token.Token
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)

	ctx.Cookie(cookie)

	ctx.Status(fiber.StatusOK)

	return nil
}

func (h *Handler) SignOut(ctx *fiber.Ctx) error {
	cookie := new(fiber.Cookie)

	cookie.Name = "token"
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(-24 * time.Hour)

	ctx.Cookie(cookie)
	ctx.Status(fiber.StatusOK)

	return nil
}
