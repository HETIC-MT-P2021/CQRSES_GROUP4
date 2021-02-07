package domain

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands_handler"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events_handler"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries_handler"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

// EventBus Allow to stores all event on database
var eventBus *event.EventBus

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
	eventBus = event.NewEventBus()
	_ = eventBus.AddHandler(events_handler.NewArticleCreatedEventHandler(), &events.ArticleCreatedEvent{})
	_ = eventBus.AddHandler(events_handler.NewArticleUpdatedEventHandler(), &events.ArticleUpdatedEvent{})
}

func initCommandBus() {
	// Initialize command bus and all commands available in application
	CommandBus = cqrs.NewCommandBus()
	_ = CommandBus.AddHandler(commands_handler.NewCreateArticleCommandHandler(eventBus), &commands.CreateArticleCommand{})
	_ = CommandBus.AddHandler(commands_handler.NewUpdateArticleCommandHandler(eventBus), &commands.UpdateArticleCommand{})
}

func initQueryBus() {
	// Initialize query bus and all queries available in application
	QueryBus = cqrs.NewQueryBus()
	_ = QueryBus.AddHandler(queries_handler.NewReadArticlesQueryHandler(), &queries.ReadArticleQuery{})

}
