package router

import (
	"github.com/food-siam-si/food-siam-si-gateway/src/app/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberRouter struct {
	*fiber.App
	Hello      fiber.Router
	User       fiber.Router
	Restaurant fiber.Router
}

func NewFiberRouter(authMiddleware middlewares.IAuthMiddleware) *FiberRouter {
	r := fiber.New(fiber.Config{
		BodyLimit: 16 * 1024 * 1024,
	})

	r.Use(cors.New(cors.ConfigDefault))

	hello := r.Group("/hello")
	user := r.Group("/user")
	restaurant := r.Group("/restaurant", authMiddleware.AuthGuard)

	return &FiberRouter{r, hello, user, restaurant}
}
