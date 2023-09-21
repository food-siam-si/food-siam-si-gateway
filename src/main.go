package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/client"
	helloHdr "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/hello"
	restaurantHdr "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/restaurant"
	userHdr "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/middlewares"
	helloSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/hello"
	restaurantSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/restaurant"
	userSrv "github.com/food-siam-si/food-siam-si-gateway/src/app/services/user"
	"github.com/food-siam-si/food-siam-si-gateway/src/app/validator"
	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"github.com/food-siam-si/food-siam-si-gateway/src/proto"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := config.LoadEnv()
	v := validator.NewValidator()

	helloConn, err := grpc.Dial(config.HelloServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Failed to connect hello service %v", err)
		os.Exit(1)
	}

	restaurantConn, err := grpc.Dial(config.RestaurantServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Failed to connect restaurant service %v", err)
		os.Exit(1)
	}

	// Hello service
	helloPbClient := proto.NewHelloServiceClient(helloConn)

	helloService := helloSrv.NewService(helloPbClient)
	helloHandler := helloHdr.NewHandler(helloService, v)

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

	app := router.NewFiberRouter(authMiddleware)

	// Route Hello Initialize
	app.Hello.Get("/", helloHandler.HelloWorld)

	// Route User Initialize
	app.User.Get("/me", authMiddleware.AuthGuard, userHandler.GetCurrentUser)
	app.User.Post("/login", userHandler.Signin)
	app.User.Delete("/logout", userHandler.SignOut)
	app.User.Post("/register", userHandler.CreateUser)

	// Route Restaurant Initialize
	app.Restaurant.Post("/", restaurantHdr.CreateRestaurant)
	app.Restaurant.Get("/random", restaurantHdr.RandomRestaurant)
	app.Restaurant.Get("/type", restaurantHdr.ViewRestaurantType)
	app.Restaurant.Put("/:id", restaurantHdr.UpdateRestaurantInfo)
	app.Restaurant.Get("/:id", restaurantHdr.ViewRestaurantById)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")

		app.Shutdown()
		helloConn.Close()

		os.Exit(0)
	}()

	// Start server
	if err := app.Listen(fmt.Sprintf(":%v", config.Port)); err != nil {
		fmt.Println(err.Error())
	}
}
