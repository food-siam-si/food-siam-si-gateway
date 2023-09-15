package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New(fiber.Config{
		BodyLimit: 16 * 1024 * 1024,
	})

	r.Use(cors.New(cors.ConfigDefault))

	return &FiberRouter{r}
}
