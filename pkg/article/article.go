package article

import (
	"github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/events"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/state"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/event"
)

func FakeEvents() []*event.EventImpl {
	var evs []*event.EventImpl

	createdEventID1 := event.NewEventImpl(&events.ArticleCreatedEvent{
		ID:          0,
		Title:       "test",
		Description: "yes",
	})

	createdEventID2 := event.NewEventImpl(&events.ArticleCreatedEvent{
		ID:          1,
		Title:       "test",
		Description: "yes",
	})

	updatedEvent := event.NewEventImpl(&events.ArticleUpdatedEvent{
		ID:          0,
		Title:       "test updated",
		Description: "yes updated",
	})

	evs = append(evs, createdEventID1)
	evs = append(evs, createdEventID2)
	evs = append(evs, updatedEvent)

	return evs
}

func UpdateArticle(article database.Article) (*event.ArticlesAggregate, error) {
	evs := FakeEvents()
	state.CurrentArticles = event.NewFromEvents(evs)

	ev := event.NewEventImpl(&events.ArticleUpdatedEvent{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
	})
	err := domain.EventBus.Dispatch(ev)
	if err != nil {
		return &event.ArticlesAggregate{}, err
	}

	return state.CurrentArticles, nil
}
