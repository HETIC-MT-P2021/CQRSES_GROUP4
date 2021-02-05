package events_handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/state"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

type ArticleCreatedEventHandler struct{}

func (ch ArticleCreatedEventHandler) Handle(ev event.Event) error {

	switch ev := ev.Payload().(type) {
	case *events.ArticleCreatedEvent:
		log.Println(ev)
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
		fmt.Println(ev)
		/*aggr := state.CurrentArticles

		err := aggr.Update(1, ev)
		if err != nil {
			return err
		}*/

		// if no errors push on db
		return nil
	default:
		return errors.New("bad event")
	}

}

func NewArticleUpdatedEventHandler() *ArticleUpdatedEventHandler {
	return &ArticleUpdatedEventHandler{}
}
