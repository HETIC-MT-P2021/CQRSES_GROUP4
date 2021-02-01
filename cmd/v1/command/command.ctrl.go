package command

import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
	"github.com/gin-gonic/gin"
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
	case pkg.TypeOf(&commands.CreateArticleCommand{}):
		command = cqrs.NewCommandImpl(&commands.CreateArticleCommand{})
	case pkg.TypeOf(&commands.UpdateArticleCommand{}):
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
