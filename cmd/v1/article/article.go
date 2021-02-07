package article

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	//r.Use(jwt.ErrorHandler)
	r.GET("/articles" /*auth.Operator(auth.Auth),*/, GetArticles)
	r.GET("/articles/:aggregate_article_id" /*auth.Operator(auth.Auth),*/, GetArticle)
	r.POST("/articles" /*auth.Operator(auth.Auth),*/, CreateArticle)
	r.PUT("/articles/:aggregate_article_id" /*auth.Operator(auth.Auth),*/, UpdateArticle)
}
