package event

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	"github.com/gin-gonic/gin"
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
	var req database.RequestCreateEvent

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	eventID := uuid.NewV4()
	aggregateArticleID := uuid.NewV4()

	currentDate := time.Now()
	timestamp := currentDate.Unix()

	event := database.Event{
		ID:        eventID.String(),
		EventType: req.EventType,
		CreatedAt: strconv.FormatInt(timestamp, 10),
		Payload: database.Article{
			ID:          aggregateArticleID.String(),
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
		"status":               "event created",
		"event_id":             eventID.String(),
		"aggregate_article_id": aggregateArticleID.String(),
	})
}
