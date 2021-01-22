package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/CQRSES_GROUP4/cmd/v1/auth"
	"github.com/jibe0123/CQRSES_GROUP4/cmd/v1/event"
	"github.com/jibe0123/CQRSES_GROUP4/cmd/v1/health_check"
)

func ApplyRoutes(r *gin.Engine) {
	app := r.Group("/api/v1/")
	{
		health_check.ApplyRoutes(app)
		auth.ApplyRoutes(app)
		event.ApplyRoutes(app)
	}
}
