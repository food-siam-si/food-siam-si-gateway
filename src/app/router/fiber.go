package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberRouter struct {
	*fiber.App
	Hello      fiber.Router
	User       fiber.Router
	Restaurant fiber.Router
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New(fiber.Config{
		BodyLimit: 16 * 1024 * 1024,
	})

	r.Use(cors.New(cors.ConfigDefault))

	hello := r.Group("/hello")
	user := r.Group("/user")
	restaurant := r.Group("/restaurant")

	return &FiberRouter{r, hello, user, restaurant}
}
