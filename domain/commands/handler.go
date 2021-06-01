package commands

import (
	"errors"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/rabbit"
)

// CreateArticleCommandHandler allows to create an article
type CreateArticleCommandHandler struct{
	RabbitRepo rabbit.RabbitRepository
}

// Handle Creates a new article
func (cHandler CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		if cmd.Title == "" || cmd.Description == "" {
			return errors.New("Title or Description should not be empty")
		}

		message := rabbit.ConsumeMessage{
			EventType: events.ArticleCreatedEventType,
			Payload: events.ArticleCreatedEvent{
				Title:       cmd.Title,
				Description: cmd.Description,
			},
		}

		return cHandler.RabbitRepo.QueueConnector(message)
	default:
		return errors.New("bad command type")
	}
}

// NewCreateArticleCommandHandler Creates an instance
func NewCreateArticleCommandHandler(repo rabbit.RabbitRepository) *CreateArticleCommandHandler {
	return &CreateArticleCommandHandler{
		RabbitRepo: repo,
	}
}

// UpdateArticleCommandHandler allows to update an article
type UpdateArticleCommandHandler struct{
	RabbitRepo rabbit.RabbitRepository
}

// Handle Updates a new article
func (cHandler UpdateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *UpdateArticleCommand:
		if cmd.Title == "" || cmd.Description == "" {
			return errors.New("Title or Description should not be empty")
		}

		message := rabbit.ConsumeMessage{
			EventType: events.ArticleUpdatedEventType,
			Payload: events.ArticleUpdatedEvent{
				AggregateArticleID: cmd.AggregateArticleID,
				Title:              cmd.Title,
				Description:        cmd.Description,
			},
		}

		return cHandler.RabbitRepo.QueueConnector(message)
	default:
		return errors.New("bad command type")
	}
}

// NewUpdateArticleCommandHandler Creates an instance
func NewUpdateArticleCommandHandler(repo rabbit.RabbitRepository) *UpdateArticleCommandHandler {
	return &UpdateArticleCommandHandler{
		RabbitRepo: repo,
	}
}
