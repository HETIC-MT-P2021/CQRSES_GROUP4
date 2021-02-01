package event

import (
	"github.com/jibe0123/CQRSES_GROUP4/domain/events"
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
)

// ArticlesAggregate aggregate.
type ArticlesAggregate struct {
	articles []db.Article

	changes []Event
	version int
}

// NewFromEvents is a helper method that creates a new ArticleAggregate
// from a series of events.
func NewFromEvents(evs []*EventImpl) *ArticlesAggregate {
	aggr := &ArticlesAggregate{}

	for _, event := range evs {
		aggr.On(event, false)
	}

	return aggr
}

// Articles returns the ArticleAggregate's []article.
func (aggr ArticlesAggregate) Articles() []db.Article {
	return aggr.articles
}

// NewEmpty creates an empty ArticleAggregate.
func NewEmpty() *ArticlesAggregate {
	aggr := &ArticlesAggregate{
		articles: []db.Article{},
		changes:  []Event{},
		version:  0,
	}
	return aggr
}

// Events returns the uncommitted events from the ArticleAggregate aggregate.
func (aggr ArticlesAggregate) Events() []Event {
	return aggr.changes
}

// Version returns the last version of the aggregate before changes.
func (aggr ArticlesAggregate) Version() int {
	return aggr.version
}

// Update an article.
func (aggr *ArticlesAggregate) Update(id int, newArticle *events.ArticleUpdatedEvent) error {
	// business logic here

	aggr.raise(NewEventImpl(&events.ArticleUpdatedEvent{
		ID:          newArticle.ID,
		Title:       newArticle.Title,
		Description: newArticle.Description,
	}))

	return nil
}

// Add an article.
func (aggr *ArticlesAggregate) Add(newArticle *events.ArticleCreatedEvent) error {
	aggr.raise(NewEventImpl(&events.ArticleCreatedEvent{
		ID:          newArticle.ID,
		Title:       newArticle.Title,
		Description: newArticle.Description,
	}))

	return nil
}

func (aggr *ArticlesAggregate) raise(ev *EventImpl) {
	aggr.changes = append(aggr.changes, ev)
	aggr.On(ev, true)
}

// On handles ArticleAggregate events on the ArticleAggregate aggregate.
func (aggr *ArticlesAggregate) On(ev *EventImpl, new bool) {
	switch e := ev.Payload().(type) {
	case *events.ArticleCreatedEvent:
		article := db.Article{
			ID:          e.ID,
			Title:       e.Title,
			Description: e.Description,
		}
		aggr.articles = append(aggr.articles, article)

	case *events.ArticleUpdatedEvent:
		index := e.ID
		article := db.Article{
			ID:          e.ID,
			Title:       e.Title,
			Description: e.Description,
		}
		aggr.ReplaceElementAt(index, article)
	}

	if !new {
		aggr.version++
	}
}

// ReplaceElementAt Replace article by another at a specified index
func (aggr *ArticlesAggregate) ReplaceElementAt(index int, newArticle db.Article) {
	for i, article := range aggr.articles {
		if article.ID == index {
			aggr.articles[i] = newArticle
			break
		}
	}
}
