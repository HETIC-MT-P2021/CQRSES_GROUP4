package commands_handler

import (
	"errors"
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/cqrs"
	commands "github.com/jibe0123/CQRSES_GROUP4/domain/commands"
	"github.com/jibe0123/CQRSES_GROUP4/domain/events"
	"github.com/jibe0123/CQRSES_GROUP4/event"
)

type CreateArticleCommandHandler struct {
	eventBus *event.EventBus
}

func (eh CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *commands.CreateArticleCommand:
		fmt.Println("Handler.")
		fmt.Println(cmd)

		event := event.NewEventImpl(&events.ArticleCreatedEvent{})
		err := eh.eventBus.Dispatch(event)
		if err != nil {
			return err
		}

		return nil
	default:
		return errors.New("bad command type")
	}
}

func NewCreateArticleCommandHandler(eventBus *event.EventBus) *CreateArticleCommandHandler {
	return &CreateArticleCommandHandler{
		eventBus: eventBus,
	}
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
