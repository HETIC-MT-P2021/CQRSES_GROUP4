package health_check

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	log.Print(database.DbConn)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
