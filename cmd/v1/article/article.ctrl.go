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

// GetArticles returns array of article from elastic search
// @Summary Get all articles from elastic search
// @Description Get an array of article struct
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"GET /articles"
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"route": "GET /articles",
	})
}

// GetArticle returns article from read-model index from elastic search
// @Summary Get an article from elastic search
// @Description Get article struct
// @Tags articles
// @Accept  json
// @Produce  json
// @Param aggregate_article_id path int true "Article ID"
// @Success 200 {object} database.Article
// @Failure 404 {object} pkg.HTTPError "Article Not found"
// @Router /articles/{aggregate_article_id} [get]
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

type Request struct {
	Title string
	Description string
}

// CreateArticle will generate a command CreateArticleCommand
// @Summary Create article in elastic search
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body Request true "Add article"
// @Success 201 {object} pkg.HTTPStatus "created"
// @Failure 500 {object} pkg.HTTPError "Internal Server Error"
// @Router /articles [post]
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
			"status": 500,
			"error":  err,
		})
	}
}

// UpdateArticle will generate a command UpdateArticleCommand
// @Summary Update article in elastic search
// @Tags articles
// @Accept  json
// @Produce  json
// @Param aggregate_article_id path int true "Article ID"
// @Param article body Request true "Update article"
// @Success 201 {object} pkg.HTTPStatus "updated"
// @Failure 500 {object} pkg.HTTPError "Internal Server Error"
// @Router /articles/{aggregate_article_id}} [put]
func UpdateArticle(c *gin.Context) {
	aggregateArticleID := c.Param("aggregate_article_id")

	var req database.Article
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	command := cqrs.NewCommandImpl(&commands.UpdateArticleCommand{
		AggregateArticleID: aggregateArticleID,
		Title:              req.Title,
		Description:        req.Description,
	})

	err := domain.CommandBus.Dispatch(command)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status": "updated",
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
	}
}
