package events

import (
	"errors"
	"os"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
)

// ArticleCreatedEventHandler allows to create an article
type ArticleCreatedEventHandler struct{}

// Handle Creates a new article
func (eHandler ArticleCreatedEventHandler) Handle(ev event.Event) error {
	switch evType := ev.Type(); evType {
	case ArticleCreatedEventType:
		if os.Getenv("APP_ENV") != pkg.Test {
			return ArticleCreatedEvent{}.Apply(ev)
		}
		return nil
	default:
		return errors.New("bad event")
	}
}

// NewArticleCreatedEventHandler Creates an instance
func NewArticleCreatedEventHandler() *ArticleCreatedEventHandler {
	return &ArticleCreatedEventHandler{}
}

// ArticleUpdatedEventHandler allows to update an article
type ArticleUpdatedEventHandler struct{}

// Handle Updates a new article
func (eHandler ArticleUpdatedEventHandler) Handle(ev event.Event) error {
	switch evType := ev.Type(); evType {
	case ArticleUpdatedEventType:
		return ArticleUpdatedEvent{}.Apply(ev)
	default:
		return errors.New("bad event")
	}
}

// NewArticleUpdatedEventHandler Creates an instance
func NewArticleUpdatedEventHandler() *ArticleUpdatedEventHandler {
	return &ArticleUpdatedEventHandler{}
}
