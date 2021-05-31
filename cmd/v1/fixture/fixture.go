package fixture

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

// ApplyRoutes All routes for fixtures
func ApplyRoutes(r *gin.RouterGroup, jwtAuth jwt.Auth) {
	r.Use(jwt.ErrorHandler)
	r.POST("/fixtures/event-store", jwt_auth.Operator(jwtAuth), CreateEventStore)
	r.POST("/fixtures/read-model", jwt_auth.Operator(jwtAuth), CreateReadModel)
}
