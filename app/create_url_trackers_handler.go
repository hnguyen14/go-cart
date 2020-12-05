package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tracker ...
type Tracker struct {
	Name     string `json:"name"`
	Selector string `json:"selector"`
}

// CreateURLTrackersData ...
type CreateURLTrackersData struct {
	TrackerList []Tracker `json:"trackers"`
}

// CreateURLTrackersHandler ...
func (app *App) CreateURLTrackersHandler(c *gin.Context) {
	var data CreateURLTrackersData

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Read Request Body %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Unmarshal Request Payload %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, t := range data.TrackerList {
		stmt := "insert into trackers (name, selector) values ($1, $2) returning id"
		var trackerID string
		app.DB.QueryRow(stmt, t.Name, t.Selector).Scan(&trackerID)
		fmt.Printf("------ assoc tracker %s, %s", c.Param("urlID"), trackerID)
		stmt = "insert into url_tracker (url_id, tracker_id) values($1, $2)"
		app.DB.Exec(stmt, c.Param("urlID"), trackerID)
	}

	c.String(http.StatusCreated, "")
}
