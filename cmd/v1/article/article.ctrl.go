package article

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/cqrs"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/queries"
)

// GetArticles from pkg/state/
func GetArticles(c *gin.Context) {
	query := cqrs.NewQueryImpl(&queries.ReadArticlesQuery{})
	articles, err := domain.QueryBus.Dispatch(query)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Articles not fount",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"articles": articles.([]database.Article),
	})
}
