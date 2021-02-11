package rabbit

import (
	"encoding/json"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

func deserializeToMAP(serialized string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(serialized), &data); err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}

//EventProcessorParams Stores Body (events.ConsumeMessage type) and EventBus
type EventProcessorParams struct {
	Body     []byte
	EventBus *event.EventBus
}

//EventProcessor Receives serialized event and dispatch it to event handler
func EventProcessor(eProcessor *EventProcessorParams) error {
	bodyMapped, err := deserializeToMAP(string(eProcessor.Body))
	if err != nil {
		return err
	}

	eventType := bodyMapped["event_type"].(string)
	eventPayload := bodyMapped["payload"]
	eventImpl := event.NewEventImpl(eventType, eventPayload)
	err = eProcessor.EventBus.Dispatch(eventImpl)
	if err != nil {
		return err
	}

	return nil
}
