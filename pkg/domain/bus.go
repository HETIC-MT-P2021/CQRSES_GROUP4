package domain

import (
	"github.com/jibe0123/CQRSES_GROUP4/pkg/cqrs"
	commands "github.com/jibe0123/CQRSES_GROUP4/pkg/domain/commands"
	commands_handler "github.com/jibe0123/CQRSES_GROUP4/pkg/domain/commands_handler"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events_handler"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/queries"
	queries_handler "github.com/jibe0123/CQRSES_GROUP4/pkg/domain/queries_handler"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

// EventBus Allow to stores all event on database
var EventBus *event.EventBus

// CommandBus Allows to manage write model
var CommandBus *cqrs.CommandBus

// QueryBus Allows to manage read model
var QueryBus *cqrs.QueryBus

// InitBusses Init event, command and query busses
func InitBusses() {
	initEventBus()
	initCommandBus()
	initQueryBus()
}

func initEventBus() {
	// Initialize event bus and all events available in application
	EventBus = event.NewEventBus()
	_ = EventBus.AddHandler(events_handler.NewArticleCreatedEventHandler(), &events.ArticleCreatedEvent{})
	_ = EventBus.AddHandler(events_handler.NewArticleUpdatedEventHandler(), &events.ArticleUpdatedEvent{})
}

func initCommandBus() {
	// Initialize command bus and all commands available in application
	CommandBus = cqrs.NewCommandBus()
	_ = CommandBus.AddHandler(commands_handler.NewCreateArticleCommandHandler(), &commands.CreateArticleCommand{})
	_ = CommandBus.AddHandler(commands_handler.NewCreateArticleCommandHandler(), &commands.UpdateArticleCommand{})
}

func initQueryBus() {
	// Initialize query bus and all queries available in application
	QueryBus = cqrs.NewQueryBus()
	_ = QueryBus.AddHandler(queries_handler.NewReadArticlesQueryHandler(), &queries.ReadArticlesQuery{})

}
