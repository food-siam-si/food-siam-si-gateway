package menu

import (
	"log"
	"strconv"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/services/menu"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/dto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	menuService menu.IService
	v           validator.IValidator
}

func NewHandler(menuService menu.IService, v validator.IValidator) *Handler {
	return &Handler{
		menuService: menuService,
		v:           v,
	}
}

func (h *Handler) CreateMenu(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*dto.UserToken)

	restaurantId := ctx.Params("restaurantId")

	restaurantIdUint, err := strconv.ParseInt(restaurantId, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	var req dto.CreateMenuRequestBody

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid request body",
		})

		return nil
	}

	if err := h.v.Validate(req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *err,
		})

		return nil
	}

	_err := h.menuService.CreateMenu(uint32(restaurantIdUint), &req, user.Id)

	if _err != nil {
		log.Println(_err)
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})

		return nil
	}

	ctx.Status(fiber.StatusCreated)

	return nil
}

func (h *Handler) UpdateMenu(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*dto.UserToken)

	restaurantId := ctx.Params("restaurantId")
	menuId := ctx.Params("menuId")

	restaurantIdUint, rerr := strconv.ParseInt(restaurantId, 10, 32)
	menuIdUint, merr := strconv.ParseInt(menuId, 10, 32)

	if merr != nil || rerr != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	var req dto.UpdateMenuRequestBody

	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid request body",
		})

		return nil
	}

	if err := h.v.Validate(req); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOErrorArray{
			Message: *err,
		})

		return nil
	}

	_err := h.menuService.UpdateMenu(uint32(menuIdUint), uint32(restaurantIdUint), &req, user.Id)

	if _err != nil {
		log.Println(_err)
		ctx.Status(_err.Code)
		ctx.JSON(dto.DTOError{
			Message: _err.Message,
		})

		return nil
	}

	ctx.Status(fiber.StatusOK)

	return nil
}

func (h *Handler) DeleteMenu(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*dto.UserToken)

	restaurantId := ctx.Params("restaurantId")
	menuId := ctx.Params("menuId")

	restaurantIdUint, rerr := strconv.ParseInt(restaurantId, 10, 32)
	menuIdUint, merr := strconv.ParseInt(menuId, 10, 32)

	if rerr != nil || merr != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id or menu id",
		})
		return nil
	}

	err := h.menuService.DeleteMenu(uint32(restaurantIdUint), uint32(menuIdUint), user.Id)

	if err != nil {
		log.Println(err)
		ctx.Status(err.Code)
		ctx.JSON(dto.DTOError{
			Message: err.Message,
		})
		return nil
	}

	ctx.Status(fiber.StatusOK)

	return nil
}

func (h *Handler) RandomMenu(ctx *fiber.Ctx) error {
	restaurantId := ctx.Params("restaurantId")

	restaurantIdUint, err := strconv.ParseInt(restaurantId, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
	}

	res, _err := h.menuService.RandomMenu(uint32(restaurantIdUint))

	if _err != nil {
		log.Println(_err)
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

func (h *Handler) GetMenus(ctx *fiber.Ctx) error {
	restaurantId := ctx.Params("restaurantId")

	restaurantIdUint, err := strconv.ParseInt(restaurantId, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	res, _err := h.menuService.GetMenus(uint32(restaurantIdUint))

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

func (h *Handler) GetRecommendMenu(ctx *fiber.Ctx) error {
	restaurantId := ctx.Params("restaurantId")

	restaurantIdUint, err := strconv.ParseInt(restaurantId, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id",
		})
		return nil
	}

	res, _err := h.menuService.GetRecommendMenu(uint32(restaurantIdUint))

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

func (h *Handler) UpdateRecommendMenu(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*dto.UserToken)

	restaurantId := ctx.Params("restaurantId")
	menuId := ctx.Params("menuId")

	restaurantIdUint, rerr := strconv.ParseInt(restaurantId, 10, 32)
	menuIdUint, merr := strconv.ParseInt(menuId, 10, 32)

	if rerr != nil || merr != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(dto.DTOError{
			Message: "Invalid restaurant id or menu id",
		})
		return nil
	}

	err := h.menuService.UpdateRecommendMenu(uint32(restaurantIdUint), uint32(menuIdUint), user.Id, true)

	if err != nil {
		log.Println(err)
		ctx.Status(err.Code)
		ctx.JSON(dto.DTOError{
			Message: err.Message,
		})
	}

	return nil
}
