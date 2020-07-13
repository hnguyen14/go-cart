package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	config *Config
)

func main() {
	config, err := NewConfig("./config.yml")
	if err != nil {
		fmt.Printf("Cannot load config %v", err)
		return
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/item", AddItemHandler)
	r.Run(fmt.Sprintf(":%v", config.Server.Port))
}
