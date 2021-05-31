package fixture

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/fixtures"
	"github.com/gin-gonic/gin"
)

// CreateEventStore Insert event in elastic search
// @Summary Create event in elastic search
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures/event-store [post]
func CreateEventStore(c *gin.Context) {
	if err := fixtures.NewEventStore(); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}

// CreateReadModel Insert read-model in elastic search
// @Summary Create read-model in elastic search
// @Tags fixtures
// @Accept  json
// @Produce  json
// @Success 201 {object} pkg.HTTPStatus "Created"
// @Failure 500 {object} pkg.HTTPStatus "Not Created"
// @Router /fixtures/read-model [post]
func CreateReadModel(c *gin.Context) {
	if err := fixtures.NewReadModel(); err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "Not created",
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
	})
}
