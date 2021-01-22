package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetArticles to do
func GetArticles(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"events": "test",
	})
}
