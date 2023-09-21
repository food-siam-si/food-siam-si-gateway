package middlewares

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/food-siam-si/food-siam-si-gateway/src/proto"
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
	token := ctx.Cookies("token")

	if token == "" {
		ctx.Status(fiber.StatusUnauthorized)
		ctx.JSON(dto.DTOError{
			Message: "Unauthorized",
		})
		return nil
	}

	err := m.service.GetCurrentUser(token)

	if err != nil {
		ctx.Status(err.Code)
		ctx.JSON(dto.DTOError{
			Message: err.Message,
		})
		return nil
	}

	ctx.Locals("user", dto.UserToken{
		Type: proto.UserType_Owner,
		Id:   1,
		Name: token,
	})
	return ctx.Next()
}

func (m *AuthMiddleware) RestaurantGuard(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

	if user.Type.String() != proto.UserType_Owner.String() {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(dto.DTOError{
			Message: "Forbidden",
		})
		return nil
	}

	return ctx.Next()
}
