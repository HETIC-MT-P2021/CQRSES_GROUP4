package events_handler

import (
	"errors"

	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/state"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

type ArticleCreatedEventHandler struct{}

func (ch ArticleCreatedEventHandler) Handle(ev event.Event) error {

	switch ev := ev.Payload().(type) {
	case *events.ArticleCreatedEvent:
		aggr := state.CurrentArticles

		err := aggr.Add(ev)
		if err != nil {
			return err
		}

		// if no errors push on db
		return nil
	default:
		return errors.New("bad event")
	}
}

func NewArticleCreatedEventHandler() *ArticleCreatedEventHandler {
	return &ArticleCreatedEventHandler{}
}

type ArticleUpdatedEventHandler struct{}

func (ch ArticleUpdatedEventHandler) Handle(ev event.Event) error {
	switch ev := ev.Payload().(type) {
	case *events.ArticleUpdatedEvent:
		aggr := state.CurrentArticles

		err := aggr.Update(1, ev)
		if err != nil {
			return err
		}

		// if no errors push on db
		return nil
	default:
		return errors.New("bad event")
	}

}

func NewArticleUpdatedEventHandler() *ArticleUpdatedEventHandler {
	return &ArticleUpdatedEventHandler{}
}
