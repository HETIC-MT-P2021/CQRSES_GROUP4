package events

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

// ArticleCreatedEventHandler allows to create an article
type ArticleCreatedEventHandler struct{}

// Handle Creates a new article
func (eHandler ArticleCreatedEventHandler) Handle(ev event.Event) error {
	switch evType := ev.Type(); evType {
	case ArticleCreatedEventType:
		event := ArticleCreatedEvent{}
		return event.Process(ev)
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
	switch ev := ev.Type(); ev {
	case ArticleUpdatedEventType:
		fmt.Println(ev)
		return nil
	default:
		return errors.New("bad event")
	}
}

// NewArticleUpdatedEventHandler Creates an instance
func NewArticleUpdatedEventHandler() *ArticleUpdatedEventHandler {
	return &ArticleUpdatedEventHandler{}
}
