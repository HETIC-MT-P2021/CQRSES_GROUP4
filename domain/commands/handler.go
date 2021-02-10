package commands

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/rabbit"
)

// CreateArticleCommandHandler allows to create an article
type CreateArticleCommandHandler struct{}

// Handle Creates a new article
func (cHandler CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		message := events.ConsumeMessage{
			EventType: pkg.TypeOf(&events.ArticleCreatedEvent{}),
			Payload: events.ArticleCreatedEvent{
				ID:          cmd.ID,
				Title:       cmd.Title,
				Description: cmd.Description,
			},
		}

		return rabbit.QueueConnector(message)
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
