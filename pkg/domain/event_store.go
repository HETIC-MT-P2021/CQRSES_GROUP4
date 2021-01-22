package event

import (
	"errors"

	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/commands"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events"
	ev "github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

type EntityImpl struct {
	Changes []ev.Event
}

func (entity EntityImpl) AddEvent(eventProcessor ev.EventProcessor) {
	event := ev.Event{
		Process: eventProcessor,
	}
	entityImpl.Changes = append(entity.Changes, event)
}

var entityImpl EntityImpl = EntityImpl{
	Changes: []ev.Event{
		ev.Event{
			Process: events.ArticleCreatedEvent{
				db.Article{
					ID:          "1",
					Title:       "First created article",
					Description: "Here is first article",
				},
			},
		},
	},
}

// SaveArticle in database
func SaveArticle(command commands.Command) error {
	switch command.Type {
	case "CreateArticle":
		event := events.ArticleCreatedEvent{
			command.Payload.(commands.CreateArticleCommand).Article,
		}

		entityImpl.AddEvent(event)
		break
	case "UpdateArticle":
		articleID := command.Payload.(commands.UpdateArticleCommand).ArticleID
		newArticle := command.Payload.(commands.UpdateArticleCommand).Article
		event := events.ArticleUpdatedEvent{
			ArticleID:  articleID,
			NewArticle: newArticle,
		}

		entityImpl.AddEvent(event)
		break
	default:
		return errors.New("Nothing happened...")
	}

	return nil
}

// LoadArticles from database
func LoadArticles() []db.Article {
	var articles []db.Article = []db.Article{}

	changes := entityImpl.Changes
	for _, change := range changes {
		articles = change.Process.Apply(articles)
	}

	return articles
}
