package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// URL ...
type URL struct {
	ID  string
	URL string
}

// URLHandler ...
func (app *App) URLHandler(c *gin.Context) {
	var url URL
	stmt := "select * from urls where id = $1"
	err := app.DB.
		QueryRow(stmt, c.Param("urlID")).
		Scan(&url.ID, &url.URL)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "url.tmpl", gin.H{
		"urlID": url.ID,
		"URL":   url.URL,
	})
}
