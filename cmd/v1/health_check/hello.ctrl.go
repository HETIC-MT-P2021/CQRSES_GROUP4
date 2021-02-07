package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck Checks database is running
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
