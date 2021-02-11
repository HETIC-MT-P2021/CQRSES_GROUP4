package events

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
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

func storeEventToElastic(article database.Article) error {
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)
	newEvent := database.Event{
		ID:        uuid.NewV4().String(),
		EventType: ArticleCreatedEventType,
		CreatedAt: createdAt,
		Payload:   article,
	}

	return elasticsearch.StoreEvent(newEvent)
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
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	article := database.Article{
		ID:          uuid.NewV4().String(),
		Title:       payloadMapped["title"].(string),
		Description: payloadMapped["description"].(string),
	}

	if err = storeEventToElastic(article); err != nil {
		return err
	}

	if err = elasticsearch.StoreReadmodel(article); err != nil {
		return err
	}

	return nil
}

//Process To update an aggregate in read-model
//1. Get aggregate from elastic-search
//2. update article state
//3. Update to elastic-search
func (event ArticleUpdatedEvent) Process(ev event.Event) error {
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	article := database.Article{
		ID:          payloadMapped["aggregate_article_id"].(string),
		Title:       payloadMapped["title"].(string),
		Description: payloadMapped["description"].(string),
	}

	if err = storeEventToElastic(article); err != nil {
		return err
	}

	err = elasticsearch.UpdateReadmodel(payloadMapped["aggregate_article_id"].(string), article)
	if err != nil {
		return err
	}

	return nil
}
