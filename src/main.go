package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/client"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/handler/menu"
	restaurantHdr "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/restaurant"
	reviewHdr "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/review"
	userHdr "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/middlewares"
	menuSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/menu"
	restaurantSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	reviewSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/review"
	userSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"github.com/food-siam-si/food-siam-si-gateway/src/proto"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/router"
)

func main() {
	config := config.LoadEnv()
	v := validator.NewValidator()

	restaurantConn, err := client.NewRestaurantClient(config)

	if err != nil {
		log.Printf("Failed to connect restaurant service %v", err)
		os.Exit(1)
	}

	// User service
	userClient := client.NewUserClient(config)

	userService := userSrv.NewService(userClient)
	userHandler := userHdr.NewHandler(userService, v)

	// Restaurant service
	restaurantClient := proto.NewRestaurantServiceClient(restaurantConn)
	restaurantTypeClient := proto.NewRestaurantTypeServiceClient(restaurantConn)

	restaurantService := restaurantSrv.NewService(restaurantClient, restaurantTypeClient)
	restaurantHdr := restaurantHdr.NewHandler(restaurantService, v)

	authMiddleware := middlewares.NewAuthMiddleware(userService)
	restaurantMiddleware := middlewares.NewRestaurantMiddleware(restaurantService)

	// Review Service
	reviewClient := client.NewReviewClient(config)
	reviewService := reviewSrv.NewService(reviewClient)
	reviewHdr := reviewHdr.NewHandler(restaurantService, reviewService, v)

	// Menu Service
	menuClient := client.NewMenuClient(config)
	menuService := menuSrv.NewService(menuClient)
	menuHdr := menu.NewHandler(menuService, v)

	app := router.NewAppRouter(authMiddleware)

	// Route Hello Initialize
	// app.Hello.Get("/", helloHandler.HelloWorld)

	// Route User Initialize
	app.User.Get("/me", authMiddleware.AuthGuard, userHandler.GetCurrentUser)
	app.User.Post("/login", userHandler.Signin)
	app.User.Post("/register", userHandler.CreateUser)

	// Route Restaurant Initialize
	app.Restaurant.Post("/", restaurantHdr.CreateRestaurant)
	app.Restaurant.Put("/me", authMiddleware.RestaurantGuard, restaurantHdr.UpdateRestaurantInfo)
	app.Restaurant.Get("/me", authMiddleware.RestaurantGuard, restaurantHdr.GetCurrentRestaurant)
	app.Restaurant.Get("/random", authMiddleware.CustomerGuard, restaurantHdr.RandomRestaurant)
	app.Restaurant.Get("/type", restaurantHdr.ViewRestaurantType)
	app.Restaurant.Get("/:id", authMiddleware.CustomerGuard, restaurantHdr.ViewRestaurantById)

	// Route Review Initialize
	app.Review.Get("/:restaurantId", restaurantMiddleware.OwnerOrCustomerGuard, reviewHdr.GetReview)
	app.Review.Post("/:restaurantId", authMiddleware.CustomerGuard, reviewHdr.CreateReview)

	// Route Menu Initialize
	app.Restaurant.Get("/:restaurantId/menus", restaurantMiddleware.OwnerOrCustomerGuard, menuHdr.GetMenus)
	app.Restaurant.Post("/:restaurantId/menus", authMiddleware.RestaurantGuard, restaurantMiddleware.OwnerGuard, menuHdr.CreateMenu)
	app.Restaurant.Put("/:restaurantId/menus/:menuId", authMiddleware.RestaurantGuard, restaurantMiddleware.OwnerGuard, menuHdr.UpdateMenu)
	app.Restaurant.Delete("/:restaurantId/menus/:menuId", authMiddleware.RestaurantGuard, restaurantMiddleware.OwnerGuard, menuHdr.DeleteMenu)
	app.Restaurant.Get("/:restaurantId/menus/random", restaurantMiddleware.OwnerOrCustomerGuard, menuHdr.RandomMenu)
	app.Restaurant.Get("/:restaurantId/menus/recommend", restaurantMiddleware.OwnerOrCustomerGuard, menuHdr.GetRecommendMenu)
	app.Restaurant.Put("/:restaurantId/menus/:menuId/recommend", authMiddleware.RestaurantGuard, restaurantMiddleware.OwnerGuard, menuHdr.UpdateRecommendMenu)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")

		app.Shutdown()
		restaurantConn.Close()

		os.Exit(0)
	}()

	// Start server
	if err := app.Listen(fmt.Sprintf(":%v", config.Port)); err != nil {
		fmt.Println(err.Error())
	}
}
