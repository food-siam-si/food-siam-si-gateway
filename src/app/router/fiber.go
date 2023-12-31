package router

import (
	"log"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type AppRouter struct {
	*fiber.App
	Hello      fiber.Router
	User       fiber.Router
	Restaurant fiber.Router
	Review     fiber.Router
	Menu       fiber.Router
}

func NewAppRouter(authMiddleware middlewares.IAuthMiddleware) *AppRouter {
	r := fiber.New(fiber.Config{
		BodyLimit: 16 * 1024 * 1024,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Println(err.Error())
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}

			ctx.Status(code)
			ctx.JSON(fiber.Map{
				"message": message,
			})

			return nil
		},
	})

	r.Use(recover.New())
	r.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,https://food-siam-si.miello.dev",
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	hello := r.Group("/hello")
	user := r.Group("/user")
	restaurant := r.Group("/restaurant", authMiddleware.AuthGuard)
	menu := r.Group("/menu", authMiddleware.AuthGuard)
	review := r.Group("/review", authMiddleware.AuthGuard)

	return &AppRouter{r, hello, user, restaurant, review, menu}
}
