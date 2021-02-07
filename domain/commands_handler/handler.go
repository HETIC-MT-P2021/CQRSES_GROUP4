package commands_handler

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

type CreateArticleCommandHandler struct {
	eventBus *event.EventBus
}

func (cHandler CreateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *commands.CreateArticleCommand:
		fmt.Println(cmd)
		// Call QueueConnector

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

type UpdateArticleCommandHandler struct {
	eventBus *event.EventBus
}

func (cHandler UpdateArticleCommandHandler) Handle(command cqrs.Command) error {
	switch cmd := command.Payload().(type) {
	case *commands.UpdateArticleCommand:
		fmt.Println(cmd)
		// Call QueueConnector
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewUpdateArticleCommandHandler(eventBus *event.EventBus) *UpdateArticleCommandHandler {
	return &UpdateArticleCommandHandler{}
}
