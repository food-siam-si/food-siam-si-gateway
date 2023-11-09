package middlewares

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Remind that restaurantId must exist in param
type RestaurantMiddleware struct {
	service restaurant.IService
}

type IRestaurantMiddleware interface {
	OwnerGuard(ctx *fiber.Ctx) error
	OwnerOrCustomerGuard(ctx *fiber.Ctx) error
}

func NewRestaurantMiddleware(service restaurant.IService) IRestaurantMiddleware {
	return &RestaurantMiddleware{
		service: service,
	}
}

func (m *RestaurantMiddleware) OwnerGuard(ctx *fiber.Ctx) error {
	restaurantId, err := ctx.ParamsInt("restaurantId")
	user := ctx.Locals("user").(dto.UserToken)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	restaurant, _err := m.service.GetCurrentRestaurant(&user)

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	if restaurantId != int(restaurant.Id) {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(dto.DTOError{
			Message: "Forbidden",
		})
		return nil
	}

	return ctx.Next()
}

func (m *RestaurantMiddleware) OwnerOrCustomerGuard(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)

	if user.Type == "Customer" {
		return ctx.Next()
	}

	restaurantId, err := ctx.ParamsInt("restaurantId")

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	restaurant, _err := m.service.GetCurrentRestaurant(&user)

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	if restaurantId != int(restaurant.Id) {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(dto.DTOError{
			Message: "Forbidden",
		})
		return nil
	}

	return ctx.Next()
}
