package main

import (
	"io/ioutil"
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

func parseContent(fetchedDataChan <-chan FetchedData) <-chan ParsedData {
	parsed := make(chan ParsedData)
	go func() {
		defer close(parsed)
		fetchedData := <-fetchedDataChan
		if fetchedData.Error != nil {
			parsed <- ParsedData{Error: fetchedData.Error}
		}
		page, err := ParseHTML(fetchedData.Resp)
		if err != nil {
			parsed <- ParsedData{Error: err}
		}
		parsed <- ParsedData{Page: *page}

	}()
	return parsed
}

// AddItemHandler ...
func AddItemHandler(c *gin.Context) {
	var data AddItemPostData
	if err := c.BindJSON(&data); err != nil {
		return
	}

	parseChan := parseContent(fetchURL(data.URL))
	parsed := <-parseChan

	if parsed.Error != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusCreated, parsed.Page)
}
