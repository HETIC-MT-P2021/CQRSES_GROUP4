package domain

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
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
	_ = eventBus.AddHandler(events.NewArticleCreatedEventHandler(), &events.ArticleCreatedEvent{})
	_ = eventBus.AddHandler(events.NewArticleUpdatedEventHandler(), &events.ArticleUpdatedEvent{})
}

func initCommandBus() {
	// Initialize command bus and all commands available in application
	CommandBus = cqrs.NewCommandBus()
	_ = CommandBus.AddHandler(commands.NewCreateArticleCommandHandler(), &commands.CreateArticleCommand{})
	_ = CommandBus.AddHandler(commands.NewUpdateArticleCommandHandler(), &commands.UpdateArticleCommand{})
}

func initQueryBus() {
	// Initialize query bus and all queries available in application
	QueryBus = cqrs.NewQueryBus()
	_ = QueryBus.AddHandler(queries.NewReadArticleQueryHandler(), &queries.ReadArticleQuery{})
}
