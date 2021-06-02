package events

import (
	event "github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/types"
	uuid "github.com/satori/go.uuid"
)

//Apply To create an aggregate in read-model
//1. Create new article state
//2. Add event to elastic-search
//3. Add read-model to elastic-search
func (articleCreatedEvent ArticleCreatedEvent) Apply(ev event.Event) error {
	payloadMapped, err := types.GetPayloadMapped(ev)
	if err != nil {
		return err
	}

	payloadMapped["aggregate_article_id"] = uuid.NewV4().String()

	// update returns nil error, so useless to init var
	newArticle, _ := articleCreatedEvent.update(payloadMapped)

	if ev.ShouldBeStored() {
		err := articleCreatedEvent.storeEventToElastic(newArticle)
		if err != nil {
			return err
		}
	}

	return articleCreatedEvent.storeReadModel(newArticle)
}

//Apply To update an aggregate in read-model
//1. Get aggregate from elastic-search
//2. update article state
//3. Update to elastic-search
func (articleUpdatedEvent ArticleUpdatedEvent) Apply(ev event.Event) error {
	payloadMapped, err := types.GetPayloadMapped(ev)
	if err != nil {
		return err
	}

	articleUpdatedEvent.AggregateArticleID = payloadMapped["aggregate_article_id"].(string)

	articleFromElastic, err := articleUpdatedEvent.update(payloadMapped)
	if err != nil {
		return err
	}

	if ev.ShouldBeStored() {
		err := articleUpdatedEvent.storeEventToElastic(articleFromElastic)
		if err != nil {
			return err
		}
	}

	return articleUpdatedEvent.storeReadModel(articleFromElastic)
}
