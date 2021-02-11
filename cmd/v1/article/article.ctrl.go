package article

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/gin-gonic/gin"
)

// GetArticles from read-model index on elastic search
func GetArticles(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"route": "GET /articles",
	})
}

// GetArticle from read-model index on elastic search
func GetArticle(c *gin.Context) {
	aggregateArticleID := c.Param("aggregate_article_id")

	query := cqrs.NewQueryImpl(&queries.ReadArticleQuery{
		AggregateArticleID: aggregateArticleID,
	})
	article, err := domain.QueryBus.Dispatch(query)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Article not found",
		})
		return
	}

	c.JSON(http.StatusCreated, article)
}

// CreateArticle will generate a command CreateArticleCommand
func CreateArticle(c *gin.Context) {
	var req database.Article

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	command := cqrs.NewCommandImpl(&commands.CreateArticleCommand{
		Title:       req.Title,
		Description: req.Description,
	})

	err := domain.CommandBus.Dispatch(command)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status": "created",
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": 0,
			"error":  err,
		})
	}
}

// UpdateArticle will generate a command UpdateArticleCommand
func UpdateArticle(c *gin.Context) {
	aggregateArticleID := c.Param("aggregate_article_id")

	var req database.Article
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	command := cqrs.NewCommandImpl(&commands.UpdateArticleCommand{
		ID:          aggregateArticleID,
		Title:       req.Title,
		Description: req.Description,
	})

	err := domain.CommandBus.Dispatch(command)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status": "updated",
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": 0,
			"error":  err.Error(),
		})
	}
}
