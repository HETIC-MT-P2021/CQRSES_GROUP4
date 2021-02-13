package fixture

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes All routes for fixtures
func ApplyRoutes(r *gin.RouterGroup) {
	//r.Use(jwt.ErrorHandler)
	r.POST("/fixtures/event-store" /*auth.Operator(auth.Auth),*/, CreateEventStore)
	r.POST("/fixtures/read-model" /*auth.Operator(auth.Auth),*/, CreateReadModel)
}
