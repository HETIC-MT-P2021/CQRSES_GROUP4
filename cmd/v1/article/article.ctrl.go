package article

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/gin-gonic/gin"
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
