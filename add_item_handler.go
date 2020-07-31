package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddItemPostData ...
type AddItemPostData struct {
	URL string `form:"url" binding:"required"`
}

// FetchedData ...
type FetchedData struct {
	Error error
	Resp  []byte
}

// ParsedData ...
type ParsedData struct {
	Error error
	Page  Page
}

func fetchURL(url string) <-chan FetchedData {
	respChan := make(chan FetchedData)
	go func() {
		defer close(respChan)
		res, err := http.Get(url)
		if err != nil {
			respChan <- FetchedData{Error: err}
		}
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			respChan <- FetchedData{Error: err}
		}
		respChan <- FetchedData{Resp: content}
	}()
	return respChan
}

func asSHA256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// AddItemHandler ...
func AddItemHandler(c *gin.Context) {
	var data AddItemPostData
	if err := c.BindJSON(&data); err != nil {
		return
	}
	ch := fetchURL(data.URL)
	rsp := <-ch

	hashedURL := asSHA256(data.URL)
	fileName := fmt.Sprintf("public/tmp/%s.html", hashedURL)

	err := ioutil.WriteFile(fileName, rsp.Resp, 0644)

	if err != nil {
		log.Printf("Writing file %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": fmt.Sprintf("/html/%s.html", hashedURL),
	})
}
