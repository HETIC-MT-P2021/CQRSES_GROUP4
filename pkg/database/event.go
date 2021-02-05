package database

// RequestCreateEvent request to create an event on elasticsearch database
type RequestCreateEvent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EventType   string `json:"event_type"`
}

// Event data model on elasticsearch database (index called event-store)
type Event struct {
	ID        string  `json:"event_id"`
	EventType string  `json:"event_type"`
	CreatedAt string  `json:"created_at"`
	Payload   Article `json:"payload"`
}
