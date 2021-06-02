package rabbit

import (
	"encoding/json"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

// EventProcessor Stores Body (events.ConsumeMessage type) and EventBus
type EventProcessor struct {
	Message   []byte
	EventBus  *event.EventBus
	StoreInDB bool
}

// ApplyEventProcessor Unmarshal body and send it to EventProcessor
func (eProcessor EventProcessor) ApplyEventProcessor() (*ConsumeMessage, error) {
	var consumeMSG *ConsumeMessage
	err := json.Unmarshal(eProcessor.Message, &consumeMSG)
	if err != nil {
		return consumeMSG, err
	}

	eventImpl := event.NewEventImpl(consumeMSG.EventType,
		consumeMSG.Payload, eProcessor.StoreInDB)

	return consumeMSG, eProcessor.EventBus.Dispatch(eventImpl)
}

// ApplyEvents Unmarshal body and send it to EventProcessor
func (eProcessor EventProcessor) ApplyEvents(aggregateArticleID string) error {
	elasticImpl := elasticsearch.NewElasticRepository(elasticsearch.ElasticClient)
	evsFromElastic, err := elasticImpl.LoadEvents(aggregateArticleID)
	if err != nil {
		return err
	}

	for _, evElastic := range evsFromElastic {
		body, err := json.Marshal(evElastic)
		if err != nil {
			return err
		}

		eProcessor := &EventProcessor{
			Message:   body,
			EventBus:  eProcessor.EventBus,
			StoreInDB: false,
		}
		_, err = eProcessor.ApplyEventProcessor()

		fmt.Printf("elastic event : %s\n", evElastic)
	}

	return err
}
