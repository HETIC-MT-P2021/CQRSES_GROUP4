package fixture

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/fixtures"
	"github.com/gin-gonic/gin"
)

// CreateEventStore Insert event-store
func CreateEventStore(c *gin.Context) {
	if err := fixtures.NewEventStore(); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Not created",
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}

// CreateReadModel Insert read-model
func CreateReadModel(c *gin.Context) {
	if err := fixtures.NewReadModel(); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Not created",
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
