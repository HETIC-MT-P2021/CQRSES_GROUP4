package event

import (
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/pkg"
)

// EventBus Contains handlers
type EventBus struct {
	handlers map[string]EventHandler
}

// Event General event to override to create custom events
type Event interface {
	Type() string
	Payload() interface{}
}

// NewEventBus Initialize empty handlers in bus
func NewEventBus() *EventBus {
	eventBus := &EventBus{
		handlers: make(map[string]EventHandler),
	}

	return eventBus
}

// AddHandler to bus
func (eventBus EventBus) AddHandler(handler EventHandler, event interface{}) error {
	typeName := pkg.TypeOf(event)
	if _, ok := eventBus.handlers[typeName]; ok {
		return fmt.Errorf("Event handler already exists")
	}

	eventBus.handlers[typeName] = handler
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
	Content interface{}
}

// NewEventImpl Initialize an Event implementation
func NewEventImpl(eventContent interface{}) *EventImpl {
	return &EventImpl{
		Content: eventContent,
	}
}

// Type Returns event type
func (event EventImpl) Type() string {
	return pkg.TypeOf(event.Content)
}

// Payload Returns event content
func (event EventImpl) Payload() interface{} {
	return event.Content
}
