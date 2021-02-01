package v1

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/article"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/auth"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/command"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/health_check"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	app := r.Group("/api/v1/")
	{
		health_check.ApplyRoutes(app)
		auth.ApplyRoutes(app)
		command.ApplyRoutes(app)
		event.ApplyRoutes(app)
		article.ApplyRoutes(app)
	}
}
