package v1

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/article"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/auth"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/fixture"
	healthcheck "github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/health_check"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	jwtAuth, err := jwt_auth.NewAuth()
	if err != nil {
		log.Println(err)
	}

	app := r.Group("/api/v1/")
	{
		healthcheck.ApplyRoutes(app)
		auth.ApplyRoutes(app, jwtAuth)
		article.ApplyRoutes(app, jwtAuth)
		fixture.ApplyRoutes(app)
	}
}
