package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/survey/cmd/v1/article"
	"github.com/jibe0123/survey/cmd/v1/auth"
	"github.com/jibe0123/survey/cmd/v1/health_check"
)

func ApplyRoutes(r *gin.Engine) {
	app := r.Group("/api/v1/")
	{
		health_check.ApplyRoutes(app)
		auth.ApplyRoutes(app)
		article.ApplyRoutes(app)
	}
}
