package article

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cmd/v1/auth"
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

func ApplyRoutes(r *gin.RouterGroup) {
	r.Use(jwt.ErrorHandler)
	r.GET("/articles", auth.Operator(auth.Auth), GetArticles)
}
