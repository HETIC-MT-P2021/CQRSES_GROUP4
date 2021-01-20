package event

import db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"

type Event struct {
	Process EventProcessor
}

// EventProcessor Make process on event
// Key() allow to get Event name
// Payload() allow to get event content
// Apply() allow to apply event to get the new state of an entity
type EventProcessor interface {
	Key() string
	Payload() Payload
	Apply([]db.Article) []db.Article
}

// Payload allow to get event content
type Payload struct {
	Content interface{}
}
