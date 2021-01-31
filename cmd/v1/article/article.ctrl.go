package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/article"
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
)

// GetArticles to do
func GetArticles(c *gin.Context) {
	var req db.Article

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	articles, err := article.UpdateArticle(req)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"events": "error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"articles": articles.Articles(),
	})
}
