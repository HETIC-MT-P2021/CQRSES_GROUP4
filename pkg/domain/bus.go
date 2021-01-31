package domain

import (
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events_handler"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

var EventBus *event.EventBus

func InitBusses() {
	EventBus = event.NewEventBus()

	_ = EventBus.AddHandler(events_handler.NewArticleCreatedEventHandler(), &events.ArticleCreatedEvent{})
	_ = EventBus.AddHandler(events_handler.NewArticleUpdatedEventHandler(), &events.ArticleUpdatedEvent{})
}
