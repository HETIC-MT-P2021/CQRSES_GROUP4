package events

import (
	"encoding/json"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/article"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/deserialize"
	uuid "github.com/satori/go.uuid"
)

func getPayloadMapped(ev event.Event) (map[string]interface{}, error) {
	eventPayload := ev.Payload()
	buffer, err := json.Marshal(eventPayload)
	if err != nil {
		return map[string]interface{}{}, err
	}

	payload := string(buffer)
	return deserialize.ToMAP(payload)
}

//Action To make to event
type Action interface {
	Process()
}

//Process To create an aggregate in read-model
//1. Create new article state
//2. Add event to elastic-search
//3. Add read-model to elastic-search
func (event ArticleCreatedEvent) Process(ev event.Event) error {
	fmt.Println("Action creation")
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	create := article.Create{
		EventType: ArticleUpdatedEventType,
	}

	if ev.ShouldBeStored() {
		articlePayload := create.PayloadToArticle(payloadMapped)
		create.StoreEventToElastic(articlePayload)
	}

	// GetOne returns nil error, so useless to init var
	newArticle, _ := create.GetOne()

	aggregateArticleID := payloadMapped["aggregate_article_id"].(string)
	if aggregateArticleID == "" {
		aggregateArticleID = uuid.NewV4().String()
	}
	newArticle.ID = aggregateArticleID
	newArticle.Title = payloadMapped["title"].(string)
	newArticle.Description = payloadMapped["description"].(string)

	return create.Store(newArticle)
}

//Process To update an aggregate in read-model
//1. Get aggregate from elastic-search
//2. update article state
//3. Update to elastic-search
func (event ArticleUpdatedEvent) Process(ev event.Event) error {
	fmt.Println("Action update")
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	aggregateArticleID := payloadMapped["aggregate_article_id"].(string)
	update := article.Update{
		AggregateArticleID: aggregateArticleID,
		EventType:          ArticleUpdatedEventType,
	}

	if ev.ShouldBeStored() {
		articlePayload := update.PayloadToArticle(payloadMapped)
		update.StoreEventToElastic(articlePayload)
	}

	articleFromElastic, err := update.GetOne()
	if err != nil {
		return err
	}

	articleFromElastic.Title = payloadMapped["title"].(string)
	articleFromElastic.Description = payloadMapped["description"].(string)

	return update.Store(articleFromElastic)
}
