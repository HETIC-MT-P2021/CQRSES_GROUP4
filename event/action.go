package event

import (
	"encoding/json"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/deserialize"
)

func GetPayloadMapped(ev Event) (map[string]interface{}, error) {
	eventPayload := ev.Payload()
	buffer, err := json.Marshal(eventPayload)
	if err != nil {
		return map[string]interface{}{}, err
	}

	payload := string(buffer)
	return deserialize.ToMAP(payload)
}

//EventApplyer To make to event
type EventApplyer interface {
	Apply(Event) error

	payloadToArticle(map[string]interface{}) database.Article
	getOne(map[string]interface{}) (database.Article, error)
	storeReadModel(database.Article) error
	storeEventToElastic(database.Article) error
}
