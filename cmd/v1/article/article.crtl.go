package article

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	dmn "github.com/jibe0123/CQRSES_GROUP4/pkg/domain"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/commands"
)

// GetArticle in database
func GetArticle(c *gin.Context) {
	articles := dmn.LoadArticles()

	c.JSON(http.StatusCreated, gin.H{
		"articles": articles,
	})
}

var index int = 1

// CreateArticle in database
func CreateArticle(c *gin.Context) {
	article := db.Article{
		ID:          2,
		Title:       "second article",
		Description: "un deuxieme",
	}

	fmt.Println(index)

	command := commands.Command{}
	if index == 1 {
		command.Type = "CreateArticle"
		command.Payload = commands.CreateArticleCommand{
			Article: article,
		}

		index++
	} else if index == 2 {
		article = db.Article{
			ID:          2,
			Title:       "article edited",
			Description: "edited",
		}

		command.Type = "UpdateArticle"
		command.Payload = commands.UpdateArticleCommand{
			Article:   article,
			ArticleID: 2,
		}
		index = 1
	} else {
		command.Type = "Nothing"
	}

	err := dmn.SaveArticle(command)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "yes",
	})
}
