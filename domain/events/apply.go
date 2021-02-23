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

	// GetOne returns nil error, so useless to init var
	newArticle, _ := event.getOne()

	newArticle.ID = payloadMapped["aggregate_article_id"].(string)
	newArticle.Title = payloadMapped["title"].(string)
	newArticle.Description = payloadMapped["description"].(string)

	if ev.ShouldBeStored() {
		articlePayload := event.payloadToArticle(payloadMapped)
		event.storeEventToElastic(articlePayload)
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

	articleFromElastic, err := event.getOne()
	if err != nil {
		return err
	}
	articleFromElastic.Title = payloadMapped["title"].(string)
	articleFromElastic.Description = payloadMapped["description"].(string)

	if ev.ShouldBeStored() {
		articlePayload := event.payloadToArticle(payloadMapped)
		event.storeEventToElastic(articlePayload)
	}

	return event.storeReadModel(articleFromElastic)
}
