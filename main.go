package main

import (
	"fmt"

	"github.com/hnguyen14/go-cart/app"
)

func main() {
	a := app.NewApp("./config.yml")
	server := a.NewServer()

	server.Run(fmt.Sprintf(":%v", a.Config.Server.Port))
}
