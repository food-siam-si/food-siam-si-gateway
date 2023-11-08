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
	return nil
}

func (h *Handler) UpdateMenu(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) DeleteMenu(ctx *fiber.Ctx) error {
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

	err := h.menuService.DeleteMenu(uint32(restaurantIdUint), uint32(menuIdUint))

	log.Print(err)
	ctx.Status(fiber.StatusNotImplemented)

	return nil
}

func (h *Handler) GetMenus(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetRecommendMenu(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateRecommendMenu(ctx *fiber.Ctx) error {
	return nil
}
