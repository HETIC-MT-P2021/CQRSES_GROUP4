package healthcheck

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/gin-gonic/gin"
)

// HealthCheck Checks database is running
func HealthCheck(c *gin.Context) {
	log.Print(database.DbConn)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
