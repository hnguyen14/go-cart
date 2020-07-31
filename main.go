package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	config *Config
)

func main() {
	config, err := NewConfig("./config.yml")
	if err != nil {
		log.Fatal("Cannot load config file")
		return
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("/html", "./public/tmp")
	r.LoadHTMLGlob("templates/*")
	r.GET("/addItem", RenderAddItemHandler)
	api := r.Group("api")
	{
		api.POST("/items", AddItemHandler)
	}

	r.Run(fmt.Sprintf(":%v", config.Server.Port))
}
