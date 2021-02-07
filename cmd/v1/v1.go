package v1

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/article"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/auth"
	healthcheck "github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/health_check"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for v1
func ApplyRoutes(r *gin.Engine) {
	app := r.Group("/api/v1/")
	{
		healthcheck.ApplyRoutes(app)
		auth.ApplyRoutes(app)
		article.ApplyRoutes(app)
	}
}
