package commands

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
)

type CreateArticleCommandHandler struct{}

func (cHandler CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		fmt.Println(cmd)
		// Call QueueConnector

		return nil
	default:
		return errors.New("bad command type")
	}
}

func NewCreateArticleCommandHandler() *CreateArticleCommandHandler {
	return &CreateArticleCommandHandler{}
}

type UpdateArticleCommandHandler struct{}

func (cHandler UpdateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *UpdateArticleCommand:
		fmt.Println(cmd)
		// Call QueueConnector
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewUpdateArticleCommandHandler() *UpdateArticleCommandHandler {
	return &UpdateArticleCommandHandler{}
}
