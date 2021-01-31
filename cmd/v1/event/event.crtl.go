package event

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/database/elasticsearch"
	uuid "github.com/satori/go.uuid"
)

// GetEvents from elasticsearch database
func GetEvents(c *gin.Context) {
	events, err := elasticsearch.LoadEvents()
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"events": events,
	})
}

// CreateEvent in elasticsearch database
func CreateEvent(c *gin.Context) {
	var req db.RequestCreate

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	id := uuid.NewV4()

	currentDate := time.Now()
	timestamp := currentDate.Unix()

	event := db.Event{
		ID:        id.String(),
		EventName: "ArticleCreatedEvent",
		CreatedAt: strconv.FormatInt(timestamp, 10),
		Payload: db.Article{
			ID:          2,
			Title:       req.Title,
			Description: req.Description,
		},
	}

	err := elasticsearch.StoreEvent(event)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"id":     id.String(),
	})
}
