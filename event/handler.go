package event

// EventHandler Allows to manage event
type EventHandler interface {
	Handle(Event) error
}
