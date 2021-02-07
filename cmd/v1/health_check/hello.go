package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for health check
func ApplyRoutes(r *gin.RouterGroup) {
	ping := r.Group("/health_check")
	{
		ping.GET("/", HealthCheck)
	}
}
