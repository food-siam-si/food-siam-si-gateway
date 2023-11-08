package review

import (
	"strconv"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/review"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	restaurantSrv restaurant.IService
	reviewSrv     review.IService
	v             validator.IValidator
}

func NewHandler(restaurantSrv restaurant.IService, reviewSrv review.IService, v validator.IValidator) *Handler {
	return &Handler{restaurantSrv, reviewSrv, v}
}

func (h *Handler) CreateReview(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.UserToken)
	restaurantId := ctx.Params("restaurantId")

	restaurantIdInt, err := strconv.Atoi(restaurantId)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	body := dto.CreateReviewRequest{}

	if err := ctx.BodyParser(&body); err != nil {
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

	_, _err := h.restaurantSrv.ViewRestaurantById(uint32(restaurantIdInt))

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
		return nil
	}

	restaurant, _ := h.restaurantSrv.GetCurrentRestaurant(&user)

	if restaurant != nil && restaurant.Id != uint32(restaurantIdInt) {
		ctx.Status(fiber.StatusForbidden)
		ctx.JSON(dto.DTOError{
			Message: "You are not allowed to review restaurant",
		})
		return nil
	}

	_err = h.reviewSrv.CreateReview(&body, &user, uint32(restaurantIdInt))

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

func (h *Handler) GetReview(ctx *fiber.Ctx) error {
	restaurantId := ctx.Params("restaurantId")

	restaurantIdInt, err := strconv.Atoi(restaurantId)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	res, _err := h.reviewSrv.GetReview(uint32(restaurantIdInt))

	if _err != nil {
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(res)

	return nil
}
