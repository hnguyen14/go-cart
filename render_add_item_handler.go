package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderAddItemHandler ...
func RenderAddItemHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "add_item.tmpl", gin.H{})
}
