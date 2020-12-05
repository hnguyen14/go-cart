package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateURLPostData ...
type CreateURLPostData struct {
	URL string `form:"url" binding:"required"`
}

// FetchedData ...
type FetchedData struct {
	Error error
	Resp  []byte
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

// CreateURLHandler ...
func (app *App) CreateURLHandler(c *gin.Context) {
	var data CreateURLPostData
	if err := c.BindJSON(&data); err != nil {
		return
	}

	urlID := ""
	stmt := "insert into urls (val) values ($1) returning id"
	err := app.DB.QueryRow(stmt, data.URL).Scan(&urlID)
	if err != nil {
		log.Printf("Adding URL to DB %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ch := fetchURL(data.URL)
	rsp := <-ch

	fileName := fmt.Sprintf("public/tmp/%s.html", urlID)

	err = ioutil.WriteFile(fileName, rsp.Resp, 0644)

	if err != nil {
		log.Printf("Writing file %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        urlID,
		"cachedURL": fmt.Sprintf("/html/%s.html", urlID),
	})
}
