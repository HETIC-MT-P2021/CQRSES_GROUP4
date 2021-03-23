package events

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

//Apply To create an aggregate in read-model
//1. Create new article state
//2. Add event to elastic-search
//3. Add read-model to elastic-search
func (event ArticleCreatedEvent) Apply(ev event.Event) error {
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	payloadMapped["aggregate_article_id"] = event.AggregateArticleID

	// update returns nil error, so useless to init var
	newArticle, _ := event.update(payloadMapped)

	if ev.ShouldBeStored() {
		event.storeEventToElastic(newArticle)
	}

	return event.storeReadModel(newArticle)
}

//Apply To update an aggregate in read-model
//1. Get aggregate from elastic-search
//2. update article state
//3. Update to elastic-search
func (event ArticleUpdatedEvent) Apply(ev event.Event) error {
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	event.AggregateArticleID = payloadMapped["aggregate_article_id"].(string)

	articleFromElastic, err := event.update(payloadMapped)
	if err != nil {
		return err
	}

	if ev.ShouldBeStored() {
		event.storeEventToElastic(articleFromElastic)
	}

	return event.storeReadModel(articleFromElastic)
}
