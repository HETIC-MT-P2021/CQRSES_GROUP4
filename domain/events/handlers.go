package events

import (
	"errors"
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

type ArticleCreatedEventHandler struct{}

func (eHandler ArticleCreatedEventHandler) Handle(ev event.Event) error {

	switch ev := ev.Payload().(type) {
	case *ArticleCreatedEvent:
		log.Println(ev)
		return nil
	default:
		return errors.New("bad event")
	}
}

func NewArticleCreatedEventHandler() *ArticleCreatedEventHandler {
	return &ArticleCreatedEventHandler{}
}

type ArticleUpdatedEventHandler struct{}

func (eHandler ArticleUpdatedEventHandler) Handle(ev event.Event) error {
	switch ev := ev.Payload().(type) {
	case *ArticleUpdatedEvent:
		fmt.Println(ev)
		return nil
	default:
		return errors.New("bad event")
	}

}

func NewArticleUpdatedEventHandler() *ArticleUpdatedEventHandler {
	return &ArticleUpdatedEventHandler{}
}
