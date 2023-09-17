package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	helloHandler "github.com/food-siam-si/food-siam-si-gateway/src/app/handler/hello"
	helloService "github.com/food-siam-si/food-siam-si-gateway/src/app/services/hello"
	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"github.com/food-siam-si/food-siam-si-gateway/src/proto"

	"github.com/food-siam-si/food-siam-si-gateway/src/app/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := config.LoadEnv()

	app := router.NewFiberRouter()

	helloConn, err := grpc.Dial(config.HelloServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Failed to connect hello service %v", err)
		os.Exit(1)
	}

	helloPbClient := proto.NewHelloServiceClient(helloConn)

	helloService := helloService.NewService(helloPbClient)
	helloHandler := helloHandler.NewHandler(helloService)

	app.Hello.Get("/", helloHandler.HelloWorld)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")

		app.Shutdown()
		os.Exit(0)
	}()

	// Start server
	if err := app.Listen(fmt.Sprintf(":%v", config.Port)); err != nil {
		fmt.Println(err.Error())
	}
}
