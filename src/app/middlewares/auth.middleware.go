package middlewares

import (
	"strings"

	"time"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	service user.IService
}

type IAuthMiddleware interface {
	AuthGuard(ctx *fiber.Ctx) error
	RestaurantGuard(ctx *fiber.Ctx) error
}

func NewAuthMiddleware(service user.IService) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}

func (m *AuthMiddleware) AuthGuard(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authorization, ok := headers["Authorization"]

	if !ok || !strings.HasPrefix(authorization, "Bearer ") {
		ctx.Status(fiber.StatusUnauthorized)
		ctx.JSON(dto.DTOError{
			Message: "Unauthorized",
		})
		return nil
	}

	token := authorization[7:]

	if token == "" {
		ctx.Status(fiber.StatusUnauthorized)
		ctx.JSON(dto.DTOError{
			Message: "Unauthorized",
		})
		return nil
	}

	user, err := m.service.GetCurrentUser(token)

	if err != nil {
		ctx.Status(err.Code)
		ctx.JSON(dto.DTOError{
			Message: err.Message,
		})

		cookie := new(fiber.Cookie)

		cookie.Name = "token"
		cookie.HTTPOnly = true
		cookie.Expires = time.Now().Add(-24 * time.Hour)

		ctx.Cookie(cookie)
		return nil
	}

	ctx.Locals("user", dto.UserToken{
		Type: user.Data.UserType,
		Id:   user.Data.Id,
		Name: user.Data.Username,
	})
	return ctx.Next()
}

func (m *AuthMiddleware) RestaurantGuard(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

	if user.Type != "Owner" {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(dto.DTOError{
			Message: "Forbidden",
		})
		return nil
	}

	return ctx.Next()
}

func (m *AuthMiddleware) CustomerGuard(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

	if user.Type != "Customer" {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(dto.DTOError{
			Message: "Forbidden",
		})
		return nil
	}

	return ctx.Next()
}
