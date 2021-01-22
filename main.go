package main

import (
	"fmt"

	"github.com/hnguyen14/go-cart/app"
)

func main() {
	a := app.NewApp("./config.yml")
	server := a.NewServer()

	fmt.Sprintf("Starting Go Server now")
	fmt.Sprintf("Starting Go Server now")
	fmt.Sprintf("Starting Go Server now")
	fmt.Sprintf("Starting Go Server now")
	fmt.Sprintf("Starting Go Server now")
	fmt.Sprintf("Starting Go Server now")
	fmt.Sprintf("Starting Go Server now")
	server.Run(fmt.Sprintf(":%v", a.Config.Server.Port))
}
