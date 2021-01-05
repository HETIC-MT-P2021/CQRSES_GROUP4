package health_check

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/survey/pkg/database"
	"log"
)

func HealthCheck(c *gin.Context) {
	log.Print(database.DbConn)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
