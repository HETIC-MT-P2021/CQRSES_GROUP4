package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type article struct {
}

// CreateArticle in database
func CreateArticle(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "article_created",
	})
}
