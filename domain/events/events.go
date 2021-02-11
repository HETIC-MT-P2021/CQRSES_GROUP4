package events

// ConsumeMessage sds
type ConsumeMessage struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
}

var ArticleCreatedEventType = "ArticleCreatedEvent"
var ArticleUpdatedEventType = "ArticleUpdatedEvent"

// ArticleCreatedEvent Event to create an article
type ArticleCreatedEvent struct {
	ID          string `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ArticleUpdatedEvent Event to update an article
type ArticleUpdatedEvent struct {
	ID          string `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
