package rabbit

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/deserialize"
)

//EventProcessorParams Stores Body (events.ConsumeMessage type) and EventBus
type EventProcessorParams struct {
	Body     []byte
	EventBus *event.EventBus
}

//EventProcessor Receives serialized event and dispatch it to event handler
func EventProcessor(eProcessor *EventProcessorParams) error {
	bodyMapped, err := deserialize.ToMAP(string(eProcessor.Body))
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
