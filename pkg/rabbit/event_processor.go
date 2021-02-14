package rabbit

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

//EventProcessorParams Stores Body (events.ConsumeMessage type) and EventBus
type EventProcessorParams struct {
	EventType    string
	EventPayload interface{}
	EventBus     *event.EventBus
	StoreInDB    bool
}

//EventProcessor Receives serialized event and dispatch it to event handler
func EventProcessor(eProcessor *EventProcessorParams) error {
	eventImpl := event.NewEventImpl(eProcessor.EventType,
		eProcessor.EventPayload, eProcessor.StoreInDB)

	return eProcessor.EventBus.Dispatch(eventImpl)
}
