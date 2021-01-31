package commands_

import (
	"errors"
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/pkg/cqrs"
	commands "github.com/jibe0123/CQRSES_GROUP4/pkg/domain/commands"
)

type CreateArticleCommandHandler struct{}

func (ch CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *commands.CreateArticleCommand:
		fmt.Println("Handler.")
		fmt.Println(cmd)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateArticleCommandHandler() *CreateArticleCommandHandler {
	return &CreateArticleCommandHandler{}
}

type UpdateArticleCommandHandler struct{}

func (ch UpdateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *commands.CreateArticleCommand:
		fmt.Println("Handler.")
		fmt.Println(cmd)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewUpdateArticleCommandHandler() *UpdateArticleCommandHandler {
	return &UpdateArticleCommandHandler{}
}
