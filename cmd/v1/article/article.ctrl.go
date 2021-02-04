package article

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetArticle from read-model index on elastic search
func GetArticle(c *gin.Context) {
	query := cqrs.NewQueryImpl(&queries.ReadArticlesQuery{})
	article, err := domain.QueryBus.Dispatch(query)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Article not fount",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"articles": article,
	})
}

// CreateArticle from pkg/state/
func CreateArticle(c *gin.Context) {
	id := uuid.NewV4()
	article := database.Article{
		ID:          id.String(),
		Title:       "hello",
		Description: "NIce",
	}
	err := elasticsearch.StoreReadmodel(article)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"aggregate_id": id.String(),
	})
}
