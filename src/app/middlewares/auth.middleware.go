package middlewares

import (
	"strings"

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
	authorization := ctx.GetRespHeader("Authorization", "")

	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
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
