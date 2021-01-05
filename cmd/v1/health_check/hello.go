package health_check

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	ping := r.Group("/health_check")
	{
		ping.GET("/", HealthCheck)
	}
}
