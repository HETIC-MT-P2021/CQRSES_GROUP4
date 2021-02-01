package article

import (
	"github.com/gin-gonic/gin"
	"github.com/jibe0123/CQRSES_GROUP4/cmd/v1/auth"
	jwt "github.com/kyfk/gin-jwt"
)

func ApplyRoutes(r *gin.RouterGroup) {
	r.Use(jwt.ErrorHandler)
	r.GET("/articles", auth.Operator(auth.Auth), GetArticles)
}
