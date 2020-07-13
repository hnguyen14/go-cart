package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddItemPostData ...
type AddItemPostData struct {
	URL string `form:"url" binding:"required"`
}

func fetchURL(url string) (*Page, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Fetch URL %w", err)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read page content %w", err)
	}

	page, err := ParseHTML(content)
	if err != nil {
		return nil, fmt.Errorf("Parse content %w", err)
	}
	page.URL = url

	return page, nil
}

// AddItemHandler ...
func AddItemHandler(c *gin.Context) {
	var data AddItemPostData
	if err := c.BindJSON(&data); err != nil {
		return
	}

	page, err := fetchURL(data.URL)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusCreated, page)
}
