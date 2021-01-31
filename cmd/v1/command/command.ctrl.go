package command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/cqrs"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain"
	commands "github.com/jibe0123/CQRSES_GROUP4/pkg/domain/commands"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/helper"
)

type requestCommandName struct {
	Name string
}

// CreateNewCommand Allows to calls Command to manage article
func CreateNewCommand(c *gin.Context) {
	var req requestCommandName

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var command *cqrs.CommandImpl

	switch req.Name {
	case helper.TypeOf(&commands.CreateArticleCommand{}):
		command = cqrs.NewCommandImpl(&commands.CreateArticleCommand{})
	case helper.TypeOf(&commands.UpdateArticleCommand{}):
		command = cqrs.NewCommandImpl(&commands.UpdateArticleCommand{})
	}

	err := domain.CommandBus.Dispatch(command)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status": 1,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": 0,
			"error":  err,
		})
	}

}
