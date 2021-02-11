package events

import (
	"encoding/json"

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

//Action To make to event
type Action interface {
	Process()
}

//Process To create an aggregate in read-model
//Create new article state
//Add to elastic-search
func (event ArticleCreatedEvent) Process(ev event.Event) error {
	payloadMapped, err := getPayloadMapped(ev)
	if err != nil {
		return err
	}

	aggregateArticleID := uuid.NewV4()
	article := database.Article{
		ID:          aggregateArticleID.String(),
		Title:       payloadMapped["title"].(string),
		Description: payloadMapped["description"].(string),
	}

	err = elasticsearch.StoreReadmodel(article)
	if err != nil {
		return err
	}

	return nil
}

func (event ArticleUpdatedEvent) Process() {

}
