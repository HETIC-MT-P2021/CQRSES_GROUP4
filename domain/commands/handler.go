package commands

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/messager"
)

// CreateArticleCommandHandler allows to create an article
type CreateArticleCommandHandler struct{}

// Handle Creates a new article
func (cHandler CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		fmt.Println(cmd)
		// Call QueueConnector
		connector, err := messager.ConnectToRabbitMQ()
		if err != nil {
			return err
		}

		return connector.Publish()
	default:
		return errors.New("bad command type")
	}
}

// NewCreateArticleCommandHandler Creates an instance
func NewCreateArticleCommandHandler() *CreateArticleCommandHandler {
	return &CreateArticleCommandHandler{}
}

// UpdateArticleCommandHandler allows to update an article
type UpdateArticleCommandHandler struct{}

// Handle Updates a new article
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

// NewUpdateArticleCommandHandler Creates an instance
func NewUpdateArticleCommandHandler() *UpdateArticleCommandHandler {
	return &UpdateArticleCommandHandler{}
}
