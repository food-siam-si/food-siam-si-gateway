package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/food-siam-si/food-siam-si-gateway/src/config"
	"github.com/food-siam-si/food-siam-si-gateway/src/router"
)

func main() {
	config, err := config.LoadEnv()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	app := router.NewFiberRouter()

	if err := app.Listen(fmt.Sprintf(":%v", config.Port)); err != nil {
		fmt.Println(err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")

		app.Shutdown()
		os.Exit(0)
	}()

}
