package event

import (
	"fmt"
)

// EventBus Contains handlers
type EventBus struct {
	handlers map[string]EventHandler
}

// Event General event to override to create custom events
type Event interface {
	Type() string
	Payload() interface{}
	ShouldBeStored() bool
}

// NewEventBus Initialize empty handlers in bus
func NewEventBus() *EventBus {
	eventBus := &EventBus{
		handlers: make(map[string]EventHandler),
	}

	return eventBus
}

// AddHandler to bus
func (eventBus EventBus) AddHandler(handler EventHandler, eventType string) error {
	if _, ok := eventBus.handlers[eventType]; ok {
		return fmt.Errorf("Event handler already exists")
	}

	eventBus.handlers[eventType] = handler
	return nil
}

// Dispatch Calls good event process
func (eventBus EventBus) Dispatch(event Event) error {
	if handler, ok := eventBus.handlers[event.Type()]; ok {
		return handler.Handle(event)
	}
	return fmt.Errorf("Handler doesn't exist")
}

// EventImpl Overrides Event
type EventImpl struct {
	EventType string
	Content   interface{}
	StoreInDB bool
}

// NewEventImpl Initialize an Event implementation
func NewEventImpl(eventType string, eventContent interface{}, storeInDB bool) *EventImpl {
	return &EventImpl{
		EventType: eventType,
		Content:   eventContent,
		StoreInDB: storeInDB,
	}
}

//Type Returns event type
func (event EventImpl) Type() string {
	return event.EventType
}

//Payload Returns event content
func (event EventImpl) Payload() interface{} {
	return event.Content
}

//ShouldBeStored Returns bool to be abble to know if should be stored on DB
func (event EventImpl) ShouldBeStored() bool {
	return event.StoreInDB
}
