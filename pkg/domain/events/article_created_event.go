package events

import (
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

// ArticleCreatedEvent Overrides /pkg/event/event.go/EventProcessor interface
type ArticleCreatedEvent struct {
	db.Article
}

// Key is event name
func (articleCreatedEvent ArticleCreatedEvent) Key() string {
	return "ArticleCreatedEvent"
}

// Payload is event content
func (articleCreatedEvent ArticleCreatedEvent) Payload() event.Payload {
	return event.Payload{
		Content: db.Article{
			ID:          articleCreatedEvent.Article.ID,
			Title:       articleCreatedEvent.Article.Title,
			Description: articleCreatedEvent.Article.Description,
		},
	}
}

// Apply an event to get new article
func (articleCreatedEvent ArticleCreatedEvent) Apply(articles []db.Article) []db.Article {
	articles = append(articles, articleCreatedEvent.Payload().Content.(db.Article))
	return articles
}
