package events

import (
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

// ArticleUpdatedEvent Overrides /pkg/event/event.go/EventProcessor interface
type ArticleUpdatedEvent struct {
	ArticleID  int
	NewArticle db.Article
}

// Key is event name
func (articleUpdatedEvent ArticleUpdatedEvent) Key() string {
	return "ArticleUpdatedEvent"
}

// Payload is event content
func (articleUpdatedEvent ArticleUpdatedEvent) Payload() event.Payload {
	return event.Payload{
		/*Content: db.Article{
			ID:          articleUpdatedEvent.Article.ID,
			Title:       articleUpdatedEvent.Article.Title,
			Description: articleUpdatedEvent.Article.Description,
		},*/
	}
}

// Apply an event to get edited article
func (articleUpdatedEvent ArticleUpdatedEvent) Apply(articles []db.Article) []db.Article {
	for index, article := range articles {
		if article.ID == articleUpdatedEvent.ArticleID {
			articles[index] = articleUpdatedEvent.NewArticle
			break
		}
	}

	return articles
}
