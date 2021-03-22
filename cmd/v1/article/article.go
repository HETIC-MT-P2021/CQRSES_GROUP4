package article

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

// ApplyRoutes All routes for articles
func ApplyRoutes(r *gin.RouterGroup, jwtAuth jwt.Auth) {
	r.Use(jwt.ErrorHandler)
	r.GET("/articles", jwt_auth.Operator(jwtAuth), GetArticles)
	r.GET("/articles/:aggregate_article_id" /*auth.Operator(auth.Auth),*/, GetArticle)
	r.POST("/articles" /*auth.Operator(auth.Auth),*/, CreateArticle)
	r.PUT("/articles/:aggregate_article_id" /*auth.Operator(auth.Auth),*/, UpdateArticle)
}
